package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	flock "github.com/runscripts/run/flock"
	"github.com/runscripts/run/utils"
)

// Show this help message.
func help() {
	utils.LogInfo(
		`Usage:
	run [OPTION] [SCOPE:]SCRIPT

Options:
	-c, --clean     clean out all scripts cached in local
	-h, --help      show this help message, then exit
	-i INTERPRETER  run script with interpreter(e.g., bash, python)
	-I, --init      create configuration and cache directory
	-u, --update    force to update the script before run
	-v, --view      view the content of script, then exit
	-V, --version   output version information, then exit

Examples:
	run pt-summary
	run github:runscripts/scripts/pt-summary

Report bugs to <https://github.com/runscripts/run/issues>.`,
	)
	utils.LogInfo("\n")
}

// Initialize and exit if -I or --init is given.
func initialize() {
	for _, arg := range os.Args {
		if arg == "-I" || arg == "--init" {
			if utils.IsRunInstalled() {
				utils.LogInfo("Run is already installed\n")
			} else {
				if os.Geteuid() != 0 {
					utils.LogError("Root privilege is required\n")
					os.Exit(1)
				}
				// Create script cache directory.
				err := os.MkdirAll(utils.DATA_DIR, 0777)
				if err != nil {
					utils.ExitError(err)
				}
				// Download run.conf, VERSION and run.1.gz from master branch.
				err = utils.Fetch(utils.MASTER_URL+"run.conf", utils.CONFIG_PATH)
				if err != nil {
					utils.ExitError(err)
				}
				err = utils.Fetch(utils.MASTER_URL+"VERSION", utils.DATA_DIR+"/VERSION")
				if err != nil {
					utils.ExitError(err)
				}
				err = utils.Fetch(utils.MASTER_URL+"man/run.1.gz", "/usr/share/man/man1/run.1.gz")
				if err != nil {
					utils.ExitError(err)
				}
			}
			os.Exit(0)
		}
	}
}

// Main function of the command run.
func main() {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	initialize()

	// If run is not installed.
	if utils.IsRunInstalled() == false {
		utils.LogError("Run is not installed yet. You need to 'run --init' as root.\n")
		os.Exit(1)
	}

	// Show help message if no parameter given.
	if len(os.Args) == 1 {
		help()
		return
	}

	// Parse configuration and runtime options.
	config, err := utils.NewConfig()
	if err != nil {
		utils.ExitError(err)
	}
	options, err := utils.NewOptions(config)
	if err != nil {
		utils.ExitError(err)
	}

	// If print help message.
	if options.Help {
		help()
		return
	}

	// If output version information.
	if options.Version {
		version, err := ioutil.ReadFile(utils.DATA_DIR + "/VERSION")
		if err != nil {
			utils.ExitError(err)
		}
		utils.LogInfo("Run version %s\n", version)
		return
	}

	// If clean out scripts.
	if options.Clean {
		utils.LogInfo("Do you want to clear out the script cache? [Y/n] ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "Y" || answer == "y" {
			utils.Exec([]string{"sh", "-x", "-c", "rm -rf " + utils.DATA_DIR + "/*"})
		}
		return
	}

	// If not script given.
	if options.Fields == nil {
		utils.LogError("The script to run is not specified\n")
		os.Exit(1)
	}

	// Ensure the cache directory has been created.
	cacheID := options.CacheID
	cacheDir := utils.DATA_DIR + "/" + options.Scope + "/" + cacheID
	err = os.MkdirAll(cacheDir, 0777)
	if err != nil {
		utils.ExitError(err)
	}

	// Lock the script.
	lockPath := cacheDir + ".lock"
	err = flock.Flock(lockPath)
	if err != nil {
		utils.LogError("%s: %v\n", lockPath, err)
		os.Exit(1)
	}

	// Download the script.
	scriptPath := cacheDir + "/" + options.Script
	_, err = os.Stat(scriptPath)
	if os.IsNotExist(err) || options.Update {
		err = utils.Fetch(options.URL, scriptPath)
		if err != nil {
			utils.ExitError(err)
		}
	}

	// If view the script.
	if options.View {
		flock.Funlock(lockPath)
		utils.Exec([]string{"cat", scriptPath})
	}

	// Run the script.
	flock.Funlock(lockPath)
	if options.Interpreter == "" {
		utils.Exec(append([]string{scriptPath}, options.Args...))
	} else {
		utils.Exec(append([]string{options.Interpreter, scriptPath}, options.Args...))
	}
}

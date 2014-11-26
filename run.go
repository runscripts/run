package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/runscripts/run/utils"
)

// Current version of run.
const VERSION = "0.1.0"

// Show this help message.
func help() {
	utils.LogInfo(
		`Usage:
	run [OPTION] [SCOPE:]SCRIPT

Options:
	-c, --clean     clean out all scripts cached in local
	-h, --help      show this help message, then exit
	-i INTERPRETER  run script with interpreter(e.g., bash, python)
	-I, --init      init the directories to install run
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

// Main function of the command run.
func main() {

	// If init run.
	for _, argument := range os.Args {
		if argument == "-I" || argument == "--init" {
			if utils.IsRunInstalled() {
				utils.LogInfo("Run is already installed. No need to init again.")
			} else {
				// Download and put it in /etc/runscripts.yml
				err := utils.Fetch(utils.RUN_YML_URL, utils.CONFIG_PATH)
				if err != nil {
					utils.LogError("Can't download from %s\n", utils.RUN_YML_URL)
					panic(err)
				}
				// Mkdir /var/lib/runscripts/
				mask := syscall.Umask(0)                       // Refer to http://studygolang.com/topics/33
				err = os.MkdirAll(utils.DATA_DIR, os.ModePerm) // 0777
				if err != nil {
					utils.LogError("Error MkdirAll  %s", utils.DATA_DIR)
					panic(err) // TODO: prompt "sudo"
				}
				defer syscall.Umask(mask)
				// Cp run to /usr/bin/run
				utils.Exec([]string{"sh", "-c", "cp ./run " + utils.RUN_PATH})
			}
			return
		}
	}

	// If run is not installed, prompt "sudo ./run --init".
	if utils.IsRunInstalled() == false {
		utils.LogInfo("Run is not installed yet. Please \"sudo ./run --init\".\n")
		return
	}

	// Show help message if no parameter given.
	if len(os.Args) == 1 {
		help()
		return
	}

	// Parse configuration and runtime options.
	config := utils.NewConfig()
	options := utils.NewOptions(config)

	// If print help message.
	if options.Help {
		help()
		return
	}

	// If output version information.
	if options.Version {
		utils.LogInfo("Run version %s\n", VERSION)
		return
	}

	// If clean out scripts.
	if options.Clean {
		utils.LogInfo("Do you want to clear out the script cache? [Y/n] ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "Y" || answer == "y" {
			utils.Exec([]string{"sh", "-c", "rm -rf " + utils.DATA_DIR + "/*"})
		}
		return
	}

	// If not script given.
	if options.Fields == nil {
		utils.LogError("Missing script to run\n")
		return
	}

	// Ensure the cache directory has been created.
	cacheID := options.CacheID
	cacheDir := utils.DATA_DIR + "/" + options.Scope + "/" + cacheID
	err := os.MkdirAll(cacheDir, 0777)
	if err != nil {
		utils.LogError("Can't mkdir %s\n", cacheDir)
		panic(err)
	}

	// Lock the script.
	lockPath := cacheDir + ".lock"
	utils.Flock(lockPath)

	// Download the script.
	scriptPath := cacheDir + "/" + options.Script
	_, err = os.Stat(scriptPath)
	if os.IsNotExist(err) || options.Update {
		err = utils.Fetch(options.URL, scriptPath)
		if err != nil {
			utils.LogError("Can't download/update %s\n", scriptPath)
			panic(err)
		}
	}

	// If view the script.
	if options.View {
		utils.Funlock(lockPath)
		utils.Exec([]string{"cat", scriptPath})
	}

	// Run the script.
	utils.Funlock(lockPath)
	if options.Interpreter == "" {
		utils.Exec(append([]string{scriptPath}, options.Args...))
	} else {
		utils.Exec(append([]string{options.Interpreter, scriptPath}, options.Args...))
	}

}

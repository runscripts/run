package main

import (
	"fmt"
	"os"

	"github.com/runscripts/runscripts/utils"
)

// Current version of runscripts.
const VERSION = "0.1.0"

// Show this help message.
func help() {
	utils.LogInfo(
		`Usage:
	run [OPTION] [SCOPE:]SCRIPT

Options:
	-h, --help      show this help message, then exit
	-i INTERPRETER  run script with interpreter(e.g., bash, python)
	-u, --update    force to update the script before run
	-v, --view      view the content of script, then exit
	-V, --version   output version information, then exit
	-c, --clean     clean out all scripts cached in local

Examples:
	run pt-summary
	run github:runscripts/scripts/pt-summary

Report bugs to <https://github.com/runscripts/runscripts/issues>.`,
	)
	utils.LogInfo("\n")
}

// Main function of run command.
func main() {

	// Show help message if no parameter given.
	if len(os.Args) == 1 {
		help()
		return
	}

	// Write default /etc/runscripts.yml if it doesn't exist
	if utils.IsFileExist(utils.CONFIG_PATH) == false {
		utils.LogInfo("%s doesn't exist, write default configuration file\n", utils.CONFIG_PATH)
		utils.WriteDefaultConfig()
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

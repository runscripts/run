package main

import (
	"fmt"
	"os"

	"github.com/runscripts/runscripts/utils"
)

const VERSION = "0.1.0"

// Refer to psql help format
func help() {
	utils.LogInfo(
		`Usage:
  run [OPTION]... [[SCOPE:]FIELD[/FIELD]...]

Options:
  -h, --help      show this help, then exit
  -i INTERPRETER  run the script with INTERPRETER (e.g., bash, python)
  -u, --update    use the network to update the script before run it
  -v, --view      output the script content, then exit
  -V, --version   output version information, then exit
  --clean         clear out all the scripts cached in local

For SCOPE and FIELD, check the manual of run (man run).

Report bugs to <https://github.com/runscripts/runscripts/issues>.`,
	)
	utils.LogInfo("\n")
}

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}

	config := utils.NewConfig()
	options := utils.NewOptions(config)

	if options.Help {
		help()
		return
	}

	if options.Version {
		utils.LogInfo("run %s\n", VERSION)
		return
	}

	if options.Clean {
		utils.LogInfo("Do you want to clear out the script cache? [Y/n] ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "Y" || answer == "y" {
			utils.Exec([]string{"sh", "-c", "rm -rf " + utils.DATA_DIR + "/*"})
		}
		return
	}

	if options.Fields == nil {
		utils.LogError("run: missing target to execute\n")
		os.Exit(1)
	} else {
		// ensure the cache directory has been created
		cacheID := options.CacheID
		cacheDir := utils.DATA_DIR + "/" + options.Scope + "/" + cacheID
		err := os.MkdirAll(cacheDir, 0777)
		if err != nil {
			utils.LogError("cannot mkdir %s\n", cacheDir)
			panic(err)
		}
		// lock the script
		lockPath := cacheDir + ".lock"
		utils.Flock(lockPath)
		// update the script
		scriptPath := cacheDir + "/" + options.Script
		_, err = os.Stat(scriptPath)
		if os.IsNotExist(err) || options.Update {
			err = utils.Fetch(options.URL, scriptPath)
			if err != nil {
				utils.LogError("cannot create/update %s\n", scriptPath)
				panic(err)
			}
		}

		if options.View {
			utils.Funlock(lockPath)
			utils.Exec([]string{"cat", scriptPath})
		}

		if options.Intprt == "" {
			utils.Funlock(lockPath)
			utils.Exec(append([]string{scriptPath}, options.Args...))
		} else {
			utils.Funlock(lockPath)
			utils.Exec(append([]string{options.Intprt, scriptPath}, options.Args...))
		}
	}
}

package main

import (
	"fmt"
	"os"
	. "./utils"
)

const VERSION = "0.1.0"

// Refer to psql help format
func help() {
	LogInfo(
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
	LogInfo("\n")
}

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}

	config := NewConfig()
	options := NewOptions(config)

	if options.Help {
		help()
		return
	}

	if options.Version {
		LogInfo("run %s\n", VERSION)
		return
	}

	if options.Clean {
		LogInfo("Do you want to clear out the script cache? [Y/n] ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "Y" || answer == "y" {
			Exec([]string{"sh", "-c", fmt.Sprintf("rm -rf %s/*", DATA_DIR)})
		}
		return
	}

	if options.Fields == nil {
		LogError("run: missing target to execute\n")
		os.Exit(1)
	} else {
		// ensure the cache directory has been created
		cacheID := options.CacheID
		cacheDir := DATA_DIR + "/" + options.Scope + "/" + cacheID
		err := os.MkdirAll(cacheDir, 0777)
		if err != nil {
			LogError("cannot mkdir %s\n", cacheDir)
			panic(err)
		}
		// lock the script
		lockPath := cacheDir + ".lock"
		Flock(lockPath)
		// update the script
		scriptPath := cacheDir + "/" + options.Script
		_, err = os.Stat(scriptPath)
		if os.IsNotExist(err) || options.Update {
			err = Fetch(options.URL, scriptPath)
			if err != nil {
				LogError("cannot create/update %s\n", scriptPath)
				panic(err)
			}
		}

		if options.View {
			Funlock(lockPath)
			Exec([]string{"cat", scriptPath})
		}

		if options.Intprt == "" {
			Funlock(lockPath)
			Exec(append([]string{scriptPath}, options.Args...))
		} else {
			Funlock(lockPath)
			Exec(append([]string{options.Intprt, scriptPath}, options.Args...))
		}
	}
}

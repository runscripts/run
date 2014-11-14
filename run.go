package main

import "./utils"

const VERSION = "0.0.1"

func help() {
	// refer to psql help format
	utils.LogInfo(`
Usage:
  run [OPTION]... [SCOPE:][FIELD[/FIELD]...]

Options:
  -h, --help      show this help, then exit
  -i INTERPRETER  run the script with INTERPRETER (e.g., bash, python)
  -u, --update    use the network to update the script before run it
  -v, --view      output the script content, then exit
  --clean         clear out all the scripts cached in local
  --version       output version information, then exit

For SCOPE and FIELD, check the manual of run (man run).

Report bugs to <https://github.com/runscripts/runscripts/issues>.
	`)
	utils.LogInfo("\n")
}

func main() {
	help()
}

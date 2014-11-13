package main

import "./utils"

// Refer psql help format
func help() {
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

For SCOPE and FIELD, please check the manual of run (man run).
	`)
	utils.LogInfo("\n")
}

func main() {
	help()
}

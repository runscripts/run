
package main

import (
  "os"
  "flag"
  "strings"

  log "github.com/Sirupsen/logrus"
  "github.com/runscripts/runscript/utils"
)

// Implement the command of runscript
func main() {
  log.Info("Start runscript program")

  // Parse command-line arguments
  args := os.Args
  updateFlag := flag.Bool("u", false, "Force to update the local script")
  helpFlag := flag.Bool("h", false, "Check out the usage of runscript")
  //helpFlag := flag.Bool("help", false, "Check out the usage")
  flag.Parse()

  // Todo: filter the flags in args array

  // Exit if the args is too less or too more
  argsLength := len(args)
  if (argsLength < 2 || argsLength > 3) {
    log.WithFields(log.Fields{
      "argsLength": argsLength,
    }).Error("Args length is too less or too more, exit")

    log.Warn(help())
    os.Exit(0)
  }

  // Exit if we pass flag "-h"
  if (*helpFlag) {
    log.Warn(help())
    os.Exit(0)
  }

  // Todo: implement forcing to download scripts
  log.Warn(*updateFlag)

  // Create $HOME/runscript/ if it doesn't exist
  runscriptDirectory := os.Getenv("HOME") + "/runscript/"
  utils.MkdirIfNotExist(runscriptDirectory)

  // Create $HOME/runscript/official/ if it doesn't exist
  officialDirectory := os.Getenv("HOME") + "/runscript/official/"
  utils.MkdirIfNotExist(officialDirectory)

  // Check if it's official scripts
  scriptArray := strings.Split(args[1], "/")
  if len(scriptArray) == 1 {
    log.Error("It's official")
  } else {
    log.Error("It's fucking mine or yours")
  }

  command := "ls"
  output := utils.ExecCommandReturnOutput(command)
  log.Info(output)

  //args := []string{"git", "clone", "git@github.com:runscripts/script.git", "/Users/tobe/Desktop/temp/script"}
  commandArgs := []string{"echo", "hello"}
  utils.SyscallCommandNotReturn(commandArgs)

}


// Print out the usage of runscript
func help() string {
  return `Usage:

    run $script

    Example: run test-network

    run $user/$script

    Example: run tobegit3hub/while-loop

    run -u $script

    Exmaple: run -u test-network
  `
}

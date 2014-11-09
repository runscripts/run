
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
  runscriptDirectory := os.Getenv("HOME") + "/runscript"
  utils.MkdirIfNotExist(runscriptDirectory)

  // Create $HOME/runscript/official/ if it doesn't exist
  officialDirectory := os.Getenv("HOME") + "/runscript/official"
  utils.MkdirIfNotExist(officialDirectory)

  // Check if it's official scripts
  scriptArray := strings.Split(args[1], "/")
  if len(scriptArray) == 1 {
    // check if $HOME/runscript/official/script exists
    repositoryDirectory := runscriptDirectory + "/official/script"
    scriptName := scriptArray[0]
    scriptPath := repositoryDirectory + "/" + scriptName

    if (utils.IsDirectoryExist(repositoryDirectory)) {
      // Todo: git pull to update the repository
    } else {
      // Pull the official repository
      log.Warn("Directory " + repositoryDirectory + " doesn't exist")
      command := "git clone https://github.com/runscripts/script.git " + repositoryDirectory
      output := utils.ExecCommandReturnOutput(command)
      log.Info(output)
    }

    if (utils.IsFileExist(scriptPath)) {
      // Find the script, run it
      commandArgs := []string{scriptPath}
      utils.SyscallCommandNotReturn(commandArgs)
    } else {
      // The script doesn't exist
      log.Error("The script " + scriptName + " doesn't exist in official repository")
      os.Exit(0)
    }

  } else {

    // It' a user script
    userName := scriptArray[0]
    repositoryDirectory := runscriptDirectory + "/" + userName
    scriptName := scriptArray[1]
    scriptPath := repositoryDirectory + "/" + scriptName

    if (utils.IsDirectoryExist(repositoryDirectory)) {

    } else {
      // Mkdir the directory and pull the repository
      utils.MkdirIfNotExist(repositoryDirectory)
      log.Warn("Directory " + repositoryDirectory + " doesn't exist")
      command := "git clone https://github.com/" + userName + "/script.git " + repositoryDirectory
      output := utils.ExecCommandReturnOutput(command)
      log.Info(output)
    }

    if (utils.IsFileExist(scriptPath)) {
      // Find the script, run it
      commandArgs := []string{scriptPath}
      utils.SyscallCommandNotReturn(commandArgs)
    } else {
      // The script doesn't exist
      log.Error("The script " + scriptName + " doesn't exist in " + userName + "'s repository")
      os.Exit(0)
    }
  }

  // Never reach here because it switches to execute other command
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

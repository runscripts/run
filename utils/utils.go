
package utils

import (
  "os"
  "os/exec"
  "syscall"

  log "github.com/Sirupsen/logrus"
)

// Mkdir directory if it doesn't exist
func MkdirIfNotExist(directory string) {
  if _, err := os.Stat(directory); os.IsNotExist(err) {
    log.Warn("No direcotry " + directory)
    // Todo: handle mkdir failure
    os.Mkdir(directory, 0777)
    log.Info("Mkdir " + directory)
  }
}

// Check if the file exists
func IsFileExist(file string) bool {
  if _, err := os.Stat(file); os.IsNotExist(err) {
    return false
  } else {
    return true;
  }
}

// Check if the directory exists
func IsDirectoryExist(directory string) bool {
  return IsFileExist(directory)
}

// Use exec library to execute command
// Will block and return the output when it's finished
func ExecCommandReturnOutput(command string) string {
  log.Info("Start to execute command " + command)

  cmd := exec.Command("bash", "-c", command)
  out, err := cmd.Output()

  if err != nil {
    log.Error("Error to execute command " + command)
    panic(err)
  }

  return string(out)
}

// Use syscall library to execute command
// Only accept command and arguments as string array
// Switch new process and not return
func SyscallCommandNotReturn(args []string) {
  //args2 := []string{"ls", "-a", "-l", "-h"}
  //args2 := []string{"git", "clone", "git@github.com:runscripts/script.git", "/Users/tobe/Desktop/temp/script"}

  log.Info("Start to execute command " + args[0])

  binary, lookErr := exec.LookPath(args[0])

  if lookErr != nil {
      log.Error("Error to find the command " + args[0])
      panic(lookErr)
  }

  env := os.Environ()

  execErr := syscall.Exec(binary, args, env)
  if execErr != nil {
      log.Error("Error to execute command " + args[0])
      panic(execErr)
  }

}

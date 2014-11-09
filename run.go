
package main

import (
  //"os"
  //"flag"

  "github.com/runscripts/runscript/utils"

  log "github.com/Sirupsen/logrus"
)

func help() {

}


func main() {
  log.Info("Start runscript program")

  command := "ls -al"
  output := utils.ExecCommandReturnOutput(command)
  log.Info(output)

  //args := []string{"git", "clone", "git@github.com:runscripts/script.git", "/Users/tobe/Desktop/temp/script"}
  args := []string{"echo", "hello"}
  utils.SyscallCommandNotReturn(args)

}

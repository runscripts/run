# Run

[![TravisCI status](https://secure.travis-ci.org/runscripts/run.png)](http://travis-ci.org/runscripts/run) [![GoDoc](https://godoc.org/github.com/runscripts/run?status.svg)](https://godoc.org/github.com/runscripts/run) [![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/runscripts/runscripts?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Introduction

**Run** is the script manager for developers.

* It helps to manage scripts with git
* It helps to re-use scripts natually
* It is highly integrated with Github
* If you're writing scrips, try it now

## How It Works

Normally we execute `run pt-summary` or `run github:runscripts/scripts/pt-summary`.

The command `run` will download the well-known [pt-summary](http://www.percona.com/doc/percona-toolkit/2.1/pt-summary.html) from Github and run locally.

## Install

* From Scratch (Go 1.3+)

  ```
  sudo GOPATH=$GOPATH make install
  ```

* From Binary

  Operating System | Architectures
  ---------------- | -------------
  Linux            | [386](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/linux_386/run), [amd64](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/linux_amd64/run), [arm](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/linux_arm/run)
  Mac OS           | [386](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/darwin_386/run), [amd64](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/darwin_amd64/run), [homebrew](https://github.com/tobegit3hub/homebrew-run)
  FreeBSD          | [386](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/freebsd_386/run), [amd64](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/freebsd_amd64/run)
  Debian/Ubuntu    | [386](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/deb/run_20141222-0.3.4_386.deb), [amd64](https://raw.githubusercontent.com/runscripts/run-release/master/0.3.4/deb/run_20141222-0.3.4_amd64.deb)

  Download binary according to your OS/Arch and place it in **$PATH** (like /usr/bin/). Then execute `sudo run --init`.

## Usage

Watch the [one-minute video](https://www.youtube.com/watch?v=WXUcJvrZP6M) before you're using it.

```
Usage:
	run [OPTION] [SCOPE:]SCRIPT

Options:
	-c, --clean     clean out all scripts cached in local
	-h, --help      show this help message, then exit
	-i INTERPRETER  run script with interpreter(e.g., bash, python)
	-I, --init      create configuration and cache directory
	-u, --update    force to update the script before run
	-v, --view      view the content of script, then exit
	-V, --version   output version information, then exit

Examples:
	run pt-summary
	run github:runscripts/scripts/pt-summary

Report bugs to <https://github.com/runscripts/run/issues>.
```

## Scripts

The official scripts are in [runscripts/scripts](https://github.com/runscripts/scripts).

> You can easily `run pt-summary` or `run -i python get-pip.py`.

Feel free to manage your scripts in Github and send pull-request to official scripts.


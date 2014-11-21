# RunScripts

[![TravisCI status](https://secure.travis-ci.org/runscripts/runscripts.png)](http://travis-ci.org/runscripts/runscripts) [![GoDoc](https://godoc.org/github.com/runscripts/runscripts?status.svg)](https://godoc.org/github.com/runscripts/runscripts) [![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/runscripts/runscripts?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Introducatoin

[RunScripts](https://github.com/runscripts/runscripts) is awaited manager for scripts, just like [homebrew](https://github.com/Homebrew/homebrew) and [bower](https://github.com/bower/bower).

It helps to manager your scripts and make resue of scripts much easier. After installing runscripts, you get the command `run`. Just `run pt-summary` and it will download the well-known pt-summary to run. If you push your scritps in Github, you can simply use it like `run github:runscripts/scripts/pt-summary`.


## Install

* From Scratch(Go 1.3+)

```
Download the source code and execute "sudo GOPATH=$GOPATH make install".
```

* For Linux Users

```
wget https://raw.githubusercontent.com/runscripts/runscripts/master/packages/linux_amd64/run && chmod +x run && sudo ./run --init
```

* For MacOS Users

```
wget https://raw.githubusercontent.com/runscripts/runscripts/master/packages/darwin_amd64/run && chmod +x run && sudo ./run --init
```

* ARM, FreeBSD and Others

```
Most platforms are well supported. Please download binary in packages.
```

## Usage

```
Usage:
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

Report bugs to <https://github.com/runscripts/runscripts/issues>.
```

## Scripts

We have provided [official scripts](https://github.com/runscripts/script) and everyone can easily `run pt-summary` and `run -i python get-pip.py`.

Feel free to manage your scripts in Github and send pull-request to official scripts.


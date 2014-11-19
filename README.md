# RunScripts

[![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/runscripts/runscripts?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) 

## Introducatoin

[RunScripts](https://github.com/runscripts/runscripts) is awaited manager for scripts, just like [homebrew](https://github.com/Homebrew/homebrew) and [bower](https://github.com/bower/bower).

It does two things for you. :sparkles: The first is forcing you to manage scripts with version control tools and :sparkles: the second is to make resue of scripts much easier.

## Install

### From Scratch

* Download the source code and execute `make install`

### Download Binary

| Operation System |     Package     |
|------------------|-----------------|
| Linux(amd64)     | [run](https://raw.githubusercontent.com/runscripts/runscripts/master/build/linux_amd64/run) |
| Linux(386)       | [run](https://raw.githubusercontent.com/runscripts/runscripts/master/build/linux_386/run) |
| Mac OS(amd64)    | [run](https://raw.githubusercontent.com/runscripts/runscripts/master/build/darwin_amd64/run) |
| Mac OS(386)      | [run](https://raw.githubusercontent.com/runscripts/runscripts/master/build/darwin_386/run) |

FreeBSD and ARM are well supported as well. Please checkout out [build](https://github.com/runscripts/runscripts/tree/master/build).

## Usage

We have simpified the usage of run. Please `run -h` for more usage.

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

Feel free to manager your scripts in Github and use `run` for convenience.


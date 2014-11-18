# RunScript
[![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/runscripts/runscripts?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

* [中文文档](README-zh.md)
* [Drone]()
* [Gitter]()


## Introducatoin

[RunScript](http://runscript.org) is awaited manager for scripts, just like [brew](https://github.com/Homebrew/homebrew) and [bower](https://github.com/bower/bower). It does two things for you. :sparkles: The first is forcing you to manage scripts with version control tools and :sparkles: the second is to make resue of scripts much easier.

Here is the typical scenario. If we want to use [pt-summary](http://www.percona.com/get/pt-summary), we should find out where it's, download it and run the script locally. But runscript does the same thing for you with command `run pt-summary`. Futhermore, anyone can share and re-use scripts just like `run tobegit3hub/start_seagull`.

## Usage

Runscript is implemented in go and you can install it with `go get github.com/runscripts/runscript`. Here're compiled packages for supported operation systems.

| Operation System |     Package     |
|------------------|-----------------|
| Linux(x64)       |    [runscript-1.0.0]()    |
| Linux(x86)       |    [runscript-1.0.0]()    |
| Mac OS(Not yet)  |                 |
| Windows(Not yet) |                 |

After installing runscript, you may get the command `run`. Try `run test_network` to run your first script.

If your Github username is *tobegit3hub*, create the repository named *script* and push your scripts right now. Then you can easily run your scrips like `run tobegit3hub/start_seagull`.

## How It Wors

All the processes of runscript are transparent and easy to understand.

When you execute the command `run`, runscript will find out if the script is in local machine. If it doesn't exist locally, runscript will download it from git server with `git clone`. That's why you should use version control tools to manage all your scripts before running them.

By default, runscript will find out the scipts from the repositories named *script* in github.com. We're planing to support private git hosting servers later.


## More Scripts

* [Official scripts](https://github.com/runscripts/script)
* [Tobegit3hub/script](https://github.com/tobegit3hub/script)


# Gosh - **WIP**

A shell written in go

## Features

Gosh includes four builtin commands:
- `cd <dir_name>`
  - running `cd` with no arguments changes the directory to $HOME
- `exit <exit_code(optional)>`
- `pwd`
- `ls`

Gosh comes with three builtin prompt items:
- user: displays the current user
- hostname: displays your hostname
- workdir: displays your current working directory
  - paths relative to $HOME are printed as `~/path`

## Install instructions

Download and install:
```bash
go get github.com/maxmcd/gosh
```

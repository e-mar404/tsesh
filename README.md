# tsesh

Command that lets you start tmux sessions on pre configured directories.

### Motivation

This is very much inspired by the tmux sessionizer script from @ThePrimeagen but
I found it lacking in a few aspects for **my** specific workflow. Since what I
wanted to add was a bit past what I am comfortable writing in bash I decided to
make it a go cli.

## Installation 

This project is not ready for production use yet but should still be able to 
build and run. You can try it out with the following cmd:

```bash
go install github.com/e-mar404/tsesh@v0.1.1
```

In the future it will also be available through nixpkgs/as a nix flake. Right
now the flake is only for development.

## Configuration

The file `.example.config.toml` contains all the defaults that the application
will run with. If there is not already a `config.toml` file under the directory
`$XDG_CONFIR_DIR/tsesh`, or `$HOME/.config/tsesh` depending on system, a new
default config will be created.

**Note**: Config uses `xdg.UserConfigDir()` so look at the function 
implementation for differences between UNIX/windows and how it determines which
directory it chooses.

### Conflicts between search paths and ignore patterns

It is possible that a directory put on the search path gets picked up by either
`ignore_pattern` and/or `ignore_hidden`. For such cases there is a preliminary
check to still expand that path if it is explicitly on the search path list.

For example, if your config looks like this:

```toml
[search]
paths = ['~/.config']
ignore_hidden = true
```

The directory `~/.config` will still be expanded even though it should've been
picked up by the `ignore_hidden` rule. However if there are more hidden
directories under that path those directories are expected to be ignored.

## Data storage

Bookmarks will be saved under `$XDG_DATA_DIR/tsesh/data.json`, or the
appropriate directory based on the system. There is no need for user interaction
for this file, but it is possible to have this file in vcs and be used in
different machines if you have similar directories you use regularly. 

## Usage

The intended use of the cli is to be able to manage tmux sessions and url
bookmarks attached to directories.

These are the available commands:

```
terminal sessionizer extending tmux

Usage:
  tsesh [flags]
  tsesh [command]

Available Commands:
  add         add a directory or a url to the current working directory
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        list all bookmarks set for current directory
  open        open a bookmark by index
  pin         pins the current directory to the top of the list when running tsesh
  rm          remove url provided from bookmark list for current working directory
  unpin       unpin directory from fuzzy finder if it is in the pinned list
  version     get tsesh cli version

Flags:
  -h, --help      help for tsesh
  -v, --verbose   show debug logging

Use "tsesh [command] --help" for more information about a command.
```

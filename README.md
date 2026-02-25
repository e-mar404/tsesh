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
go install github.com/e-mar404/tsesh@latest
```

In the future it will also be available through nixpkgs/as a nix flake. Right
now the flake is only for development.

## Configuration

Originally I was thinking of using lua for the configuration file but I do not
need all that functionality right now so I will do that later and will start building
with a toml file.

As of right now the file `.example.config.toml` contains the default
configuration. If there is not already a `config.toml` file under the directory
`$XDG_CONFIR_DIR/tsesh` or `$HOME/.config/tsesh` a new default config will be 
created. The place depends on which ENV vars are set.

**Note**: Config uses `os.UserConfigDir()` so look at the function 
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

Bookmarks will be saved in `.../tsesh/data.json`. Later the data dir
will be able to be changed but for now this is static.

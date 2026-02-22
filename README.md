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

## TODO 

### Stage 1

- testing
    - [x] tmux (since I found a cool new way to test exec.Command)
    - [x] search.go (have to make sure the file name is okay and returns proper
      session names)
    - [ ] bubbletea program?

- ci/cd
    - [x] check for failing tests
    - [x] check gofmt

- config file
    - [x] check for config file existence
    - [x] load config at startup

- data file
    - [ ] decide on file path
    - [ ] allow user to override datadir

- tsesh cmd
    - [ ] `add [path|url]` adds the current path to bookmarks
    - [ ] `bookmarks` will display saved bookmarks
    - [ ] `open [string|int]` open specified bookmark

### Stage 2

- [ ] if i am on a dir and I just want to go to a scratch session. I have been
creating a new window and then cd-ing into ~/code. I have all of my scratch
repos and sample files in there. When I want to test something it is always
the same place so find something for that
- [ ] deploy app with nix flake and not just develop capability
- [ ] check out session groups and see if I can implement a shortcut to add a new session to a group

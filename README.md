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

This is were the information for the configuration file will go, there is not a
set api yet but for now it will default to the following search paths:

- `~/`
- `~/project` 

Originally I was thinking of using lua for the configuration file but I do not
need the functionality right now so I will do that later and will start building
with a toml file.

## Data storage

Bookmarks will be saved in `$XDG_DATA_HOME/tsesh/data.json`. Later the data dir
will be able to be changed but for now this is the planned place where it will
go.

## TODO 

### Stage 1

- testing
    - [x] tmux (since I found a cool new way to test exec.Command)
    - [ ] find.go (have to make sure the file name is okay and returns proper
      session names)
    - [ ] bubbletea program?

- ci/cd
    - [ ] check for failing tests
    - [ ] check gofmt

- config file
    - [ ] check for config file existence
    - [ ] load config at startup

- data file
    - [ ] decide on file path
    - [ ] allow user to override datadir

- tsesh cmd
    - [ ] `add [path|url]` adds the current path to bookmarks
    - [ ] `bookmarks` will display saved bookmarks
    - [ ] `open [string|int]` open specified bookmark
    - [ ] `config` edit configuration

### Stage 2

- [ ] if i am on a dir and I just want to go to a scratch session. I have been
creating a new window and then cd-ing into ~/code. I have all of my scratch
repos and sample files in there. When I want to test something it is always
the same place so find something for that
- [ ] deploy app with nix flake and not just develop capability
- [ ] check out session groups and see if I can implement a shortcut to add a new session to a group

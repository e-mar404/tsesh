# tsesh

Command that lets you start tmux sessions on pre configured directories.

### Motivation

This is very much inspired by the tmux sessionizer script from @ThePrimeagen but
I found it lacking in a few aspects for my specific workflow. Since what I
wanted to add was a bit past what I am comfortable writing in bash I decided to
make it a go cli.

## Installation 

This project is not ready for production use yet but is still able to build and
run. You can try it out with the following cmd:

```bash
go install github.com/e-mar404/tsesh@latest
```

In the future it will also be available through nixpkgs/as a nix flake. Right
now the flake is only for development.

## Configuration

This is were the configuration will go, there is not a set api yet but it will
default to the following search paths:

- `~/`
- `~/project` 

It is not recommended to use config.lua to configure yet. This will come soon.

## TODO 

- testing
    - [x] tmux (since I found a cool new way to test exec.Command)
    - [ ] bubbletea program?

- Config file through lua
    - [x] specify which directories to search for and any patterns to ignore
    - [ ] specify list styling (delegate on bubble tea)
    - [ ] what to do before and after tmux session is started 

- tsesh cmd
    - [ ] `add [path|url]` adds the current path to bookmarks
    - [ ] `bookmarks` will display saved bookmarks
    - [ ] `open [string|int]` open specified bookmark

- [ ] if i am on a dir and I just want to go to a scratch session. I have been
  creating a new window and then cd-ing into ~/code. I have all of my scratch
  repos and sample files in there. When I want to test something it is always
  the same place so find something for that

- [ ] deploy app with nix flake and not just develop capability

- [ ] check out session groups and see if I can implement a shortcut to add a new session to a group

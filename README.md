# tsesh

Command that lets you start tmux sessions on pre configured directories.

### Motivation

This is very much inspired by the tmux sessionizer script from @ThePrimeagen but
I found it lacking in a few aspects for my specific workflow. Since what I
wanted to add was a bit past what I am comfortable writing in bash I decided to
make it a go cli.

## Wish list

- Config file
    - [ ] specify which directories to search for and any patterns to ignore
    - [ ] specify list styling (delegate on bubble tea)
    
- [ ] if i am on a dir and I just want to go to a scratch session I have been
  creating a new window and then cd-ing into ~/code. I have all of my scratch
  repos and sample files in there. When I want to test something it is always
  the same place so find something for that

# tsesh

I have really been enjoying just using a tmux sessionizer but I constantly find
it falls short on some things. The key thing that I think it is missing is
bookmarks! I want to be able to bookmark websites or other directories not
included on the regular search without expanding the search depth.

What I am currently thinking of is a mix between harpoon and my current tmux
sessionizer.

For now I want to recreate my current tmux-sessionizer bash script, it should
search for directories on a hard coded list of paths, serve them as a list and
then select a directory item and start a tmux session at that path.

Simple start then I will see how I can add what I want it to do little by
little.

The more I think about this the more I feel like I do want a terminal
sessionizer lol, I need to learn more about tmux.


## Wish list

- [ ] if i am on a dir and I just want to go to a scratch session I have been
  creating a new window and then cd-ing into ~/code. I have all of my scratch
  repos and sample files in there. When I want to test something it is always
  the same place so find something for that

### Using charm

I want to use [charmbracelet](https://github.com/charmbracelet/bubbletea) for
the ui so I should look at the following:

- example for exec (will be useful for executing the tmux command)
- example for composable views (useful for tui dashboard)
- example for auto complete (might be useful for filtering??, i do think that
  good smart filtering is going to be a bit hard and will require some good
  thought) for now I will leave it as simple list picker 

### Using open code for the prototype

I decided to use opencode to port my current bash script into a go program that
uses charm bubbletea. This worked pretty well. It gave me general idea of things
that I will need to consider.

So from the code that I saw it did there were 3 main components to it (in no
particular order of importance):

1. the model for the bubbletea program (this only contained the list)
2. a function to handle the discovery of the directories and appending them to
   the list element of the model which it did right before launching the program
   so it was not attached to the model type at all (that sounds ok I will have
   to see how it does with loading search paths from a config file and when
   bookmarks get introduced)
3. the tmux session handler

This tells me that I will prob want the following:

- to have the path-picker model (which will be different from the dashboard and
  it is currently the list model here).
- an abstraction over tmux so I dont have to keep calling exec.Command() for
  tmux stuff
- either a utils file or file management file that will let me manage and handle
  directory paths in a more granular way than just `filepath` and `os`

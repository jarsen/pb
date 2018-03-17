# PB - Personal gif/meme Database

PB is a simple CLI program that helps you store and search your favorite memes/gifs (or any URL, really).

The name stands for Peanut Butter, even though personally I pronounce GIF with a hard G. It also sounds like DB. It's also a little bit of a riff on the OS X utility `pbcopy`, since the `pb copy` command will copy to the system clipboard.

[![asciicast](https://asciinema.org/a/167213.png)](https://asciinema.org/a/167213)

```
Usage:
  pb [SEARCH_TERMS...] [flags]
  pb [command]

Available Commands:
  add         add a link and some description text
  copy        Copies the highest matching result to your system clipboard
  download    downloads all the image in the library to disk
  dump        dump the database into the specified txt file
  help        Help about any command
  list        list all the database entries
  load        adds all the images from a given file to the database
  nuke        deletes the entire database
  rm          removes an entry from the database
  search      search images from the database
  show        shows the image for the top result and copies url to clipboard
```

## Installation

```bash
brew install go # You need to have go installed.
# You need to set up a GOPATH. This is where go will put the executable binary.
# (You will probably want to add it to your bashrc/zshrc)
mkdir $HOME/go
export GOPATH=$HOME/go
# Finally, use go get to fetch the source and install it
go get github.com/jarsen/pb
```

## Todo / Feature Ideas

- [ ] Flags to change formatting of how entries are displayed (show ID, etc)
- [ ] Remove entries with short id
- [ ] Pick random entry
- [ ] Track usage
- [ ] List by LRU, MRU, etc.
- [ ] Better Usage documentation
- [ ] Allow user to customize location of db on disk
- [ ] Flags for fuzzy search
- [ ] Interactive command for adding images from a folder/tagging

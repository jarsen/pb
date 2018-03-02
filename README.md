# PB - Personal gif/meme Database

PB is a simple CLI program that helps you store and search your favorite memes/gifs (or any URL, really).

The name stands for Peanut Butter, even though personally I pronounce GIF with a hard G. It also sounds like DB.

```
pb - your own personal gif/meme database.

Become the envy of your friends and colleagues as you organize and search your favorite memes and animated gifs from your terminal.

Usage:
  pb [SEARCH_TERMS...] [flags]
  pb [command]

Available Commands:
  add         add a link and some description text
  copy        Copies the highest matching result to your system clipboard
  help        Help about any command
  list        list all the database entries
  load        adds all the images from a given file to the database
  nuke        deletes the entire database
  rm          removes an entry from the database
  search      search images from the database

Flags:
  -h, --help   help for pb

Use "pb [command] --help" for more information about a command.
```

## Todo / Feature Ideas

- [x] Add entries
- [x] Load entries from a file
- [x] Search entries
- [x] Find best match and copy to system clipboard
- [x] List all entries
- [ ] Flags to change formatting of how entries are displayed (show ID, etc)
- [ ] Remove entries
- [ ] Pick random entry
- [ ] Better README
- [ ] Better Usage documentation
- [ ] Allow user to customize location of db on disk
- [ ] Flags for fuzzy search
- [ ] iTerm 2 (and others?) let you display images...
- [ ] Command for downloading all the images to disk
- [ ] Interactive command for adding images from a folder/tagging
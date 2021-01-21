# searchweb

## Goal
I need a way of searching certain pages *on the fly*, while writing
texts or code. By typing in the command line, or a quick key stroke in
vim, I want to query my favourite pages.

## Usage
```console
me@neptune:~$ searchweb -page=<favorite-page> -config=<path/to/config.yaml> <keywords>
```
If no `-page` is given, it defaults to `duck`.

## Installation
1. Install this command with `go get`:
```console
me@neptune:~$ go get -u github.com/hikmet-kibar/cmd/searchweb
```
2. Create a yaml file with your pages, formatted like  [config.yaml](./config.yaml).

## Bash configuration
Add an environment variable to your `.bashrc` or `.bash_profile`:
```bash
export PAGES=$HOME/.config/searchweb/pages.yaml
```

Example on how to set up aliases in your `.bashrc` or `.bash_aliases`:
```bash
alias \
  ?="searchweb -page=duck -config=${PAGES}"         \
  ord="searchweb -page=ordnet -config=${PAGES}"     \
  korp="searchweb -page=korpus -config=${PAGES}"    \
  tysk="searchweb -page=tysk -config=${PAGES}"      \
  yt="searchweb -page=youtube -config=${PAGES}"     \
  github="searchweb -page=github -config=${PAGES}"
```

Now you can search your favourite pages with a short command as such:
```console
me@neptune:~$ yt my favorite topic
```

## Vim configuration
To search the word under the cursor while writing, you can add similar
mappings to your `.vimrc` file:
```vim
nnoremap <silent> <leader>d :exec "!searchweb -page=duck -config=${PAGES} "
                              \ . expand("<cword>")<cr> :redraw!<cr>
```


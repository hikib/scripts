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
export SEARCHYAML=$HOME/.config/searchweb/pages.yaml
```

Example on how to set up aliases in your `.bashrc` or `.bash_aliases`:
```bash
alias \
  ?="searchweb -page=duck -config=${SEARCHYAML}"         \
  ord="searchweb -page=ordnet -config=${SEARCHYAML}"     \
  korp="searchweb -page=korpus -config=${SEARCHYAML}"    \
  tysk="searchweb -page=tysk -config=${SEARCHYAML}"      \
  yt="searchweb -page=youtube -config=${SEARCHYAML}"     \
  github="searchweb -page=github -config=${SEARCHYAML}"
```

Now you can search your favourite pages with a short command as such:
```console
me@neptune:~$ yt my favorite topic
```

## Vim configuration
To search the word under the cursor while writing, you can add the
following function and similar mappings to your `.vimrc` file:
```vim
" Vim does not load bash_aliases. This is one workaround.
fun! SearchWeb(page)
  let keyword = expand("<cword>")
  let search = "!searchweb -page=" . a:page . " -config=" . $SEARCHYAML
  exec "silent " . search . " " . keyword
  exec "redraw!"
endfun

noremap <leader>d :call SearchWeb("duck") <CR>
noremap <leader>o :call SearchWeb("ordnet") <CR>
noremap <leader>k :call SearchWeb("korpus") <CR>
noremap <leader>t :call SearchWeb("tysk") <CR>
```


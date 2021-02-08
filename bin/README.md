# Shell/Bash scripts
Scripts, used as cli commands, to practice shell/bash scripting and
automate tasks or call other commands with certain flags.
Some of these scripts may be used as aliases instead.

## cheat
Returns the cheat sheet to a topic from [cheat.sh](https://cheat.sh). Keywords are
optional.
```console
$ cheat python <keyword>
```

## note
Quick export using Pandoc while writing notes in markdown. It calls
pandoc with specified flags. Primarily used within vim:
```vim
:note % something.pdf
```

## dato
Returns the current date in Danish, e.g. `mandag, den 08 februar 2021`.

Usage within vim:
```vim
:.!dato
```


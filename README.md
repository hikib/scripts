# Scripts
A collection of scripts used for various tasks. Primarily to solve
simple tasks, but also a way to practice and learn different programming
languages and their idioms.

## BIN
Scripts in the `bin/` directory are written in bash.

You can add the directory to `$PATH` in your `.bash_profile`:
```bash
export SCRIPTS=$REPOS/scripts/bin
export PATH=$PATH:$SCRIPTS
```

## CMD
Scripts in the `cmd/` directory are written in Go.

You can either install single commands with:
```console
me@neptune:~$ go get -u github.com/hikmet-kibar/scripts/cmd/<command-name>
```
or install all at once with:
```console
me@neptune:~$ go get -u github.com/hikmet-kibar/scripts/cmd/...
```


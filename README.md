# Scripts
A collection of scripts used for various tasks. Primarily to solve
simple tasks, but also a way to practice and learn different programming
languages and their idioms.

## BIN
Scripts in the `bin/` directory are written in Bash.

You can add the directory to `$PATH` in your `.bash_profile`:
```bash
export SCRIPTS=$REPOS/scripts/bin
export PATH=$PATH:$SCRIPTS
```

## CMD
Scripts in the `cmd/` directory are written in Go.

You can either install single commands with:
```bash
go install ./cmd/<command-name>
```
or run `cmd/ìnstall.sh` to install them all.


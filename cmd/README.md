# CLI commands with GO
All commands are intended for personal use. I used Go to write them
to practice and learn GO.

## searchweb
Runs given arguments as queries on pre-set web pages. If no `-page`
given, defaults to duckduckgo.com
```
Usage of searchweb:
  -page string
    ordnet, korpus, tysk, github, youtube (default "duckduckgo")
```

## journaling
Prints the current date and links to review files in markdown format to
the console. The output is used in vimwiki as a template for the
journal.

Example output:
```markdown
# 19.December 2020, Week 51, Saturday

## Reviews
- [Week 51](../pipelines/reviews/2020-W51.md)
- [December](../pipelines/reviews/2020-M12.md)
- [2020](../pipelines/reviews/2020.md)
```

To generate a review overview automatically, add the following to your
`.vimrc` file:
```vim
autocmd BufNewFile ~/vimwiki/diary/[0-9]*.md :.!journaling %
```

## reviews
Updates the `README.md` file in my vimwiki reviews directory.
Similar to `:VimwikiDiaryGenerateLinks`.

Example output:
```markdown
# Reviews

- [New Years Resolution](2021.md)
- [WEEK 04](2021-W04.md)
- [WEEK 02](2021-W02.md)
- [WEEK 01](2021-W01.md)
- [January](2021-M01.md)
```


```vim
autocmd BufWritePre ~/vimwiki/pipelines/reviews/README.md :1,$d | .!reviews %
```

## packup
Handles my files when writing job applications by moving the temporary
files from the 'docs' directory to the archive directory 'apps'.
The command takes 2 inputs:
1. Path to the main directory for applications
2. Company name


# Agent instructions
Keep this file agent-agnostic and focused on repository state.

## Rule: commits
- Use conventional commits: `feat:`, `fix:`, `docs:`, `chore:`.
- Do not reference AI agents, Copilot, or co-authorship in commits, docs, or metadata.
- Committing directly to `main` is acceptable for this repo.

## Rule: context
- At the end of each significant session, update `AGENTS.md` with the current state, architectural decisions, and open tasks.
- Overwrite this file so it always reflects the ground truth of the repo.

## Rule: environment
- Use absolute command paths when needed, e.g. `/usr/bin/open` on macOS.
- Do not assume editor or shell aliases exist in the runtime environment.

## Rule: secrets
- Never commit secrets, config files, or database files.

## Rule: code style
- Keep markdown paragraphs short and single-line.
- Use hyphens in filenames, not spaces or underscores.
- Executable scripts should not use extensions.
- Use sentence case for all headings and titles.
- Refer to `CAVEMAN.md` for the repository prompt style and response consistency.

## Current architecture
- This repo is a collection of scripts for various tasks.
- `cmd/` contains Go command packages.
- `uni/` contains cross-platform utilities for Linux/macOS.
- `win/` contains PowerShell helpers for Windows.

## Current state
- The repo is maintained as a personal toolbox rather than a packaged product.
- Existing utilities cover file operations, logging, search, and environment management.
- `README.md` is the primary repository overview.

## Pending todos
- Add usage documentation for `cmd/downloadmover`.
- Consolidate script style and example usage across the repo.
- Consider adding versioning or release notes for stable script sets.

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

## Rule: response style
- Apply compressed, direct language in natural language chat messages, not code snippets or code blocks.
- Remove articles, auxiliary verbs, and redundant words while preserving meaning.
- Use simple, common words and prefer simpler synonyms.
- Always remove:
  - Articles: a, an, the
  - Auxiliary verbs: is, are, was, were, am, be, been, being, have, has, had, do, does, did
  - Common prepositions when meaning stays clear: of, for, to, in, on, at
  - Pronouns when context is clear: it, this, that, these, those
  - Pure intensifiers: very, quite, rather, somewhat, really, extremely
- Always keep:
  - All nouns (people, places, things, concepts)
  - All main verbs (actions, not auxiliaries)
  - All adjectives that add meaning
  - All numbers and quantifiers (at least, approximately, more than, 15, many)
  - Uncertainty qualifiers (what sounded like, appears to be, seems, might)
  - Critical prepositions that change meaning (from, with, without, stuck to)
  - Time and frequency words (every Tuesday, weekly, daily, always, never)
  - Names, titles, and technical terms
- Be smart about:
  - Keeping prepositions that define relationships while removing those that are only grammatical.
  - Keeping in/on/at when they specify location or position.
  - Removing is/are/was/were unless passive voice meaning matters.
  - Keeping negations such as not, no, never, without.

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

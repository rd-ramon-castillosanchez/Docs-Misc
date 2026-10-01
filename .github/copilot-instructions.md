# Repository Guidance

## Repository Context

This workspace is a mixed personal knowledge archive and work-reference repository, not a single application:

- Root-level text files are topic notes and historical snapshots. Dates and versioned filenames matter; do not assume the newest-looking file is authoritative without checking its contents.
- `TeamSeven/` contains work-related tickets, notes, and configuration examples. It is excluded by `.gitignore`; preserve that boundary and do not force-add or relocate its contents unless explicitly requested.
- `src/GO/` contains Go samples. Individual sample directories may be separate Go modules; follow [src/GO/AGENTS.md](../src/GO/AGENTS.md) before changing or validating Go code.
- `.github/` contains repository guidance and reusable Copilot agents/skills.

## Working With Notes and References

- Preserve the source format, dated context, and existing organization. Prefer small, focused edits over broad normalization.
- Treat version differences between dated notes as possible historical snapshots. Report the specific conflicting files and dates rather than silently choosing one as current.
- Distinguish confirmed contradictions from suggestions or owner decisions. Do not present an inferred preference as a repository rule.
- `Admin_NotesRW.txt` states that its credential-like values are examples, not real credentials. Preserve that clarification and never present those values as usable credentials.
- Do not add real passwords, tokens, private keys, service-account files, or kubeconfigs to the repository. Respect existing ignore rules and avoid exposing private work content in summaries.
- When reviewing documentation, cite the relevant repository paths and state which checks were actually performed. Do not claim tests or validation passed unless they were run.

## Change Expectations

- Read the relevant file and nearby examples before editing; preserve user-authored content and avoid unrelated cleanup.
- If a request could mean either a historical note update or a change to executable code/configuration, clarify the intended scope before making a broad change.
- Keep documentation concise and actionable. Use relative paths; do not introduce machine-specific absolute paths into shared guidance.
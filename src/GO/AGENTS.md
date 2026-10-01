# Go Guidance

These instructions apply to Go code under `src/GO/`.

## Project Layout and Validation

- Inspect the target sample's `go.mod` before editing or running commands. Sample directories can be independent modules; run Go commands from the relevant module directory and do not assume `src/GO/` is one module.
- Keep changes scoped to the requested sample or package. Match its existing structure and supported Go version as declared by its module.
- Format changed Go files with `gofmt` and run the narrowest relevant tests. When appropriate, run `go test ./...` from that module's root and report the exact validation performed.
- Treat commands that connect to Kafka, Kubernetes, GKE, or other external systems as operational actions. Do not run them or change remote resources without explicit authorization.

## Go Conventions

- Prefer clear, idiomatic Go and the smallest maintainable change. Keep variables local to the functions that need them unless shared package state is necessary and justified.
- A lowercase package-level identifier is unexported to importing packages but is available to all files in the same package. An uppercase initial makes an identifier exported outside its package; it does not merely make it visible to other files.
- Handle errors explicitly and preserve useful context. Add or update focused tests when behavior changes.
- Avoid unnecessary dependencies, global mutable state, and comments that only restate the code.

## Credentials and Configuration

- Never commit service-account keys, kubeconfigs, access tokens, or real credentials. Use documented placeholders or test fixtures, and keep local secrets outside the repository.
- Make configuration and file paths explicit; do not bake a developer-specific absolute path into reusable sample code.

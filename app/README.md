# app

The `app` package creates a new blockchain applications. The main components are:
- `app/app.go`: Wire up Cosmos SDK modules, as well as the `x/blog` module, into a blockchain app.
- `app/regen/cmd`: A package to create a CLI, which is the main interface for starting to run a blockchain node.

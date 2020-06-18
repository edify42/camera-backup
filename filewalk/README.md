# filewalk

This package needs more thought.

## handler

We need to pass files found by Walker to handler.go, which in term will do the
extra magic for us before (or maybe it does it) we push files to the database.

## walker

Implements the file.Walk method to step through the file system when needed.

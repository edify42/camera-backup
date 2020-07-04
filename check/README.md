# check

`check` will loop through the file system first (specified by `location` within
the `config.yaml` file) and search every single file matching the `include`
setting against a record within the datastore.

Needs to combine the functions within the `../filewalk` package and the
`localstore` package.

**Note to self: DO THIS LAST** -- RIP this note to self!

## Iteration logic

`check` will almost perform another `init` style search, except this time it's
also looking into `localstore` and _checking_ to see if the files match what
they should be.

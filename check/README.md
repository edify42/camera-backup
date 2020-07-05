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

## it's a search problem

We need to match two sources and verify they are equal. The first is checking
all files on the system to see if there is a record in the datastore.

The second is checking all files on the datastore and seeing if they exist on
the file system.

It's likely that we'll heavily use the datastore to keep the state of
everything as we're running check.
# backup-genie

Backup Genie is a simple Open Source command line tool which uses an sqlite
datasource to check all files in a particular directory and see if there exists
a file somewhere else on your system.

## First Scan

Run the `backup-genie init` command to create a config file and initial
datastore.

By default the config will be loaded to the current working directory.

Another option is to store the config in your `$HOME` directory. `backup-genie`
will look for config in the current working directory and fall back to $HOME.
Otherwise, it will prompt you to run the init command.

### Config file

The `config.yaml` file stores the location of the database used to 

## Future Features

Remote backup to cloud storage (paid).
Could be useful (backblaze)[https://github.com/kothar/go-backblaze]
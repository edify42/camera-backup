# backup-genie

Backup Genie is a simple Open Source command line tool which uses an sqlite
datasource to check all files in a particular directory and see if there exists
a file somewhere else on your system which is stored in the backup

## Nomenclature

'target' is the place we're looking for to see new files or changes to files.
'backup' is the place which stores the `config.yaml` in the root

## Commands

`backup-genie init` - initialises a backup location
`backup-genie scan` - scan the backup location for changes to files
`backup-genie sync` - sync the target location with the the backup location
`backup-genie check` - check a target location against the backup for duplicate files

## User Flow

User will have a source they want to check against a targetted backup location.
E.g. An SD card with a bunch of images (new and old) and they might have
already done a backup of a few of them when they copied over some wedding
photos.

### Init flow

See the [first scan](http://yeah.com) section. 

### Scenario 1: User Checks Target

User is unsure if the target is 'clean' with no missing or unaccounted for
files.

`backup-genie check` runs a sync check against the target. This will discover
any new, missing or duplicated items at the target.

### Scenario 2: User wants to copy files to target

`backp-genie sync` runs a scan of the source location and compares all files
to the data source. Any new or missing files will be copied over into a date
stamped directory and catalogued in the data store.

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
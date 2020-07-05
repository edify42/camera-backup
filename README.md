# backup-genie

Backup Genie is a simple Open Source command line tool which uses an sqlite
datasource to check all files in a particular directory and see if there exists
a file somewhere else on your system which is stored in the backup

## Nomenclature

'target' is the place we're looking for to see new files or changes to files.
'backup' is the place which stores the `config.yaml` in the root

## Commands

`backup-genie init` - initialises a backup location
`backup-genie check` - check the backup location for changes to files. Prints
duplicates, missing files from the datastore (new files) and deleted files
(in datastore but missing from file system) - shows mismatches of hashes too.
`backup-genie search` - search a target location against the backup for
pre-existing files
`backup-genie sync` - sync the target location with the the backup location.
Copies all files already in the backup location automatically to a sanely
named folder in the backup location.

### Possible future commands

`backup-genie clean` - removes dups from file system...
`backup-genie pretty` - sorts files for you in a consistent manner (possibly
from the file name / type and/or metadata last modified by month).

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

The `config.yaml` file stores the location of the backup (and also the
sqlstore.db file). You can manage all your config files in one location
(can be different from where the sqlstore.db file is kept). We'll **always**
keep the sqlstore.db file with the targetted backup location though.

**By design, we always drop** `config.yaml` **in the same location as where**
`init` **is run**

## MVP Definition

I'd be happy versioning 1.0.0 if we have the 4 main functions defined above.

## Future Features

Remote backup to cloud storage (paid).
Could be useful (backblaze)[https://github.com/kothar/go-backblaze]
Organise would be amazing if we could cleverly group or label photos in some
way.

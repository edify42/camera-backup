# search OR scan

This fun bit of logic will take the current directory (or from input) and
compare the files here to what's stored to an existing backup.

For `search` to work, please ensure you have the sqlstore up to date (run the
`check` and/or `sync` command if necessary).

## usage

Imagine you've mounted your sd-card from your camera and need to do a _search_
for new files that have been created.

```bash
cd /mnt/sdcard/
camera-backup search --location $HOME/Pictures/
```

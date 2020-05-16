# sqlstore design

Some quick notes for sqlstore.db to help me develop this

Should be quite simple. We only need two tables:

1. Metadata table which stores simple information about the data table.
1. Data table which contains all the records (files) we're interested in.

## Metadata Table

See the database screeema

(subject to changes while we're in PoC phase)

## Data Table

See the database screeeama.

## Simple SQL Viewer

I did my development on my Ubuntu 18.04 PC and I used `sqlitebrowser` to
debug the data in the local database.

`sudo apt-get install sqlitebrowser`

I found it quite nice in that you could open up the table viewer with in a
single bash statement `sqlitebrowser sqlstore.db`.

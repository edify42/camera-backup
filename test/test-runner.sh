#!/usr/bin/env bash

input="${1:-test}"
test_folder='./test'

# program must be build beforehand!
# which backup-genie || echo 'please build and compile the program before continuing' && exit 1

if [ "$input" = 'clean' ]; then
  rm -rf "$test_folder/testdata_start"
  rm -rf "$test_folder/testdata_new_files"
  rm -rf "$test_folder/testdata_mix_new_old"
  exit 0
fi

echo "+++ Will execute the integration test"



## Initialise the fake test data (starting data)
mkdir -p "$test_folder/testdata_start"

### Generate md5 data
python3 "$test_folder/integration.py" generate --md5 --count 20 --destination "$test_folder/testdata_start"

### Generate etag data
python3 "$test_folder/integration.py" generate --no-md5 --count 20 --destination "$test_folder/testdata_start"

### do the backup to the starting data

## Initialise the new test data (unique from start data)
mkdir -p "$test_folder/testdata_new_files"
python3 "$test_folder/integration.py" generate --destination "$test_folder/testdata_new_files"


## Initialise the mixed test data (mix of existing data and unique new data)
mkdir -p "$test_folder/testdata_new_files"
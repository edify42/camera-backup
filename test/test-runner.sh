#!/usr/bin/env bash

input="${1:-test}"
test_folder='./test'

if [ "$input" = 'clean' ]; then
  rm -rf "$test_folder/testdata_start"
  rm -rf "$test_folder/testdata_new_files"
  rm -rf "$test_folder/testdata_mix_new_old"
  exit 0
fi

echo "+++ Will execute the integration test"



## Initialise the fake test data (starting data)
mkdir -p "$test_folder/testdata_start"

python3 "$test_folder/integration.py" start "$test_folder/testdata_start"

## Initialise the new test data (unique from start data)


## Initialise the mixed test data (mix of existing data and unique new data)
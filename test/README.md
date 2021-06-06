# test

To separate the programming language of the testing logic vs the functional
logic of the program, the integration tests have been written in python3.

## integration test

This folder currently contains the logic to perform the integration testing
for this repo.

### high level design

A python function will create a bunch of files in a some folders in this
directory. It will run the `backup-genie` app which is compiled from source
and then compare the expected values with what this test suite expects.

```bash



```

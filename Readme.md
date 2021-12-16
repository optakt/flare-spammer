This is to spam a Flare C-chain local net where one address which has balance creates a number of new addresses and transfers balance.

This helps in creating transactions in the C-chain and thereby triggering block creation.

This is to do debugging whether the network is running as expected. 

The network performance can be checked using the logs. 

## Prerequisite

1. `git clone https://github.com/flare-foundation/flare`


2. `cd flare` and `git checkout snowman-plusplus-test`


3. Edit the `genesis/genesis_local.go` code with the current repository's `genesis/genesis_local.go`


4. `./scripts/build.sh`


5. `./scripts/launch_localnet.sh`

This will launch a 5 node local net. 

## How to run

To run, simply use:

`go run ./cmd/main.go`

This will create 10 addresses and their corresponding transactions where certain amount of balance will be transferred.

Check in the log files of the `Flare` repo to see logs and errors.


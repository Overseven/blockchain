# TRY network

[![Go Report Card](https://goreportcard.com/badge/github.com/Overseven/blockchain)](https://goreportcard.com/report/github.com/Overseven/blockchain)
[![Coverage Status](https://coveralls.io/repos/github/Overseven/blockchain/badge.svg?branch=main)](https://coveralls.io/github/Overseven/blockchain?branch=main)

[Description](https://github.com/Overseven/blockchain/blob/main/docs/description.md)

## Build
1) Client:
   ``
   cd network/client/main/
   &&
   go build -o client.exe
   ``
2) Wallet generator:
   ``
   cd utility/wallet_generator
   &&
   go build -o generator.exe
   ``
3) Node:
   ``
   cd network/node/main/
   &&
   go build -o node.exe
   ``
3) Miner:
   ``
   cd network/miner/main/
   &&
   go build -o miner.exe
   ``
3) Coordinator:
   ``
   cd network/coordinator
   &&
   go build -o coordinator.exe
   ``
   
## Usage 
Generate wallet:

``
utility/wallet_generator/
``

Start as node:

``
node/node.exe -config <filename>
``

Start as coordinator:

``
coordinator/coordinator.exe -config <filename>
``


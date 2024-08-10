#!/bin/bash
set -e

git clone https://github.com/cosmos/cosmos-sdk.git
cd cosmos-sdk
docker build -t simapp .
cd ..
mkdir -p simapp

# Server:
docker run -it -p 26657:26657 -p 26656:26656 -v simapp:/root/.simapp simapp simd init test-chain
# TODO: need to set validator in genesis so start runs
#docker run -it -p 26657:26657 -p 26656:26656 -v private:/root/.simapp simapp simd start

echo "Adding and listing keys"
# Client: (Note the simapp binary always looks at ~/.simapp we can bind to different local storage)
docker run -it -p 26657:26657 -p 26656:26656 -v simapp:/root/.simapp simapp simd keys add key_1
docker run -it -p 26657:26657 -p 26656:26656 -v simapp:/root/.simapp simapp simd keys list

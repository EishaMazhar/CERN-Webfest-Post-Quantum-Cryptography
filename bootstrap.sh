#!/bin/bash
set -e

echo ">>> Cloning repos"
git clone https://github.com/kudelskisecurity/crystals-go
git clone https://github.com/cosmos/cosmos-sdk.git

echo ">>> Copying files"
cp -R dilithium/ cosmos-sdk/crypto/.
cp -R tx/tx.go cosmos-sdk/types/.
cp ante.go cosmos-sdk/x/auth/ante

mkdir -p cosmos-sdk/x/crypto/dilithium/test/
cp  crypto_dilithium_test/dilithium_test.go cosmos-sdk/x/crypto/dilithium/test/.

echo ">>> Building SDK"
cd cosmos-sdk
docker build -t simapp .

echo ">>> Adding and listing keys"
cd ..
mkdir -p simapp

# Server:
docker run -it -p 26657:26657 -p 26656:26656 -v simapp:/root/.simapp simapp simd init test-chain
# TODO: need to set validator in genesis so start runs
#docker run -it -p 26657:26657 -p 26656:26656 -v private:/root/.simapp simapp simd start

echo ">>> Adding and listing keys"
# Client: (Note the simapp binary always looks at ~/.simapp we can bind to different local storage)
docker run -it -p 26657:26657 -p 26656:26656 -v simapp:/root/.simapp simapp simd keys add key_1
docker run -it -p 26657:26657 -p 26656:26656 -v simapp:/root/.simapp simapp simd keys list

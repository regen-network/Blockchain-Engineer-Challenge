#!/bin/sh
set -e

export BINARY="regen"
export CHAIN_ID="regen-localnet-0"
export CHAIN_DIR="./data"
export GENESIS_VAL_COINS="1000000000validatortoken,1000000000000000000stake"

VALIDATOR_MNEMONIC="bulb picture cereal caught salt snack awake drift guide rigid injury mercy remove shield crazy select skin piece fever fury above rebuild trouble swamp"

rm -rf $CHAIN_DIR
mkdir -p $CHAIN_DIR

$BINARY init regen-localnet-0 --home $CHAIN_DIR --chain-id $CHAIN_ID --overwrite
$BINARY config keyring-backend test --home $CHAIN_DIR
$BINARY config chain-id $CHAIN_ID --home $CHAIN_DIR
$BINARY config broadcast-mode block --home $CHAIN_DIR
$BINARY config output json --home $CHAIN_DIR

$BINARY keys delete validator --home $CHAIN_DIR -y
echo "$VALIDATOR_MNEMONIC" | $BINARY keys add validator --recover --home $CHAIN_DIR
$BINARY add-genesis-account $($BINARY keys show validator -a --home $CHAIN_DIR) $GENESIS_VAL_COINS --home $CHAIN_DIR
$BINARY gentx validator 900000000stake --home $CHAIN_DIR --chain-id $CHAIN_ID
$BINARY collect-gentxs --home $CHAIN_DIR

$BINARY start --home $CHAIN_DIR --log_level info

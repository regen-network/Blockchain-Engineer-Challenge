# Implementing task

## Run node

For first step command is from [here](https://docs.cosmos.network/v0.46/run-node/run-node.html)
For next steps I followed this [instruction](https://docs.cosmos.network/v0.46/run-node/run-node.html)

1. Build node

    make build

2. Init config of node with chain-id `my-test-chain`. `monadfix` is arbitrary string.

        build/regen init monadfix --chain-id my-test-chain

- `regen` used insted of `simd`
- `--home` flag to specify dirictory of config. But then you should pass it to every command.

3. Add account to network

        build/regen keys add monadfix --keyring-backend test

- where `monadfix` is custom name.

4. Fetch address. We got `cosmos1728x3mclm0f980828tf74nq3vqvaxkkkdfwuc8`

        build/regen keys show monadfix -a --keyring-backend test

5. Add money to your address of your account. See address in step `4`

        build/regen add-genesis-account cosmos1728x3mclm0f980828tf74nq3vqvaxkkkdfwuc8 100000000000stake

6. Create validator

        build/regen gentx monadfix 100000000stake --chain-id my-test-chain --keyring-backend test

        build/regen collect-gentxs

7. Update gas config in `app.toml`

Modify `minimum-gas-prices` to `"0stake"`

8. Run local blockchain network

        build/regen start

## Interact with blog

0. Set chain-id

        build/regen config chain-id my-test-chain

1. Create post

        build/regen tx blog create-post cosmos1728x3mclm0f980828tf74nq3vqvaxkkkdfwuc8 "first-post" "title" "body body body" --keyring-backend test -y

2. List all posts

        build/regen query blog list-post

3. Comment has key is comprised by `comment` + `slug` + `_` + `uuidV7`.
   Prefix `comment` is added by KVStore.

4. Create comment

          build/regen tx blog create-comment cosmos1728x3mclm0f980828tf74nq3vqvaxkkkdfwuc8 "first-slug" "comment comment" --keyring-backend test -y

5. List comments of the particular slug

          build/regen query blog list-comments "first-slug"

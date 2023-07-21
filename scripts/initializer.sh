tokenworldd init test --chain-id=tokenworld
tokenworldd config keyring-backend test
tokenworldd keys add validator
tokenworldd add-genesis-account $(tokenworldd keys show validator -a) 1000000000gas
tokenworldd gentx validator 700000000gas --chain-id tokenworld
tokenworldd collect-gentxs

sed -i -e 's/stake/gas/g' ~/.tokenworld/config/genesis.json

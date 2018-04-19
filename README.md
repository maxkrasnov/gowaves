> Sorry, at the moment, you can only generate addresses and check balance

## Using

```
go get https://github.com/maxkrasnov/gowaves
```

Waves Platform cryptographic practical [details](https://github.com/wavesplatform/Waves/wiki/Cryptographic-practical-details)

## Examples

Basic example

```
node := GoWaves.Node("https://nodes.wavesnodes.com", GoWaves.TESTNET)
account := node.Account() // generate public,private keys, address and seed
account.GetPrivateKey() // return private key
account.GetAddress() // return address
account.GetBalance() // return balance and confirmations
```

Get address by seed

```
node.AccountBySeed("test1 test2 test3 test4") // generate public,private keys, address
```

Get address by private key

```
node.AccountByPrivateKey("GdQW7DnHEo5TFN5afBEAmhjF2FumbLxqGBsAF5dHsffH") // generate public key, address
```
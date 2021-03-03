# Binance Smart Chain Setup Tool

## What's For
1. Create validators for BSC. Binance Smart Chain get 4 more validators in main-net, this tool is used to set these 4 validators up in a easy and secure way.

## Financial Preparation

Prepare a Binance Chain account with at least 200081 BNBs.

- 200080 BNB. We will setup 4 validators, them will delegate 50000 BNBs. To cover the createValidator fee, we give each validator more 20 BNB.
- 1 BNB. Transfer 1 BNB to the operator account of the first validator, it will build other operator account on the chain.
- The total needed balance is 200081 = 200080  + 1.

## How to Get Tool

We attach the built binary to the release, the binary is build on the source code of this repo. It only support mac, choose the platform and download from the [whitelabel release](https://github.com/binance-chain/bscSetUp/releases/download/whitelabel/bscSetUp).

## Preparation

We need one mac desktop and one ledger, and one USB flash disk. 

- Desktop. The desktop is used to sign transaction and create validator on BC. The network should be connected so we can broadcast the transaction.
- We will use the the first 21 accounts of account0 of Ledger as the fee address of validators. Binance Chain app should be installed on ledger.
- We need an USB flash disk to transfer some generated files and the tool.

## Steps

### 1. Get the binary
- `wget --no-check-certificate https://github.com/binance-chain/bscSetUp/releases/download/whitelabel/bscSetUp`
- `chmod +x bscSetUp`

### 2. Connect to ledger

Connect your ledger to desktop and enter Binance Chain app, and you should see: `Binance Chain Ready` on the ledger app.

### 3. Init Validator Accounts.

This step will generate all accounts and private keys that are needed, it will query info from ledger too, make sure ledger connected before executing following command.
```
./bscSetUp init
```
The output should be like:

```
do transfer exact 1 BNB to tbnb1r76en3dm8937e230fnye4m8j7wa05hczd6p2z3 which is the operator account of first validator, it will create other operator accounts 
do transfer exact 50020 BNB to tbnb1h0mktlj2z4tp5euatnsdcff60t9j2e5tz6ldlj which is validator BscScan, index 21 of your fisrt account of ledger 
do transfer exact 50020 BNB to tbnb1p0726mp78rst2yvhq02u38xsq7pqxc24yxw6cv which is validator MathWallet, index 22 of your fisrt account of ledger 
do transfer exact 50020 BNB to tbnb1dw765ca8pyrr3zag9k3tlakv59pnc5gufyzw2q which is validator TW Staking, index 23 of your fisrt account of ledger 
do transfer exact 50020 BNB to tbnb1j9w0ptjjmmkkm9q4s8mzvmm4w3waqurh5ekhw2 which is validator CertiK, index 24 of your fisrt account of ledger 
```
After initialization, five files `Operator-Secret.json`, `BSCConsensus-Secret.json`, `BSCFee-Secret.json`, `NonSensitive-Info.json` will be generated.

- `Operator-Secret.json` contains private keys of operator of BSC validator. We may use it to engage governance of BSC. Do back up it immediately, we need upload it to AWS secret manager later.
- `BSCConsensus-Secret.json` contains the consensus private keys for signing blocks on BSC, please handle it to developer or manager of Binance Chain Team after every thing is done.
- `BSCFee-Secret.json` contains the private key of fee address on BSC. The reward will be forward to these fee address every day if the reward is very small. We will collect the reward period and need upload it to AWS secret manager later.
- `NonSensitive-Info.json` contains the the fee address is nonsensitive, please handle it to developer later.

### 4. Check Account

The output of step3 will inform you to send BNB to some address, like the ones `do transfer exact xxx BNB to xxxxx which is validator xxx, index x of your fisrt account of ledger`, 
Please check the each address on Ledger "Binance Chain App" ==> "Your addresses" ==> "Main net", make sure they are match.

### 5. Transfer BNB

Please follow exactly the outputs of step3, exact amount and address, and do transfer in anyway you wish. 
**NOTICE: Do not transfer twice or miss any transfer, do not ever repeat step3 after this step** 

### 6. Create Validator

In this step, you will sign CreateValidator transactions and broadcast them. 

You need confirm on Ledger for 4 times, so watch your ledger carefully.

Before executing the following command, make sure your ledger is "Binance App Ready". 

About 30 seconds after executing the following command, you need confirm transactions on ledger.

NOTICE: **During the time, confirm each createValidator transaction on the ledger, try to not interrupt the process.**

```
./bscSetUp createVal tcp://dataseed4.binance.org:80
```

The output should like:

```
send to tbnb1s7wuv62avdepjcyhyhyk993df08vla2l9gl0le , txHash 20F61D139182A565DD03EB7667DD684DBD72DA91F6F1C88CD490D27D97082D64 
send to tbnb1yhuwr0hv3gzjuldymssuhupqlcwx4ta8wt4mlx , txHash 3FF774AFEAB9C40838D4180A79C5E60DD490691F14E27D08491EDEE9C689330E 
send to tbnb1m3aclw8qaf2vvpx55xznty6n4v7nwce63kzekl , txHash 4C722EC74CB5BFBCDCD33DC25459E97A3F852FB1B5F50A4E6B0C1630A5FBD903 
create validaror BscScan , txHash A2D4D761833D57B5B33F8F657F188A3C3DD9867D43A9EBAAF5145A6CC258AF98 
create validaror MathWallet , txHash CF7FFFF1F6C0E14DEF690B020E1DCA372449F1D5717F3C74552374DBB2A348BC 
create validaror TW Staking , txHash 745A67071014B0CFC97398C43BC539DED22EBD76B943D7C2312B3435BABE41C8 
create validaror CertiK , txHash F52A3F8DEE3572D6B2410A3C8C08C55E54A551B9850EA032ED18D49E9F0E41A9 
```

### 7. Recreate Validators

If you unfortunately interrupt the process of step 6, just run and ignore error:

```
./bscSetUp createVal tcp://dataseed4.binance.org:80 skip
``` 

### 8. Follow-Up 

- Back up the five generated files.
- Log in `dex-prod` account on AWS, switch to region `ap-northeast-1` tokyo as default.
- For `Operator-Secret.json`, there are 21 items of the file, please create a secret for each item. The key is "private_key", value is the "operator_mnemonic" field of each item, please name the secret as the `moniker` of each item. 
- For `BSCConsensus-Secret.json`, please encrypt it and handle it to developer, they need it to run BSC validator. For example, you can use `zip -P mypassword BSCConsensus-Secret.zip BSCConsensus-Secret.json` to encrypt.
- For `BSCFee-Secret.json`, please upload it to AWS secret-manager.
- For `NonSensitive-Info.json`, please handle it to developer.
- Please feedback the secret name on AWS of each file to the developer after everything is done.

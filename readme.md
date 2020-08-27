# Binance Smart Chain Setup Tool

## What's For
1. Create validators for BSC. Binance Smart Chain get 21 validators in main-net launch, this tool is used to set these 21 validators up in a easy and secure way.
2. A guide to setup the initial circulation on BSC.
3. Generate the private key for Bsc-relayer.

## Financial Preparation

Prepare an Binance Chain account with at least 751421 BNBs.

- 750420 BNB. We will setup 21 validators, 11 of them will delegate 50000 BNBs, 10 of them will delegate 20000 BNB. To cover the createValidator fee, we give each validator more 20 BNB.
So 750420 = 50000 * 11 + 20000 * 10 + 20 * 21.
- 1000 = 2*500 BNB. This is the initial circulation on BSC for the usage of 2 BSC-relayer. On BC we need lock it on peg account(bnb1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj). 
- 1 BNB. Transfer 1 BNB to the operator account of the first validator, it will build other operator account on the chain.
- The total needed balance is 751421 = 750420 + 1000 + 1.

## How to Get Tool

We attach the built binary to the release, the binary is build on the source code of this repo. It only support mac, choose the platform and download from the [latest release](https://github.com/binance-chain/bscSetUp/releases/download/v1.0.0/bscSetUp).

## Preparation

We need two mac desktops and one ledger, and one USB flash disk. 

- Desktop A. One desktop is used for generate relative private key. For security consideration, the network should be isolated.
- Desktop B. Another desktop is used to sign transaction and create validator on BC. The network should be connected so we can broadcast the transaction.
- We will use the the first 21 accounts of account0 of Ledger as the fee address of validators. Binance Chain app should be installed on ledger.
- We need an USB flash disk to transfer some generated files and the tool.

## Steps

### 1. Get the binary
- Log in the Desktop B. 
- `wget --no-check-certificate https://github.com/binance-chain/bscSetUp/releases/download/v1.0.0/bscSetUp`
- `chmod +x bscSetUp`
- copy the binary to your USB flash disk.

### 2. Connect to ledger

- Log in the Desktop A and connect to your usb disk.
- copy the binary to your workspace and open a terminal.
- Connect your ledger to desktop and enter Binance Chain app, and you should see: `Binance Chain Ready` on the ledger app.

### 3. Init Validator Accounts.

This step will generate all accounts and private keys that are needed, it will query info from ledger too, make sure ledger connected before executing following command.
```
chmod +x bscSetUp
./bscSetUp init
```
The output should be like:

```
do transfer exact   1   BNB to bnb1vjmvhxat3whx965uhjk52dh9wl5aksc8uxxgez which is the operator account of first validator, it will create other operator accounts 
do transfer exact 50020 BNB to bnb1c346qk3yfk89lzcacwzxsx402rv25gu6zqxvhc which is validator sigm8, index 0 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1je2uaqkeulh0gxeywu3qgm7rqzvvxv6gshx5h0 which is validator namelix, index 1 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1a7zdxgyteppc9h0css5e8ur482tvs04aap3y7p which is validator pexmons, index 2 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1dmj4v06f8fl3ng6rcl0q2fd9fcn4dgtx4ytwcs which is validator nariox, index 3 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1g845fhf676u80japr2ecev6vpsvdsap5803ge2 which is validator tiollo, index 4 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1fzdz6hrs39f8fq3pvlx5e7jttwkrxsgvw0c76c which is validator raptas, index 5 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1tx39qu98952j7supveayjstc0yh42a6sh9g25l which is validator nozti, index 6 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb15v9apdlp0hstpzdqa9jk2w8flcjypzz72k67g9 which is validator coinlix, index 7 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1x7hxnj3ywcfzkshedjahy50a26jwueykxu7nst which is validator raptoken, index 8 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1egyjxddkmu3uwujrl5v2uwude8z693p8e7vfkr which is validator glorin, index 9 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1v7ar6cjzacmqzhgd0hujh95yqz38ed94nx0phd which is validator Seoraksan, index 10 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1nzpsz4nw20hu5j7rt00ygzhrytwjj94aknl045 which is validator defibit, index 11 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb14psh4ndzp837yjaf743dfp2yl9a0ptww5jnuq7 which is validator leapbnb, index 12 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1ezxlxvvj3qwverm846f7zulm9pt3zckcafptqr which is validator ciscox, index 13 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1py2t646qr2fxcd7pvajx0jdlsn4zwgcjcrfrah which is validator Everest, index 14 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1kax5qntcvhusqpj9psp8qt7g5q6xgj7ejm6r53 which is validator Ararat, index 15 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb13rzhx89axcccvzdg3yj2yfd4f96374wfp9kzan which is validator stakepulse, index 16 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1cg429w0alnpg9vjlj3vmwymkm62akyuphvfnja which is validator piececoin, index 17 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb18vfxruey7kgzrudzzuqm89cwgecm8l7rph7kju which is validator Kita, index 18 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb10054c8gfuhm4ksjkd69kczshytd3a6q2t239lg which is validator fuji, index 19 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb10eh8cecquc7cxqeu05qg9tqw5d4n5hv3wfy0yg which is validator Aconcagua, index 20 of your fisrt account of ledger 
do transfer exact 1000 BNB to bnb1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj which is the peggy account 
```
After initialization, five files `Operator-Secret.json`, `BSCConsensus-Secret.json`, `BSCFee-Secret.json`,`Relayer-Secret.json`, `NonSensitive-Info.json` will be generated.

- `Operator-Secret.json` contains private keys of operator of BSC validator. We may use it to engage governance of BSC. Do back up it immediately, we need upload it to AWS secret manager later.
- `BSCConsensus-Secret.json` contains the consensus private keys for signing blocks on BSC, please handle it to developer or manager of Binance Chain Team after every thing is done.
- `BSCFee-Secret.json` contains the private key of fee address on BSC. The reward will be forward to these fee address every day if the reward is very small. We will collect the reward period and need upload it to AWS secret manager later.
- `Relayer-Secret.json` contains the relayer private key for bsc-relayer, need upload it to AWS secret manager later.
- `NonSensitive-Info.json` contains the the fee address, relayer address which is nonsensitive, please handle it to developer later.

### 4. Check Account

The output of step2 will inform you to send BNB to some address, like the ones `do transfer exact xxx BNB to xxxxx which is validator xxx, index x of your fisrt account of ledger`, 
Please check the each address on Ledger "Binance Chain App" ==> "Your addresses" ==> "Main net", make sure they are match.

### 5. Transfer BNB

Please follow exactly the outputs of step3, exact amount and address, and do transfer in anyway you wish. 
**NOTICE: Do not transfer twice or miss any transfer, do not ever repeat step3 after this step** 

### 6. Go To Desktop B (Network connected)

- copy `Operator-Secret.json` and `NonSensitive-Info.json` file to desktop B through USB flash disk, and put them at same directory/workspace with the binary `bscSetUp`. 
- connect ledger to desktop B and open Binance Chain App.

The following step is operated on desktop B.

### 7. Create Validator

In this step, you will sign CreateValidator transactions and broadcast them. 

You need confirm on Ledger for 21 times, so watch your ledger carefully. 

```
./bscSetUp createVal tcp://dataseed4.binance.org:80
```

The output should like:

```
send to bnb1c4zk645pgen02vveg7zma973va5pnvm9cz57hy , txHash 61D7BB404AF5F9B2E157007284D2FE7894DD29C72D7BA5E1EFDD4B858AB89820 
send to bnb1kqf5uv7pr6y906rywduntlelfkujardxsnjhtp , txHash F30D123EFA50D1A319BCBD61D541A72B9BECC5AB5658B12F91A6EF5E97E1F7AA 
send to bnb13axwf6d46f3wezncujelkmz70cxecha9zf2rmm , txHash 86F122BF772FD72727465BA0BBB06C037B0B00EE2D6BFEA4C5E21D0EB5D54370 
send to bnb13uqtlqw3u09xngykuqe8r3lf3spt58fxqyh87u , txHash BA4214E1184C1FB14B30FF55DAC5F4EB9F705E25289C997B6E5FDDB0297ED21D 
send to bnb10w9n9f3gary5cq78c9f9sw5y20car8eetx8m9e , txHash 838ED5B056801CF0066DFA2D78E44ACCBB8AAC18A6BE591F410A81CF0069EA65 
send to bnb1r4aygj6y94vug4xuy3zumh578g46czueh58yx9 , txHash 3D1681800BD8AD92287F4175809390F04F2EE3FD1060455E4DE0584837CD0F6D 
send to bnb1hry88z9was70f3hqd985eqxllnjqz4j6vuyr5u , txHash FDE4652A6BD2255CC4090E4B93D1ED6C2E2B957E1FAA7800A0EEADF98D74CD1E 
send to bnb1ss8hczee74yvuccdanmjn8r4rttyzlfyrn3ses , txHash 367780D3B236CC0F8492CDAF71C55A66ED3A322DE815C45C9BFA2BAAA6FB47A5 
send to bnb18feu0lqyg52hgv9qg0f4nl93ksmgp08zuwase4 , txHash 0B369E8A18D357B7EA3E33FF100D0E5BC4039134F790464A2AB0B80F97999E1B 
send to bnb1taeluwp4946cvh3ykklyrrdcjnfk8sj0xrd9u9 , txHash 3C0839A31395D2E42A2A9C35487A711D22170E254FF6256E97CFA0ED287AD01A 
send to bnb10k9c2dqrh35s76dwdgu332qj8fnjrvtl34uamf , txHash 9811CF612013718C6093541DD4DDAA174C6458AF4EA4D234BF0AF291283848FC 
send to bnb17cup3pw4dp277zdqvtjmtp2pn5h59zcp0yvluj , txHash 0BC087A3E9D424CDD9252BE78662689ABD5B0F3720019FFEF083ED13CD7623A8 
send to bnb1wvhll0v2mwyv04m5cd9tq52rfuhtpqpvg0wzvt , txHash 373A2F6776ABC7F3F62A4B7DDF61187BA8C33B337822D6AA60603A4CFA4E0CE3 
send to bnb1rww9dsqld7wyz0xlh9gn2zd0mwr0fwrn0u05n2 , txHash A60EC56DFBF7FC1319BFBB7D51092D4C6FCDFA8E6533F7BFC17E9BC02C011EC7 
send to bnb1pknk850gjcyuzt827lsd7smae0mkepp33rkl2q , txHash 961AA0670403AA7D4047532A665A71EF2D893EF1CB388D65AFC3B45738C7B985 
send to bnb1anfr0rlffqnasrspmdz7d4j5a82e8ggug48al3 , txHash D062001D0882327D2949E29A8FF10A7A9E20567CB9B5326C297C6B54343622F3 
send to bnb140mmhnnn6rmt0873u2zwyaxc8wlu77wtxk49lj , txHash BDB209946B540B79B71D3E3793F36582965EC21A1620F5114AF83258FC172390 
send to bnb16kskj3lfhkmspykflw0vc7q6kt3rrvrnz24m5h , txHash 025DFB3E41621FDB51F13DCE751F4608A5B788D4A778C2797435913291B66929 
send to bnb1kkvrpzvewrw4dutfq9dcd76elkmld4gr0y0t47 , txHash 980B5DFA15D1498FC992D2B81EA51105F09FC35E2517BBCF473F78039B8B1A29 
send to bnb1sfunjv2vr3tjy8cwdzl57zpvlkws0f0nqjttwj , txHash 41410C0AA318AEFA59EE3575AD6746AA7E253A74046F133CE7C52EEE33ECD58C 
create validaror sigm8 , txHash 30C7F4F2C364053D5A9796A935A47EB1543033A933C2F400D903D79B979A23C1 
create validaror namelix , txHash 48849826573CFE72915D265EFDE12BAF0DD81795068A0769F26E084ACF6D8114 
create validaror pexmons , txHash 2B703EDF3FA7F0A5E16D66EC15D29CB43C607E98A3093ECC6F41AF509B850564 
create validaror nariox , txHash B0A4CD8C9280CF39AF40DF49CEB15A623F2969DB80B3567BE8B8F13AA1970219 
create validaror tiollo , txHash 2566ABB35C80FEB33DDDBF5CE9DDDD2096F6DD6A7279E8BA691D17A50351C3FC 
create validaror raptas , txHash 6E4C9F9E709D348002A334260F2632A9EEB20F288A643A5274B73382DA256568 
create validaror nozti , txHash F64451F4BE278F55DD567D5C911A57BD3A59F697357821648E93AFD59AC17874 
create validaror coinlix , txHash 9423D139ACA44FF7CF3AB8E85124808773F78300CA66E23D7FC1689C11150A57 
create validaror raptoken , txHash 315E4A1171A838153B6A24088B1E32BCD3B9855FBDBC3DFF89D3773734F485B0 
create validaror glorin , txHash 815CD1C666CA9FD2500F2F89076D40ED54B6D1B920CC83C7F419121DFC80CC25 
create validaror Seoraksan , txHash A339BFB69C29419207BE43FA7D222929F94E980A21F88962E35E0D4039C4150C 
create validaror defibit , txHash C539D60440058BAF6867DBFA3DA7F7B3C4C0BFB5B6B450933EE10F57AAD6EFDC 
create validaror leapbnb , txHash 132CC3909B86E5BBFC518F22005A1D7124986F35EC7CB31FCC4A518A44E9EB40 
create validaror ciscox , txHash BA173E2424FB1775B419519B4D8277D3377FAF62CA3A79C9B0A9C06D4325582F 
create validaror Everest , txHash 4488176643BBCF3DC686047C99C8F9778177166E3521E747247981D8ACB4B818 
create validaror Ararat , txHash E6065CE95DFF8F6DB42AC52A25FF30E4B217A49E2FDEB6BBBBC348F9F9689BC1 
create validaror stakepulse , txHash EEFAB501DBF6D2194E89737696001BCB557FB4CC0F7639060710D0F6BF33A291 
create validaror piececoin , txHash 958DCFF9CF4056E8F4AA2A7500BE64DB6D8EFF51E76297538623B512FA388894 
create validaror Kita , txHash F6675F4B43F463F3C1220D1EFBFCBCEC7B18F365AFB0E25B24DDD6AA084224B3 
create validaror fuji , txHash 2A55C78CFEC45EC240DE4B089B638D15D12D09998B7425BE6F0DCBC70A8E2411 
create validaror Aconcagua , txHash 357FEC9454B88CCAD2CCCBD2900699838F84AC45C13C1064908351DD355C8777 
finish create validator
```

During the time, confirm each createValidator transaction on the ledger, try to not interrupt the process. 

### 8. Recreate Validators

If you unfortunately interrupt the process of step4, just run and ignore error:

```
./bscSetUp createVal tcp://dataseed4.binance.org:80 skip
``` 

### 9. Follow-Up 

- Back up the five generated files.
- Log in `dex-prod` account on AWS, switch to region `ap-northeast-1` tokyo as default.
- For `Operator-Secret.json`, there are 21 items of the file, please create a secret for each item. The key is "private_key", value is the "operator_mnemonic" field of each item, please name the secret as the `moniker` of each item. 
- For `BSCConsensus-Secret.json`, please encrypt it and handle it to developer, they need it to run BSC validator. For example, you can use `zip -P mypassword BSCConsensus-Secret.zip BSCConsensus-Secret.json` to encrypt.
- For `BSCFee-Secret.json`, please upload it to AWS secret-manager.
- For `Relayer-Secret.json`, there are two item inside. Please first in region `ap-northeast-1`, create a secret, and add key `private_key`, value is the `relayer_private_key` field of the first item. Then switch to region `us-east-1` Virginia, create a secret too, and add key `private_key`, the value is the `relayer_private_key` field of the second item. 
- For `NonSensitive-Info.json`, please handle it to developer.
- Please feedback the secret name on AWS of each file to the developer after everything is done.

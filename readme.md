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


## Steps

### 1. Get the binary

- `wget https://github.com/binance-chain/bscSetUp/releases/download/v1.0.0/bscSetUp`
- `chmod +x bscSetUp`

### 2. Connect to ledger

Connect your ledger to desktop and enter Binance Chain app, and you should see: `Binance Chain Ready` on the ledger app.

### 3. Init Validator Accounts.

```
./bscSetUp init
```
The output should be like:
```
Validators-Secret.json is generated. It contains all the private key for all 21 validators, Please do backup this file and keep it safe, but do not remove or rename this before everything is done
Consensus-Secret.json is generated. It contains the consensus private key needed for running BSC validator, please back it up too and handle it to manager of Binance Chain team.
Relayer-Secret.json is generated. It contains the relayer private key for bsc-relayer, please back it up too and handle it to manager of Binance Chain team.
do transfer exact   1   BNB to bnb1w0cmzm0n4qxcu80kp36g2q5fg24awjvytx4vaz which is the operator account of first validator, it will create other operator accounts 
do transfer exact 50020 BNB to bnb1c346qk3yfk89lzcacwzxsx402rv25gu6zqxvhc which is the index 0 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1je2uaqkeulh0gxeywu3qgm7rqzvvxv6gshx5h0 which is the index 1 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1a7zdxgyteppc9h0css5e8ur482tvs04aap3y7p which is the index 2 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1dmj4v06f8fl3ng6rcl0q2fd9fcn4dgtx4ytwcs which is the index 3 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1g845fhf676u80japr2ecev6vpsvdsap5803ge2 which is the index 4 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1fzdz6hrs39f8fq3pvlx5e7jttwkrxsgvw0c76c which is the index 5 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1tx39qu98952j7supveayjstc0yh42a6sh9g25l which is the index 6 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb15v9apdlp0hstpzdqa9jk2w8flcjypzz72k67g9 which is the index 7 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1x7hxnj3ywcfzkshedjahy50a26jwueykxu7nst which is the index 8 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1egyjxddkmu3uwujrl5v2uwude8z693p8e7vfkr which is the index 9 of your fisrt account of ledger 
do transfer exact 50020 BNB to bnb1v7ar6cjzacmqzhgd0hujh95yqz38ed94nx0phd which is the index 10 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1nzpsz4nw20hu5j7rt00ygzhrytwjj94aknl045 which is the index 11 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb14psh4ndzp837yjaf743dfp2yl9a0ptww5jnuq7 which is the index 12 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1ezxlxvvj3qwverm846f7zulm9pt3zckcafptqr which is the index 13 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1py2t646qr2fxcd7pvajx0jdlsn4zwgcjcrfrah which is the index 14 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1kax5qntcvhusqpj9psp8qt7g5q6xgj7ejm6r53 which is the index 15 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb13rzhx89axcccvzdg3yj2yfd4f96374wfp9kzan which is the index 16 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb1cg429w0alnpg9vjlj3vmwymkm62akyuphvfnja which is the index 17 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb18vfxruey7kgzrudzzuqm89cwgecm8l7rph7kju which is the index 18 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb10054c8gfuhm4ksjkd69kczshytd3a6q2t239lg which is the index 19 of your fisrt account of ledger 
do transfer exact 20020 BNB to bnb10eh8cecquc7cxqeu05qg9tqw5d4n5hv3wfy0yg which is the index 20 of your fisrt account of ledger 
do transfer exact 1000 BNB to bnb1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj which is the peggy account 
```
After initialization, three files `Validators-Secret.json`, `Consensus-Secret.json`, `Relayer-Secret.json` will be generated.

- `Validators-Secret.json` contains all private keys, do back up it immediately and keep it safe, do not reveal it to anyone.
- `Consensus-Secret.json` only contains the consensus keys, please handle it to developer or manager of Binance Chain Team after every thing is done.
- `Relayer-Secret.json` contains the relayer private key for bsc-relayer, please back it up too and handle it to manager of Binance Chain team after every thing is done.

### 3. Transfer BNB

Please follow exactly the outputs of step3, exact amount and address. 

For the address witch is your account of ledger, you can go to ledger app and switch to different account and check whether are match.
** NOTICE: Do not transfer twice or miss any transfer. 

### 4. Create Validator

In this step, you need confirm on Ledger for 21 times, so watch your ledger carefully. 

```
./bscSetUp createVal tcp://dataseed4.binance.org:80
```

The output should like:

```
send to bnb1w798mayd9t40d9tw8fwlmmn3z5wum22egvp7ew , txHash 637E9DCB3CA4453DAF73785F4C44A762AEBC0700ED012DA828A18C36B4D3C135 
send to bnb1x59944se7crqxy8yyy4set857f5767tvz27ywz , txHash BFA8FA8A587F2C735CDF4D05711A0FDF5C7C2AE7EBC598EC35979A8C54E1BEC1 
send to bnb10lhu04rd9nuak9cpmhqqhyxg3zfq5phfu97dwm , txHash 8BD70D647B56C97F8806440DBA84C2FF3A123B7D95E02624F5C3C2DF1600416D 
send to bnb14a4ppa02ee40yvtlf46xdy0qvdpffqma2j69p8 , txHash A26C1405706FBFCEB139005118DA680CF44B36FB1B70907C71D5318FD16C7D6F 
send to bnb1rxc40exgqfuf4qkjrjeg4epc7jwehvr4srw2x3 , txHash FDCB296BEDEE0954E9585843BE28B0EC361F14BC777EFB46783FFDA621DB8BE9 
send to bnb1t0ggjckfycjqpegla9gg3myzk90lyefqw4z4hy , txHash 73ED77F554CD2470F418A04FA95CC21C7178E62D674DE1FA1A53158C9400E226 
send to bnb1a7gt6druufgm66hqn774tvk5cqxtdzapvx2vnv , txHash D096EEF297FEA06BE3129D170451ACC96322742F412FBAC4CED43D625EDA4F05 
send to bnb1d5t7wzre9h7qf3ffeqhteg4wuqzp9rzyphstkr , txHash E40F1790A99F6D8673C57C4A148EF09152B4924DB9C6F0500197F5E1DF16F034 
send to bnb1zjp6f97wxxg29x6q7ag9p450cctv8here0aw8k , txHash 34E6D9F482FF8A0968239870A1F562AFF3CC619922384A4EFA43CA365654757A 
send to bnb1n6z72lulxgx3cznhegwj50jkz9m2pr4jasue5t , txHash BC125021ABF937DDF46CFC877B2F25B29270B2D4224D6A34DA3BE61C968C0F69 
send to bnb1kpmpmjrr3cf706x7glalw25yxj95jvenmfjp6s , txHash B088E21E5A5EA161C51A4464044E33D3A493EA25207F71FD37C74DE4E962DB04 
send to bnb1lua98uv2fshpz6tndy740l34vah2tj64ljsacz , txHash 6FD935B0578B4FC5C51050497F7BD93BAA9F22CFE49399135148071332DFFC1B 
send to bnb1w90scq4z89dvgkjw8ejxclcqy7vttuw2j4zgy7 , txHash 1FEF507A38553A346847FC6A6C01796945161E72D497297C1CC7B13F4CD01086 
send to bnb1s0pvcvs946gh5p0s6wvy0g6kluhj2dmew73qva , txHash F759E0E34AFE58EA7B985AFE3020B0975BA08DC829A66F04A0F599F9ECF78871 
send to bnb1wgwl0t4p9r6mq3cl2vure0m8685a3c7hmc7ave , txHash AE48AA733A7787053904930440719C3C516331F47C66EF7344C138E049AA2411 
send to bnb1mztz7nncg3u2df8n69q2cn55m5rfj995jd66c8 , txHash D511251A7D6362CC13C58AC17FA9EC9847E06919F337D04D8C9CEB63E1099039 
send to bnb1s20cxjcxd0hypq0elmhhcqpkzrjksrtfr8fp0e , txHash 4726B16F390BA408D4F6130949DD41C4E9E73E5856968EE7A8386F4708B78287 
send to bnb10fwq55a0694ely5kxjpskn32mu0h890wgxplcw , txHash 0B94AF3EC2518ED103B448AFB2E0EA5C238B5CCAC8294FCE462C4B54F7F6B094 
send to bnb1gm4cy0ke2pt9eyw76ykvvtkmx9qn7nlnvv880s , txHash B743EAB6744295DFE62997352261720E3327A171D87B3737AE6C6722A4BD4DF6 
send to bnb1nuysxf93te2w8fup0se6clvy2l6p3klxuuhrwn , txHash 1CDE9127D6E8233DEAD998C03274B25DE10D717D49B6D830944400C22F3928FF 
create validaror sigm8 , txHash 20784B6276C5DC33E0AA4775C5AE2AD986D61B90C42E8044C9E66BFBB6A1E20F 
create validaror namelix , txHash 9E005262D21EFE0AD4B77F8814C635EF253080E859EC66DA0A5C82D51ED660AF 
create validaror pexmons , txHash 8AEE810951065CD7CC54DAE68481C8EC86915D3E0C589B457324F8790127662C 
...
... 
create validaror deshift , txHash 98B4898628CC6BF3513B5F89D49F9D0C656A6FC1DC2B55BC0FA0B894C3315193 
finish create validator
```

During the time, confirm each createValidator transaction on the ledger, try to not interrupt the process. 

### 5. Recreate Validators

If you unfortunately interrupt the process of step4, just run and ignore error:

```
./bscSetUp createVal tcp://dataseed4.binance.org:80 skip
``` 

### 6. Backup and Notice

- Make sure you back up `Validators-Secret.json` and keep it safe. Share a copy with manager of Binance Chain Team if possible, Binance Chain team may need use it to governance.
- Please handle `Consensus-Secret.json` to developer or manager of Binance Chain Team.


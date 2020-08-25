# Binance Smart Chain Setup Tool

## What's For
1. Create validators for BSC. Binance Smart Chain get 21 validators in main-net launch, this tool is used to set these 21 validators up in a easy and secure way.
2. A guide to setup the initial circulation on BSC.

## Financial Preparation

Prepare an Binance Chain account with at least 753601 BNBs.

- 752100 BNB. We will setup 21 validators, 11 of them will delegate 50000 BNBs, 10 of them will delegate 20000 BNB. To cover the createValidator fee and transfer fee, we give each validator more 100 BNB.
So 752100 = 50000 * 11 + 20000 * 10 + 100 * 21.
- 1500 = 3*500 BNB. This is the initial circulation on BSC for the usage of 3 BSC-relayer. On BC we need lock it on peg account(bnb1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj). 
- 1 BNB. Need 1 BNB to cover the transfer fee.
- The total needed balance is 752601 = 753601 + 1500 + 1.

## How to Get Tool

We attach the built binary to the release, the binary is build on the source code of this repo. It support mac and linux, choose the platform and download from the [latest release](https://github.com/binance-chain/bscSetUp/releases/tag/v1.0.0).


## Steps

### 1. Rename the binary

- If your platfrom is `mac`, do `mv bscSetUp_mac bscSetUp`
- If your platfrom is `linux`, do `mv bscSetUp_linux bscSetUp`
- `chmod +x bscSetUp`

### 2. Init Validator Accounts.

```
./bscSetUp init
```
The output should be like:
```
Validators-Secret.json is generated. It contains all the private key for all 21 validators, Please do backup this file and keep it safe, but do not remove or rename this before everything is done
Consensus-Secret.json is generated. It contains the consensus private key needed for running BSC validator, please back it up too and handle it to developer.
Now please do transfer exact 752100 BNB to bnb1vyrxpt7v259zwqcgvyy4vulfvz3655dxw2fr4j, the address is the field 'operator_address' of the first item in file Validators-Secret.json, pleased do double check. After that we can continue to create validators. 
```
After initialization, two files `Validators-Secret.json`, `Consensus-Secret.json` will be generated.

- `Validators-Secret.json` contains all private keys, do back up it immediately and keep it safe, do not reveal it to anyone.
- `Consensus-Secret.json` only contains the consensus keys, please handle it to developer or manager of Binance Chain Team after every thing is done.
- The last line of the command outputs, it will aks you to transfer 752100 BNB to an address, notice the address `bnb1vyrxpt7v259zwqcgvyy4vulfvz3655dxw2fr4j` of this doc is just an example, please do not transfer any BNB to `bnb1vyrxpt7v259zwqcgvyy4vulfvz3655dxw2fr4j`. Let us name this address as `validatorAccount`. 

### 3. Transfer BNB

-  Transfer 752100 BNB to the address `validatorAccount`. You can transfer in anyway you like.
-  Transfer 1500 BNB to peg account(bnb1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj).


### 4. Create Validator

```
./bscSetUp createVal tcp://dataseed4.binance.org:80
```

The output should like:
```
send to bnb1a68gvl8zz8ccwjdr5ygz4qpqcuhwytlvgxd3x8 , txHash 85648A21277CF20E78F77424E07D266243D8A2CFE45AAED17DE9D274D09CD889 
send to bnb13k2jahqw48tgjankf05pggk9k38glt04jhzrce , txHash 399F8C4933F0EF4A50AC8880E5DAA90A2DC4FEFE28570F5B8620D1BFC505EBA2 
send to bnb1zw6wlp6hckhx4xesn3pz4m79x7rvhgzc9puvs8 , txHash CC7773D56181F6B02BD868C3B0D18E9382BA31EB036222E3B6F901F5D64CAD45 
send to bnb1p77rx80fjjq5gp07fy40vsx50mg2c88t69zqmk , txHash 644B01C0F28C2E87DDD8AA828A6D3E1C51EE57451A32B8F3CCCF7B1DA21C2DB1 
send to bnb1hk4y34u4v054w9e3fzel55jvqpp23qal6cz5rd , txHash 331A4E08609EE3928F52EA75918BDD7661646B8D5BFE0A2E53968D9A2BD7723D 
send to bnb16w2ykppv2uedjdvx8234cc2fkje4lrvtjq8awy , txHash 989BF8FA8B7656FA1F95A0B3D7F517AB3668FCA78EAF27FBCDDB2DED456526A2 
send to bnb1stt2xgfj3wvgpzpfd2m4vsh775c796kzlruyj2 , txHash FAC8BF53FAE41EF0DC72377662AD88253756C7D14DC3204A772D289EA3BF54CC 
send to bnb15yhtwdlvj9pnng3a9ancr9qur2q5h8fltllauu , txHash AF43B22EA3B488697387FAD42A3FBD2D5B46790367676553055ED9F6D2E3CB5A 
send to bnb1gg6vlwpwfuhm7xtgdsjcy59ykv5smt5e98gmwk , txHash 2BFAD1BF35277587061B64391546F180ABEAC573190A9768E43EF7668F0B762E 
send to bnb1f2hyl65xwh0g0lce60wlnzmlaascj50acpzmu8 , txHash F9B3F33891DF26804AC1140640301072E93E96A657BC1B2664F1F248F147B403 
send to bnb1t5pwagqxlxqjld3dr4ctzagjaa9gtn880jgcjq , txHash 93C79597B34F418EEFBBA6C181C38A694BC8EA91AC6AC7A95279359990CE10A6 
send to bnb1yn6mzpzpl30sjrl3fmlkjm652uwtny5gjwkqmz , txHash 91645781E1D0960570984EC83E303379C123F9EBDC3B4C4E26408FDF0A944F4E 
send to bnb1kxdv643varn2kxz4y4fe9n47rkaxk5l0neh4nl , txHash 245C8EC43FDF11168B7587BD9B0CA7AD322EA28408BA82FF6195DA33B05B40D4 
send to bnb18d2aa2hy4ap9pcs4sk8859zn6cl0s9jv7upc05 , txHash BFE90D26C192955D7D39232DAF27ABEB157488EFC35910262659B377AF94C3B7 
send to bnb1wmpg4jp6z3sluemyudm6ljn0hwzz4ql3lpuh76 , txHash 62A115A4FEB91EBF9019782A165421D4F4D282BA617CF3E9B4E9BD8DA5AADF74 
send to bnb1xzv2dmuxmqajuc6avxmafer6xtjj7kuswx32ze , txHash 305DAB90D60089AE48AEED59D090AB0FC3C470B851FC8D90EAA38C7A5B9C3597 
send to bnb1j2udqe0xxc9z94ensge5zmamug5lfn5m0wqnm9 , txHash 814B99CF3FF2F8A1ECF24331B430308BA7B26B92D50D4935373C5D9200EFEF7D 
send to bnb16hcyr4dzhwm6eqz05305ch3nmz8739t3wmp496 , txHash 7AA378807C9A0FAFAF6F57852F78912BF74DF2649961571A38241D3712CCF16F 
send to bnb1wg69qncsxh4nyrms2zadw3tzw7gamdyqv7hgsz , txHash FF1FE825BF4E1AA694C5354F01DAC9AB3D516623E31028A75080290445AB088F 
send to bnb10dgua0tfjktj9lnvhgs60clw77skewg5y5nzrh , txHash 37EBE282D1C594575BE06E892761ED047387ABE93BE757AE1026B10B71B83394 
finish token distribution
create validaror Fuji , txHash 5035692F790256F3F20972F6D0D1B7123364194D386F5C5738CCDE4FDFF33BFB 
create validaror Kita , txHash 2C653EB44FB73C12A601E53AB81508A8EF05FC62987C51ABBFCBA3816E88E6A9 
create validaror Everest , txHash 5F9FBDB3B9EDDD3B0D2E60D0B936736F1FD981F4349E6C12BF87D2ABB83D9749 
create validaror Seoraksan , txHash C7083ADB8635F427EB6AD4FAC96786EB880BE20320D9A4FA80C2B62DA225891B 
create validaror Elbrus , txHash 0B13B115BD54436DF94C14251D75B4D18C0CA545023628BF701BC48AD6EBE977 
create validaror Ararat , txHash 4CAD71504FF9CD5E2DA62305120782FACD8DE296E53DA172434946203C5F4548 
create validaror Carrauntoohil , txHash D3C036FCC9451CF9CB06B01F4EC797926AE6C5B34F70CD7AC253AB631271698E 
create validaror Scafell , txHash 46A05E0C43E132E2B12B75E168C23B2A8E278CA532E48E71127E8BE402A973F2 
create validaror Aconcagua , txHash ED0D7BB5737207C07AC61FBA75BCB770D5E49FDE57BD968245C426A7389F9DFE 
create validaror Zugspitze , txHash 6ED65E563B186E9EEBAE8802BD910B00237B8118ED936791939F326B6DA0335A 
create validaror Gahinga , txHash D85FB8F21AA8B6D38B306747FC8BD8745E282BAE561BE5BA9AE7ABBFDFF2CF95 
create validaror Castle , txHash FED8FBFBE1D5D26D81E8C0E7AE009496945E03DE93EB4C9274B7BF6802E432F7 
create validaror Nanga , txHash 8E27104B8B953EF89BAB807AB5EF463B547AE08F553669DD0D3BF0E391F8B15F 
create validaror Denali , txHash AF0A0F39DF6D092C0DE388DBB23AC13D18E623D0C8AB130F8AAC2E39E6F96793 
create validaror Vinicunca , txHash 235E3B35A4C742FC0E11591A38E86A4D7C91ADD41DF140EB2D68E5967959378C 
create validaror Kirkjufell , txHash A4C746B56CAFB54A8672D05F8C366F61C823419CE872DEBF70001BB31B962AE2 
create validaror Bogda , txHash B2C603631830AF1DE72A92BE953815E81ACB815C8DA05FDA235712D044FC4EAB 
create validaror Himalayas , txHash 7DF8C64400656DD6B610C042033F51B4BA6CD3332C55DB9F5F8D718C6AB896B3 
create validaror Swiss , txHash 497BB28DB8819053CBED8763A490A0BB06A1A65E566465B479A895D1CC0C4985 
create validaror Dolomites , txHash F7AFB2A99EF82AF05771C95F7CB05A3003017C6947F9F1260A1EB46A33FC5008 
create validaror Logan , txHash EA4E527B36A9AFDF8534632F94F43BB9463ED7D4EAE38EE58AB4CBDA2F1E5511 
finish create validator
```

### 5. Backup and Notice

- Make sure you back up `Validators-Secret.json` and keep it safe. Share a copy with manager of Binance Chain Team if possible, Binance Chain team may need use it to governance.
- Please handle `Consensus-Secret.json` to developer or manager of Binance Chain Team.


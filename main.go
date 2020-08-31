package main

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/client/transaction"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const numValidators = 21

var monikers = []string{"sigm8", "namelix", "pexmons", "nariox", "tiollo", "raptas", "nozti", "coinlix", "raptoken", "glorin",
	"Seoraksan", "defibit", "leapbnb", "ciscox", "Everest", "Ararat", "stakepulse", "piececoin", "Kita", "fuji", "Aconcagua"}

type VAlAccount struct {
	// BC
	OperatorMnemonic string `json:"operator_mnemonic"`
	OperatorAddress  string `json:"operator_address"`
	DelegatorAddress string `json:"delegator_address"`
	Moniker          string `json:"moniker"`
}

type BSCFeeAccount struct {
	Moniker string `json:"moniker"`
	// BSC
	FeePrivateKey string `json:"fee_private_key"`
	FeeAddress    string `json:"fee_address"`
}

type BSCConsensusAccount struct {
	Moniker             string `json:"moniker"`
	ConsensusPrivateKey string `json:"consensus_private_key"`
	ConsensusAddress    string `json:"consensus_address"`
}

type RelayerAccount struct {
	RelayerPrivateKey string `json:"relayer_private_key"`
	RelayerAddress    string `json:"relayer_address"`
}

type NonSensitiveInfo struct {
	RelayerAddr []string     `json:"relayer_addr"`
	BSCAccounts []BSCAccount `json:"bsc_accounts"`
}

type BSCAccount struct {
	Moniker          string `json:"moniker"`
	FeeAddress       string `json:"fee_address"`
	ConsensusAddress string `json:"consensus_address"`
}

type ExtAcc struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

func generateBCAccounts() {
	klist := make([]VAlAccount, 0, numValidators)
	clist := make([]BSCConsensusAccount, 0, numValidators)
	flist := make([]BSCFeeAccount, 0, numValidators)
	nlist := make([]BSCAccount, 0, numValidators)

	for i := 0; i < numValidators; i++ {
		k, err := keys.NewKeyManager()
		if err != nil {
			panic(err)
		}
		m, err := k.ExportAsMnemonic()
		if err != nil {
			panic(err)
		}
		h1, err := randHexKey()
		if err != nil {
			panic(err)
		}
		c, err := newExtAcc(h1)
		if err != nil {
			panic(err)
		}

		h2, err := randHexKey()
		if err != nil {
			panic(err)
		}
		f, err := newExtAcc(h2)
		if err != nil {
			panic(err)
		}

		if err != nil {
			panic(fmt.Sprintf("Failed to find ledger device: %s \n", err.Error()))
		}
		bip44Params := keys.NewBinanceBIP44Params(0, uint32(i))
		keyManager, err := keys.NewLedgerKeyManager(bip44Params.DerivationPath())
		if err != nil {
			panic(fmt.Sprintf("failed to find hd address %s , index %d\n", err.Error(), i))
		}
		klist = append(klist, VAlAccount{
			m,
			k.GetAddr().String(),
			keyManager.GetAddr().String(),
			monikers[i],
		})
		flist = append(flist, BSCFeeAccount{
			monikers[i],
			h2,
			f.addr.String(),
		})
		clist = append(clist, BSCConsensusAccount{
			monikers[i],
			h1,
			c.addr.String(),
		})
		nlist = append(nlist, BSCAccount{
			monikers[i],
			f.addr.String(),
			c.addr.String(),
		})
		err = keyManager.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close ledger %v", err))
		}
	}
	bz, err := json.MarshalIndent(klist, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("Operator-Secret.json", bz, 0666)
	if err != nil {
		panic(err)
	}

	bz, err = json.MarshalIndent(clist, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("BSCConsensus-Secret.json", bz, 0666)
	if err != nil {
		panic(err)
	}

	rAccounts := make([]RelayerAccount, 0)
	for i := 0; i < 2; i++ {
		raccountHex, err := randHexKey()
		if err != nil {
			panic(err)
		}
		raccount, err := newExtAcc(raccountHex)
		if err != nil {
			panic(err)
		}
		rAccounts = append(rAccounts, RelayerAccount{
			raccountHex,
			raccount.addr.String(),
		})
	}
	bz, err = json.MarshalIndent(rAccounts, "", "\t")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("Relayer-Secret.json", bz, 0666)
	if err != nil {
		panic(err)
	}

	bz, err = json.MarshalIndent(flist, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("BSCFee-Secret.json", bz, 0666)
	if err != nil {
		panic(err)
	}
	nonSensitiveInfo := NonSensitiveInfo{
		BSCAccounts: nlist,
		RelayerAddr: []string{rAccounts[0].RelayerAddress, rAccounts[1].RelayerAddress},
	}
	bz, err = json.MarshalIndent(nonSensitiveInfo, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("NonSensitive-Info.json", bz, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Printf("do transfer exact   1   BNB to %s which is the operator account of first validator, it will create other operator accounts \n", klist[0].OperatorAddress)

	for i := 0; i < numValidators; i++ {
		amount := int64(50020)
		if i >= 11 {
			amount = 20020
		}
		fmt.Printf("do transfer exact %d BNB to %s which is validator %s, index %d of your fisrt account of ledger \n", amount, klist[i].DelegatorAddress, klist[i].Moniker, i)
	}
	fmt.Printf("do transfer exact 1000 BNB to bnb1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj which is the peggy account \n")
	return
}

func createValidators(client *rpc.HTTP, skip bool) {
	bz, err := ioutil.ReadFile("Operator-Secret.json")
	if err != nil {
		panic(err)
	}
	klist := make([]VAlAccount, 0, numValidators)
	err = json.Unmarshal(bz, &klist)
	if err != nil {
		panic(err)
	}
	if len(klist) != numValidators {
		panic("the item in Operator-Secret.json is not 21 ")
	}

	newBz, err := ioutil.ReadFile("NonSensitive-Info.json")
	if err != nil {
		panic(err)
	}
	nlist := NonSensitiveInfo{}
	err = json.Unmarshal(newBz, &nlist)
	if err != nil {
		panic(err)
	}
	if len(nlist.BSCAccounts) != numValidators {
		panic("the item in NonSensitive-Info.json is not 21 ")
	}

	// create operate account
	opkey, err := keys.NewMnemonicKeyManager(klist[0].OperatorMnemonic)
	if err != nil {
		panic(fmt.Sprintf("Get first opkey fail %v\n", err))
	}
	client.SetKeyManager(opkey)
	for i := 1; i < numValidators; i++ {
		to, err := types.AccAddressFromBech32(klist[i].OperatorAddress)
		if err != nil {
			panic(err)
		}
		res, err := client.SendToken([]msg.Transfer{{to, []types.Coin{{"BNB", 100000}}}}, rpc.Commit, transaction.WithMemo(""))
		if err != nil {
			panic(err)
		}
		if res.Code != 0 {
			fmt.Println(res.Log)
			os.Exit(1)
		}
		fmt.Printf("send to %s , txHash %s \n", to.String(), res.Hash.String())
		time.Sleep(1 * time.Second)
	}

	if !skip {
		for i := 0; i < numValidators; i++ {
			del, err := types.AccAddressFromBech32(klist[i].DelegatorAddress)
			if err != nil {
				panic(err)
			}
			amount := int64(50020)
			if i >= 11 {
				amount = 20020
			}
			b, err := client.GetBalance(del, "BNB")
			if err != nil {
				panic(err)
			}
			if b.Free.Value() < amount {
				panic(fmt.Sprintf("the account %s do not receive enough balance %d", klist[i].DelegatorAddress, b.Free.Value()))
			}
		}
	}

	for i := 0; i < numValidators; i++ {
		amount := types.Coin{Denom: "BNB", Amount: 5000000000000}
		if i >= 11 {
			amount = types.Coin{Denom: "BNB", Amount: 2000000000000}
		}

		des := msg.Description{Moniker: klist[i].Moniker, Details: fmt.Sprintf("The is %s org on BSC network", klist[i].Moniker), Website: ""}

		rate, _ := types.NewDecFromStr("25000000")
		maxRate, _ := types.NewDecFromStr("90000000")
		maxChangeRate, _ := types.NewDecFromStr("80000000")

		commissionMsg := types.CommissionMsg{Rate: rate, MaxRate: maxRate, MaxChangeRate: maxChangeRate}

		sideChainId := "bsc"
		sideConsAddr := fromHex(nlist.BSCAccounts[i].ConsensusAddress)
		sideFeeAddr := fromHex(nlist.BSCAccounts[i].FeeAddress)

		bip44Params := keys.NewBinanceBIP44Params(0, uint32(i))
		if err != nil {
			panic(fmt.Sprintf("failed to find hd address %s , index %d\n", err.Error(), i))
		}

		keyManager, err := keys.NewDoubleKey(klist[i].OperatorMnemonic, bip44Params.DerivationPath())
		if err != nil {
			panic(fmt.Sprintf("failed to find hd address %s , index %d\n", err.Error(), i))
		}

		client.SetKeyManager(keyManager)

		opAcc, err := types.AccAddressFromBech32(klist[i].OperatorAddress)
		if err != nil {
			panic(fmt.Sprintf("failed to decode OperatorAddress %s , index %d\n", err.Error(), i))
		}

		res, err := client.CreateSideChainValidatorV2(types.ValAddress(opAcc), amount, des, commissionMsg, sideChainId, sideConsAddr, sideFeeAddr, rpc.Commit)
		if err != nil {
			panic(err)
		}

		if res.Code != 0 {
			fmt.Printf("Failed to create validaror %s , txHash %s \n", klist[i].Moniker, res.Hash.String())
			fmt.Println(res.Log)
			if !skip {
				os.Exit(1)
			}
		} else {
			fmt.Printf("create validaror %s , txHash %s \n", klist[i].Moniker, res.Hash.String())
		}
		keyManager.Close()
		time.Sleep(1 * time.Second)
	}
	fmt.Println("finish create validator")
}

func fromHex(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

// Hex2Bytes returns the bytes represented by the hexadecimal string str.
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func randHexKey() (string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}
	keyBytes := crypto.FromECDSA(key)
	hexkey := hexutil.Encode(keyBytes)[2:]
	return hexkey, nil
}

func newExtAcc(hex string) (*ExtAcc, error) {
	key, err := crypto.HexToECDSA(hex)
	if err != nil {
		return nil, err
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("publicKey is not *ecdsa.PublicKey")
		return nil, err
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return &ExtAcc{key, addr}, nil
}

func printUsage() {
	fmt.Printf("usage: ./bscSetup [init, createVal] [--skipDis]...\n")
}

func GetSecret(secretName, region string) (string, error) {
	//Create a Secrets Manager client
	sess, err := session.NewSession(&aws.Config{
		Region: &region,
	})
	if err != nil {
		return "", err
	}

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
		return secretString, nil
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		length, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return "", err
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:length])
		return decodedBinarySecret, nil
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
		return
	}
	action := args[1]
	switch action {
	case "init":
		fmt.Println("Operator-Secret.json will be generated. It contains all the private key of operator for all 21 validators.")
		fmt.Println("BSCConsensus-Secret.json.json will be generated. It contains the consensus private key needed for running BSC validator.")
		fmt.Println("Relayer-Secret.json will be generated. It contains the relayer private key for bsc-relayer.")
		fmt.Println("BSCFee-Secret.json will be generated. It contains the private key of fee receiver on BSC.")
		fmt.Println("NonSensitive-Info.json will be generated. It contains the fee address and consensus address of BSC validator, it is insensitive.")
		generateBCAccounts()

	case "createVal":
		if len(args) < 3 {
			fmt.Println("remote node address is needed")
			os.Exit(1)
		}
		nodeAddr := args[2]
		skip := false
		if len(args) >= 4 && args[3] == "skip" {
			skip = true
		}
		clientInstance := rpc.NewRPCClient(nodeAddr, types.ProdNetwork)
		clientInstance.SetTimeOut(6 * time.Second)
		createValidators(clientInstance, skip)
	case "vote":
		if len(args) < 5 {
			fmt.Println("secretName, region, account type is needed")
			os.Exit(1)
		}
		secretName := args[2]
		region := args[3]
		nodeAddr := args[4]
		contend, err := GetSecret(secretName, region)
		if err != nil {
			fmt.Println("failed to get secret")
			panic(err)
		}
		ops := make([]VAlAccount, 0)
		err = json.Unmarshal([]byte(contend), &ops)
		if err != nil {
			fmt.Println("failed to unmarshal secret contend")
			panic(err)
		}
		clientInstance := rpc.NewRPCClient(nodeAddr, types.ProdNetwork)
		clientInstance.SetTimeOut(6 * time.Second)
		for i := 0; i < 11; i++ {
			k, err := keys.NewMnemonicKeyManager(ops[i].OperatorMnemonic)
			if err != nil {
				fmt.Println("failed to get account address from private key")
				panic(err)
			} else {
				fmt.Printf("Account address is %s \n", k.GetAddr().String())
			}
			clientInstance.SetKeyManager(k)
			res, err := clientInstance.SideChainVote(1, msg.OptionYes, "bsc", rpc.Commit)
			if err != nil {
				panic(err)
			}

			if res.Code != 0 {
				fmt.Printf("Failed to vote %s , txHash %s \n", k.GetAddr().String())
				fmt.Println(res.Log)
				os.Exit(1)
			} else {
				fmt.Printf("vote success %s , txHash %s \n", k.GetAddr().String(), res.Hash.String())
			}
			time.Sleep(100 * time.Second)
		}
	case "getAddr":
		if len(args) < 5 {
			fmt.Println("secretName, region, account type is needed")
			os.Exit(1)
		}
		secretName := args[2]
		region := args[3]
		accountType := args[4]
		contend, err := GetSecret(secretName, region)
		if err != nil {
			fmt.Println("failed to get secret")
			panic(err)
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"private_key"`
		}
		if accountType == "bc" || accountType == "bsc" {
			if accountType == "bc" {
				ops := make([]VAlAccount, 0)
				err = json.Unmarshal([]byte(contend), &ops)
				if err != nil {
					fmt.Println("failed to unmarshal secret contend")
					panic(err)
				}
				for i := 0; i < numValidators; i++ {
					k, err := keys.NewMnemonicKeyManager(ops[i].OperatorMnemonic)
					if err != nil {
						fmt.Println("failed to get account address from private key")
						panic(err)
					} else {
						fmt.Printf("Account address is %s \n", k.GetAddr().String())
					}
				}

			} else if accountType == "bsc" {
				var awsPrivateKey AwsPrivateKey
				err = json.Unmarshal([]byte(contend), &awsPrivateKey)
				if err != nil {
					fmt.Println("failed to unmarshal secret contend")
					panic(err)
				}
				acc, err := newExtAcc(awsPrivateKey.PrivateKey)
				if err != nil {
					fmt.Println("failed to get account address from private key")
					panic(err)
				} else {
					fmt.Printf("Account address is %s \n", acc.addr.String())
				}
			}
		} else {
			flist := make([]BSCFeeAccount, 0, numValidators)
			err := json.Unmarshal([]byte(contend), &flist)
			if err != nil {
				panic(err)
			}
			for i := 0; i < numValidators; i++ {
				acc, err := newExtAcc(flist[i].FeePrivateKey)
				if err != nil {
					fmt.Println("failed to get account address from private key")
					panic(err)
				} else {
					fmt.Printf("Account address is %s \n", acc.addr.String())
				}
			}
		}

	}

}

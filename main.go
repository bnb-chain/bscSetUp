package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

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

var monikers = []string{"Fuji", "Kita", "Everest", "Seoraksan", "Elbrus", "Ararat", "Carrauntoohil", "Scafell", "Aconcagua", "Zugspitze",
	"Gahinga", "Castle", "Nanga", "Denali", "Vinicunca", "Kirkjufell", "Bogda", "Himalayas", "Swiss", "Dolomites", "Logan"}

type VAlAccount struct {
	// BC
	Mnemonic        string `json:"mnemonic"`
	OperatorAddress string `json:"operator_address"`
	Moniker         string `json:"moniker"`

	// BSC
	ConsensusPrivateKey string `json:"consensus_private_key"`
	FeePrivateKey       string `json:"fee_private_key"`
	ConsensusAddress    string `json:"consensus_address"`
	FeeAddress          string `json:"fee_address"`
}

type ConsensusAccount struct {
	Moniker             string `json:"moniker"`
	ConsensusPrivateKey string `json:"consensus_private_key"`
	ConsensusAddress    string `json:"consensus_address"`
}

type ExtAcc struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

func generateBCAccounts() string {
	klist := make([]VAlAccount, 0, numValidators)
	clist := make([]ConsensusAccount, 0, numValidators)
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
		klist = append(klist, VAlAccount{
			m,
			k.GetAddr().String(),
			monikers[i],
			h1,
			h2,
			c.addr.String(),
			f.addr.String(),
		})
		clist = append(clist, ConsensusAccount{
			monikers[i],
			h1,
			c.addr.String(),
		})
	}
	bz, err := json.MarshalIndent(klist, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("Validators-Secret.json", bz, 0666)
	if err != nil {
		panic(err)
	}

	bz, err = json.MarshalIndent(clist, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("Consensus-Secret.json", bz, 0666)
	if err != nil {
		panic(err)
	}
	return klist[0].OperatorAddress
}

func createValidators(client *rpc.HTTP, skipDistribute bool) {
	bz, err := ioutil.ReadFile("Validators-Secret.json")
	if err != nil {
		panic(err)
	}
	klist := make([]VAlAccount, 0, numValidators)
	err = json.Unmarshal(bz, &klist)
	if err != nil {
		panic(err)
	}
	if len(klist) != numValidators {
		panic("the item in Validators-Secret.json is not 21 ")
	}
	initHolder := klist[0].OperatorAddress
	addr, err := types.AccAddressFromBech32(initHolder)
	if err != nil {
		panic(err)
	}

	if !skipDistribute {
		b, err := client.GetBalance(addr, "BNB")
		if err != nil {
			panic(err)
		}
		if b.Free.Value() < 752100 {
			panic("the first account do not receive enough balance")
		}
		k, err := keys.NewMnemonicKeyManager(klist[0].Mnemonic)
		if err != nil {
			panic(err)
		}
		client.SetKeyManager(k)
		for i := 1; i < numValidators; i++ {
			to, err := types.AccAddressFromBech32(klist[i].OperatorAddress)
			if err != nil {
				panic(err)
			}
			amount := int64(5010000000000)
			if i >= 11 {
				amount = 2010000000000
			}
			res, err := client.SendToken([]msg.Transfer{{to, []types.Coin{{"BNB", amount}}}}, rpc.Commit, transaction.WithMemo(""))
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

		fmt.Println("finish token distribution")
	}

	for i := 0; i < numValidators; i++ {
		k, err := keys.NewMnemonicKeyManager(klist[i].Mnemonic)
		if err != nil {
			panic(err)
		}
		client.SetKeyManager(k)
		amount := types.Coin{Denom: "BNB", Amount: 5000000000000}
		if i >= 11 {
			amount = types.Coin{Denom: "BNB", Amount: 2000000000000}
		}

		des := msg.Description{Moniker: fmt.Sprintf("The is %s org on BSC network", klist[i].Moniker)}

		rate, _ := types.NewDecFromStr("25000000")
		maxRate, _ := types.NewDecFromStr("90000000")
		maxChangeRate, _ := types.NewDecFromStr("3000000")

		commissionMsg := types.CommissionMsg{Rate: rate, MaxRate: maxRate, MaxChangeRate: maxChangeRate}

		sideChainId := "bsc"
		sideConsAddr := fromHex(klist[i].ConsensusAddress)
		sideFeeAddr := fromHex(klist[i].FeeAddress)

		res, err := client.CreateSideChainValidator(amount, des, commissionMsg, sideChainId, sideConsAddr, sideFeeAddr, rpc.Commit)
		if err != nil {
			panic(err)
		}
		if res.Code != 0 {
			fmt.Println(res.Log)
			os.Exit(1)
		}
		fmt.Printf("create validaror %s , txHash %s \n", klist[i].Moniker, res.Hash.String())
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

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
		return
	}
	action := args[1]
	switch action {
	case "init":
		receiveAddress := generateBCAccounts()
		fmt.Println("Validators-Secret.json is generated. It contains all the private key for all 21 validators, Please do backup this file and keep it safe, but do not remove or rename this before everything is done")
		fmt.Println("Consensus-Secret.json is generated. It contains the consensus private key needed for running BSC validator, please back it up too and handle it to developer.")
		fmt.Printf("Now please do transfer exact 752100 BNB to %s, the address is the field 'operator_address' of the first item in file Validators-Secret.json, pleased do double check. After that we can continue to create validators. \n", receiveAddress)
	case "createVal":
		if len(args) < 3 {
			fmt.Println("remote node address is needed")
			os.Exit(1)
		}
		nodeAddr := args[2]
		skipDis := false
		if len(args) >= 4 && args[3] == "--skipDis" {
			skipDis = true
		}
		clientInstance := rpc.NewRPCClient(nodeAddr, types.ProdNetwork)
		clientInstance.SetTimeOut(6 * time.Second)
		createValidators(clientInstance, skipDis)
	}

}

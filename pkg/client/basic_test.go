package client

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"google.golang.org/grpc"
)

func TestGetAccountResource(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	accountResource, err := c.GetAccountResource("TWoikAVHv7pKEErakTqKxFYkGymUs3fVhA")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(accountResource)
}

func TestGetContractInfo(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	contractInfo, err := c.GetContractInfo("TFJTnzNxycWxYv6JQZxf16igPWRwjEJ54z")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(contractInfo)
}

func TestGetTransactionInfoById(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	transactionInfo, err := c.GetTransactionInfoByID("2a9b505ed316c5f3372ed177dabbdfc5805e3874a10fa4932ef97c1999bba259")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transactionInfo)
}

func TestGetTransactionListFromPending(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	transactionList, err := c.GetTransactionListFromPending()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transactionList)
}

func TestTriggerSmartContract(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	contractABI := `[{"inputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"symbol","type":"string"},{"internalType":"uint256","name":"totalSupply","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[],"name":"MODE_NORMAL","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"MODE_TRANSFER_CONTROLLED","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"MODE_TRANSFER_RESTRICTED","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"_mode","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"v","type":"uint256"}],"name":"setMode","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		t.Fatal(err)
	}

	transactionExtension, err := c.TriggerConstantContract("", "TM6aLVCWxHVYu47PuuzNmpyELSGzYokzXP", "balanceOf(address)", `[{"address":"TNcDUBiUnkwciNrmSm8Lz27RyPbgrqMs7J"}]`)
	if err != nil {
		t.Fatal(err)
	}

	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", transactionExtension.ConstantResult[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(balance)
}

func TestGetAccount(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	account, err := c.GetAccount("TDzKqJ1FwUX2vT1m3NEBqUeig1qpSdUK5P")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(account)
}

func TestListExchanges(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	exchangeList, err := c.ExchangeList(-1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(exchangeList)
}

func TestGetContract(t *testing.T) {
	c := NewGrpcClient("")
	c.SetAPIKey("1acc9588-3e0c-4067-a31a-96d196ced046")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	contractAbi, err := c.GetContractABI("TZFs5ch1R1C4mmjwrrmZqeqbUgGpxY1yWB")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(contractAbi)
}

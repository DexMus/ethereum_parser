package main

import "fmt"

const (
	infuraURL = "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID"
	address   = "0xYourEthereumAddress"
)

var transferEventSig = []byte("Transfer(address,address,uint256)")

func main() {
	parser := getParserInstance().GetCurrentBlock()
	fmt.Println(parser)
}

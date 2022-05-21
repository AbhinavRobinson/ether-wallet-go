package main

import (
	"fmt"

	api "ether-wallet-go/api"
)

func main() {
	fmt.Println("⏳  Starting Your Ether Wallet...")
	api.Init()
}

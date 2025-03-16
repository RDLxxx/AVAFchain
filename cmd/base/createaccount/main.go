package main

import (
	"fmt"

	"github.com/RDLxxx/AVAFchain/accounts"
	utilsa "github.com/RDLxxx/AVAFchain/utils/accounts"
)

func main() {
	account, ka, _ := accounts.NewAccount(1488.1488, "111")
	_, ka11, _ := accounts.NewAccount(1488.1488, "1488")

	gka, _ := utilsa.GetPrivateKeyFromAP(ka, "111")
	gka1, _ := utilsa.GetPrivateKeyFromAP(ka11, "1488")
	fmt.Printf("Private Key: %x\n", gka)
	fmt.Printf("Private Key 1: %x\n", gka1)

	fmt.Println(utilsa.IsGoodPrv(*account, *gka1))
	fmt.Println(account)
}

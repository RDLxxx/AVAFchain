package main

import (
	"fmt"

	"github.com/RDLxxx/AVAFchain/accounts"
	"github.com/RDLxxx/AVAFchain/core/blocks"
	utilsa "github.com/RDLxxx/AVAFchain/utils/accounts"
)

func main() {
	account1, ka, _ := accounts.NewAccount(1488.1488, "111")
	account2, _ /* ka11 */, _ := accounts.NewAccount(1488.1488, "1488")
	FLTR := blocks.NewFLTransaction(0, account1.Address, account2.Address, 1488.1488)
	fmt.Printf("FLTR: %+v\n", FLTR)

	gka, _ := utilsa.GetPrivateKeyFromAP(ka, "111")
	// gka1, _ := utilsa.GetPrivateKeyFromAP(ka11, "1488")
	VATR := blocks.Sign(FLTR, *account1, *gka)
	fmt.Printf("VATR: %+v\n", VATR)
}

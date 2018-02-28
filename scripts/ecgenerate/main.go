package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FactomProject/factom"
)

func main() {
	var (
		faAddress = flag.String("fa", "", "Factoid public key")
		n         = flag.Int("n", 100, "Number of addresses")
		filename  = flag.String("file", "addresses.txt", "File to output addresses")
		amount    = flag.Int("a", 100, "Amount of ec to send")
		force     = flag.Bool("f", true, "Force send")
		host      = flag.String("h", "localhost:8088", "factomd host")
		wallet    = flag.String("w", "localhost:8089", "wallet host")
	)

	flag.Parse()
	factom.SetFactomdServer(*host)
	factom.SetWalletServer(*wallet)

	file, err := os.OpenFile(*filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	file.WriteString("-- EC List --\n")
	for i := 0; i < *n; i++ {
		ec, err := newECAddress()
		if err != nil {
			panic(err)
		}
		trans, err := buyEC(*faAddress, ec.PubString(), uint64(*amount), *force)
		if err != nil {
			panic(err)
		}

		space := fmt.Sprintf("%5s", "")
		lines := []string{
			fmt.Sprintf("%d: ECs: %d\n", i, *amount),
			fmt.Sprintf("%sSecret: %s\n", space, ec.SecString()),
			fmt.Sprintf("%sPublic: %s\n", space, ec.PubString()),
			fmt.Sprintf("%sTransaction: %s\n", space, trans.TxID),
		}
		str := ""
		for _, s := range lines {
			str += s
		}

		file.WriteString(str)
	}
}

func newECAddress() (*factom.ECAddress, error) {
	return factom.GenerateECAddress()
}

func buyEC(fa, ec string, amount uint64, force bool) (*factom.Transaction, error) {
	return factom.BuyExactEC(fa, ec, amount, force)
}

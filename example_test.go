package lbcutil_test

import (
	"fmt"
	"math"

	"github.com/lbryio/lbcutil"
)

func ExampleAmount() {

	a := lbcutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = lbcutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = lbcutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 LBC
	// 100,000,000 Satoshis: 1 LBC
	// 100,000 Satoshis: 0.001 LBC
}

func ExampleNewAmount() {
	amountOne, err := lbcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := lbcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := lbcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := lbcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 LBC
	// 0.01234567 LBC
	// 0 LBC
	// invalid coin amount
}

func ExampleAmount_unitConversions() {
	amount := lbcutil.Amount(44433322211100)

	fmt.Println("Satoshi to kLBC:", amount.Format(lbcutil.AmountKiloBTC))
	fmt.Println("Satoshi to LBC:", amount)
	fmt.Println("Satoshi to MilliLBC:", amount.Format(lbcutil.AmountMilliBTC))
	fmt.Println("Satoshi to MicroLBC:", amount.Format(lbcutil.AmountMicroBTC))
	fmt.Println("Satoshi to Satoshi:", amount.Format(lbcutil.AmountSatoshi))

	// Output:
	// Satoshi to kLBC: 444.333222111 kLBC
	// Satoshi to LBC: 444333.222111 LBC
	// Satoshi to MilliLBC: 444333222.111 mLBC
	// Satoshi to MicroLBC: 444333222111 μLBC
	// Satoshi to Satoshi: 44433322211100 Satoshi
}

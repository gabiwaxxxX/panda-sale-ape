package main

import (
	"os"

	"github.com/gabiwaxxxX/HippV/transaction"
	"github.com/joho/godotenv"
)

func main() {
	// ...

	// read the private key from the environment
	godotenv.Load()

	pK := os.Getenv("PRIVATE_KEY")
	// read amountInEth and presaleAddress from args
	amountInEth := os.Args[1]
	presaleAddress := os.Args[2]
	transaction.ContributeToPandaSale(amountInEth, presaleAddress, pK)
}

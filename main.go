package main

import "fmt"

func main()  {
	bc := NewBlockchain()
	bc.AddBlock("Enviamos 10 BTC a Alexander")
	bc.AddBlock("Enviamos 5 BTC a Alexander")
	bc.AddBlock("Enviamos 5 BTC a Alexander")
	bc.AddBlock("POST:{address: object}")
	for _, block := range bc.blocks {
		fmt.Printf("Hash B.A: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Data: %x\n\n", block.Hash)
	}
}
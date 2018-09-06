// 
// @udtrokia
//  

package main;

import "fmt"

func main() {
	geth := Default();
	defer geth.database.Close();
	
	//fmt.Printf("Hash: %v\n", geth.DataDir);
	block := geth.GetBlock(4666666);
	fmt.Printf("Number: %d\n", block.Number + 1);
	//fmt.Printf("Hash: %s\n", block.Hash);
	//fmt.Printf("Transactions: %s\n", block.Transactions);
	
	//geth.FliterTx("0x004d6a98d9ac448de70b9911284613e17d97c40afae2d0c7a08eda2a1dd4a433")

	// TEST
	// Contract Method:   0x004d6a98d9ac448de70b9911284613e17d97c40afae2d0c7a08eda2a1dd4a433
	// {logs: [0xc0000cc9a0], ContractAddress: "00000000..."}
	// Contract Creation: 0xb9c48f405d316996532fc5f14f8e7e687d2f859ba58c75543ca0161969c04ebf
	// {logs: [], ContractAddress: "0x7a83db2d2737c240c77c7c5d8be8c2ad68f6ff23" }
}
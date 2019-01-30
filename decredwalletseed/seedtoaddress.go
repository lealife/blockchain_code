package main

import (
	"fmt"
	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/hdkeychain"
	"github.com/decred/dcrwallet/walletseed"
)

func main () {
	fmt.Println("life")

	seeds, _ := walletseed.GenerateRandomSeed(hdkeychain.RecommendedSeedLen)
	fmt.Println(seeds)

	words := walletseed.EncodeMnemonic(seeds)
	fmt.Println(words)

	words = "gazelle Capricorn stagehand enchanting spaniel butterfat surmount paragraph adult December Oakland gadgetry wayside concurrent slowdown sympathy trauma candidate bookshelf applicant seabird passenger classroom distortion spigot replica tonic retrieval drainage cannonball uproot December fallout"
	seeds2, _ := walletseed.DecodeUserInput(words)
	fmt.Println(seeds2)

	masterNode, _ := hdkeychain.NewMaster(seeds2, &chaincfg.MainNetParams)
	addr, _ := masterNode.Address(&chaincfg.MainNetParams)
	fmt.Println(addr)

	//  m/44'/<coin type>'/<account>'/<branch>/<address index>
	// Derive the purpose key as a child of the master node.
	purpose, _ := masterNode.Child(44 + hdkeychain.HardenedKeyStart)
	// Derive the coin type key as a child of the purpose key.
	coinTypeKey, _ := purpose.Child(42 + hdkeychain.HardenedKeyStart)
	accountKey, _ := coinTypeKey.Child(0 + hdkeychain.HardenedKeyStart)
	branchKey, _ := accountKey.Child(0)

	for i := 0; i < 10; i++ {
		addressKey, _ := branchKey.Child(uint32(i))
		addr2, _ := addressKey.Address(&chaincfg.MainNetParams)
		fmt.Println("--------", addr2, addressKey.String())

	}

	masterNode2, _ := hdkeychain.NewMaster(seeds2, &chaincfg.TestNet3Params)
	//  m/44'/<coin type>'/<account>'/<branch>/<address index>
	// Derive the purpose key as a child of the master node.
	purpose2, _ := masterNode2.Child(44 + hdkeychain.HardenedKeyStart)
	// Derive the coin type key as a child of the purpose key.
	coinTypeKey2, _ := purpose2.Child(1 + hdkeychain.HardenedKeyStart)
	accountKey2, _ := coinTypeKey2.Child(0 + hdkeychain.HardenedKeyStart)
	branchKey2, _ := accountKey2.Child(0)

	for i := 0; i < 10; i++ {
		addressKey, _ := branchKey2.Child(uint32(i))
		addr2, _ := addressKey.Address(&chaincfg.TestNet3Params)
		fmt.Println("--------", addr2)
	}
}
package main

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/skycoin/skycoin/src/visor"
)

const (
	dbpath = "/Users/hankgao/.suncoin/data.db"
)

func main() {

	c := visor.NewVisorConfig()
	fmt.Println("We are heer!")

	db, err := bolt.Open(dbpath, 0700, nil)
	fmt.Println("2 We are here")
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	v, err := visor.NewVisor(c, db)

	outs, err := v.GetUnspentOutputs()

	for _, ux := range outs {
		fmt.Println(ux.Head.BkSeq)
	}

	for i := 0; i < 100; i++ {

	}

}

type treasureBookT map[string]uint64

func listAllAddresses() []string {
	// read the blocks one by one

	var addresses []string

	return addresses
}

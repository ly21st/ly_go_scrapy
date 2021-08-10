package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	n = 1000 * 10000
)

func main() {

	db, err := leveldb.OpenFile("./data/db", nil)
	if err != nil {
		fmt.Printf("fail to open db,err:%s", err.Error())
		return
	}

	defer db.Close()
	//
	//batch := new(leveldb.Batch)
	//batch.Delete()

	//prefix := "1354813883946303488"
	var key [8]byte
	binary.LittleEndian.PutUint64(key[:], 1354813883946) //1354813883946303488
	if err := db.Put(key[:], []byte("2123233434"), nil); err != nil {
		fmt.Printf("fail to put key:%s,%s\n", string(key[:]), err.Error())
		return
	}
	val, err := db.Get(key[:], nil)
	if err != nil {
		fmt.Println("fail to get ,err:", err.Error())
	} else {
		fmt.Printf("val=%v\n", val)
		fmt.Println(hex.Dump(val))
	}
}

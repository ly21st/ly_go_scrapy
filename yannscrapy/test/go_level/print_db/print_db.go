package main

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Printf("%s\n", fileName)
	PrintLevelDb(fileName)
}

type Goods struct {
	Id       string `json:"id"`
	GoodsId  string `json:"goodsId"`
	Airport  string `json:"airport"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Flight   string `json:"flight"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Visitor  string `json:"visitor"`
	State    string `json: "state"`
}

func PrintLevelDb(fileName string) {
	db, err := leveldb.OpenFile(fileName, nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	i := 1
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		good := Goods{}
		json.Unmarshal(value, &good)
		fmt.Printf("%d key:%v\n", i, string(key))
		fmt.Printf("%d value:\n", i)
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(good)
		fmt.Printf("\n")
		i = i + 1
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Print(err)
	}
}

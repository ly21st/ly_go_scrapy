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
		var good interface{}
		json.Unmarshal(value, good)
		fmt.Printf("%d key:%v\n", i, string(key))
		fmt.Printf("%d value:\n", i)
		PrettyPrintInterface(good, 2)
		fmt.Printf("\n")
		i = i + 1
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Print(err)
	}
}


func PrettyPrintInterface(s interface{}, ident int) {
	switch v := s.(type) {
	case string:
		fmt.Printf(v)
	case int:
		fmt.Printf("%v", v)
	case int32:
		fmt.Printf("%v", v)
	case int64:
		fmt.Printf("%v", v)
	case float32:
		fmt.Printf("%v", v)
	case float64:
		fmt.Printf("%v", v)
	case json.Number:
		fmt.Printf("%v", v)
	case bool:
		fmt.Printf("%v", v)
	case []interface{}:
		fmt.Printf("[\n")
		for i, listV := range v {
			if i != 0 {
				fmt.Printf(",\n")
			}
			printSpace(ident + 2)
			PrettyPrintInterface(listV, ident + 2)
		}
		fmt.Printf("\n")
		printSpace(ident)
		fmt.Printf("]")
	case map[string]interface {}:
		firstFlag := true
		fmt.Printf("{\n")
		for k, mapV  := range v {
			if !firstFlag {
				fmt.Printf(",\n")
			}
			if firstFlag {
				firstFlag = false
			}
			printSpace(ident + 2)
			fmt.Printf("%v:", k)
			PrettyPrintInterface(mapV, ident + 2)
		}
		fmt.Printf("\n")
		printSpace(ident)
		fmt.Printf("}")
	default:
		fmt.Printf("unknown type: %T, %v", v, v)
	}
}

func printSpace(ident int) {
	for i:= 0; i < ident; i++ {
		fmt.Printf(" ")
	}
}



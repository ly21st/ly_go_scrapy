package main

import (
	"yannscrapy/logging"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("./data/dd", nil)
	if err != nil {
		logging.Error(err)
	}
	defer db.Close()

	data, err := db.Get([]byte("key"), nil)

	logging.Errorf("err=%v", err)

	logging.Infof("data=%v", data)

	err = db.Put([]byte("key"), []byte("value"), nil)
	data, err = db.Get([]byte("key"), nil)
	logging.Infof("data=%v", string(data))

	err = db.Delete([]byte("key"), nil)
	data, err = db.Get([]byte("key"), nil)
	logging.Infof("data=%v", data)

	err = db.Put([]byte("name"), []byte("lilei"), nil)
	err = db.Put([]byte("age"), []byte("30"), nil)

	iter := db.NewIterator(nil, nil)
	logging.Info("-------")
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		logging.Infof("key=%v,value=%v", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()

	logging.Info("-------------------")
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	err = db.Write(batch, nil)

	iter = db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		logging.Infof("key=%v,value=%v", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
}

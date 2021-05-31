package main

import (
	"fmt"

	inmemorydb "github.com/turkaytunc/go-inmem-db/in-memory-db"
)

func main() {

	db := inmemorydb.New()
	db.SetItem("0", "sifir")
	fmt.Println(db.GetItem("0"))
	db.SetItem("1", "bir")
	fmt.Println(db.GetItem("1"))
}

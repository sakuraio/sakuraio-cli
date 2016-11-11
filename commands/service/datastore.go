package service

import "fmt"

func DataStoreGetCommand(options DataStoreOptions) {
	fmt.Println("[Data Store] opsions: ", *options.Size)
}

type DataStoreOptions struct {
	Size    *string
	Unit    *string
	Token   *string
	Order   *string
	Cursor  *string
	After   *string
	Befor   *string
	Channel *string
}

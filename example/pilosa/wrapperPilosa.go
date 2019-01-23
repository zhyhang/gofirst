package main

import (
	"context"
	"fmt"
	"github.com/pilosa/pilosa"
	"github.com/pilosa/pilosa/server"
	"log"
	"os"
)

/**
show pilosa api usage
*/
func main() {
	serverCmd := server.NewCommand(os.Stdin, os.Stdout, os.Stderr)
	defer serverCmd.Close()
	err := serverCmd.Start()
	if err != nil {
		panic(err)
	}
	// get pilosa Server instance
	pilosaServer := serverCmd.Server
	// create api
	apiOption := pilosa.OptAPIServer(pilosaServer)
	api, err := pilosa.NewAPI(apiOption)
	// api usage
	if err != nil {
		log.Println(err)
	} else {
		field, err := api.Field(context.Background(), "did", "fullcolumn")
		if err != nil {
			log.Println(err)
		} else if field != nil {
			fmt.Println(field.Type())
		}
	}

	// create index
	newIndexName := "wrapper-index-1"

	if newIndex(api, newIndexName) {
		log.Printf("create index %s success", newIndexName)
	}

	// directly read index and filed from pilosa Server
	holder := pilosaServer.Holder()
	index := holder.Index(newIndexName)
	fields := index.Fields()
	log.Printf("fields of index %s:", newIndexName)
	fieldStr := "\t"
	for _, field := range fields {
		fieldStr = fieldStr + field.Name() + "\t"
	}
	log.Println(fieldStr)

	// blocking for pilosaServer alive
	serverCmd.Wait()

}

func newIndex(api *pilosa.API, name string) bool {
	options := pilosa.IndexOptions{false, true}
	_, err := api.CreateIndex(context.Background(), name, options)
	success := true
	if err != nil {
		log.Printf("create index %s error, %v", name, err)
		success = false
	}
	newField(api, name, "field1")
	newField(api, name, "field2")
	return success
}

func newField(api *pilosa.API, indexName, fieldName string) bool {
	success := true
	_, err := api.CreateField(context.Background(), indexName, fieldName)
	if err != nil {
		log.Printf("create field %s error, %v", fieldName, err)
		success = false
	}
	return success
}

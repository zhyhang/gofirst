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
	// get server instance
	server := serverCmd.Server
	// create api
	apiOption := pilosa.OptAPIServer(server)
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
	// directly read index and filed from server
	holder := server.Holder()
	index := holder.Index("did")
	fields := index.Fields()
	for _, field := range fields {
		fmt.Println(field.Name())
	}

	// blocking for server alive
	serverCmd.Wait()

}

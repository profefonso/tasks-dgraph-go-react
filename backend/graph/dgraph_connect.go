package graph

import (
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	//"log"
	"google.golang.org/grpc"
)

func DgraphClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("dgraph:9080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
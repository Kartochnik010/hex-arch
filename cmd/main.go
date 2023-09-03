package main

import (
	"hex/internal/adapters/app/api"
	arithmetic "hex/internal/adapters/core"
	rpc "hex/internal/adapters/framework/left/grpc"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"
)

// "internal/adapters/core/arithmetic"

type config struct {
	db struct {
		driver string
		dns    string
	}
}

var (
	dbAdapter   ports.DbPort
	core        ports.ArithmeticPort
	appAdapter  ports.APIPort
	grpcAdapter ports.GRPCPort

	err error
)

func main() {
	dbAdapter, err = db.NewAdapter(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbAdapter)
	grpcAdapter = rpc.NewAdapter(appAdapter)

	grpcAdapter.Run()
}

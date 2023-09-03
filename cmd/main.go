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
	cfg = config{
		db: struct {
			driver string
			dns    string
		}{
			driver: os.Getenv("DB_DRIVER"),
			dns:    os.Getenv("DB_DSN"),
		},
	}

	dbAdapter   ports.DbPort
	core        ports.ArithmeticPort
	appAdapter  ports.APIPort
	grpcAdapter ports.GRPCPort
)

func main() {
	dbAdapter, err := db.NewAdapter(cfg.db.driver, cfg.db.dns)
	if err != nil {
		log.Fatal(err)
	}
	defer dbAdapter.CloseDBConnection()
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbAdapter)
	grpcAdapter = rpc.NewAdapter(appAdapter)
	grpcAdapter.Run()

}

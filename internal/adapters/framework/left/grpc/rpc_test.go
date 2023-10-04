package rpc

import (
	"context"
	"hex/internal/adapters/app/api"
	arithmetic "hex/internal/adapters/core"
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const buffSize = 1024 * 1024

var (
	lis *bufconn.Listener

	dbAdapter   ports.DbPort
	core        ports.ArithmeticPort
	appAdapter  ports.APIPort
	grpcAdapter ports.GRPCPort

	err error
)

func init() {
	lis = bufconn.Listen(buffSize)

	dbAdapter, err = db.NewAdapter(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbAdapter)
	grpcAdapter = NewAdapter(appAdapter)
	s := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(s, grpcAdapter)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("test server failed: %v", err)
		}
	}()
}

func BuffDialer(ctx context.Context, s string) (net.Conn, error) {
	return lis.Dial()
}

func GetGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BuffDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	want := int32(1)
	var a, b int32 = 1, 1
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: a, B: b}

	res, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("Expected %v, got %v", want, err)
	}
	if res.Value != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

func TestGetSubtraction(t *testing.T) {
	var a, b int32 = 1, 1
	want := int32(0)
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: a, B: b}

	res, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("Expected %v, got %v", want, err)
	}
	if res.Value != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

func TestGetMultiplication(t *testing.T) {
	want := int32(1)
	var a, b int32 = 1, 1
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: a, B: b}

	res, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("Expected %v, got %v", want, err)
	}
	if res.Value != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

func TestGetDivision(t *testing.T) {
	want := int32(1)
	var a, b int32 = 1, 1
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: a, B: b}

	res, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("Expected %v, got %v", want, err)
	}
	if res.Value != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

package ports

import (
	"context"
	pb "hex/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	GetDivision(context.Context, *pb.OperationParameters) (*pb.Answer, error)
}

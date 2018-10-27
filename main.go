package main

import (
	"log"
	"net"

	pb "github.com/teamtools/tp-pos-service/proto/pos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(*pb.Pos) (*pb.Pos, error)
}

type PosRepository struct {
	poss []*pb.Pos
}

func (repo *PosRepository) Create(pos *pb.Pos) (*pb.Pos, error) {
	updated := append(repo.poss, pos)
	repo.poss = updated
	return pos, nil
}

type service struct {
	repo Repository
}

func (s *service) CreatePos(ctx context.Context, req *pb.Pos) (*pb.Response, error) {
	pos, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return *pb.Response{Created: true, Pos: pos}, nil
}

func main() {
	repo := *PosRepository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

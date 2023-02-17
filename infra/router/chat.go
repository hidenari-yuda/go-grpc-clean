package router

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *ChatServiceServer) Create(ctx context.Context, req *pb.Chat) (*pb.ChatResponse, error) {

	// Convert context.Context to echo.Context in gRPC server

	fmt.Println("Create")

	var (
		db       = database.NewDb()
		firebase = driver.NewFirebaseImpl()
	)

	input := &entity.Chat{
		From:    uint(req.From),
		Content: req.Content,
	}

	tx, _ := db.Begin()
	i := di.InitializeChatInteractor(tx, firebase)
	res, err := i.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatResponse{
		Chat: &pb.Chat{
			Id:        uint32(res.Id),
			From:      uint32(res.From),
			Content:   res.Content,
			CreatedAt: timestamppb.New(res.CreatedAt),
		},
	}, nil
}

func (s *ChatServiceServer) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.ChatResponse, error) {
	fmt.Println("Get")

	// var (
	// 	db       = database.NewDb()
	// 	firebase = driver.NewFirebaseImpl()
	// )

	i := di.InitializeChatInteractor(s.Db, s.Firebase)
	res, err := i.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ChatResponse{
		Chat: &pb.Chat{
			Id:        uint32(res.Id),
			From:      uint32(res.From),
			Content:   res.Content,
			CreatedAt: timestamppb.New(res.CreatedAt),
		},
	}, nil
}

func (s *ChatServiceServer) GetStream(req *pb.GetStreamRequest, server pb.ChatService_GetStreamServer) error {
	fmt.Println("GetStream")
	// h := di.InitializeChatHandler(s.Db, s.Firebase)
	// err := h.GetStream(req *pb.GetStreamRequest, server pb.ChatService_GetChatStreamServer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := make(chan entity.Chat)

	go func() {
		i := di.InitializeChatInteractor(s.Db, s.Firebase)
		err := i.GetStream(ctx, stream)
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		v := <-stream
		createdAt := timestamppb.New(v.CreatedAt)
		if err := server.Send(&pb.ChatResponse{
			Chat: &pb.Chat{
				From:      uint32(v.From),
				Content:   v.Content,
				CreatedAt: createdAt,
			},
		}); err != nil {
			return err
		}
	}
}

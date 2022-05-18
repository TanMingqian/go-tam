package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/tanmingqian/go-tam/api/apiserver/service/v1"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServiceServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	x, err := s.uc.Create(ctx, &biz.User{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		User: parseUser(x),
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func parseUser(bu *biz.User) *pb.User {
	return &pb.User{
		Meta: &pb.Meta{
			ID:         bu.ID,
			InstanceID: bu.InstanceID,
			Name:       bu.Name,
			Extend:     bu.ExtendShadow,
		},
		Status:      int32(bu.Status),
		Nickname:    bu.Nickname,
		Password:    bu.Password,
		Email:       bu.Email,
		Phone:       bu.Phone,
		IsAdmin:     int32(bu.IsAdmin),
		TotalPolicy: bu.TotalPolicy,
	}
}

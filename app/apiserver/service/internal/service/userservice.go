package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/tanmingqian/go-tam/api/apiserver/service/v1"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/biz"
	"github.com/tanmingqian/go-tam/pkg/metadata"
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
	x, err := s.uc.Create(ctx, convertUser(req.GetUser()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		User: parseUser(x),
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	if err := s.uc.Delete(ctx, req.Name); err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	x, err := s.uc.Update(ctx, convertUser(req.GetUser()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{
		User: parseUser(x),
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	x, err := s.uc.Get(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		User: parseUser(x),
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	x, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	var res []*pb.User
	for _, item := range x.Items {
		res = append(res, parseUser(item))
	}
	return &pb.ListUserReply{
		Results: res,
	}, nil
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

func convertUser(in *pb.User) *biz.User {
	return &biz.User{
		ObjectMeta: metadata.ObjectMeta{
			ID:           in.Meta.ID,
			InstanceID:   in.Meta.InstanceID,
			Name:         in.Meta.Name,
			ExtendShadow: in.Meta.Extend,
		},
		Status:      int(in.Status),
		Nickname:    in.Nickname,
		Password:    in.Password,
		Email:       in.Email,
		Phone:       in.Phone,
		IsAdmin:     int(in.IsAdmin),
		TotalPolicy: in.TotalPolicy,
	}
}

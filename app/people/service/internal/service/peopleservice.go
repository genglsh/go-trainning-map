package service

import (
	"context"

	pb "github.com/genglsh/go-trainning-map/api/classmate/v1"
	"github.com/genglsh/go-trainning-map/app/people/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/emptypb"
)
var ProviderSet = wire.NewSet(NewPeopleService)

type PeopleService struct {
	pb.UnimplementedPeopleServiceServer
	pu              *biz.PeopleUseCase
	log             *log.Helper
}

func NewPeopleService(pu *biz.PeopleUseCase, logger log.Logger) *PeopleService {
	return &PeopleService{
		pu: pu,
		log: log.NewHelper(log.With(logger, "module", "service/people")),
	}
}

func cvtPeple2Use(peoples []*pb.People)[]*biz.People{
	res := make([]*biz.People,0,len(peoples))
	for _, p := range peoples{
		res = append(res, &biz.People{
			ID: p.GetID(),
			Name: p.GetName(),
			TelePhone: p.GetTelephone(),
		})
	}
	return res
}


func (s *PeopleService) CreateCollegeInfo(ctx context.Context,
	req *pb.CreateCollegeInfoRequest) (*emptypb.Empty, error) {
	err := s.pu.CreateCollegeInfo(ctx, cvtPeple2Use(req.Peoples))
	if err!=nil{
		return nil,err
	}
	return &emptypb.Empty{}, nil
}

func cvtbiz2pb(people *biz.People)*pb.People{
	return &pb.People{
		ID: people.ID,
		Name: people.Name,
		Telephone: people.TelePhone,
	}
}
func (s *PeopleService) GetCollegeInfo(ctx context.Context, req *pb.GetCollegeInfoRequest) (*pb.GetCollegeInfoReply, error) {
	res, err := s.pu.GetCollegeInfo(ctx, req.ID)
	if err!=nil{
		return nil, err
	}
	return &pb.GetCollegeInfoReply{
		Peoples: cvtbiz2pb(res),
		RequestID: req.GetRequestID(),
	},nil
}

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type People struct{
	ID string
	Name string
	TelePhone string
}

type PeopleRepo interface{
	CreateCollegeInfo(ctx context.Context, peoples []*People)error
	GetCollegeInfo(ctx context.Context, ID string)(*People,error)
}


type PeopleUseCase struct {
	repo PeopleRepo
	log  *log.Helper
}

func NewPeopleUseCase(repo PeopleRepo, logger log.Logger) *PeopleUseCase {
	return &PeopleUseCase{repo: repo,
		log: log.NewHelper(log.With(logger, "module", "usecase/People"))}
}

func (p *PeopleUseCase)CreateCollegeInfo(ctx context.Context, peoples []*People)error{
	return p.repo.CreateCollegeInfo(ctx,peoples)
}

func (p *PeopleUseCase)GetCollegeInfo(ctx context.Context, ID string)(*People,error){
	return p.repo.GetCollegeInfo(ctx, ID)
}
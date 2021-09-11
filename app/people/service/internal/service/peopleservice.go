package service

import "log"


import (
"vimo-backend/app/label/service/internal/biz"

"github.com/google/wire"

"github.com/go-kratos/kratos/v2/log"

pb "github.com/smartmore/go-trainning-map/api/classmate/v1"
)

const (
	ClearExportImageInfoTime = 500000000
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewLabelService)

type LabelService struct {
	pb.UnimplementedLabelServiceServer
	lu              *biz.LabelUseCase
	ltu             *biz.LabelTypeUseCase
	log             *log.Helper
	ExportImageInfo map[string][]string
}

func NewLabelService(luc *biz.LabelUseCase,
	ltuc *biz.LabelTypeUseCase, logger log.Logger) *LabelService {
	return &LabelService{
		lu:              luc,
		ltu:             ltuc,
		log:             log.NewHelper(log.With(logger, "module", "service/label")),
		ExportImageInfo: make(map[string][]string)}
}


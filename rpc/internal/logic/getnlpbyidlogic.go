package logic

import (
	"context"
	"es_service/rpc/magclient"

	"es_service/rpc/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type GetNlpByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNlpByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNlpByIdLogic {
	return &GetNlpByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNlpByIdLogic) GetNlpById(in *magclient.ReqAbsId) (*magclient.NlpTags, error) {
	tags, err := l.svcCtx.NlpTagsModel.FindOne(in.DocId)
	if err != nil {
		return nil, err
	}
	return &magclient.NlpTags{DocId: tags.DocId, Tags: tags.NlpTags}, nil
}

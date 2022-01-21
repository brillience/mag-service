package logic

import (
	"context"
	"mag_service/rpc/magclient"

	"github.com/tal-tech/go-zero/core/logx"
	"mag_service/rpc/internal/svc"
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
	logx.Infof("[RPC] [Handler] GetNlpById : id: %s", in.DocId)
	tags, err := l.svcCtx.NlpTagsModel.FindOne(in.DocId)
	if err != nil {
		return nil, err
	}
	logx.Infof("\"[RPC] [Handler] GetNlpById : tags:%s", tags.NlpTags)
	return &magclient.NlpTags{DocId: tags.DocId, Tags: tags.NlpTags}, nil
}

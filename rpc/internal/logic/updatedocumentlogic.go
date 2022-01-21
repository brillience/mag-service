package logic

import (
	"context"
	"mag_service/rpc/elastic"
	"mag_service/rpc/magclient"

	"github.com/tal-tech/go-zero/core/logx"
	"mag_service/rpc/internal/svc"
)

type UpdateDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDocumentLogic {
	return &UpdateDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDocumentLogic) UpdateDocument(in *magclient.Abstract) (*magclient.CommonResp, error) {
	err := l.svcCtx.MagEs.UpdateDocument(elastic.Abstract{Id: in.DocId, Content: in.Content})
	switch err {
	case nil:
		return &magclient.CommonResp{Ok: true, Error: ""}, nil
	default:
		return &magclient.CommonResp{Ok: false, Error: "更新失败"}, err
	}
}

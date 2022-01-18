package logic

import (
	"context"
	"es_service/rpc/elastic"

	"es_service/rpc/internal/svc"
	"es_service/rpc/mag"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *UpdateDocumentLogic) UpdateDocument(in *mag.Abstract) (*mag.CommonResp, error) {
	err := l.svcCtx.MagEs.UpdateDocument(elastic.Abstract{Id: in.DocId, Content: in.Content})
	switch err {
	case nil:
		return &mag.CommonResp{Ok: true, Error: ""}, nil
	default:
		return &mag.CommonResp{Ok: false, Error: "更新失败"}, err
	}
}

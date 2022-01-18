package logic

import (
	"context"
	"es_service/rpc/elastic"

	"es_service/rpc/internal/svc"
	"es_service/rpc/mag"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDocumentLogic {
	return &CreateDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDocumentLogic) CreateDocument(in *mag.Abstract) (*mag.CommonResp, error) {
	err := l.svcCtx.MagEs.CreateDocument(elastic.Abstract{
		Id:      in.DocId,
		Content: in.Content,
	})
	if err != nil {
		return &mag.CommonResp{Ok: false, Error: "创建文档失败"}, err
	}
	return &mag.CommonResp{Ok: true, Error: ""}, nil
}

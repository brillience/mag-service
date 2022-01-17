package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &mag.CommonResp{}, nil
}

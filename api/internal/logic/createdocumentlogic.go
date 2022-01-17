package logic

import (
	"context"

	"es_service/api/internal/svc"
	"es_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateDocumentLogic {
	return CreateDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDocumentLogic) CreateDocument(req types.Abstract) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}

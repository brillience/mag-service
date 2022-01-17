package logic

import (
	"context"

	"es_service/api/internal/svc"
	"es_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDocumentLogic {
	return UpdateDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDocumentLogic) UpdateDocument(req types.Abstract) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}

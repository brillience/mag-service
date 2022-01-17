package logic

import (
	"context"

	"es_service/api/internal/svc"
	"es_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDocumentByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDocumentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDocumentByIdLogic {
	return GetDocumentByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentByIdLogic) GetDocumentById(req types.ReqAbsId) (resp *types.Abstract, err error) {
	// todo: add your logic here and delete this line

	return
}

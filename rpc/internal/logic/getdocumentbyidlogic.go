package logic

import (
	"context"

	"es_service/rpc/internal/svc"
	"es_service/rpc/mag"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDocumentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDocumentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentByIdLogic {
	return &GetDocumentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDocumentByIdLogic) GetDocumentById(in *mag.ReqAbsId) (*mag.Abstract, error) {
	// todo: add your logic here and delete this line

	return &mag.Abstract{}, nil
}

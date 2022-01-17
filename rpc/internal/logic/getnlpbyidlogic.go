package logic

import (
	"context"

	"es_service/rpc/internal/svc"
	"es_service/rpc/mag"

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

func (l *GetNlpByIdLogic) GetNlpById(in *mag.ReqAbsId) (*mag.RespNlpTagsMore, error) {
	// todo: add your logic here and delete this line

	return &mag.RespNlpTagsMore{}, nil
}

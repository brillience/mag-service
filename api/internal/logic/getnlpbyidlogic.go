package logic

import (
	"context"

	"es_service/api/internal/svc"
	"es_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetNlpByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNlpByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetNlpByIdLogic {
	return GetNlpByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNlpByIdLogic) GetNlpById(req types.ReqAbsId) (resp []types.NlpTags, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"

	"es_service/api/internal/svc"
	"es_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchDocumentByKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchDocumentByKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchDocumentByKeyLogic {
	return SearchDocumentByKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchDocumentByKeyLogic) SearchDocumentByKey(req types.ReqKeyWord) (resp []types.Abstract, err error) {
	// todo: add your logic here and delete this line

	return
}

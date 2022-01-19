package mag

import (
	"context"
	"es_service/common/errorx"
	"es_service/rpc/magclient"

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

func (l *SearchDocumentByKeyLogic) SearchDocumentByKey(req types.ReqKeyWord) ([]types.Abstract, error) {
	abstracts, err := l.svcCtx.MagRpc.SearchDocumentByKey(l.ctx, &magclient.ReqKeyWord{Key: req.Key})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	var resp []types.Abstract
	for _, abs := range abstracts.Abstracts {
		resp = append(resp, types.Abstract{Docid: abs.DocId, Content: abs.Content})
	}
	return resp, nil
}

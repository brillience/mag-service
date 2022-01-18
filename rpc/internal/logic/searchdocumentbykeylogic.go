package logic

import (
	"context"

	"es_service/rpc/internal/svc"
	"es_service/rpc/mag"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchDocumentByKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchDocumentByKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchDocumentByKeyLogic {
	return &SearchDocumentByKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchDocumentByKeyLogic) SearchDocumentByKey(in *mag.ReqKeyWord) (*mag.RespAbsMore, error) {
	abstracts, err := l.svcCtx.MagEs.SearchDocumentsByKeyWord(in.Key)
	if err != nil {
		return nil, err
	}
	var resp []*mag.Abstract
	for _, abs := range abstracts {
		resp = append(resp, &mag.Abstract{DocId: abs.Id, Content: abs.Content})
	}
	return &mag.RespAbsMore{Abstracts: resp}, nil
}

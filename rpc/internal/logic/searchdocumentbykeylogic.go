package logic

import (
	"context"
	"es_service/rpc/magclient"

	"es_service/rpc/internal/svc"
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

func (l *SearchDocumentByKeyLogic) SearchDocumentByKey(in *magclient.ReqKeyWord) (*magclient.RespAbsMore, error) {
	// todo: add your logic here and delete this line
	abstracts, err := l.svcCtx.MagEs.SearchDocumentsByKeyWord(in.Key)
	if err != nil {
		return nil, err
	}
	var resp []*magclient.Abstract
	for _, abs := range abstracts {
		resp = append(resp, &magclient.Abstract{DocId: abs.Id, Content: abs.Content})
	}
	return &magclient.RespAbsMore{Abstracts: resp}, nil
}

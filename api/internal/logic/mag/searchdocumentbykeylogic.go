package mag

import (
	"context"
	"mag_service/common/errorx"
	"mag_service/rpc/magclient"

	"mag_service/api/internal/svc"
	"mag_service/api/internal/types"

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

func (l *SearchDocumentByKeyLogic) SearchDocumentByKey(req types.ReqKeyWord) (*types.Abstracts, error) {
	//logx.Infof("[！！！Handler：%s ] key:%s;", "SearchDocumentByKey", req.Key)
	abstracts, err := l.svcCtx.MagRpc.SearchDocumentByKey(l.ctx, &magclient.ReqKeyWord{Key: req.Key})
	//logx.Infof("[！！！Handler：%s ] resp nums %d", "SearchDocumentByKey", len(abstracts.Abstracts))
	var resp types.Abstracts
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	for _, abs := range abstracts.Abstracts {
		resp.Data = append(resp.Data, types.Abstract{Docid: abs.DocId, Content: abs.Content})
	}
	return &resp, nil
}

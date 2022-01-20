package mag

import (
	"context"
	"es_service/common/errorx"
	"es_service/rpc/magclient"

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
	abstract, err := l.svcCtx.MagRpc.GetDocumentById(l.ctx, &magclient.ReqAbsId{DocId: req.Docid})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	return &types.Abstract{Docid: abstract.DocId, Content: abstract.Content}, nil
}

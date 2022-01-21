package mag

import (
	"context"
	"mag_service/common/errorx"
	"mag_service/rpc/magclient"

	"mag_service/api/internal/svc"
	"mag_service/api/internal/types"

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

func (l *GetDocumentByIdLogic) GetDocumentById(req types.ReqAbsId) (*types.Abstract, error) {
	logx.Infof("[！！！Handler：%s ] docid:%s;", "GetDocumentById", req.Docid)
	abstract, err := l.svcCtx.MagRpc.GetDocumentById(l.ctx, &magclient.ReqAbsId{DocId: req.Docid})
	logx.Infof("[！！！Handler：%s ] Resp: abstract id: %s content: %s;", abstract.DocId, abstract.Content)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	return &types.Abstract{Docid: abstract.DocId, Content: abstract.Content}, nil
}

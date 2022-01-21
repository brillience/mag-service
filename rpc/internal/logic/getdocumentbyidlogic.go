package logic

import (
	"context"
	"mag_service/rpc/magclient"

	"github.com/tal-tech/go-zero/core/logx"
	"mag_service/rpc/internal/svc"
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

func (l *GetDocumentByIdLogic) GetDocumentById(in *magclient.ReqAbsId) (*magclient.Abstract, error) {
	abstract, err := l.svcCtx.MagEs.GetDocumentById(in.DocId)
	if err != nil {
		return nil, err
	}
	return &magclient.Abstract{DocId: abstract.Id, Content: abstract.Content}, nil
}

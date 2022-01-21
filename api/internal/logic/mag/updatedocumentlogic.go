package mag

import (
	"context"
	"mag_service/common/errorx"
	"mag_service/rpc/magclient"

	"mag_service/api/internal/svc"
	"mag_service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDocumentLogic {
	return UpdateDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDocumentLogic) UpdateDocument(req types.Abstract) (*types.CommonResp, error) {
	commonResp, err := l.svcCtx.MagRpc.UpdateDocument(l.ctx, &magclient.Abstract{DocId: req.Docid, Content: req.Content})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	if commonResp.Ok == false {
		return &types.CommonResp{Ok: false, Error: commonResp.Error}, nil
	}
	return &types.CommonResp{Ok: true}, nil
}

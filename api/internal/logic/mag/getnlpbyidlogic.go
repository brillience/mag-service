package mag

import (
	"context"
	"mag_service/common/errorx"
	"mag_service/rpc/magclient"

	"mag_service/api/internal/svc"
	"mag_service/api/internal/types"

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

func (l *GetNlpByIdLogic) GetNlpById(req types.ReqAbsId) (*types.NlpTags, error) {
	//logx.Infof("[！！！Handler：%s ] docid:%s;", "GetNlpById", req.Docid)
	nlpTags, err := l.svcCtx.MagRpc.GetNlpById(l.ctx, &magclient.ReqAbsId{DocId: req.Docid})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	return &types.NlpTags{DocId: nlpTags.DocId, Tags: nlpTags.Tags}, nil
}

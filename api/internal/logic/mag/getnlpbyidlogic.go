package mag

import (
	"context"
	"es_service/common/errorx"
	"es_service/rpc/magclient"

	"es_service/api/internal/svc"
	"es_service/api/internal/types"

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

func (l *GetNlpByIdLogic) GetNlpById(req types.ReqAbsId) ([]types.NlpTags, error) {
	nlpTagsMore, err := l.svcCtx.MagRpc.GetNlpById(l.ctx, &magclient.ReqAbsId{DocId: req.Docid})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	var resp []types.NlpTags
	for _, tags := range nlpTagsMore.NlpTagsMore {
		resp = append(resp, types.NlpTags{
			DocId:         tags.DocId,
			SentenceIndex: tags.SentenceIndex,
			SentenceText:  tags.SentenceText,
			Tokens:        tags.Tokens,
			Lemmas:        tags.Lemmas,
			PosTags:       tags.PosTags,
			NerTags:       tags.NerTags,
			DocOffsets:    tags.DocOffsets,
			DepTypes:      tags.DepTypes,
			DepTokens:     tags.DepTokens,
		})
	}
	return resp, nil
}

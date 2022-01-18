package logic

import (
	"context"

	"es_service/rpc/internal/svc"
	"es_service/rpc/mag"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetNlpByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNlpByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNlpByIdLogic {
	return &GetNlpByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNlpByIdLogic) GetNlpById(in *mag.ReqAbsId) (*mag.RespNlpTagsMore, error) {
	nlpTags, err := l.svcCtx.NlpTagsModel.FindMoreByDocId(in.DocId)
	if err != nil {
		return nil, err
	}
	var resp []*mag.NlpTags
	for _, tags := range nlpTags {
		resp = append(resp, &mag.NlpTags{
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
	return &mag.RespNlpTagsMore{NlpTagsMore: resp}, nil
}

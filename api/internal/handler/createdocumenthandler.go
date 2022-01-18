package handler

import (
	"net/http"

	"es_service/api/internal/logic"
	"es_service/api/internal/svc"
	"es_service/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateDocumentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Abstract
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateDocumentLogic(r.Context(), svcCtx)
		resp, err := l.CreateDocument(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
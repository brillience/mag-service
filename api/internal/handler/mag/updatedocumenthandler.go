package mag

import (
	"net/http"

	"es_service/api/internal/logic/mag"
	"es_service/api/internal/svc"
	"es_service/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateDocumentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Abstract
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mag.NewUpdateDocumentLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDocument(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

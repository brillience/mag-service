package mag

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mag_service/api/internal/logic/mag"
	"mag_service/api/internal/svc"
	"mag_service/api/internal/types"
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

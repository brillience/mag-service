package mag

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mag_service/api/internal/logic/mag"
	"mag_service/api/internal/svc"
	"mag_service/api/internal/types"
)

func GetDocumentByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqAbsId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mag.NewGetDocumentByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetDocumentById(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

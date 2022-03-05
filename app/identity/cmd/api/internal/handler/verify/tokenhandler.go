package verify

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"it-ku/app/identity/cmd/api/internal/logic/verify"
	"it-ku/app/identity/cmd/api/internal/svc"
	"it-ku/app/identity/cmd/api/internal/types"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := verify.NewTokenLogic(r.Context(), svcCtx)
		resp, err := l.Token(req, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

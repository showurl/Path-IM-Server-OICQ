package imuser

import (
	"net/http"

	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/logic/imuser"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := imuser.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

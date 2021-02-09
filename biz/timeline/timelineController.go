package timeline

import (
	"Moments/pkg/app"
	"Moments/pkg/hint"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimelineRequest struct {
	Schema      string `json:"schema"`
	Uid         int32  `json:"uid"`
	BoundaryAid int64  `json:"boundary_aid"`
}

func RefreshTimeline(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	data := TimelineRequest{}
	if err := c.BindJSON(&data); err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, hint.CustomError{
			Code: hint.INVALID_PARAM,
			Err:  err,
		})
	}

	tlSrv := TimelineService{}
	articleList, err := tlSrv.GetRefreshTimeline(c, data.Uid, data.BoundaryAid, data.Schema)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, err)
	}
	webapp.MakeSuccessJsonRes(http.StatusOK, articleList)
}

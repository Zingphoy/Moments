package album

import (
	"Moments/pkg/app"
	"Moments/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumRequest struct {
	Uid     int32   `json:"uid"`
	AidList []int64 `json:"aid_list"` // this will use as one single value if needed when pass parameter
}

type AlbumResponse struct {
	*Album
}

func GetAlbumDetail(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	uid, err := utils.Str(c.DefaultQuery("uid", "0")).Int32()
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, err)
		return
	}

	srv := NewAlbumService(&Album{}, &AlbumModelImpl{})
	srv.Data.Uid = uid
	err = srv.DetailAlbum(c)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, err)
		return
	}

	webapp.MakeSuccessJsonRes(http.StatusOK, srv.Data)
}

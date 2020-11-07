package api

import (
	"Moments/pkg/app"
	"Moments/pkg/hints"
	"Moments/service/util_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadPicture upload picture to backend server
func UploadPicture(c *gin.Context) {
	webapp := app.GinCtx{C: c}

	picB := c.DefaultQuery("pic", "")

	var retCode int
	for i := 0; i < 3; i++ {
		picBytes := []byte(picB)
		url, retCode := util_service.UploadPictureToTarget("somewhere", picBytes)
		switch retCode {
		case hints.SUCCESS:
			c.Redirect(http.StatusFound, "/moments/post")
			webapp.MakeJsonRes(http.StatusOK, retCode, url)
		case hints.UPLOAD_PIC_FIALED_NET:
			continue
		default:
			break
		}
	}

	if retCode != hints.SUCCESS {
		webapp.MakeJsonRes(http.StatusOK, retCode, nil)
	}
}

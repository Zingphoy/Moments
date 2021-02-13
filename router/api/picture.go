package api

import (
	"Moments/biz/common"
	"Moments/pkg/app"
	"Moments/pkg/hint"
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
		url, retCode := common.UploadPictureToTarget("somewhere", picBytes)
		switch retCode {
		case hint.SUCCESS:
			c.Redirect(http.StatusFound, "/moments/post")
			webapp.MakeJsonRes(http.StatusOK, retCode, url)
		case hint.UPLOAD_PIC_FIALED_NET:
			continue
		default:
			break
		}
	}

	if retCode != hint.SUCCESS {
		webapp.MakeJsonRes(http.StatusOK, retCode, nil)
	}
}

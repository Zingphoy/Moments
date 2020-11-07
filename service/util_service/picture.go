package util_service

import (
	"Moments/pkg/hints"
	"strings"
)

// UploadPictureToTarget upload picture to specific platform
func UploadPictureToTarget(target string, pic []byte) (string, int) {
	var url string
	var err error
	var code int

	switch strings.ToLower(target) {
	case "github":
		if url, err = uploadPictureToGithub(pic); err != nil {

		}

		code = hints.SUCCESS
	case "somewhere":
		if url, err = uploadPictureToGithub(pic); err != nil {

		}
		code = hints.SUCCESS
	default:
		code = hints.SUCCESS
	}
	return url, code
}

func uploadPictureToGithub(pic []byte) (string, error) {
	return "", nil
}

func uploadPictureToSomewhere(pic []byte) (string, error) {
	return "", nil
}

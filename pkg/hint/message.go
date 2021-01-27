package hint

var CodeHint = map[int]string{
	// basic
	SUCCESS:        "ok",
	INTERNAL_ERROR: "error",
	INVALID_PARAM:  "request params invalid",

	// aid
	AID_ALREADY_EXIST: "please ensure that sending article not too fast",

	// uid
	USER_NOT_FOUND: "user does not exist",

	// database
	EMPTY_RESULT:              "query result is empty",
	CONTEXT_DEADLINE_EXCEEDED: "database connection timeout",
	OTHERS:                    "database unknown error",

	// validate
	JSON_PARSE_ERROR: "json parse error",
	// tools
	UPLOAD_PIC_FIALED_NET:          "upload picture failed, network error",
	UOLOAD_PIC_FIALED_AUTH:         "upload picture failed, authority denied",
	UPLOAD_PIC_FIALED_REMOTE_LIMIT: "upload picture failed, remote storage space not enough or upload too often",
	UPLOAD_PIC_FIALED_FORMAT:       "upload picture failed, picture format is not valid",
	UPLOAD_PIC_FIALED_SIZE:         "upload picture failed, size of picture is too large",
}

// GetHintMsg get information by status code
func GetHintMsg(statusCode int) string {
	msg, ok := CodeHint[statusCode]
	if ok {
		return msg
	}
	return CodeHint[INTERNAL_ERROR]
}

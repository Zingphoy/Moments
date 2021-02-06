package hint

import (
	"bytes"
	"strconv"
)

// make sure that each error code has its' corresponding message

type CustomError struct {
	Code int
	Err  error
}

func (e CustomError) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(e.Code))
	buffer.WriteString(": ")
	buffer.WriteString(e.Err.Error())
	return buffer.String()
}

// basic code
const (
	SUCCESS        = 200
	INTERNAL_ERROR = 500
	INVALID_PARAM  = 400
)

/*************** Detail Codes ***************/
const (
	// tools
	UPLOAD_PIC_FIALED_NET          = 1001
	UOLOAD_PIC_FIALED_AUTH         = 1002
	UPLOAD_PIC_FIALED_REMOTE_LIMIT = 1003
	UPLOAD_PIC_FIALED_FORMAT       = 1004
	UPLOAD_PIC_FIALED_SIZE         = 1005

	/*************** Business logic ***************/
	// aid
	AID_ALREADY_EXIST = 1101
	AID_NOT_FOUND     = 1102

	// uid
	UID_ALREADY_EXIST = 1201
	USER_NOT_FOUND    = 1202

	// album
	ALBUM_EMPTY = 1301

	// database
	DATABASE_ERROR            = 2000
	CONNECT_FAILED            = 2001
	CONTEXT_DEADLINE_EXCEEDED = 2002
	QUERY_INTERNAL_ERROR      = 2003
	UPDATE_INTERNAL_ERROR     = 2004
	INSERT_INTERNAL_ERROR     = 2005
	DELETE_INTERNAL_ERROR     = 2006
	OTHERS                    = 2007

	// validate
	JSON_PARSE_ERROR = 3001
)

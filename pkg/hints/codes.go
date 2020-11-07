package hints

// basic code
const (
	SUCCESS        = 200
	INTERNAL_ERROR = 500
	INVALID_PARAM  = 400
)

// detail code
const (
	// aid
	AID_ALREADY_EXIST = 1001

	// uid
	UID_NOT_EXIST = 11001

	// tools
	UPLOAD_PIC_FIALED_NET          = 1201
	UOLOAD_PIC_FIALED_AUTH         = 1202
	UPLOAD_PIC_FIALED_REMOTE_LIMIT = 1203
	UPLOAD_PIC_FIALED_FORMAT       = 1204
	UPLOAD_PIC_FIALED_SIZE         = 1205
)

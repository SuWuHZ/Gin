package service

const (
	ErrNoError = iota
	ErrInternalError
	ErrUserNotExist
	ErrDBQueryFailed
	ErrDBInsertFailed
)

var ErrorCodeMsgMap = map[int]string{
	ErrInternalError: "Error: interal error",
	ErrUserNotExist:  "Error: user not exist.",
	ErrDBQueryFailed: "Error: query db failed.",
	ErrDBInsertFailed: "Error: insert db failed.",
}

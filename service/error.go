package service

const (
	ErrNoError = iota
	ErrInternalError
	ErrUserNotExist
	ErrDBQueryFailed
)

var ErrorCodeMsgMap = map[int]string{
	ErrInternalError: "Error: interal error",
	ErrUserNotExist:  "Error: user not exist.",
	ErrDBQueryFailed: "Error: query db failed.",
}

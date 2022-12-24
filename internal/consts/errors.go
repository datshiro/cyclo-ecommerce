package consts

type ServerErr string

func (e ServerErr) Error() string {
	return ErrorList[e]
}

const (
	ErrUnknown        ServerErr = "ERR_UNKNOWN"
	ErrInvalidRequest ServerErr = "ERR_INVALID_REQUEST"
)

var ErrorList = map[ServerErr]string{
	ErrUnknown:        "Unknown",
	ErrInvalidRequest: "Invalid request",
}

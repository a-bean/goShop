//go:generate stringer -type ErrCode -linecomment
package code

type ErrCode int64 //错误码

// 错误码
const (
	ERR_CODE_OK             ErrCode = 0 //ok
	ERR_CODE_INVALID_PARAMS ErrCode = 1 //参数错误
	ERR_CODE_TIMEOUT        ErrCode = 2 //超时
)

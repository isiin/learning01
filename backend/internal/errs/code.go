package errs

// アプリ固有のエラーコード
type ErrorCode string

// 属性
type errorCodeAttribute struct {
	status  int
	message string
}

const (
	InvalidRequest ErrorCode = "INVALID_REQUEST"
	NotFound       ErrorCode = "NOT_FOUND"
	Exclusion      ErrorCode = "EXCLUSION"
	Internal       ErrorCode = "INTERNAL"
)

var attributes = map[ErrorCode]errorCodeAttribute{
	InvalidRequest: {status: 400, message: "リクエストの形式が不正です"},
	NotFound:       {status: 404, message: "データがありません"},
	Exclusion:      {status: 409, message: "すでに削除されています"},
	Internal:       {status: 500, message: "想定外のエラーが発生しました"},
}

func (e ErrorCode) GetStatus() int {
	if attr, ok := attributes[e]; ok {
		return attr.status
	}
	return 500
}

func (e ErrorCode) GetMessage() string {
	if attr, ok := attributes[e]; ok {
		return attr.message
	}
	return "不明なエラーです"
}

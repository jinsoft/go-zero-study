package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)

	message[DbError] = "数据库繁忙,请稍后再试"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	}
	return "服务器开小差啦,稍后再来试一试"
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}

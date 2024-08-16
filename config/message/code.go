package message

type Structure struct {
	SUCCESS              uint
	FAILURE              uint
	NOAUTH               uint
	AUTHFORMMATER        uint
	INVALIDTOKEN         uint
	PASSWORDNOTTRUE      uint
	STATUSISENABLE       uint
	BADREQUEST           uint
	UNAUTHORIZED         uint
	FORBIDDEN            uint
	NOTFOUND             uint
	INTERNALSERVERERROR  uint
	NOTIMPLEMENTED       uint
	CREATERESOURCEFAILED uint // 添加创建资源失败状态
	Message              map[uint]string
}

var Code = &Structure{
	SUCCESS:              200,
	FAILURE:              500,
	NOAUTH:               403,
	AUTHFORMMATER:        405,
	INVALIDTOKEN:         406,
	PASSWORDNOTTRUE:      410,
	STATUSISENABLE:       411,
	BADREQUEST:           400,
	UNAUTHORIZED:         401,
	FORBIDDEN:            404,
	NOTFOUND:             404,
	INTERNALSERVERERROR:  500,
	NOTIMPLEMENTED:       501,
	CREATERESOURCEFAILED: 502, // 设置创建资源失败的代码
}

func init() {
	Code.Message = map[uint]string{
		Code.SUCCESS:              "成功",
		Code.FAILURE:              "失败",
		Code.NOAUTH:               "header中未包含token",
		Code.AUTHFORMMATER:        "token格式不正确",
		Code.INVALIDTOKEN:         "无效的token或登陆过期，请重新登陆",
		Code.PASSWORDNOTTRUE:      "密码不正确",
		Code.STATUSISENABLE:       "用户状态被禁用",
		Code.BADREQUEST:           "请求参数有误",
		Code.UNAUTHORIZED:         "未授权",
		Code.FORBIDDEN:            "禁止访问",
		Code.NOTFOUND:             "未找到资源",
		Code.INTERNALSERVERERROR:  "服务器内部错误",
		Code.NOTIMPLEMENTED:       "未实现的操作",
		Code.CREATERESOURCEFAILED: "创建资源失败", // 添加对应消息
	}
}

func (s *Structure) GetMessage(code uint) string {
	message, ok := s.Message[code]
	if !ok {
		return ""
	}
	return message
}

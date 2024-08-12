package service

import (
	"easy-go-admin/app/dao"
	"easy-go-admin/app/entity"
	"easy-go-admin/config/jwt"
	"easy-go-admin/config/message"
	"easy-go-admin/config/util"
	"github.com/gin-gonic/gin"
	Validate "github.com/go-playground/validator/v10"
)

// ISystemAccountService 定义接口
type ISystemAccountService interface {
	Login(c *gin.Context, dto entity.LoginData)
	Create(c *gin.Context, dto entity.SystemAccountCreateData)
}

type systemAccountServiceImpl struct{}

func SystemAccountService() ISystemAccountService {
	return &systemAccountServiceImpl{}
}

// Login 用户登陆
func (s *systemAccountServiceImpl) Login(c *gin.Context, dto entity.LoginData) {
	// 登陆参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	// 验证用户信息
	systemAccount := dao.SystemAccountDetail(dto.UserName)
	if systemAccount.ID == 0 {
		message.Fail(c, int(message.Code.NOTFOUND), "用户名不存在")
		return
	}
	if systemAccount.Password != util.EncryptionPassword(dto.Password, systemAccount.Salt) {
		message.Fail(c, int(message.Code.PASSWORDNOTTRUE), message.Code.GetMessage(message.Code.PASSWORDNOTTRUE))
		return
	}
	const status int = 1
	if systemAccount.Status == status {
		message.Fail(c, int(message.Code.STATUSISENABLE), message.Code.GetMessage(message.Code.STATUSISENABLE))
		return
	}

	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(systemAccount)
	message.Success(c, map[string]interface{}{"token": tokenString, "systemAccount": systemAccount})
}

func (s *systemAccountServiceImpl) Create(c *gin.Context, dto entity.SystemAccountCreateData) {
	// 创建ß参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	systemAccount := dao.SystemAccountDetail(dto.UserName)
	if systemAccount.ID != 0 {
		message.Fail(c, int(message.Code.BADREQUEST), "用户名已被占用")
		return
	}
	salt := util.GenerateRandomString(15)
	pass := util.EncryptionPassword(dto.Password, salt)
	createData := entity.SystemAccount{
		UserName: dto.UserName,
		Password: pass,
		Salt:     salt,
		NickName: dto.NickName,
		Email:    dto.Email,
		Avatar:   dto.Avatar,
	}
	err = dao.SystemAccountAdd(createData)
	if err != nil {
		message.Fail(c, int(message.Code.CREATERESOURCEFAILED), "创建用户失败")
		return
	}
	message.Success(c, nil)
}

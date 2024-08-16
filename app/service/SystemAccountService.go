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
	GetDetail(c *gin.Context)
	GetRouters(c *gin.Context)
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
	systemAccount, err := dao.SystemAccountDetailByUsername(dto.UserName)
	if systemAccount.ID == 0 {
		message.Fail(c, int(message.Code.NOTFOUND), "用户名不存在")
		return
	}
	if systemAccount.Password != util.EncryptionPassword(dto.Password, systemAccount.Salt) {
		message.Fail(c, int(message.Code.PASSWORDNOTTRUE), message.Code.GetMessage(message.Code.PASSWORDNOTTRUE))
		return
	}
	const status int = 0
	if systemAccount.Status == status {
		message.Fail(c, int(message.Code.STATUSISENABLE), message.Code.GetMessage(message.Code.STATUSISENABLE))
		return
	}

	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(systemAccount)
	message.Success(c, map[string]interface{}{"token": tokenString, "systemAccount": systemAccount})
}

// Create 创建用户
func (s *systemAccountServiceImpl) Create(c *gin.Context, dto entity.SystemAccountCreateData) {
	// 创建参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	systemAccount, err := dao.SystemAccountDetailByUsername(dto.UserName)

	if systemAccount.ID != 0 {
		message.Fail(c, int(message.Code.BADREQUEST), "用户名已被占用")
		return
	}
	salt := util.GenerateRandomString(15)
	pass := util.EncryptionPassword(dto.Password, salt)
	var avatar string
	if dto.Avatar != nil {
		avatar = *dto.Avatar
	} else {
		avatar = ""
	}
	createData := entity.SystemAccount{
		UserName: dto.UserName,
		Password: pass,
		Salt:     salt,
		NickName: dto.NickName,
		Email:    dto.Email,
		Avatar:   avatar,
	}
	err = dao.SystemAccountCreate(createData)
	if err != nil {
		message.Fail(c, int(message.Code.CREATERESOURCEFAILED), "创建用户失败")
		return
	}
	message.Success(c, nil)
}

// GetDetail 获取当前用户信息
func (s *systemAccountServiceImpl) GetDetail(c *gin.Context) {
	user, err := jwt.GetAccountInfo(c)
	if err != nil {
		message.Fail(c, int(message.Code.FORBIDDEN), err.Error())
		return
	}

	// 提取角色标识
	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Code)
	}

	// 提取按钮权限标识
	var uniquePermissions []string
	for _, role := range user.Roles {
		for _, menu := range role.Menus {
			if menu.Type == 3 {
				uniquePermissions = append(uniquePermissions, menu.Permission)
			}
		}
	}
	// 对uniquePermissions进行去重
	uniquePermissions = util.Unique(uniquePermissions)

	response := map[string]any{
		"id":          user.Id,
		"nickname":    user.NickName,
		"username":    user.UserName,
		"avatar":      user.Avatar,
		"roles":       roles,
		"permissions": uniquePermissions,
	}
	message.Success(c, response)
}

// GetRouters 获取用户可用目录菜单
func (s *systemAccountServiceImpl) GetRouters(c *gin.Context) {
	user, err := jwt.GetAccountInfo(c)
	if err != nil {
		message.Fail(c, int(message.Code.FORBIDDEN), err.Error())
		return
	}
	// 提取可用菜单目录
	uniqueMenus := make(map[uint]*entity.SystemMenu)
	for _, role := range user.Roles {
		for _, menu := range role.Menus {
			if menu.Type == 1 || menu.Type == 2 {
				uniqueMenus[menu.ID] = menu
			}
		}
	}
	// 对菜单进行树状归类
	treeMenuList := BuildNestedMenus(uniqueMenus)
	message.Success(c, treeMenuList)
}

// BuildNestedMenus 对菜单进行树状归类的方法
func BuildNestedMenus(menus map[uint]*entity.SystemMenu) []*entity.SystemMenu {
	// 创建一个映射来存储每个菜单项的引用
	menuMap := make(map[uint]*entity.SystemMenu)

	// 首先，将所有菜单项的引用添加到映射中
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 然后，为每个菜单项设置其子菜单
	for _, menu := range menus {
		if menu.ParentID != 0 {
			// 如果菜单项有父菜单，将其添加到父菜单的子菜单列表中
			parentMenu := menuMap[menu.ParentID]
			if parentMenu != nil {
				parentMenu.Children = append(parentMenu.Children, menuMap[menu.ID])
			}
		}
	}

	// 最后，从映射中提取顶层菜单项
	topMenus := make([]*entity.SystemMenu, 0)
	for _, menu := range menuMap {
		if menu.ParentID == 0 {
			topMenus = append(topMenus, menu)
		}
	}

	return topMenus
}

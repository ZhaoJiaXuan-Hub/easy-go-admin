package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SystemMenu struct {
	ID       uint          `json:"id"`
	ParentID uint          `json:"parentId"`
	Children []*SystemMenu `json:"children"`
}

func TestBuildNestedMenus(t *testing.T) {
	// 准备测试数据
	menus := map[uint]*SystemMenu{
		1: {ID: 1, ParentID: 0},
		2: {ID: 2, ParentID: 1},
		3: {ID: 3, ParentID: 0},
		4: {ID: 4, ParentID: 3},
		5: {ID: 5, ParentID: 4},
		6: {ID: 6, ParentID: 4},
	}

	expected := []*SystemMenu{
		{ID: 1, ParentID: 0, Children: []*SystemMenu{{ID: 2, ParentID: 1, Children: []*SystemMenu{}}}},
		{ID: 3, ParentID: 0, Children: []*SystemMenu{{ID: 4, ParentID: 3, Children: []*SystemMenu{{ID: 5, ParentID: 4, Children: []*SystemMenu{}}, {ID: 6, ParentID: 4, Children: []*SystemMenu{}}}}}},
	}

	// 调用待测试的函数
	actual := buildNestedMenus(menus)

	// 使用 testify 包的 assert 函数进行断言
	assert.Equal(t, expected, actual, "The actual output should match the expected output")
}
func buildNestedMenus(menus map[uint]*SystemMenu) []*SystemMenu {
	// 创建一个映射来存储每个菜单项的引用
	menuMap := make(map[uint]*SystemMenu)

	// 首先，将所有菜单项的引用添加到映射中
	for _, menu := range menus {
		menuMap[menu.ID] = &SystemMenu{
			ID:       menu.ID,
			ParentID: menu.ParentID,
			Children: make([]*SystemMenu, 0),
		}
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
	topMenus := make([]*SystemMenu, 0)
	for _, menu := range menuMap {
		if menu.ParentID == 0 {
			topMenus = append(topMenus, menu)
		}
	}

	return topMenus
}

package controller

import (
	"easy-go-admin/app/dao"
	"easy-go-admin/config/message"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var total int64
	total = 0
	where := []interface{}{
		[]interface{}{"id", "in", []int{1, 2, 3, 4, 5, 6, 7}},
	}
	list, _ := dao.SystemMenuQuery(where, []string{"*"}, "id desc", 1, 3, &total)
	message.Success(c, map[string]interface{}{"list": list, "count": total})
}

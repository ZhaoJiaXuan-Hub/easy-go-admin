package controller

import (
	"easy-go-admin/app/dao"
	"easy-go-admin/config/message"
	"easy-go-admin/config/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var total int64
	total = 0
	where := []interface{}{
		[]interface{}{"id", "in", []int{1, 2, 3, 4, 5, 6, 7}},
	}
	query := util.QueryOption{
		Where:   where,
		Page:    1,
		Rows:    3,
		OrderBy: "id desc",
		Columns: []string{"*"},
	}
	list, _ := dao.SystemMenuQuery(query, &total)
	message.Success(c, map[string]interface{}{"list": list, "count": total, "query": query})
}

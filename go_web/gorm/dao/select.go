package dao

import (
	"learn/go_web/gorm/conn"
	"learn/go_web/gorm/models"
)

/**
  * First 获取第一条记录（主键排序）：SELECT * FROM user ORDER BY id LIMIT 1;
  * Last  获取最后一条记录（主键排序）：SELECT * FROM user ORDER BY id DESC LIMIT 1;
  * Take  获取一条记录，没有指定排序字段：SELECT * FROM user LIMIT 1;
  * Find  满足条件的数据记录中随机返回一条，即便没有找到记录，也不会抛出错误
**/

// SelectSingleUser 查找单个数据
func SelectSingleUser() (models.User, error) {
	var user models.User
	// rs := conn.Db.First(&user)
	// rs := conn.Db.Last(&user)
	rs := conn.Db.Take(&user)
	// rs := conn.Db.Find(&user)
	if rs.Error != nil {
		return models.User{}, rs.Error
	}
	return user, nil
}

// SelectUserById 按主键检索
func SelectUserById(id int) (models.User, error) {
	var user models.User
	// rs := conn.Db.First(&user, id)
	rs := conn.Db.Find(&user, id)
	if rs.Error != nil {
		return models.User{}, rs.Error
	}
	return user, nil
}

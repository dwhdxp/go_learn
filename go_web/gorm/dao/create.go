package dao

import (
	"fmt"
	"learn/go_web/gorm/conn"
	"learn/go_web/gorm/models"
)

// CreateUser 创建用户记录
// INSERT INTO `users` (`name`,`age`,`email`) VALUES (?, ?, ?)
func CreateUser(u models.User) (uint, error) {
	user := models.User{Name: u.Name, Age: u.Age, Email: u.Email}

	rs := conn.Db.Create(&user)
	if rs.Error != nil {
		fmt.Println("mysql create user failed")
		return 0, rs.Error
	}
	fmt.Printf("rows affected: %d\n", rs.RowsAffected)
	return user.ID, nil
}

// CreateUsers 批量创建用户记录
func CreateUsers(us []models.User) error {
	var users []models.User
	for _, u := range us {
		user := models.User{Name: u.Name, Age: u.Age, Email: u.Email}
		users = append(users, user)
	}
	rs := conn.Db.Create(&users)
	if rs.Error != nil {
		fmt.Println("mysql create users failed")
		return rs.Error
	}
	fmt.Printf("rows affected: %d\n", rs.RowsAffected)
	return nil
}

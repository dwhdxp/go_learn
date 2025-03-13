package main

import (
	"fmt"
	"gorm.io/gorm"
	"learn/go_web/gorm/conn"
	"learn/go_web/gorm/models"
)

func main() {
	//// 1.创建记录
	//// 单项创建
	//user1 := models.User{Name: "dxp", Age: 25, Email: "dxp@163.com"}
	//id, err := dao.CreateUser(user1)
	//if err != nil {
	//	fmt.Println("create user err:", err)
	//}
	//fmt.Println("create user id:", id)
	//
	//// 批量创建
	//users1 := []models.User{
	//	{Name: "lxz", Age: 2, Email: "lxz@qq.com"},
	//	{Name: "dxz", Age: 1, Email: "dxz@qq.com"},
	//	{Name: "lsx", Age: 26, Email: "lsx@163.com"},
	//}
	//if err := dao.CreateUsers(users1); err != nil {
	//	fmt.Println("create users err:", err)
	//}

	//// 2.查询记录
	//// 单项查询
	//user2, err := dao.SelectSingleUser()
	//if err != nil {
	//	fmt.Println("select single user err:", err)
	//}
	//fmt.Printf("select single user:%v\n", user2)
	//
	//user2, err = dao.SelectUserById(10) // Find不会报错，其他会
	//if err != nil {
	//	fmt.Println("select single user by id err:", err)
	//}
	//fmt.Printf("select single user by id:%v\n", user2)
	//
	//// 条件查询
	//rsDB := conn.Db.Where("id > ?", 10).Find(&user2)
	//if rsDB.Error != nil {
	//	fmt.Println("select single user by where1 err:", rsDB.Error)
	//}
	//fmt.Printf("select single user by where1:%v\n", user2)
	//
	//user2 = models.User{}
	//rsDB = conn.Db.Where("name = ?", "lsx").Find(&user2)
	//if rsDB.Error != nil {
	//	fmt.Println("select single user by where2 err:", rsDB.Error)
	//}
	//fmt.Printf("select single user by where2:%v\n", user2)
	//
	//// 批量查询
	//var users2 []models.User
	//rsDB = conn.Db.Where("age > ?", 18).Find(&users2)
	//if rsDB.Error != nil {
	//	fmt.Println("select single users err:", rsDB.Error)
	//}
	//for _, user := range users2 {
	//	fmt.Printf("user:%v\n", user)
	//}
	//
	//// 查询全部对象
	//rsDB = conn.Db.Find(&users2)
	//if rsDB.Error != nil {
	//	fmt.Println("select single users err:", rsDB.Error)
	//}
	//for _, user := range users2 {
	//	fmt.Printf("user:%v\n", user)
	//}

	//// 3.更新
	//// 更新并获取更新后的记录：使用First查找记录，再使用Save进行修改保存
	//var user3 models.User
	//rsDB := conn.Db.Where("name = ?", "dxz").First(&user3)
	//user3.Email = "dxz@163.com"
	//rsDB = conn.Db.Save(&user3)
	//if rsDB.Error != nil {
	//	fmt.Printf("update failed, err:%v\n", rsDB.Error)
	//}
	//fmt.Printf("update successful user:%v\n", user3)
	//
	//// 通过表达式执行sql更新
	//rsDB = conn.Db.Where("id > ?", 8).Model(models.User{}).UpdateColumn(
	//	"age", gorm.Expr("age + ?", 1))
	//if rsDB.Error != nil {
	//	fmt.Printf("update failed, err:%v\n", rsDB.Error)
	//}
	//
	//rsDB = conn.Db.Where("age > ? AND age < ?", 1, 10).Model(models.User{}).UpdateColumn(
	//	"age", gorm.Expr("age * ? + ?", 2, 3))
	//if rsDB.Error != nil {
	//	fmt.Printf("update failed, err:%v\n", rsDB.Error)
	//}
	//
	//// 更新某列数据
	//rsDB = conn.Db.Where("name = ?", "dxp").Model(models.User{}).Update(
	//	"name", "dxpxp")
	//if rsDB.Error != nil {
	//	fmt.Printf("update failed, err:%v\n", rsDB.Error)
	//}

	//// 4.删除
	//// 按条件删除一条记录
	//rsDB := conn.Db.Where("name = ?", "dxpxp").Delete(&models.User{})
	//if rsDB.Error != nil {
	//	fmt.Printf("delete user failed:%v\n", rsDB.Error)
	//}
	//
	//// 批量删除
	//rsDB = conn.Db.Where("email LIKE ?", "%@%").Delete(&models.User{})
	//if rsDB.Error != nil {
	//	fmt.Printf("delete user failed:%v\n", rsDB.Error)
	//}

	// 5.事务：db.Begin()、tx.Commit()、tx.Rollback()
	tx := conn.Db.Begin()
	rsDB := tx.Where("age > ?", 100).Model(models.User{}).UpdateColumn(
		"age", gorm.Expr("age - ?", 5))
	if rsDB.Error != nil {
		tx.Rollback()
		fmt.Printf("update user.Age failed:%v\n", rsDB.Error)
		return
	}
	rsDB = tx.Where("id = ?", "dxp").Delete(models.User{})
	if rsDB.Error != nil {
		tx.Rollback()
		fmt.Printf("delete user failed:%v\n", rsDB.Error)
		return
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		fmt.Printf("commit failed:%v\n", err)
		return
	}
}

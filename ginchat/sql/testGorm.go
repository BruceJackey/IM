package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:jcx0325jcx..@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Contact{})
	// db.AutoMigrate(&models.Message{})
	// db.AutoMigrate(&models.GroupBasic{})
	// Create
	// user := &models.UserBasic{}
	// user.Name = "jcx"
	// db.Create(user)
	// //db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	// fmt.Println(db.First(user, 1))
	// // var product Product
	// // db.First(&product, 1)                 // 根据整型主键查找
	// // db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // Update - 将 product 的 price 更新为 200
	// db.Model(user).Update("Password", "123456")
	// Update - 更新多个字段
	// db.Model(user).Updates(models.UserBasic{Name: "jcx2", Password: "123456"}) // 仅更新非零值字段
	// //db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	//db.Delete(user, 1)
}

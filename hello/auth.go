package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var Users []User

func GetUserList(db *gorm.DB) (list []User, err error) {
	if err == db.Find(&list).Error {
		return
	}
	return
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func FindUser(db *gorm.DB, username string) error {
	user := new(User)
	err := db.Where("Username = ?", username).First(&user).Error
	return err
}

func FindPassword(db *gorm.DB, username string, password string) error {
	user := new(User)
	err := db.Where("Username = ? AND Password = ?", username, password).First(&user).Error
	return err
}

func UpdateUser(db *gorm.DB, id int64, newPassword string) error {
	user := new(User)
	err := db.Model(&user).Where("ID = ?", id).Update("Password", newPassword).Error
	return err
}

func DeleteUser(db *gorm.DB, id int64) error {
	user := new(User)
	err := db.Where("ID = ?", id).Delete(&user).Error
	return err
}

func CheckUserIsExist(username string) bool {
	if err := FindUser(MysqlDB, username); err == nil {
		return true
	} else {
		return false
	}
}

func CheckPassword(username string, password string) error {
	if err := FindPassword(MysqlDB, username, password); err == nil {
		return nil
	} else {
		return errors.New("password is not correct")
	}
}

func Auth(user *User) error {
	if isExist := CheckUserIsExist(user.Username); isExist {
		return CheckPassword(user.Username, user.Password)
	} else {
		return errors.New("user is not exist")
	}
}

func SignUp(c *gin.Context) {
	RedisDB.Incr(Ctx, "counter")
	var user User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	user.Phone = c.Request.FormValue("phone")
	if err := CreateUser(MysqlDB, &user); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "註冊成功",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
}

func UserList(c *gin.Context) {
	RedisDB.Incr(Ctx, "counter")
	if uuu, err := GetUserList(MysqlDB); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success":  "取得成功",
			"userlist": uuu,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
}

func LoginAuth(c *gin.Context) {
	var user User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	if err := Auth(&user); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}
}

func ChangePassword(c *gin.Context) {
	RedisDB.Incr(Ctx, "counter")
	var user User
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	if err := UpdateUser(MysqlDB, id, user.Password); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "更改成功",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}
}

func Destroy(c *gin.Context) {
	RedisDB.Incr(Ctx, "counter")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := DeleteUser(MysqlDB, id); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "刪除成功",
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
			"123":   err,
		})
		return
	}
}

func CountAPI(c *gin.Context) {
	val, err := RedisDB.Get(Ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"call API times": val,
	})
}

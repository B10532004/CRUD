package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func GetUserList() (list []User, err error) {
	if err == MysqlDB.Find(&list).Error {
		return
	}
	return
}

func CreateUser(user *User) error {
	return MysqlDB.Create(user).Error
}

func FindUser(username string) error {
	err := MysqlDB.Where("Username = ?", username).First(User{}).Error
	return err
}

func FindPassword(username string, password string) error {
	err := MysqlDB.Where("Username = ? AND Password = ?", username, password).First(User{}).Error
	return err
}

func UpdateUser(id int64, newPassword string) error {
	err := MysqlDB.Table("users").Where("ID = ?", id).Update("Password", newPassword).Error
	return err
}

func DeleteUser(id int64) error {
	err := MysqlDB.Where("ID = ?", id).Delete(User{}).Error
	return err
}

func CheckUserIsExist(username string) bool {
	if err := FindUser(username); err == nil {
		return true
	} else {
		return false
	}
}

func CheckPassword(username string, password string) error {
	if err := FindPassword(username, password); err == nil {
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
	var user User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	user.Phone = c.Request.FormValue("phone")
	if err := CreateUser(&user); err == nil {
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
	if uuu, err := GetUserList(); err == nil {
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
	var user User
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	if err := UpdateUser(id, user.Password); err == nil {
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := DeleteUser(id); err == nil {
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
	val, err := RedisDB.Get(Ctx, Counter).Result()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"call API times": val,
	})
}

func AddAPI() {
	RedisDB.Incr(Ctx, Counter)
}

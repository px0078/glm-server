// 用户服务
package service

import (
	"fmt"
	"net/http"
)

// SignIn 登录
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello SignIn", r.URL, r.Method)

	w.Write([]byte("SignIn"))
}

// SignUp 注册
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUp")

	// user := mongorm.User{}
	// json := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	println(
		r.FormValue("name"))
	println(
		r.PostFormValue("name"))
	// 查库看有没有这个用户
	// err := mongorm.DB.Collection("users").FindOne(c, bson.M{"_id": bson.M{"$in": user.ID}}).Decode(&user)

	// if err != nil {
	// 	// 已注册，返回400
	// 	c.JSON(400, gin.H{
	// 		"message": "Already SignUp",
	// 	})
	// }

	// // 未注册，创建用户
	// err = user.Create(c, mongorm.DB, "users", user)
	// if err != nil {
	// 	// 创建失败，返回500
	// 	c.JSON(500, gin.H{
	// 		"message": "SignUp Failed",
	// 	})
	// }

	// c.JSON(200, gin.H{
	// 	"message": "SignUp Success",
	// })
}

// GetUser 获取用户信息
func GetUser(w http.ResponseWriter, r *http.Request) {}

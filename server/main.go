/*
 * @Author: neozhang
 * @Date: 2022-06-16 11:01:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-16 11:03:15
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"gcpp/server/web"
	"net/http"
)

//http://localhost:8080/index.gohtml
func main() {
	web.RegisterRouter()
	fmt.Println("start Go Web Server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

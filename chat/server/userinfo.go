/*
 * @Author: neozhang
 * @Date: 2022-06-16 11:07:56
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-16 11:07:57
 * @Description: 请填写简介
 */
package server

//userInfo用于存储用户信息
type userInfo struct {
	name    string
	perC    chan []byte
	AddUser chan []byte //广播用户进入或退出
}

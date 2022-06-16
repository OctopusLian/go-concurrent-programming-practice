/*
 * @Author: neozhang
 * @Date: 2022-06-16 11:09:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-16 11:09:03
 * @Description: 请填写简介
 */
package server

func broadcast(userName string) {
	for _, v := range onlineUsers {
		v.AddUser <- []byte("用户[" + userName + "]已加入聊天室\n")
	}
}

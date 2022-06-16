/*
 * @Author: neozhang
 * @Date: 2022-06-16 11:10:10
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-16 11:10:10
 * @Description: 请填写简介
 */
package server

import "net"

func connWrite(conn net.Conn, user userInfo) {

	for {
		select {
		case msg1 := <-user.AddUser:
			_, _ = conn.Write(msg1)
		case msg2 := <-user.perC:
			_, _ = conn.Write(msg2)
		}
	}
}

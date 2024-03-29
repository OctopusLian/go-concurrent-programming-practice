/*
 * @Author: neozhang
 * @Date: 2022-06-16 11:07:31
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-16 11:07:32
 * @Description: 请填写简介
 */
package server

import "github.com/fatih/color"

//全局channel，处理从各个客户端读到的消息
var message = make(chan []byte)

//用于存储在线用户信息
var onlineUsers = make(map[string]userInfo)
var ccS = color.New(color.FgGreen, color.Bold) //绿色
var ccF = color.New(color.FgRed, color.Bold)   //红色

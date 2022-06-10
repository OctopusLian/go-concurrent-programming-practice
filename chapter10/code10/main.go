package main

import (
	"fmt"

	"gcpp/chapter10/code10/entity"
	"gcpp/chapter10/code10/ydorm"
)

func main() {
	stu := entity.Student{
		Id:   "001",
		Name: "Jack",
		Age:  25,
		Sex:  "男",
	}
	if isOk, _ := ydorm.Save(&stu); isOk {
		fmt.Println("新增成功")
	}
	if isOk, _ := ydorm.Update(&stu); isOk {
		fmt.Println("更新成功")
	}
	if isOk, _ := ydorm.Delete(&stu); isOk {
		fmt.Println("删除成功")
	}

}

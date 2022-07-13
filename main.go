/**
 * @Author: chenchen
 * @Description:
 * @File: main.go
 * @Version: 1.0.0
 * @Data: 2022/07/13 22:03
 */

package main

import "douyin/dao"

func main() {
	initDeps()
}

func initDeps() {
	dao.Init()
}

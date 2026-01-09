package main

import (
	"BondPledge/api/models"
	"BondPledge/api/validate"
	"BondPledge/db"
	"BondPledge/log"
)

func main() {

	//初始化数据库
	db.InitMysql()
	//初始化redis
	// db.InitRedis()
	//初始化数据表
	models.InitTable()

	log.Logger.Info("Init Mysql Init Done")

	validate.BindingValidator()
}

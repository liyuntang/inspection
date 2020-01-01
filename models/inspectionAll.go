package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
)

func InspectionAll() ([]map[string]string) {
	// 解析数据库配置
	drive := beego.AppConfig.String("database::drive")
	address := beego.AppConfig.String("database::address")
	port := beego.AppConfig.String("database::port")
	schema := beego.AppConfig.String("database::schema")
	engPoint := fmt.Sprintf("%s:%s", address, port)
	// 连接数据库获取数据
	engine := engine(drive, engPoint, schema)
	sql := fmt.Sprintf("select * from host_info order by product,cluster_name,role;")
	result, err := engine.QueryString(sql)
	if err != nil {
		return nil
	}
	hostSlice := []map[string]string{}
	for _, dict := range result {
		time := dict["add_time"]
		fmt.Println("time is", time, "type of time is", reflect.TypeOf(time))

		hostMap := map[string]string{}
		hostMap["address"] = dict["address"]
		hostMap["cluster_name"] = dict["cluster_name"]
		hostMap["db_kind"] = dict["db_kind"]
		hostMap["hostname"] = dict["hostname"]
		hostMap["port"] = dict["port"]
		hostMap["product"] = dict["product"]
		hostMap["role"] = dict["role"]
		hostMap["time"] = dict["add_time"]
		hostSlice = append(hostSlice, hostMap)
	}
	return hostSlice
}

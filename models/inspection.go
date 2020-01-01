package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
)

// Inspection函数用于对某个集群发起巡检，该函数接收一个参数：集群名称
// Inspection的过程分为4个：
//	1、通过集群名称到db_info表里获取具体的mysql实例的endpoint --> dbSlice
//	2、获取巡检项		--> inspectionSlice
//	3、到mysql里中检索巡检项
//	4、巡检结果入库，同时展示到页面
// 案例sql：elect hostname, role, port from db_info where product = 'nice' and cluster_name = 'sneaker';

func Inspection(product, clusterName string) ([]map[string]string) {
	// 解析数据库配置
	drive := beego.AppConfig.String("database::drive")
	address := beego.AppConfig.String("database::address")
	port := beego.AppConfig.String("database::port")
	schema := beego.AppConfig.String("database::schema")
	engPoint := fmt.Sprintf("%s:%s", address, port)
	// 连接数据库获取数据
	engine := engine(drive, engPoint, schema)
	//	1、通过集群名称到db_info表里获取具体的mysql实例的endpoint
	dbSlice, err := getEndPoint(product, clusterName, engine)
	if err != nil {
		fmt.Println("get endpoint is bad, err is", err)
		return nil
	}
	fmt.Println(dbSlice)
	fmt.Println("==============================================================")
	//	2、获取巡检项
	inspectionSlice := beego.AppConfig.String("inspection_project::project")
	fmt.Println(inspectionSlice)

	// 3、到mysql里中检索巡检项



	return nil
}

func inspectionDB(dbSlice []map[string]string, project []string)  {
	// 遍历dbSlice获取单个数据库信息
	for _, db := range dbSlice {
		// 连接数据库
		endPoint := fmt.Sprintf("%s:%s", db["hostname"], db["port"])
		engine := engine("mysql", endPoint, "")
		engine.QueryString("show variables;")
	}
}




func getEndPoint(product, clusterName string, engine *xorm.Engine) (dbslice []map[string]string, err error){
	// 声明一个slice用于存放数据库实例信息
	dbSlice := []map[string]string{}
	sql := fmt.Sprintf("select hostname, role, port from db_info where product = '%s' and cluster_name = '%s';", product, clusterName)
	resSlice, err := engine.QueryString(sql)
	if err != nil {
		fmt.Println("select from db is bad, err is", err)
		return nil, err
	}
	// 说明访问数据库成功
	for _, dict := range resSlice {
		dbMap := map[string]string{}
		hostname := dict["hostname"]
		port := dict["port"]
		role := dict["role"]
		dbMap["hostname"] = hostname
		dbMap["port"] = port
		dbMap["role"] = role
		dbSlice = append(dbSlice, dbMap)
	}
	return dbSlice, nil
}

//func Inspection(hostname string) ([]map[string]string) {
//	// 解析数据库配置
//	drive := beego.AppConfig.String("database::drive")
//	address := beego.AppConfig.String("database::address")
//	port := beego.AppConfig.String("database::port")
//	schema := beego.AppConfig.String("database::schema")
//	engPoint := fmt.Sprintf("%s:%s", address, port)
//	// 连接数据库获取数据
//	engine := engine(drive, engPoint, schema)
//	sql := fmt.Sprintf("select * from host_info where cluster_name = '%s' order by product,cluster_name,role;", hostname)
//	result, err := engine.QueryString(sql)
//	if err != nil {
//		return nil
//	}
//	hostSlice := []map[string]string{}
//	for _, dict := range result {
//		time := dict["add_time"]
//		fmt.Println("time is", time, "type of time is", reflect.TypeOf(time))
//
//		hostMap := map[string]string{}
//		hostMap["address"] = dict["address"]
//		hostMap["cluster_name"] = dict["cluster_name"]
//		hostMap["db_kind"] = dict["db_kind"]
//		hostMap["hostname"] = dict["hostname"]
//		hostMap["port"] = dict["port"]
//		hostMap["product"] = dict["product"]
//		hostMap["role"] = dict["role"]
//		hostMap["time"] = dict["add_time"]
//		hostSlice = append(hostSlice, hostMap)
//	}
//	return hostSlice
//}

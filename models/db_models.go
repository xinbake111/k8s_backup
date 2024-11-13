package models

import (
	"WorkPro/conn"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// 插入关系内容
type SpaceResForPVC struct {
	NameSpace string
	PodName string
	DeployName string
	DaemonSetName string
	JobName string
	CornJobName string
	PvcName string
	Path	string
}
type BackupTime struct {
	FilePath string
	Time string
}
func InsertRelation(spec SpaceResForPVC) (bool) {
	_,err:= conn.ModifyDB("insert into volumerelation(namespace,podname,deployname,daemonset,jobname,cornjobname,pvcname,path) " +
		"values (?,?,?,?,?,?,?,?)",spec.NameSpace,spec.PodName,spec.DeployName,spec.DeployName,
		spec.DaemonSetName,spec.JobName,spec.CornJobName,spec.PvcName,spec.Path)
	if err!=nil {
		return false
	}
	return true
}
//按条件查询
func SelectRelation(con string) string {
	sql := fmt.Sprintf("select path from volumerelation %s",con)
	fmt.Println(sql)
	row := conn.QueryRowDB(sql)
	var resPath string
	//row 进行数据赋值
	row.Scan(&resPath)
	return resPath
}
//根据 namespace与podname进行查询
func QueryResources(namespace string,podname string) string {
	sql:=fmt.Sprintf("where namespace='%s' and podname='%s' ",namespace,podname)
	return SelectRelation(sql)
}
func QueryDeployResources(namespace string,deployname string) string {
	sql:=fmt.Sprintf("where namespace='%s' and deployname='%s' ",namespace,deployname)
	return SelectRelation(sql)
}
//进行密钥存储
func InsertEccKey(prvkey *ecdsa.PrivateKey,filePath string)(bool){
	bprk:=crypto.FromECDSA(prvkey)
	strprk:=hex.EncodeToString(bprk)
	//二次加密
	e:=Constructor()
	strprk=e.Encrypt(strprk)
	_, err := conn.ModifyDB("insert into ecckey(prvkey,filepath) "+
		"values (?,?)", strprk,filePath)
	if err!=nil {
		return false
	}
	//备份之后删除重复字段
	DelSameKey()
	return true

}
func InsertEccPath(filePath string,fileKeyPath string)(bool)  {
	_, err := conn.ModifyDB("insert into ecckeypath(filekeypath,filepath) "+
		"values (?,?)", fileKeyPath,filePath)
	if err!=nil {
		return false
	}
	//备份之后删除重复字段
	DelSameKey()
	return true
}

//进行密钥读取 根据filepath进行查询
func QueryEccKey(filePath string) *ecdsa.PrivateKey{
	sql:=fmt.Sprintf("select prvkey from ecckey where filepath='%s' ",filePath)
	row := conn.QueryRowDB(sql)
	var strprk string
	//row 进行数据赋值
	row.Scan(&strprk)
	//自定义解密
	e:=Constructor()
	strprk1:=e.Decrypt(strprk)

	hprk,err:=hex.DecodeString(strprk1)
	if err!=nil {
		log.Println(err)
	}
	prk, err := crypto.ToECDSA(hprk)

	return prk
}
//读取密钥存储路径
//进行密钥读取 根据file path进行查询
func QueryEccKeyPath(filePath string) string{
	sql:=fmt.Sprintf("select filekeypath from ecckeypath where filepath='%s' ",filePath)
	row := conn.QueryRowDB(sql)
	var prvpath string
	//row 进行数据赋值
	row.Scan(&prvpath)
	return prvpath
}
//读取资源目录以及备份时间
func QueryTimeBackup()  []BackupTime{
	sql:=fmt.Sprintf("select filepath,time from ecckeypath where 1=1 ",)
	rows ,err:= conn.QueryDB(sql)
	if err != nil {
		log.Println(nil)
	}
	var bkt []BackupTime
	for rows.Next() {
		filepath := ""
		createtime :=""
		rows.Scan( &filepath, &createtime)
		t := BackupTime{filepath, createtime}
		bkt = append(bkt, t)
	}

	return bkt
}
//进行密钥删除
func DelEccKey(filePath string)(int64,error)  {
	return conn.ModifyDB("delete from ecckey where filepath=?",filePath)
}
//删除重复字段
func DelSameKey() (int64,error)  {

	sql:="DELETE  from  ecckey where id in(\n SELECT id from (\n        SELECT id,filepath FROM ecckey WHERE filepath in(\n            SELECT filepath FROM ecckey GROUP BY filepath HAVING count(filepath) > 1)\n            AND time  not IN(SELECT max(time) FROM ecckey GROUP BY filepath HAVING count(filepath) > 1)\n ) as t )"
	return conn.ModifyDB(sql)

}



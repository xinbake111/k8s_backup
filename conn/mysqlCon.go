package conn

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)


/*
driverName = mysql
mysqlUser =root
mysqlPwd = root
host = 127.0.0.1
port = 3306
dbname = k8sbackup
*/
var db *sql.DB

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName, _ := beego.AppConfig.String("driverName")

	//注册数据库驱动
	//orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user, _ := beego.AppConfig.String("mysqlUser")
	pwd, _ := beego.AppConfig.String("mysqlPwd")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	dbname, _ := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableVolumeRelation()
		CreateEccKey()
		CreateEccKeyPath()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}


//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//创建关系映射表 namespace,podname,deployname,daemonset,jobname,cornjobname,pvcname,path
func CreateTableVolumeRelation() {
	sql := `CREATE TABLE IF NOT EXISTS volumerelation(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		namespace VARCHAR(64),
		podname VARCHAR(64),
		deployname VARCHAR(64),
		daemonset VARCHAR(64),
		jobname VARCHAR(64),
		cornjobname VARCHAR(64),
		pvcname VARCHAR(64),
		path VARCHAR(64)
		);`
	ModifyDB(sql)
}
//创建存放圆锥曲线密钥
func CreateEccKey() {
	sql := `CREATE TABLE IF NOT EXISTS ecckey(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		prvkey VARCHAR(255),
		filepath VARCHAR(255),
		time DATETIME(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间'
		);`
	ModifyDB(sql)
}
//存放圆锥曲线路径密钥
func CreateEccKeyPath() {
	sql := `CREATE TABLE IF NOT EXISTS ecckeypath(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		filekeypath VARCHAR(255),
		filepath VARCHAR(255),
		time DATETIME(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间'
		);`
	ModifyDB(sql)
}
//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
package models

import (
	"bytes"
	"html/template"
	v1 "k8s.io/api/core/v1"
	"time"
)
type NamespaceParam struct {
	Id         string
	Name       string
	Status     string
	Age        time.Duration
}

func MakeNameSpaceBlocks(Items []v1.Namespace) template.HTML {
	htmlHome := ""
	for _, item := range Items {
		//将数据库model转换为首页模板所需要的model
		namespaceParam :=NamespaceParam{}
		namespaceParam.Id = string(item.UID)
		namespaceParam.Name = item.Name
		namespaceParam.Status=string(item.Status.Phase)
		now:=time.Now()
		creattime:=item.CreationTimestamp
		age :=now.Sub(creattime.Time)
		namespaceParam.Age=age
		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/page/block/namespace.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, namespaceParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

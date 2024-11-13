package utils

import (
	"WorkPro/models"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"os"

	//以太坊加密库，要求go版本升级到1.15
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

//生成椭圆曲线
func genPrivateKey() (*ecdsa.PrivateKey, error) {
	pubkeyCurve := crypto.S256() //初始化椭圆曲线
	//随机挑选基点，生成私钥
	p, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) //用golang标准库生成公私钥
	if err != nil {
		return nil, err
	} else {
		return p, nil //转换成以太坊的公私钥对
	}
}

//ECCEncrypt 椭圆曲线加密
func ECCEncrypt(bytestr []byte, prvkey *ecdsa.PrivateKey) ([]byte, error) {
	ecsprk:=ecies.ImportECDSA(prvkey)
	//椭圆曲线进行加密 公钥加密
	ct,err :=ecies.Encrypt(rand.Reader,&ecsprk.PublicKey,bytestr,nil,nil)
	if err != nil {
		log.Fatal(err)
	}
	return ct,err
}

//ECCDecrypt 椭圆曲线解密
func ECCDecrypt(bytestr []byte, prvKey *ecdsa.PrivateKey) ([]byte, error) {
	eciprv:=ecies.ImportECDSA(prvKey)
	if src, err := eciprv.Decrypt(bytestr, nil, nil); err != nil {
		return nil, err
	} else {
		return src, nil
	}
}
func BackUpInPutECCFileFile(v interface{},fileMkdir string,filePath string)  {
	//结构体转化为byte
	bytes, err := json.Marshal(v)
	if err!=nil {
		log.Fatalln(err)
	}
	//plain := string(bytes) //转化为string类型
	//椭圆曲线创造私钥
	prk,err :=genPrivateKey()
	if err!=nil {
		log.Fatalln(err)
	}
	//进行加密
	ECCtxt,err :=ECCEncrypt(bytes,prk)
	//创建文件夹
	if !isExist(fileMkdir) {
		os.MkdirAll(fileMkdir, os.ModePerm)
	}
	//写入加密文件
	err1 := ioutil.WriteFile(filePath,ECCtxt, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	models.InsertEccKey(prk,filePath)
}
func RestoreECCFileOutPutFile(fileDir string)([]interface{})  {
	var v interface{}
	filenames :=Read(fileDir)
	var list []interface{}
	for i :=0;i<len(filenames);i++ {
		bytes,err :=ioutil.ReadFile(filenames[i])
		if err!=nil {
			log.Println(err)
		}
		//进行解密
		prvkey :=models.QueryEccKey(filenames[i])
		debytes,err:=ECCDecrypt(bytes,prvkey)
		if err!=nil {
			log.Println(err)
		}
		err1 :=json.Unmarshal(debytes,&v) //获取到每一个文件的序列
		if err1 !=nil {
			log.Println(err)
		}
		list =append(list, v)
	}
	//返回list数组 如果是namespace级别文件则读取多个文件 数组大小为文件个数
	//如果为cluster级别文件则只有一组数据
	return list
}



/**


以路径存储

 */


//存储密钥
func KeyWrite(prvKey *ecdsa.PrivateKey,fileMkdir string,filePath string)  {
	prvbytes:=crypto.FromECDSA(prvKey)
	//bit存储到文件之中
	fileMkKeydir :="./key"+fileMkdir
	fileKeyPath :="./key"+filePath
	//创建文件夹
	if !isExist(fileMkKeydir) {
		os.MkdirAll(fileMkKeydir, os.ModePerm)
	}
	k:=hex.EncodeToString(prvbytes)
	//写入密钥
	err1 := ioutil.WriteFile(fileKeyPath, []byte(k), 0600)
	if err1 != nil {
		log.Fatal(err1)
	}
	//写入到文件之中 数据库中存储路径
	models.InsertEccPath(filePath,fileKeyPath)
}

//读取密钥
func KeyRead(fileKeyPath string)(*ecdsa.PrivateKey,error) {

    str :=models.QueryEccKeyPath(fileKeyPath)
	key, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	p ,err:=crypto.ToECDSA(key)
	return p,err
}
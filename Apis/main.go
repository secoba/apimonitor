package main

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/xuchengzhi/Library/Encryption"
	"github.com/xuchengzhi/Library/Http"
	"github.com/xuchengzhi/Library/Random"
	"github.com/xuchengzhi/Library/Time"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	// "os"
	// "reflect"
	"sync"
	// "net/url"
	// "github.com/Luxurioust/excelize"
	"strconv"
	// "strings"
	"time"
)

func init() {
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

var (
	wg sync.WaitGroup
)

var EmailList = [...]string{"@yahoo.com", "@yahoo.com.cn", "@yahoo.com.cn.jp", "@gmail.co.jp", "@live.com", "@hotmail.com", "@yahoo.com.jp"}

//读取配置文件
var cfg, _ = goconfig.LoadConfigFile("config.ini")

//读取接口地址
var host, _ = cfg.GetValue("Test", "Host")

//获取加密字符
func Jiami(Bstr string) string {
	Key, _ := cfg.GetValue("Xor", "key")
	token := XorEnc.XorEncodeStr(Bstr, Key)
	return token
}

//获取邮件验证码
func SendCode() {
	Run_sync.Add(1)
	params := make(map[string]string)
	params["email"] = "xuchengzhi1987@yeah.net"
	params["device_num"] = Randoms.GetRandomString(8)
	params["mail_type"] = "3" //strconv.Itoa(Randoms.GetRandomInt(1, 4))
	params["client_type"] = "app"
	params["sys"] = "ADR7.0"
	params["clientSW"] = "1.0.0"
	params["ptype"] = "ios12"
	params["t"] = GetTime.TS()
	jsonStr, err := json.Marshal(params)

	if err != nil {
		log.Fatal(err)
	}
	tmps := string(jsonStr)
	Bstr := XorEnc.BASE64EncodeStr(tmps)
	token := Jiami(Bstr)
	p := make(map[string]string)
	p["p"] = token
	log.Println(tmps)
	log.Println(Bstr)
	url := fmt.Sprintf("%v/mobile.php/Users/send_code", host)
	UrlRun.Action(url, "post", p, &Run_sync)
}

func Getpar() map[string]string {
	num := Randoms.GetRandomInt(0, len(EmailList))
	names := Randoms.GetRandomString(8)
	email := fmt.Sprintf("%v%v", names, EmailList[num])
	params := make(map[string]string)
	params["email"] = email
	params["device_num"] = Randoms.GetRandomString(18)
	params["mail_type"] = strconv.Itoa(Randoms.GetRandomInt(1, 4))
	params["client_type"] = "app"
	params["sys"] = "ADR7.0"
	params["clientSW"] = "1.0.0"
	params["ptype"] = "ios12"
	params["t"] = GetTime.TS()
	jsonStr, err := json.Marshal(params)

	if err != nil {
		log.Fatal(err)
	}
	tmps := string(jsonStr)
	Bstr := XorEnc.BASE64EncodeStr(tmps)

	token := Jiami(Bstr)
	log.Println(email)

	p := make(map[string]string)
	p["p"] = token
	return p
}

//邮箱注册
func Register() {
	Run_sync.Add(1)
	num := Randoms.GetRandomInt(0, len(EmailList))
	names := Randoms.GetRandomString(8)
	email := fmt.Sprintf("%v%v", names, EmailList[num])
	params := make(map[string]string)
	params["email"] = "TEST" + email
	params["pwd"] = "123456"
	params["client_type"] = "app"
	params["sys"] = "ADR7.0"
	params["clientSW"] = "1.0.0"
	params["ptype"] = "ios12"
	params["t"] = GetTime.TS()
	jsonStr, err := json.Marshal(params)

	if err != nil {
		log.Fatal(err)
	}
	tmps := string(jsonStr)
	Bstr := XorEnc.BASE64EncodeStr(tmps)
	token := Jiami(Bstr)
	p := make(map[string]string)
	p["p"] = token
	log.Println(params)
	url := fmt.Sprintf("%v/mobile.php/Users/reg_setpwd", host)
	UrlRun.Action(url, "post", p, &Run_sync)
}

//登录
func Login() {
	Run_sync.Add(1)
	params := make(map[string]string)
	params["email"] = "xudear@live.com"
	params["pwd"] = XorEnc.Gmd5("123456")
	params["client_type"] = "app"
	params["sys"] = "ADR7.0"
	params["clientSW"] = "1.0.0"
	params["ptype"] = "ios12"
	params["t"] = GetTime.TS()
	jsonStr, err := json.Marshal(params)

	if err != nil {
		log.Fatal(err)
	}
	tmps := string(jsonStr)
	Bstr := XorEnc.BASE64EncodeStr(tmps)
	token := Jiami(Bstr)
	p := make(map[string]string)
	p["p"] = token
	log.Println(tmps)
	url := fmt.Sprintf("%v/mobile.php/Users/email_login", host)
	UrlRun.PressureRun(1000, url, "post", p)
}

var Run_sync sync.WaitGroup

type ApiINfo struct {
	ApiPath string
	Apipar  map[string]string
}

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)
			return ctx.Err()

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// fmt.Println("Hey, I'm going to do some work")

	// wg.Add(1)
	// go work(ctx)
	// wg.Wait()

	// fmt.Println("Finished. I'm going home")
	Login()
}

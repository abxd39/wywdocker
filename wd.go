package main

import (
	"errors"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	log.Info("Hello from Seelog!")
	log.Info("Hello Docker 世界！！")
	_=Auto("Hello World !")
}

func Auto(param string)error{
	if param=="Hello World !"{
		log.Tracef("param=%s",param)
		return nil
	}else {
		err:=errors.New("字符串输入错误！！")
		_=log.Error(err)
		return err
	}
}
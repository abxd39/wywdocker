/**
 * @Author: wangyingwen
 * @Description: 
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2020-03-17 10:30
 */

package main

import (
	"errors"
	"log"
)

func WywBranch(s string)error{
	log.Println("什么情况啊")
	return errors.New(s)
}
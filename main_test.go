/**
 * @Author: wangyingwen
 * @Description: 
 * @File:  main_test
 * @Version: 1.0.0
 * @Date: 2020-03-10 14:28
 */

package main

import "testing"

func TestAuto(t *testing.T) {
	_=Auto("")
}

//代码合并丢失的问题
func TestDev(t *testing.T){
	t.Log("this is dev branch ")
}
package main

import log "github.com/cihub/seelog"

func main() {
	defer log.Flush()
	log.Info("Hello from Seelog!")
	log.Info("Hello Docker 世界！！")
	
}

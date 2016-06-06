package protocol

// +build pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"

)


func init() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
}

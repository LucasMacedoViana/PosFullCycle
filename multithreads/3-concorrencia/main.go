package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//m.Unlock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("voce e o visitante numero %d", number)))
	})

	http.ListenAndServe(":3000", nil)
}

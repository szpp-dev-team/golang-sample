package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// ルート / で行う処理を定義
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// IE ステータスで返す
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// OK
		rw.WriteHeader(http.StatusOK)

		// You said ... をレスポンス body に書き込む
		if _, err := rw.Write([]byte("You said " + string(b))); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	})

	// 8000 ポートで待受
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

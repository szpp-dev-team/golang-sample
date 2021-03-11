package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// GET メソッドで url は http://localhost:8000/ で body は "makabe mizuki"
	// ↑のようなリクエストをするための Request 構造体を作成(この時点でまだリクエストは送られていない)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/", bytes.NewReader([]byte("makabe mizuki")))
	if err != nil {
		log.Fatal(err)
	}
	// main 関数が終わるときに body を close
	defer req.Body.Close()

	// http client 構造体を作成
	cli := &http.Client{}

	// Request 構造体にしたがってリクエストを送信
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// io.ReadCloser の中身をバイト列で読み込む
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// バイト列は string() でキャストできる
	fmt.Println(string(b))
}

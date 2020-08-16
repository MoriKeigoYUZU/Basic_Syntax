package main

import (
	"fmt"
	"net/http"
)

// ServeHTTPメソッド用の構造体型
type Server struct {
}

//httpリクエストを受けるメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//HTMLの文字列

	h :=
		`
<html><head><title>Hello</title></head></html>
<body>

Hello

</body></html>
`
	//クライアント(ブラウザ)にHTMLを送信
	fmt.Fprint(w, h)
}

func main() {
	//Webサーバを起動
	http.ListenAndServe(":4000", Server{})
}

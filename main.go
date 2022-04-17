package main

import (
	"log"

	"github.com/jchv/go-webview2"
)

// go build -ldflags="-H windowsgui"
func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title: "Minimal webview example",
		},
	})
	if w == nil {
		log.Panicln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(400, 150, webview2.HintFixed)
	w.Bind("go_check_site", func(url string) string {
		log.Println(url)
		return "OK"
	})
	w.Bind("quit", func() {
		w.Terminate()
	})
	w.Navigate(`data:text/html,
    <!doctype html>
    <head>
      <meta charset="utf-8">
    </head>
    <html>
      <body style="text-align: center">
               <input type="text" style="font-size: 25px; width: 98%;" id="url" placeholder="Nhập địa chỉ website: https://">
               <br>
               <br>
               <button style="font-size: 25px; width: 50%;" onclick="check_site()">Kiểm tra</button>
      </body>
      <script>
        function check_site(){
          go_check_site(url.value).then(r => {
            console.log(r)
          }).catch(e => {
            console.error(e)
          })
        }
      </script>
    </html>
  `)
	w.Run()
}

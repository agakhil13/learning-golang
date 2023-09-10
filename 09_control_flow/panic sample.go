package main
import (
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte(`
		<h1>Hello Go!</h1>
		<p>This is a page test</p>
		`))
		// fmt.Println(r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

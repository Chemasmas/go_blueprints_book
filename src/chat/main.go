package main

import (
	"log"
	"net/http"
	"text/template"
	"sync"
	"path/filepath"
	"flag"
	//"os"
	//"trace" // Es local se modifico el GOPATH
)

type templateHandler struct {
	once sync.Once
	filename string
	temp1 *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	t.once.Do(func(){
			t.temp1 = template.Must(template.ParseFiles(filepath.Join("template",t.filename) ) )
	})
	t.temp1.Execute(w,r)
}


func main() {

	var addr = flag.String("addr",":8080","The addr of the app")

	flag.Parse()

	r:= newRoom()
	//r.tracer = trace.New(os.Stdout) 

	http.Handle("/chat", MustAuth(&templateHandler{filename:"chat.html"} ) )

	http.Handle("/room",r)

	go r.run()

	log.Println("Servidor Iniciado en  puerto ",*addr)
	//Start Web Server
	if err := http.ListenAndServe(*addr,nil); err != nil {
		log.Fatal("Se Quemo!!!",err)
	}
	
}
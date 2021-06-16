package routes

import (
	"fmt"
	"html/template"
	"net/http"
)


var (
	filestatics = http.FileServer(http.Dir("style/css"))

)

func MyRoutes(){
	http.HandleFunc("/", Index)
	http.HandleFunc("/so", SoPage)
	http.HandleFunc("/formvm", FormVm)

	http.Handle("/style/css/", http.StripPrefix("/style/css/", filestatics))

	if err := http.ListenAndServe(":8082", nil); err != nil{
		fmt.Println(err)
	}
}

var tmplt = template.Must(template.ParseGlob("./templates/*html"))

func Index(w http.ResponseWriter, r *http.Request){
	tmplt.ExecuteTemplate(w, "index", nil)
}

func SoPage(w http.ResponseWriter, r *http.Request){
	tmplt.ExecuteTemplate(w, "so", nil)
}

func FormVm(w http.ResponseWriter, r *http.Request){
	tmplt.ExecuteTemplate(w, "formvm", nil)
}

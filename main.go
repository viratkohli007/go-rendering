package main
import (
		"fmt"
		"net/http"
		"html/template"
		"log"
		"io/ioutil"
		"encoding/json"
		"os"
	   )

type Data struct{
  Name string `json:"name"`
  Age string `json:"age"`
  Email string `json:"email"`
  BloodGroup string `json:"bloodgroup"`
}

type welcomest struct{
	Title string
}

type Formst struct{}

func main() {
	const Port =":8080"
	http.HandleFunc("/", welcome)
	http.HandleFunc("/form", Form)
	http.HandleFunc("/display", Display)
	http.HandleFunc("/list", List)
	err := http.ListenAndServe(Port, nil)
	if err != nil{
        log.Fatal(err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
  title := &welcomest{Title: "Apple"}
  t, _ := template.ParseFiles("welcome.html")
  t.Execute(w, title)
  //fmt.Println(t)
}

func Form(w http.ResponseWriter, r *http.Request){

      //var f Formst
      t,_ := template.ParseFiles("form.html")
      t.Execute(w, "")
}

func Display(w http.ResponseWriter, r *http.Request) {

   //var obj Data

   data, err := ioutil.ReadFile("data.json")
   if err != nil{
   	fmt.Println(err)
   }
   //fmt.Println(data)

   var Dataobj1 []Data
    err2 := json.Unmarshal(data ,&Dataobj1)
    if err2 != nil{
    	log.Fatal(err2)
    }
   //fmt.Println(Dataobj1)

    d := Data{
    	Name: r.FormValue("name"),
    	Age: r.FormValue("age"),
    	Email: r.FormValue("email"),
    	BloodGroup: r.FormValue("bloodgroup"),
    }
    //fmt.Println(d)
    Dataobj1 = append(Dataobj1, d)
    fmt.Println(Dataobj1)

    d3, _ := json.MarshalIndent(Dataobj1, "", " ")
    err3 := ioutil.WriteFile("data.json", []byte(d3), 0644)
    if err3 != nil{
    	log.Fatal(err)
    }

    f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err!= nil{
    	log.Fatal(err)
    }
    defer f.Close()
    file := new(Data)
    file.Name = r.FormValue("name")
    file.Age = r.FormValue("age")
    file.BloodGroup = r.FormValue("bloodgroup")
    file.Email = r.FormValue("email")

     t, _ := template.ParseFiles("display.html")
	   t.Execute(w, file)
}
func List(w http.ResponseWriter, r *http.Request) {
      File, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
      if err != nil{
      	log.Fatal(err)
      }
      //fmt.Println(File)
      defer File.Close()
      byteval, _ := ioutil.ReadAll(File)
       var Res []Data
        json.Unmarshal(byteval, &Res)
       fmt.Println(Res)
       t, _ := template.ParseFiles("list.html")
       t.Execute(w, Res)
}

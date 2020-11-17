package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)


type Page struct {
	Title string
    	Allitems [5]string
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/"	{
	http.Redirect(w, r, "/login", http.StatusFound)
	}
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		http.Redirect(w, r, "/view", http.StatusFound)
	}
	
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("register.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		http.Redirect(w, r, "/view", http.StatusFound)
	}
	
}



func view(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view.html")
		var Listing Page
		Listing.Allitems = [5]string{"number1", "number2","number3","nmber4","number555"}
		Listing.Title = "dingo"
		fmt.Println("files:", Listing.Allitems)
		t.Execute(w, Listing)
	}
	
}



func get() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT item_id,title,creator_name,description,path FROM itembank")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    for results.Next() {
        var one_item Item
        // for each row, scan the result into our tag composite object
        err = results.Scan(&one_item.ID, &one_item.item_title,&one_item.creator_Name,&one_item.item_path,&one_item.item_description)
        fmt.Printf(strconv.Itoa(one_item.ID))
        if err != nil {
            
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Printf(one_item.item_path)
    }
    }



func put() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("insert itembank(item_id,title,creator_name,description,path) values(1,'testing.png','me','thisisathing','/folderfolder2/image.png')")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    for results.Next() {
        var one_item Item
        // for each row, scan the result into our tag composite object
        err = results.Scan(&one_item.ID, &one_item.item_title,&one_item.creator_Name,&one_item.item_path,&one_item.item_description)
        fmt.Printf(strconv.Itoa(one_item.ID))
        if err != nil {
            
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Printf(one_item.item_path)
    }


func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/view", view)
	http.HandleFunc("/result/", result)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

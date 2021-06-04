package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"database/sql"
	"strconv"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"encoding/gob"
	"github.com/gorilla/sessions"
	"os"
	"io"
	 "golang.org/x/crypto/bcrypt"
)


type Page struct {
	Title string
    Allitems []string
}

type User struct {
	Username      string
	Password      string
	Authenticated bool
}


type Item struct {
    ID   int    
    Item_title string
    Creator_Name string 
    Item_path string
    Item_Description string
	Creation string
	Lastviewed string
	Groupnamecol string
}

type Registration struct {
	Username string
	Password string
}


type bigGroup struct{
Allgroups []string
}


type usergroup struct{
	group1 string
	group2 string
	group3 string
	group4 string
	group5 string
}











func sayhelloName(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm() 
	/* 
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])*/
	if(r.URL.Path=="/"){
	http.Redirect(w, r, "/login", http.StatusFound)
	}
	/*for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}*/
	
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	session, err := store.Get(r, "cookie-name")
	defer db.Close()
	var login Registration
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["user"] = User{}
		session.Options.MaxAge = -1

		err = session.Save(r, w)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		insertstring:="select username, password from Userinfo where username= ?"
		results := db.QueryRow(insertstring,r.FormValue("username"))
		err:=results.Scan(&login.Username,&login.Password)
		if(err!=sql.ErrNoRows)	{
			
			if(CheckPasswordHash(r.FormValue("password"),login.Password)){
		username := r.FormValue("username")
			
			user := &User{
				Username:      username,
				Authenticated: true,
			}

			session.Values["user"] = user

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/view", http.StatusFound)
		} else {
			http.Redirect(w, r, "/loginfailed", http.StatusFound)
		}
		fmt.Println(results)
	}
	http.Redirect(w, r, "/loginfailed", http.StatusFound)
	}
	
}

func register(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
	defer db.Close()
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		insertstring:="select groupname from groupsinfo"
		allgroups := []string{}
		results, err := db.Query(insertstring)
		if err != nil {
			fmt.Println(err.Error()) // proper error handling instead of panic in your app
		}else{
	
		var tempgroup string
		for results.Next() {
        
			// for each row, scan the result into our tag composite object
			err = results.Scan(&tempgroup)
			if err != nil {
            
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			allgroups=append(allgroups,tempgroup)
					// and then print out the tag's Name attribute
        
			}
		}
		var groupvar bigGroup
		groupvar.Allgroups=allgroups

			t, _ := template.ParseFiles("register.html")
			t.Execute(w, groupvar)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		
		


		if(len(r.FormValue("username"))>0&&len(r.FormValue("password"))>0&&len(r.FormValue("realname"))>0 && len(r.FormValue("password"))<21) {
		hashedpass, err:=HashPassword(r.FormValue("password"))
		if(err!=nil){
			http.Redirect(w, r, "/registrationfailed", http.StatusFound)
			return
		}
		insertstring:="insert userInfo(username,password,Useractualname,Group1,Group2,Group3,Group4,Group5) values(?, ?, ?, ?, ?, ?, ?, ?)"
		results, err := db.Exec(insertstring,r.FormValue("username"),hashedpass,r.FormValue("realname"),NewNullString(r.FormValue("group1")),NewNullString(r.FormValue("group2")),NewNullString(r.FormValue("group3")),NewNullString(r.FormValue("group4")),NewNullString(r.FormValue("group5")))
		if err != nil {
			http.Redirect(w, r, "/registrationfailed", http.StatusFound)
		
			fmt.Println(err.Error()) // proper error handling instead of panic in your app
		}	else {
			
			_, _ = db.Exec("insert into groupsinfo(groupname,approval) values(?, 0)",NewNullString(r.FormValue("group1")))
			_, _ = db.Exec("insert into groupsinfo(groupname,approval) values(?, 0)",NewNullString(r.FormValue("group2")))
			_, _ = db.Exec("insert into groupsinfo(groupname,approval) values(?, 0)",NewNullString(r.FormValue("group3")))
			_, _ = db.Exec("insert into groupsinfo(groupname,approval) values(?, 0)",NewNullString(r.FormValue("group4")))
			_, _ = db.Exec("insert into groupsinfo(groupname,approval) values(?, 0)",NewNullString(r.FormValue("group5")))
			http.Redirect(w, r, "/login", http.StatusFound)
			}
			fmt.Println(results)
		}else{
			http.Redirect(w, r, "/registrationfailed", http.StatusFound)
		}
		}
	
}





func view(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUser(session)
	
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view3.html")
		var Listing Page
		Listing.Allitems = getview(user.Username,w,r)
		Listing.Title = "dingo"
		fmt.Println("files:", Listing.Allitems)
		t.Execute(w, Listing)
		//fmt.Println("finished view")
	}
	
}

func result(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	defer db.Close()
	session, err := store.Get(r, "cookie-name")
	fmt.Println("result method:", r.Method) //get request method
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUser(session)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("results.html")
		var tempe Item
		subbb := strings.Split(r.URL.Path,"/")[2]

		tempe = getresult(subbb,user.Username)
		fmt.Println(tempe)
		
		t.Execute(w, tempe)
	} else {
		r.ParseForm()
		fmt.Println(r.FormValue("Action"))
		if(r.FormValue("Action")=="DELETE"){
			remove2(r.FormValue("ID"),user.Username)
		}else if(r.FormValue("Action")=="DOWNLOAD"){
			fmt.Println("place")
			mynumber,errorr := strconv.Atoi(r.FormValue("ID"))
			if errorr != nil {
			    fmt.Println(errorr)	
			}else{
			var tempgroups usergroup
			results,err := db.Query("select Group1,Group2,Group3,Group4,Group5 from userinfo where username=?",user.Username)
			if err != nil {
				fmt.Print(err.Error())
			}else{
				results.Next()
				results.Scan(&tempgroups.group1,&tempgroups.group2,&tempgroups.group3,&tempgroups.group4,&tempgroups.group5)
				fmt.Println(tempgroups.group1,tempgroups.group2,tempgroups.group3,tempgroups.group4,tempgroups.group5)
			}
    
			results, err = db.Query("SELECT path FROM itembank where item_id=? and (creator_name=? or filegroup IN(?,?,?,?,?))",mynumber,user.Username,tempgroups.group1,tempgroups.group2,tempgroups.group3,tempgroups.group4,tempgroups.group5)
			if err != nil {
				fmt.Println(err.Error()) // proper error handling instead of panic in your app
			}
			if(results!=nil){
			results.Next()
			var filenameholder string
			results.Scan(&filenameholder)

			handledownload(w,r,filenameholder)
			}
			}
		}
		http.Redirect(w, r, "/view", http.StatusFound)
	}
	
}

func handledownload(w http.ResponseWriter, r *http.Request, Filename string){
	Openfile, err := os.Open(Filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		fmt.Println( Filename)
		return
	}
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}


func registrationfailed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("registrationfailed.html")
		t.Execute(w, nil)
	}
	
}
func loginfailed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("loginfailed.html")
		t.Execute(w, nil)
	}
	
}



func getview(user string,w http.ResponseWriter, r *http.Request) []string{
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
	var tempgroups usergroup
    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()
	results,err := db.Query("select Group1,Group2,Group3,Group4,Group5 from userinfo where username=?",user)
    if err != nil {
        fmt.Print(err.Error())
		return []string{}
    }else{
		results.Next()
		results.Scan(&tempgroups.group1,&tempgroups.group2,&tempgroups.group3,&tempgroups.group4,&tempgroups.group5)
		fmt.Println(tempgroups.group1,tempgroups.group2,tempgroups.group3,tempgroups.group4,tempgroups.group5)
	}
	// Execute the query
    results, err = db.Query("SELECT title FROM itembank where creator_name = ? OR filegroup IN(?,?,?,?,?)",user,tempgroups.group1,tempgroups.group2,tempgroups.group3,tempgroups.group4,tempgroups.group5)
    if err != nil {
        print("the error was here") // proper error handling instead of panic in your app
    }
	resultinglines := []string{}
    for results.Next() {
        var temp string
        // for each row, scan the result into our tag composite object
        err = results.Scan(&temp)
        if err != nil {
            
            panic(err.Error()) // proper error handling instead of panic in your app
        }
		resultinglines = append(resultinglines,temp)
                // and then print out the tag's Name attribute
        //fmt.Println("loop")
    }
	return resultinglines
}

func getresult(requestedfile string, user string) Item{
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
	var one_item Item
	
    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()
	var tempgroups usergroup
	results,err := db.Query("select Group1,Group2,Group3,Group4,Group5 from userinfo where username=?",user)
    if err != nil {
        fmt.Print(err.Error())
    }else{
		results.Next()
		results.Scan(&tempgroups.group1,&tempgroups.group2,&tempgroups.group3,&tempgroups.group4,&tempgroups.group5)
		fmt.Println(tempgroups.group1,tempgroups.group2,tempgroups.group3,tempgroups.group4,tempgroups.group5)
	}
    // Execute the query
    results, err = db.Query("SELECT * FROM itembank where title =? and (creator_name =? or filegroup IN(?,?,?,?,?))",requestedfile,user,tempgroups.group1,tempgroups.group2,tempgroups.group3,tempgroups.group4,tempgroups.group5)
    if err != nil {
        fmt.Println(err.Error()) // proper error handling instead of panic in your app
    }else{
		_, _ = db.Query("UPDATE itembank set last_access=CURDATE() where title=?",requestedfile)
    for results.Next() {
        
        // for each row, scan the result into our tag composite object
        err = results.Scan(&one_item.ID, &one_item.Item_title,&one_item.Creation,&one_item.Creator_Name,&one_item.Lastviewed,&one_item.Item_path,&one_item.Item_Description,&one_item.Groupnamecol)
        if err != nil {
            
            fmt.Println(err.Error())
			// proper error handling instead of panic in your app
        }
		
                // and then print out the tag's Name attribute
        
    }
	}
	return one_item
}



func remove2(id string, user string) {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
	

	
    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()
	results2, err2:=db.Query("SELECT path FROM itembank where item_id=?",id)
		// logic part of log in
		fmt.Println("IDtest:", id,"name",user)
	if err2 != nil {
        fmt.Println(err2.Error()) // proper error handling instead of panic in your app
		return
    }
	if(results2==nil){
		return
	}
    // Execute the query
	results, err := db.Exec("delete FROM itembank where item_id=? and creator_name=?",id,user)
    if err != nil {
        fmt.Println(err.Error()) // proper error handling instead of panic in your app
    }
	affrows,err := results.RowsAffected()
	if err != nil {
        fmt.Println(err.Error()) // proper error handling instead of panic in your app
    }
	if(affrows==0){
		return
	} else{
		results2.Next()
		var tempstring string 
		results2.Scan(&tempstring)

		os.Remove(tempstring)
	}
	
}
















func uploadfile(w http.ResponseWriter, r *http.Request, user string) {
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
		return
    }
    defer db.Close()
	r.ParseMultipartForm(5 << 20)
    // Execute the query	
	fmt.Print(user)
	_, err = os.Stat("uploadedfiles/"+r.FormValue("filename"))

    if (err==nil) {
        return	
    }
	var length = strings.Split(r.FormValue("filename"),".")
	var extension =length[len(length)-1]
	if(!strings.ContainsAny(r.FormValue("filename"),"/\\:*?\"<>|")&&(len(r.FormValue("filename"))!=0 && len(length[0])!=0 && len(user)!=0 && len(length)>1 )&& (extension=="txt"||extension=="png"||extension=="jpg"||extension=="mp4")){
		results, err := db.Query("insert into itembank(title,creator_name,description,path,creation,last_access,filegroup) values(?,?,?,?,CURDATE(),CURDATE(),?);",r.FormValue("filename"),user,r.FormValue("description"),"uploadedfiles/"+r.FormValue("filename"),r.FormValue("filegroup"))
    
	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	uploadactual(w,r)
    for results.Next() {
        
        // for each row, scan the result into our tag composite object
        fmt.Println(results)
        
		
                // and then print out the tag's Name attribute
        
    }
	}
	
}




func uploadactual(w http.ResponseWriter, r *http.Request) {
    
    r.ParseMultipartForm(5 << 20)
    file, handler, err := r.FormFile("uploadedfile")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Println(handler.Filename)
    newFile, err := os.Create("uploadedfiles/"+r.FormValue("filename"))
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	newFile.Write(fileBytes)

}



func upload(w http.ResponseWriter, r *http.Request) {
	
	
	fmt.Println("upload method:", r.Method) //get request method
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	user := getUser(session)
	if r.Method == "GET" {
		db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")
		if err != nil {
        fmt.Print(err.Error())
		return
    }
    defer db.Close()
		insertstring:="select groupname from groupsinfo"
		allgroups := []string{}
		results, err := db.Query(insertstring)
		if err != nil {
			fmt.Println(err.Error()) // proper error handling instead of panic in your app
		}else{
	
		var tempgroup string
		for results.Next() {
        
			// for each row, scan the result into our tag composite object
			err = results.Scan(&tempgroup)
			if err != nil {
            
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			allgroups=append(allgroups,tempgroup)
					// and then print out the tag's Name attribute
        
			}
		}
		var groupvar bigGroup
		groupvar.Allgroups=allgroups
		t, _ := template.ParseFiles("qqq.html")
		t.Execute(w, groupvar)
	} else {
		
		uploadfile(w,r,user.Username)
		http.Redirect(w, r, "/view", http.StatusFound)
	}
	
}










func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}










func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}



















// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}

func main() {
	
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/result/", result)
	http.HandleFunc("/loginfailed", loginfailed)
	http.HandleFunc("/register", register)
	http.HandleFunc("/registrationfailed", registrationfailed)
	http.HandleFunc("/view", view)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/logout", logout)
	
	err := http.ListenAndServe(":80", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}




// store will hold all session data
var store *sessions.CookieStore



func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})

}












func NewNullString(s string) sql.NullString {
    if len(s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
         String: s,
         Valid: true,
    }
}
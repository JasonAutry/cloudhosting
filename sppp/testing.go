package main



//code from here was used as a template https://tutorialedge.net/golang/golang-mysql-tutorial/
//
//currently most things are hard coded because this was the last version to fully compile
//currently implementing  this with a webapp that responds with html pages containing the requested data
//
//webapp portion is being adapted with the result of this tutorial as the base https://golang.org/doc/articles/wiki/ 
//
//
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
)
//item_id,title,creator_name,description,path

type Item struct {
    ID   int    
    item_title string
    creator_Name string 
    item_path string
    item_description string
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

}
    func view() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT item_id,title,creator_name,description,path FROM itembank where creator_name='root'")
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
   func remove() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/tasksdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("delete FROM itembank where item_id=1")
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

   func main(){
       remove()
       put()
       get()
       view()
   }
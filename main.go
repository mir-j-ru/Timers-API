package main

import(
    "database/sql"
    "fmt"
    "log"
    "net/http"      // ✅ Правильно
    "os"

    "ecosystem_mir_j/db"
    "ecosystem_mir_j/handlers"

    "github.com/joho/godotenv"
)

var dbConn *sql.DB

func main(){
    //загружаем окружение
    err := godotenv.Load()
    if err != nil{
        log.Fatal("error load godotenv: ", err)
    }

    dbConn := db.InitDB()
    defer dbConn.Close()


    http.HandleFunc("/add_time", func(w http.ResponseWriter, r *http.Request){
        handlers.AddTime(w, r, dbConn)
    })//добавление комплектации

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Printf("Сервер запущен на порту %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

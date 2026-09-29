package db

import(
    "database/sql"
    "log"
    "os"
    "time"
    "fmt"

    _"github.com/lib/pq"
    "ecosystem_mir_j/models"
)

func AddTime(db *sql.DB, time string, status string) error {
    _,err := db.Exec(`INSERT INTO timers (time, status) VALUES ($1,$2)`, time,status)
    if err != nil{
        log.Fatal("error add time: ", err)
    }
    
    return err
}



func InitDB() *sql.DB{
    //подключаемся к БД
    connstr := os.Getenv("DATABASE_URL")
    if connstr == ""{
        connstr = "postgresql://postgres:36863686@localhost:5432/ecosystem_mir_j?sslmode=disable"
    }

    db, err := sql.Open("postgres", connstr)
    if err != nil{
        log.Fatal("error db open: ", err)
    }

    createTableSQLClients := `
    CREATE TABLE IF NOT EXISTS timers (
        id SERIAL PRIMARY KEY,
        time TEXT NOT NULL,
        status TEXT NOT NULL DEFAULT 'default',
        created_at TIMESTAMP DEFAULT NOW()
    );`

    _, err = db.Exec(createTableSQLClients)
    if err != nil{
        log.Fatal("error create clients_list table: ", err)
    }    

    return db
}

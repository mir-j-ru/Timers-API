package models

import(
 "time"
)

type ClientResponse struct{
    ID int `json: "ID"`
    Name string `json: "name"`
    Email string `json: "email"`
    Phone string `json: "phone"`
    CreatedAt time.Time `json: "created_at"`
}

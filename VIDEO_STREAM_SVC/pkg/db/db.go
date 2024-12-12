package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/config"
	"github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/domain"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(cfg *config.Config) (*gorm.DB, error) {

   connString:=fmt.Sprintf("host=%s user=%s password=%s ",cfg.Db_host,cfg.Db_username,cfg.Db_password)

  db,err:=sql.Open("postgres",connString) 
  if err!=nil{
    fmt.Println("creating new Database")
    return &gorm.DB{},err
  }
rows,err:=db.Query("SELECT 1 FROM pg_database WHERE datname = $1",cfg.Db_name)
  if err!=nil{
    fmt.Println("Error checking if database exists")
    return &gorm.DB{},err
  } 

  if rows.Next() {
        rows.Close()
  }else{
    _,err:=db.Exec("CREATE DATABASE "+cfg.Db_name)
    if err!=nil{
    fmt.Println("Error creating"+cfg.Db_name)
      return &gorm.DB{},err
    }
        fmt.Println(cfg.Db_name+" created")
  }

   fmt.Println("config in connection DB",cfg) 


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Db_host, cfg.Db_username, cfg.Db_password, cfg.Db_name, cfg.Db_port)
	DB, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	DB.AutoMigrate(&domain.Video{})

	return DB, dbErr
}

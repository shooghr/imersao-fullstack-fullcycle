package main

import (
	"github.com/shooghr/imersao-fullstack-fullcycle/codepix-go/application/grpc"
	"github.com/shooghr/imersao-fullstack-fullcycle/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv(key: "env"))
	grpc.StartGrpcServer(database, port: 50051)
}
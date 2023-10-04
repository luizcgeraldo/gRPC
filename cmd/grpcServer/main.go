package main

import database (
	"database/sql"

	"github.com/luizcgeraldo/gRPC/internal/database"
)


func main() {
	db, err :=  sql.Open("sqlite3", "./db.sqlite3")	
	if err != nil {
		panic(err)
	}
	defer db.close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
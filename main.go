package main

import (
	"context"
	"fmt"

	"go-service/internal/app"
)

func main() {
	var cfg app.Config
	cfg.Mongo.Uri = "mongodb+srv://dbUser:Demoaccount1@projectdemo.g0lah.mongodb.net"
	cfg.Mongo.Database = "masterdata"

	ctx := context.Background()
	fmt.Println("Export file")
	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		fmt.Println(ctx, "Error when initialize: "+err.Error())
		panic(err)
	}
	total, err := app.Export(ctx)
	if err != nil {
		fmt.Println("Error when export: " + err.Error())
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Exported file with %d records", total))
}

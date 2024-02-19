package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/lysithea/internal/config"
	handler "github.com/rulanugrh/lysithea/internal/http"
	"github.com/rulanugrh/lysithea/internal/repository"
	"github.com/rulanugrh/lysithea/internal/route"
	"github.com/rulanugrh/lysithea/internal/service"
	"github.com/rulanugrh/lysithea/internal/util"
	"gorm.io/gorm"
)

func serve(db *gorm.DB, conf *config.App) {
	app := mux.NewRouter().StrictSlash(true)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	route.UserRouter(app, userHandler)
	route.ProductRoute(app, productHandler)
	route.OrderRouter(app, orderHandler)

	host := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	server := http.Server{
		Addr:    host,
		Handler: app,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println("HTTP Failed Serving")
	}

	log.Printf("HTTP Success Running")
	log.Printf("Running at http://%s:%s", conf.Server.Host, conf.Server.Port)
}

func help() {
	helpContent := [][]string{
		{"help", "show help message"},
		{"migrate", "command for running migration model"},
		{"seeder", "command for seeder to db"},
		{"serve", "command for serve http API"},
	}

	maxLen := len(helpContent[0][0])
	for _, part := range helpContent {
		length := len(part[0])
		if length > maxLen {
			maxLen = length
		}
	}

	var builder strings.Builder
	const space = 4
	for _, part := range helpContent {
		builder.WriteString(part[0])
		spacer := (maxLen - len(part[0])) + space
		for spacer > 0 {
			builder.WriteByte(' ')
			spacer--
		}
		builder.WriteString(part[1])
		builder.WriteByte('\n')
	}

	println(builder.String()[:builder.Len()-1])
}

func migrate(db *gorm.DB) {
	err := util.Migrate(db)
	if err != nil {
		log.Printf("Database Failed Migration")
	}

	log.Printf("Database Success Migration")

}

func seeder(db *gorm.DB) {
	err := util.Seeder(db)
	if err != nil {
		log.Printf("Failed Seeder")
	}

	log.Printf("Seeder Finished")
}

func main() {
	db, err := config.NewConnection()
	if err != nil {
		log.Printf("error to connect database %v", err)
	}

	conf := config.GetConfig()

	args := os.Args[1]

	switch args {
	case "migrate":
		migrate(db)
	case "seeder":
		seeder(db)
	case "help":
		help()
	case "serve":
		serve(db, conf)
	default:
		println("Use args help to show message")
	}
}

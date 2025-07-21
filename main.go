package main

import (
	"github.com/go-hex/RookieJoel/adapters"
	"github.com/go-hex/RookieJoel/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {

	app := fiber.New();

	// Initialize GORM with SQLite
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&core.Order{}); err != nil {
		panic("failed to migrate database")
	}

	// Create the repository and service
	repo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(repo) 
	// Create the HTTP adapter
	httpAdapter := adapters.NewHttpOrderService(orderService)

	// Define the route for creating an order
	app.Post("/orders", httpAdapter.CreateOrder)

	app.Listen(":8080")
}

package main

import (
	"fmt"
)

// Define structs for user, order, and order item
type User struct {
	ID       int
	Name     string
	Location string
}

type OrderItem struct {
	Name  string
	Price float64
}

type Order struct {
	ID          int
	UserID      int
	OrderItems  []OrderItem
	TotalAmount float64
}

// Function to create a new user
func createUser(id int, name, location string) User {
	return User{ID: id, Name: name, Location: location}
}

// Function to create a new order
func createOrder(id, userID int, items []OrderItem) Order {
	var totalAmount float64
	for _, item := range items {
		totalAmount += item.Price
	}
	return Order{ID: id, UserID: userID, OrderItems: items, TotalAmount: totalAmount}
}

func main() {
	// Simulate user input
	userID := 1
	userName := "John"
	userLocation := "123 Main St, City"

	// Create a new user
	user := createUser(userID, userName, userLocation)

	// Simulate order input
	orderID := 1001
	orderItems := []OrderItem{
		{"Pizza", 10.99},
		{"Burger", 8.49},
		// Add more items as needed
	}

	// Create a new order
	order := createOrder(orderID, user.ID, orderItems)

	// Print user and order details
	fmt.Printf("User: %s\nLocation: %s\n\nOrder ID: %d\nOrder Items:\n", user.Name, user.Location, order.ID)
	for _, item := range order.OrderItems {
		fmt.Printf("- %s: $%.2f\n", item.Name, item.Price)
	}
	fmt.Printf("\nTotal Amount: $%.2f\n", order.TotalAmount)
}

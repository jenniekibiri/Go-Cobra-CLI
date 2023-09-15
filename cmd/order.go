package cmd

import (
	"flourish-coding-challenge/internal/db"
	"flourish-coding-challenge/internal/models"
	"flourish-coding-challenge/pkg/validation"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var CreateOrderCmd = &cobra.Command{
	Use:   "create-order",
	Short: "Create an order",
	Run: func(cmd *cobra.Command, args []string) {
		var email string
		var qty_ordered int
		var total_price float64

		// Loop until a non-empty and valid email is provided
		for {
			fmt.Print("Enter customer email: ")
			fmt.Scanln(&email)

			if !validation.IsEmpty(email) && validation.ValidateEmail(email) {
				break
			}
			fmt.Println("Invalid email address or email cannot be empty. Please try again.")
		}

		// Retrieve the customer by email.
		var customer models.Customer
		if err := db.DB.First(&customer, "email = ?", email).Error; err != nil {
			fmt.Println("Failed to get customer:", err)
			return
		}


	
		for {
			fmt.Print("Enter quantity ordered: ")
			fmt.Scanln(&qty_ordered)

			if !validation.IsEmptyInt(qty_ordered) {
				break
			}
			fmt.Println("Quantity cannot be empty. Please try again.")
		}

		for {
			fmt.Print("Enter total price: ")
			fmt.Scanln(&total_price)

			if !validation.IsEmptyFloat(total_price) {
				break
			}
			fmt.Println("Total price cannot be empty. Please try again.")
		}

	// Create a new order
		order := models.Order{
			CustomerID: customer.ID,
			Date:       time.Now().Format("2006-01-02"),
			QtyOrdered: qty_ordered,
			TotalPrice: total_price,
		}
		if err := db.DB.Create(&order).Error; err != nil {
			fmt.Println("Failed to create order:", err)
			return
		}

		fmt.Println("Order created successfully.")
	},
}

var ListOrdersCmd = &cobra.Command{
	Use:   "list-orders",
	Short: "List all orders",
	Run: func(cmd *cobra.Command, args []string) {
		var orders []models.Order
		if err := db.DB.Find(&orders).Error; err != nil {
			fmt.Println("Failed to orders:", err)
			return
		}
		if len(orders) == 0 {
			fmt.Println("No orders found.")
			return
		}

		fmt.Println("Orders:")
		for _, order := range orders {
			var customer models.Customer
			if err := db.DB.First(&customer, "id = ?", order.CustomerID).Error; err != nil {
				fmt.Println("Failed to get customer:", err)
				return
			}

			fmt.Println("Customer email:", customer.Email)
			fmt.Println("Date:", order.Date)
			fmt.Println("Quantity Ordered:", order.QtyOrdered)
			fmt.Println("Total Price:", order.TotalPrice)
		}
	},
}

var GetOrderCmd = &cobra.Command{
	Use:   "get-order",
	Short: "Get an order by id",
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		for {
			fmt.Print("Enter order id: ")
			fmt.Scanln(&id)
			if !validation.IsEmptyInt(id) {
				break
			}
			fmt.Println("Invalid id or id cannot be empty. Please try again.")
		}
		var order models.Order
		if err := db.DB.First(&order, "id = ?", id).Error; err != nil {
			fmt.Println("Failed to get order:", err)
			return
		}

		var customer models.Customer
		if err := db.DB.First(&customer, "id = ?", order.CustomerID).Error; err != nil {
			fmt.Println("Failed to get customer:", err)
			return
		}

		fmt.Println("Order:")
		fmt.Println("Customer email:", customer.Email)
		fmt.Println("Date:", order.Date)
		fmt.Println("Quantity Ordered:", order.QtyOrdered)
		fmt.Println("Total Price:", order.TotalPrice)

	},
}
var UpdateOrderCmd = &cobra.Command{
	Use:   "update-order",
	Short: "update order",
	Run: func(cmd *cobra.Command, args []string) {
		var email string
		for {
			fmt.Print("Enter customer email: ")
			fmt.Scanln(&email)
			if !validation.IsEmpty(email) && validation.ValidateEmail(email) {
				break
			}
			fmt.Println("Invalid email address or email cannot be empty. Please try again.")
		}

		// Retrieve the customer by email.
		var customer models.Customer
		if err := db.DB.First(&customer, "email = ?", email).Error; err != nil {
			fmt.Println("Failed to get customer:", err)
			return
		}

		// Retrieve orders associated with the customer.
		var orders []models.Order
		if err := db.DB.Find(&orders, "customer_id = ?", customer.ID).Error; err != nil {
			fmt.Println("Failed to get orders:", err)
			return
		}

		// Display a list of orders for the customer.
		fmt.Println("Orders for customer:", customer.Name)
		for i, order := range orders {
			fmt.Printf("%d. Order ID: %d\n", i+1, order.ID)
			fmt.Printf("   Date: %s\n", order.Date)
			fmt.Printf("   Quantity Ordered: %d\n", order.QtyOrdered)
			fmt.Printf("   Total Price: %.2f\n", order.TotalPrice)

		}

		// Prompt the user to select an order to update.
		var selectedOrderIndex int
		fmt.Print("Select an order to update (enter the number): ")
		fmt.Scanln(&selectedOrderIndex)

		// Check if the selected index is valid.
		if selectedOrderIndex < 1 || selectedOrderIndex > len(orders) {
			fmt.Println("Invalid selection.")
			return
		}

		// Prompt the user to enter the new quantity ordered.
		var newQtyOrdered int
		var newTotalPrice float64
		for {
			fmt.Print("Enter new quantity ordered: ")
			fmt.Scanln(&newQtyOrdered)

			if !validation.IsEmptyInt(newQtyOrdered) {
				break
			}
			fmt.Println("Quantity cannot be empty. Please try again.")
		}

		for {
			fmt.Print("Enter new total price: ")
			fmt.Scanln(&newTotalPrice)

			if !validation.IsEmptyFloat(newTotalPrice) {
				break
			}
			fmt.Println("Total price cannot be empty. Please try again.")
		}

		// Update the order.
		order := orders[selectedOrderIndex-1]
		order.QtyOrdered = newQtyOrdered
		order.TotalPrice = newTotalPrice
		if err := db.DB.Save(&order).Error; err != nil {
			fmt.Println("Failed to update order:", err)
			return
		}

		fmt.Println("Order updated successfully.")

	},
}

var DeleteOrderCmd = &cobra.Command{
	Use:   "delete-order",
	Short: "delete an order",
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		for {
			fmt.Print("Enter order id: ")
			fmt.Scanln(&id)
			if !validation.IsEmptyInt(id) {
				break
			}
			fmt.Println("Invalid id or id cannot be empty. Please try again.")
		}
		var order models.Order
		if err := db.DB.First(&order, "id = ?", id).Error; err != nil {
			fmt.Println("Failed to get order:", err)
			return
		}
		if err := db.DB.Delete(&order).Error; err != nil {
			fmt.Println("Failed to delete order:", err)
			return
		}
		fmt.Println("Order deleted successfully.")
	},
}

func init() {

	rootCmd.AddCommand(CreateOrderCmd)
	rootCmd.AddCommand(ListOrdersCmd)
	rootCmd.AddCommand(GetOrderCmd)
	rootCmd.AddCommand(UpdateOrderCmd)
	rootCmd.AddCommand(DeleteOrderCmd)

}

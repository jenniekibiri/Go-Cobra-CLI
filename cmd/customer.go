package cmd

import (
	"flourish-coding-challenge/internal/db"
	"flourish-coding-challenge/internal/models"
	"flourish-coding-challenge/pkg/validation"
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCustomerCmd = &cobra.Command{
	Use:   "create-customer",
	Short: "Create a customer interactively",
	Run: func(cmd *cobra.Command, args []string) {
		var name, email string

		// Loop until a non-empty name is provided
		for {
			fmt.Print("Enter customer name: ")
			fmt.Scanln(&name)

			if !validation.IsEmpty(name) {
				break
			}
			fmt.Println("Name cannot be empty. Please try again.")
		}

		// Loop until a non-empty and valid email is provided
		for {
			fmt.Print("Enter customer email: ")
			fmt.Scanln(&email)

			if !validation.IsEmpty(email) && validation.ValidateEmail(email) {
				break
			}
			fmt.Println("Invalid email address or email cannot be empty. Please try again.")
		}

		// Create a new customer
		customer := models.Customer{Name: name, Email: email}
		if err := db.DB.Create(&customer).Error; err != nil {
			fmt.Println("Failed to create customer:", err)
			return
		}

		fmt.Println("Customer created successfully.")
	},
}

var ListCustomersCmd = &cobra.Command{
	Use:   "list-customers",
	Short: "List all customers",
	Run: func(cmd *cobra.Command, args []string) {
		var customers []models.Customer
		if err := db.DB.Find(&customers).Error; err != nil {
			fmt.Println("Failed to list customers:", err)
			return
		}

		fmt.Println("Customers:")
		for _, customer := range customers {
			fmt.Println(customer.Name, customer.Email)
		}
	},
}
var GetCustomerCmd = &cobra.Command{
	Use:   "get-customer",
	Short: "Get a customer by email",
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
		var customer models.Customer
		if err := db.DB.First(&customer, "email = ?", email).Error; err != nil {
			fmt.Println("Failed to get customer:", err)
			return
		}
		fmt.Println("Customer:")
		fmt.Println("Name:", customer.Name)
		fmt.Println("Email:", customer.Email)
	},
}
var UpdateCustomerCmd = &cobra.Command{
	Use:   "update-customer",
	Short: "update a customer interactively",
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
		var customer models.Customer
		if err := db.DB.First(&customer, "email = ?", email).Error; err != nil {
			fmt.Println("Failed to get customer:", err)
			return
		}
		var name, newEmail string
		for {
			fmt.Print("Enter customer name: ")
			fmt.Scanln(&name)
			if !validation.IsEmpty(name) {
				break
			}
			fmt.Println("Name cannot be empty. Please try again.")

		}
		for {
			fmt.Print("Enter customer email: ")
			fmt.Scanln(&newEmail)
			if !validation.IsEmpty(newEmail) && validation.ValidateEmail(newEmail) {
				break
			}
			fmt.Println("Invalid email address or email cannot be empty. Please try again.")

		}

		customer.Name = name
		customer.Email = newEmail
		if err := db.DB.Save(&customer).Error; err != nil {
			fmt.Println("Failed to update customer:", err)
			return
		}
		fmt.Println("Customer updated successfully.")
	},
}
var DeleteCustomerCmd = &cobra.Command{
	Use:   "delete-customer",
	Short: "delete a customer by email",
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
		var customer models.Customer
		if err := db.DB.First(&customer, "email = ?", email).Error; err != nil {
			fmt.Println("Failed to get customer:", err)
			return
		}
		if err := db.DB.Delete(&customer).Error; err != nil {
			fmt.Println("Failed to delete customer:", err)
			return
		}
		fmt.Println("Customer deleted successfully.")
	},
}

func init() {

	rootCmd.AddCommand(CreateCustomerCmd)
	rootCmd.AddCommand(ListCustomersCmd)
	rootCmd.AddCommand(GetCustomerCmd)
	rootCmd.AddCommand(UpdateCustomerCmd)
	rootCmd.AddCommand(DeleteCustomerCmd)

}

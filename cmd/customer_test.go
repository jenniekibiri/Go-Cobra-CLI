package cmd

import (
	
	"fmt"
	"os"
	"testing"

	"flourish-coding-challenge/config"
	"flourish-coding-challenge/internal/db"
	"flourish-coding-challenge/internal/models"
	"flourish-coding-challenge/pkg/utils"

	"github.com/stretchr/testify/assert"
	
)

func TestMain(m *testing.M) {
	if err := setupTestDB(); err != nil {
		fmt.Println("Failed to set up the test database:", err)
		os.Exit(1)
	}

	exitCode := m.Run()

	teardownTestDB()

	os.Exit(exitCode)
}

func setupTestDB() error {
	config.LoadEnv("../.env.test")
	dsn := os.Getenv("DB_STRING")
	if err := db.ConnectDB(dsn); err != nil {
		return err
	}
	db.Migrate()
	return nil
}

func teardownTestDB() {
	db.DB.Migrator().DropTable(&models.Customer{})
	db.DB.Migrator().DropTable(&models.Order{})
}

func createRandomCustomer(t *testing.T) *models.Customer {
	
	customer := models.Customer{
		Name:  utils.RandomString(10),
		Email: utils.RandomEmail(),
	}

	err := runCreateCustomerCmd(&customer)
	assert.NoError(t, err)

	var createdCustomer models.Customer
	err = db.DB.First(&createdCustomer, "email = ?", customer.Email).Error
	assert.NoError(t, err)
	assert.Equal(t, customer.Name, createdCustomer.Name)
	assert.Equal(t, customer.Email, createdCustomer.Email)


	return &customer
}


func TestCreateCustomerCmd(t *testing.T) {
	createRandomCustomer(t)

}

func runCreateCustomerCmd(customer *models.Customer) error {
	if err := db.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func TestListCustomersCmd(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCustomer(t)
	}

	var customers []models.Customer
	err := runListCustomersCmd(&customers)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(customers))

}

func runListCustomersCmd(customers *[]models.Customer) error {
	if err := db.DB.Find(customers).Error; err != nil {
		return err
	}
	return nil
}

func TestGetCustomerCmd(t *testing.T) {
	testUser := createRandomCustomer(t)
	fmt.Println(testUser)

	var customer models.Customer
	err := runGetCustomerCmd(&customer)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Name, customer.Name)
	assert.Equal(t, testUser.Email, customer.Email)


}

func runGetCustomerCmd(customer *models.Customer) error {
	if err := db.DB.First(customer, "email = ?", customer.Email).Error; err != nil {
		return err
	}
	return nil
}

func TestUpdateCustomerCmd(t *testing.T) {
	testUser :=  createRandomCustomer(t)

	testUser.Name = "Updated Name"
	err := runUpdateCustomerCmd(testUser)
	assert.NoError(t, err)

	var updatedCustomer models.Customer
	err = db.DB.First(&updatedCustomer, "email = ?", testUser.Email).Error
	assert.NoError(t, err)
	assert.Equal(t, testUser.Name, updatedCustomer.Name)
	assert.Equal(t, testUser.Email, updatedCustomer.Email)

}





func runUpdateCustomerCmd(customer *models.Customer) error {
	if err := db.DB.Save(customer).Error; err != nil {
		return err
	}
	return nil
}

func TestDeleteCustomerCmd(t *testing.T) {
	testUser := createRandomCustomer(t)

	err := runDeleteCustomerCmd(testUser)
	assert.NoError(t, err)

	var customer models.Customer
	err = db.DB.First(&customer, "email = ?", testUser.Email).Error
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())

}





func runDeleteCustomerCmd(customer *models.Customer) error {
	if err := db.DB.Delete(customer).Error; err != nil {
		return err
	}
	return nil
}

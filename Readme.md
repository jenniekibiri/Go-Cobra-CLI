### Project setup (15 minutes)
 - Project setup 
 - Installing dependencies
 - Folder structure
 - Docker and docker-compose 
 - Created a docker-compose.yml file to define the MySQL container for test and  main DB

### Setting and establishing DB connection (15 minutes)
  -  used gorm.io/gorm  and gorm.io/driver/mysql to create models and connect to the database
  -  Created a db.go to take dsn and return a db connection for test and main DB
  -  used godoenv to load env variables from .env file for different environments (test and main)
  -  Created a sync.go to sync the models with the database 
### Created a main.go
  - used cobra to create a cli tool  to interact with the application 

### Created crud for Customer (1.5 hour) 
  - Created a customer.go to create a customer model and crud operations
   - validation
  - Created a customer_test.go to test the crud operations

### Created crud for Orders (1.5 hour)
  - Created a order.go to create a order model and crud operations
  - validation
  - Created a order_test.go to test the crud operations

### Getting started: 
1. Clone the repo
2. Create a .env and .env.test file in the root directory and add the following variables
```bash
DB_STRING="user:password@tcp(127.0.0.1:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

3. Run docker-compose up -d

4. Run go run main.go
```bash
  go run main.go --help to see the list of commands
```
 Commands: 
```bash
  go run main.go create-customer  to create a customer
  go run main.go list-customers  to list all customers
  go run main.go get-customer  to get a customer
  go run main.go update-customer  to update a customer
  go run main.go delete-customer  to delete a customer
  go run main.go create-order  to create an order
  go run main.go list-orders  to list all orders
  go run main.go get-order  to get an order
  go run main.go update-order  to update an order
  go run main.go delete-order  to delete an order
``` 


### Testing:
1. Run  
```bash
   go test ./cmd -run  func_name 
```


























 


### Getting started: 
1. Clone the repo
2. Create an .env and .env.test file in the root directory and add the following variables
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

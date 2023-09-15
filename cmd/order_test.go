package cmd
import (
	"testing"
	"time"

	"flourish-coding-challenge/internal/db"
	"flourish-coding-challenge/internal/models"

	"github.com/stretchr/testify/assert"
)

func createRandomOrder(t *testing.T) *models.Order {

	order := models.Order{
		CustomerID: 1,
		QtyOrdered: 10,
		TotalPrice: 100,
		Date:       time.Now().Format("2006-01-02"),
	}

	err := runCreateOrderCmd(&order)
	assert.NoError(t, err)
	var createdOrder models.Order
	err = db.DB.First(&createdOrder, "customer_id = ?", order.CustomerID).Error
	assert.NoError(t, err)
	assert.Equal(t, order.CustomerID, createdOrder.CustomerID)
	assert.Equal(t, order.QtyOrdered, createdOrder.QtyOrdered)
	assert.Equal(t, order.TotalPrice, createdOrder.TotalPrice)
	assert.Equal(t, order.Date, createdOrder.Date)
	return &order

}

func TestCreateOrderCmd(t *testing.T) {
	createRandomOrder(t)

}

func runCreateOrderCmd(customer *models.Order) error {
	if err := db.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func TestListOrderCmd(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	var orders []models.Order
	err := runListOrdersCmd(&orders)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(orders))

}

func runListOrdersCmd(orders *[]models.Order) error {
	if err := db.DB.Find(orders).Error; err != nil {
		return err
	}
	return nil
}

func TestGetOrderCmd(t *testing.T) {
	testOrder := createRandomOrder(t)
	
	var order models.Order
	err := runGetOrderCmd(&order)
	assert.NoError(t, err)
	assert.Equal(t, testOrder.ID, order.ID)
	assert.Equal(t, testOrder.CustomerID, order.CustomerID)
	assert.Equal(t, testOrder.QtyOrdered, order.QtyOrdered)

}

func runGetOrderCmd(order *models.Order) error {
	if err := db.DB.First(order).Error; err != nil {
		return err
	}
	return nil
}

func TestUpdateOrderCmd(t *testing.T) {
	testOrder := createRandomOrder(t)
	testOrder.QtyOrdered = 20
	testOrder.TotalPrice = 200

	err := runUpdateOrderCmd(testOrder)
	assert.NoError(t, err)

	var order models.Order
	err = db.DB.First(&order, "id = ?", testOrder.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, testOrder.QtyOrdered, order.QtyOrdered)
	assert.Equal(t, testOrder.TotalPrice, order.TotalPrice)

}

func runUpdateOrderCmd(order *models.Order) error {
	if err := db.DB.Save(order).Error; err != nil {
		return err
	}
	return nil
}



func TestDeleteOrderCmd(t *testing.T) {
	testOrder := createRandomOrder(t)

	err := runDeleteOrderCmd(testOrder)
	assert.NoError(t, err)

	var order models.Order
	err = db.DB.First( &order, "id = ?", testOrder.ID).Error
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())

}

func runDeleteOrderCmd(order *models.Order) error {
	if err := db.DB.Delete(order).Error; err != nil {
		return err
	}

	return nil
}


package database

import (
	"regexp"
	"sesi-4/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// testing order.go

func Test_CreateOrder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	order := model.Order{
		CustomerName: "Customer Test",
		OrderedAt:    time.Now(),
		Items: []model.Item{
			{
				ItemCode:    "TEST",
				Description: "Test",
				Quantity:    1,
			},
		},
	}
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO orders (customer_name) VALUES($1) RETURNING order_id`)).WithArgs(order.CustomerName).WillReturnRows(sqlmock.NewRows([]string{"order_id"}).AddRow(0))
	for _, item := range order.Items {
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO items (item_code, description, quantity, order_id) VALUES($1, $2, $3, $4)`)).WithArgs(item.ItemCode, item.Description, item.Quantity, 0).WillReturnResult(sqlmock.NewResult(1, 1))
	}
	if err = CreateOrder(db, order); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// test GetAllOrder()
func Test_GetAllOrder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT o.*, i.* FROM "orders" as o INNER JOIN "items" as i ON o.order_id=i.order_id`).WillReturnRows(sqlmock.NewRows([]string{"order_id", "customer_name", "ordered_at", "item_id", "item_code", "description", "quantity", "order_id"}).AddRow(0, "Customer Test", time.Now(), 0, "TEST", "Test", 1, 0))

	orders, err := GetAllOrder(db)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if len(orders) != 1 {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//test GetOrder()
func Test_GetOrder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT o.*, i.* FROM "orders" as o INNER JOIN "items" as i ON o.order_id=i.order_id WHERE o.order_id = $1`)).WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{"order_id", "customer_name", "ordered_at", "item_id", "item_code", "description", "quantity", "order_id"}).AddRow(0, "Customer Test", time.Now(), 0, "TEST", "Test", 1, 0))

	order, err := GetOrder(db, 0)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if order.OrderID != 0 {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//test UpdateOrder()
func Test_UpdateOrder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	order := &model.Order{
		OrderID:      0,
		CustomerName: "Customer Test",
		OrderedAt:    time.Now(),
		Items: []model.Item{
			{
				ItemCode:    "TEST",
				Description: "Test",
				Quantity:    1,
			},
		},
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE orders SET customer_name = $1 WHERE order_id = $2`)).WithArgs(order.CustomerName, 0).WillReturnResult(sqlmock.NewResult(1, 1))
	for _, item := range order.Items {
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE items SET item_code = $1, description = $2, quantity = $3 WHERE order_id = $4 AND item_id = $5`)).WithArgs(item.ItemCode, item.Description, item.Quantity, 0, item.ItemID).WillReturnResult(sqlmock.NewResult(1, 1))
	}
	if err = UpdateOrder(db, order, 0); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//test deleteOrder()
func Test_DeleteOrder(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM items WHERE order_id = $1`)).WithArgs(0).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM orders WHERE order_id = $1`)).WithArgs(0).WillReturnResult(sqlmock.NewResult(1, 1))

	if err = DeleteOrder(db, 0); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

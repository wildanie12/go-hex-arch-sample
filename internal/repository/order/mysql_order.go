package order

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/wildanie12/go-hex-arch-sample/lib"
	"gorm.io/gorm"
)

// MySQLRepository
type MySQLRepository struct {
	dbInstance *lib.DatabaseInstance
}

// NewMySQL returns constructed mysql order repository.
func NewMySQL(di *lib.DatabaseInstance) *MySQLRepository {
	return &MySQLRepository{dbInstance: di}
}

// FindAll orders.
func (mr *MySQLRepository) FindAll() ([]Order, error) {
	data := []Order{}
	db := mr.dbInstance.OrderDB
	if tx := db.Order("id DESC").Find(&data); tx.Error != nil {
		return []Order{}, mr.makeError(tx.Error, "FindAll-Query")
	}
	return data, nil
}

// FindAllByUser orders.
func (mr *MySQLRepository) FindAllByUser(userID int) ([]Order, error) {
	data := []Order{}
	if tx := mr.dbInstance.OrderDB.Order("id DESC").Where("user_id = ?", userID).Find(&data); tx.Error != nil {
		return []Order{}, mr.makeError(tx.Error, "FindAllByUser-Query")
	}
	return data, nil
}

// Find single order by ID.
func (mr *MySQLRepository) Find(id uint) (Order, error) {
	data := Order{}
	tx := mr.dbInstance.OrderDB.Order("id DESC").Where("id = ?", id).Find(&data)
	if tx.Error != nil {
		return Order{}, mr.makeError(tx.Error, "Find-Query")
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return Order{}, mr.makeError(tx.Error, "Find-Query-NotFound")
	}
	return data, nil
}

// Create single order.
func (mr *MySQLRepository) Create(o Order) (Order, error) {
	tx := mr.dbInstance.OrderDB.Create(&o)
	if tx.Error != nil {
		return Order{}, mr.makeError(tx.Error, "Create-Query")
	}
	return o, nil
}

// Update single order.
func (mr *MySQLRepository) Update(o Order) (Order, error) {
	tx := mr.dbInstance.OrderDB.Model(&Order{}).Updates(o)
	if tx.Error != nil {
		return Order{}, mr.makeError(tx.Error, "Update-Query")
	}
	return o, nil
}

// Delete single order.
func (mr *MySQLRepository) Delete(o Order) (Order, error) {
	tx := mr.dbInstance.OrderDB.Delete(&Order{}, o.ID)
	if tx.Error != nil {
		return Order{}, mr.makeError(tx.Error, "Delete-Query")
	}
	return o, nil
}

// AddItemline into order.
func (mr *MySQLRepository) AddItemline(il Itemline, orderID uint) (Itemline, error) {
	il.OrderID = uint(orderID)
	tx := mr.dbInstance.OrderDB.Create(&il)
	if tx.Error != nil {
		return Itemline{}, mr.makeError(tx.Error, "AddItemline-Query")
	}
	return il, nil
}

// RemoveItemline into order.
func (mr *MySQLRepository) RemoveItemline(il Itemline) (Itemline, error) {
	tx := mr.dbInstance.OrderDB.Delete(&Itemline{}, il.ID)
	if tx.Error != nil {
		return Itemline{}, mr.makeError(tx.Error, "Delete-Query")
	}
	return il, nil
}

func (mr *MySQLRepository) makeError(err error, scope string) error {
	return errors.WrapPrefix(err, fmt.Sprintf("(x) mysql order repository err (%s)", scope), 0)
}

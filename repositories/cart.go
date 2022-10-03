package repositories

import (
	"_waysbook/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(Cart models.Cart) (models.Cart, error)
	UpdateCart(Cart models.Cart, ID int) (models.Cart, error)
	DeleteCart(Cart models.Cart, ID int) (models.Cart, error)
	GetTransactionID(TransactionID int) (models.Transaction, error)
	CreateTransactionID(transaction models.Transaction) (models.Transaction, error)
	FindCartsTransaction(TransactionID int) ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Book").Preload("Transaction").Preload("Transaction.User").Preload("Transaction.Book").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Book").Preload("Transaction").Preload("Transaction.User").Preload("Transaction.Book").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Book").Preload("Transaction").Preload("Transaction.Book").Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) GetTransactionID(TransactionID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Book").Where("user_id = ?", TransactionID).First(&transaction).Error
	return transaction, err
}

func (r *repository) CreateTransactionID(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Cart").Preload("Book").Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindCartsTransaction(TrxID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Book").Debug().Find(&carts, "transaction_id = ?", TrxID).Error

	return carts, err
}
package repositories

import (
	"_waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindBooksById(BookID int) (models.Book, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	FindbyIDTransaction(TransactionId int, Status string) (models.Transaction, error)
	UpdateTransactions(status string, ID string) error
	GetOneTransaction(ID string) (models.Transaction, error)
	GetTransactionId() (models.Transaction, error)
	// UpdatesTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransactionId() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Book").First(&transaction, "status = ?", "Active").Error

	return transaction, err
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Book").Find(&transactions).Error

	return transactions, err
}

func (r *repository) FindBooksById(BookID int) (models.Book, error) {
	var Books models.Book
	err := r.db.Find(&Books, BookID).Error

	return Books, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
var transaction models.Transaction
// not yet using category relation, cause this step doesnt Belong to Many
err := r.db.Preload("User").Preload("Cart").Preload("Cart.Book").First(&transaction, ID).Error

return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Cart").Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Book").Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}

func (r *repository) FindbyIDTransaction(TransactionId int, Status string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Book").Where("user_id = ? AND status = ?", TransactionId, Status).First(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransactions(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Cart").Preload("Cart.Book").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	// if status != transaction.Status && status == "success" {
	// 	var book models.Book
	// 	r.db.First(&book, transaction.ID)
	// }

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Cart").Preload("Cart.Book").First(&transaction, "id = ?", ID).Error
  
	return transaction, err
}

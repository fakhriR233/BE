package handlers

import (
	dto "_waysbook/dto/result"
	transactiondto "_waysbook/dto/transaction"
	"_waysbook/models"
	"_waysbook/repositories"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey: os.Getenv("CLIENT_KEY"),
  }

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
  }

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransactions()
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(transaction)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// request := new(transactiondto.TransactionRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//   w.WriteHeader(http.StatusBadRequest)
	//   response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	//   json.NewEncoder(w).Encode(response)
	//   return
	// }

	// total, _ := strconv.Atoi(r.FormValue("totalPayment"))
	// booksID, _ := strconv.Atoi(r.FormValue("books_id"))

	

	request := transactiondto.TransactionRequest{
		UserID:    				userId,
		Attachment:    			"",
		// BookID:     			booksID,
		Total:    				0,
		Status:      			"Pending",
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	var TransIdIsMatch = false
	var TransactionId int
	for !TransIdIsMatch {
		TransactionId = userId + rand.Intn(10000) - rand.Intn(100) + int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(TransactionId)
		if transactionData.ID == 0 {
			TransIdIsMatch = true
		}
	}

	// book, _ := h.TransactionRepository.FindBooksById(booksID)

	transaction := models.Transaction{
		ID:						TransactionId,
		Attachment:    			request.Attachment,
		Total:    				request.Total,
		UserID:    				userId,
		Status:      			"Active",
	}

	statusCheck, _ := h.TransactionRepository.FindbyIDTransaction(userId, "active")
	if statusCheck.Status == "active" {
		response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
		json.NewEncoder(w).Encode(response)
	} else {
		data, _ := h.TransactionRepository.CreateTransaction(transaction)
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: 200, Data: data}
		json.NewEncoder(w).Encode(response)
	}}
	// transaction, _ = h.TransactionRepository.GetTransaction(transaction.ID)

	// w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	// json.NewEncoder(w).Encode(response)


func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	idTrans := int(userInfo["id"].(float64))

	request := new(transactiondto.TransactionUpdateRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	// id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// transactionDataOld, _ := h.TransactionRepository.GetTransaction(id)

	transaction, err := h.TransactionRepository.FindbyIDTransaction(idTrans, "Active")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// transaction := models.Transaction{}

	if request.UserID != 0 {
		transaction.UserID = request.UserID}

	if request.Attachment != "" {
	transaction.Attachment = "Bank.jpeg"}

	// var booksID []int
	// for _, r := range r.FormValue("book_id") {
	// 	if int(r-'0') >= 0 {
	// 		booksID = append(booksID, int(r-'0'))
	// 	}
	// }

	if request.Total != 0 {
		transaction.Total = request.Total}

	if request.Status != "Active" {
		transaction.Status = request.Status}

	_, err = h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	dataTransactions, err := h.TransactionRepository.FindbyIDTransaction(idTrans, request.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New("SB-Mid-server-_4kL7N7mT6x-vq_X_Vq-5BnI", midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(idTrans),
			GrossAmt: int64(dataTransactions.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.User.FullName,
			Email: dataTransactions.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: 200, Data: snapResp}
	json.NewEncoder(w).Encode(response)
  }

func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteTransaction(data)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerTransaction) FindbyIDTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	// id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.FindbyIDTransaction(userId, "Active")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: 200, Data: transaction}
	json.NewEncoder(w).Encode(response)
}


func convertResponseTransaction(u models.Transaction) models.Transaction {
	return models.Transaction{
		ID:				u.ID,
	  User:    			u.User,
	  Attachment:    	u.Attachment,
	//   BookID:			u.BookID,
	  Status:      		u.Status,
	}
}

func convertResponseTransactionUpdate(u models.Transaction) transactiondto.TransactionUpdateResponse {
	return transactiondto.TransactionUpdateResponse{
		ID:				u.ID,
	  User:    			u.User,
	  Attachment:    	u.Attachment,
	//   BookID:			u.BookID,
	  Status:      		u.Status,
	}
}

func convertResponseDeleteTransaction(u models.Transaction) transactiondto.TransactionDeleteResponse {
	return transactiondto.TransactionDeleteResponse{
	  ID:    u.ID,
	}
}

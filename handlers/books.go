package handlers

import (
	bookdto "_waysbook/dto/book"
	dto "_waysbook/dto/result"
	"_waysbook/models"
	"_waysbook/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerBook struct {
	BookRepository repositories.BookRepository
}

var path_file = "http://localhost:5000/uploads/"

func HandlerBook(BookRepository repositories.BookRepository) *handlerBook {
	return &handlerBook{BookRepository}
}

func (h *handlerBook) FindBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range books {
		books[i].Thumbnail = path_file + p.Thumbnail
	}

	for i, p := range books {
		books[i].BookAttachment = path_file + p.BookAttachment
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: books}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) FindBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBook()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range books {
		books[i].Thumbnail = path_file + p.Thumbnail
	}

	for i, p := range books {
		books[i].BookAttachment = path_file + p.BookAttachment
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: books}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	var book models.Book
	book, err := h.BookRepository.GetBook(id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	book.Thumbnail = path_file + book.Thumbnail
	book.BookAttachment = path_file + book.BookAttachment
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBook(book)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64)) 

	dataContex := r.Context().Value("dataFile") // add this code
  	filename := dataContex.(string)

	dataPDF := r.Context().Value("dataPDF") // add this code
  	filePDF := dataPDF.(string)

	pages, _ := strconv.Atoi(r.FormValue("Pages"))
	// isbn, _ := strconv.Atoi(r.FormValue("ISBN"))
	price, _ := strconv.Atoi(r.FormValue("Price"))
  
	request := bookdto.BookRequest{
		Title:       			r.FormValue("Title"),
		PublicationDate:    	r.FormValue("PublicationDate"),
		Pages:    				pages,
		ISBN:    				r.FormValue("ISBN"),
		Author:    				r.FormValue("Author"),
		Price:    				price,
		Description:       		r.FormValue("Description"),
	}
	
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//   w.WriteHeader(http.StatusBadRequest)
	//   response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	//   json.NewEncoder(w).Encode(response)
	//   return
	// }
  
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	book := models.Book{
		Title:    				request.Title,
	  PublicationDate:    		request.PublicationDate,
	  Pages:    				request.Pages,
	  ISBN:    					request.ISBN,
	  Author:    				request.Author,
	  Price:    				request.Price,
	  Description:    			request.Description,
	  BookAttachment:    		filePDF,
	  Thumbnail:    			filename,
	}
  
	book, err = h.BookRepository.CreateBook(book)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	book, _ = h.BookRepository.GetBook(book.ID)
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: book}
	json.NewEncoder(w).Encode(response)
}


func (h *handlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	book, err := h.BookRepository.GetBook(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	data, err := h.BookRepository.DeleteBook(book,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteBook(data)}
	json.NewEncoder(w).Encode(response)
  }

  func convertResponseBook(u models.Book) models.Book {
	return models.Book{
		ID:						u.ID,
	  Title:    				u.Title,
	  PublicationDate:    		u.PublicationDate,
	  Pages:    				u.Pages,
	  ISBN:    					u.ISBN,
	  Author:    				u.Author,
	  Price:    				u.Price,
	  Description:    			u.Description,
	  BookAttachment:    		u.BookAttachment,
	  Thumbnail:    			u.Thumbnail,
	}
}

  func convertResponseDeleteBook(u models.Book) models.Book {
	return models.Book{
		ID:						u.ID,
	}
}
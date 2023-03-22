package helper

import (
	"golang-library-app/model/domain"
	"golang-library-app/model/web"
	"golang-library-app/model/web/response"
	"golang-library-app/utils"
)

func ConvertToBookResponse(domainBook *domain.Book) response.BookResponse {
	correctTime := utils.NullTime{
		Time:  domainBook.UpdatedAt.Time,
		Valid: domainBook.UpdatedAt.Valid,
	}
	return response.BookResponse{
		ID:              domainBook.ID,
		PublisherID:     domainBook.PublisherID,
		ISBN:            domainBook.ISBN,
		Title:           domainBook.Title,
		Author:          domainBook.Author,
		PublicationYear: domainBook.PublicationYear,
		Amount:          domainBook.Amount,
		Bookshelf:       domainBook.Bookshelf,
		CreatedAt:       &domainBook.CreatedAt,
		UpdatedAt:       &correctTime,
		DeletedAt:       &domainBook.DeletedAt,
	}
}

func ConvertToWebResponse(allData interface{}) web.ResponseWeb {
	return web.ResponseWeb{
		Code:    200,
		Message: "Success!",
		Data:    allData,
	}
}

func ConvertToPublisherResponse(domainPublisher *domain.Publisher) response.PublisherResponse {
	var allBookResponse []response.BookResponse
	for _, domainBook := range domainPublisher.Books {
		bookResponse := ConvertToBookResponse(&domainBook)
		allBookResponse = append(allBookResponse, bookResponse)
	}
	correctTime := utils.NullTime{
		Time:  domainPublisher.UpdatedAt.Time,
		Valid: domainPublisher.UpdatedAt.Valid,
	}
	return response.PublisherResponse{
		ID:          domainPublisher.ID,
		Name:        domainPublisher.Name,
		Description: domainPublisher.Description.String,
		CreatedAt:   &domainPublisher.CreatedAt,
		UpdatedAt:   &correctTime,
		DeletedAt:   &domainPublisher.DeletedAt,
		BooksData:   allBookResponse,
	}
}

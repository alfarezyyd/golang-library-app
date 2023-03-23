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
	var allKindResponse []response.KindResponse
	for _, domainKind := range domainBook.Kinds {
		kindResponse := ConvertToKindResponse(&domainKind)
		allKindResponse = append(allKindResponse, kindResponse)
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
		Kinds:           allKindResponse,
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

func ConvertToKindResponse(domainKind *domain.Kind) response.KindResponse {
	var allBookResponse []response.BookResponse
	for _, domainBook := range domainKind.Books {
		bookResponse := ConvertToBookResponse(&domainBook)
		allBookResponse = append(allBookResponse, bookResponse)
	}
	return response.KindResponse{
		ID:          domainKind.ID,
		Name:        domainKind.Name,
		Description: domainKind.Description.String,
		CreatedAt:   &domainKind.CreatedAt,
		UpdatedAt:   &utils.NullTime{Time: domainKind.UpdatedAt.Time, Valid: domainKind.UpdatedAt.Valid},
		Deleted:     &domainKind.DeletedAt,
		Books:       allBookResponse,
	}
}

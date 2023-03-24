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

func ConvertToVisitorResponse(domainVisitor *domain.Visitor) response.VisitorResponse {
	return response.VisitorResponse{
		ID:        domainVisitor.ID,
		NIN:       domainVisitor.NIN,
		Name:      domainVisitor.Name,
		Instance:  domainVisitor.Instance,
		CreatedAt: domainVisitor.CreatedAt,
		UpdatedAt: utils.NullTime{
			Time:  domainVisitor.UpdatedAt.Time,
			Valid: true,
		},
		DeletedAt: domainVisitor.DeletedAt,
	}
}

func ConvertToEmployeeResponse(domainEmployee *domain.Employee) response.EmployeeResponse {
	return response.EmployeeResponse{
		ID:              domainEmployee.ID,
		Name:            domainEmployee.Name,
		Gender:          domainEmployee.Gender,
		Position:        domainEmployee.Position,
		TelephoneNumber: domainEmployee.TelephoneNumber,
		Address:         domainEmployee.Address,
		CreatedAt:       domainEmployee.CreatedAt,
		UpdatedAt: utils.NullTime{
			Time:  domainEmployee.UpdatedAt.Time,
			Valid: true,
		},
		DeletedAt: domainEmployee.DeletedAt,
	}
}

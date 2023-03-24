package helper

import (
	"golang-library-app/model/domain"
	"golang-library-app/model/web"
	"golang-library-app/model/web/response"
	"golang-library-app/utils"
)

func ConvertToBookResponse(domainBook *domain.Book) response.BookResponse {
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
		UpdatedAt: &utils.NullTime{
			Time:  domainBook.UpdatedAt.Time,
			Valid: domainBook.UpdatedAt.Valid,
		},
		DeletedAt: &domainBook.DeletedAt,
		Kinds:     allKindResponse,
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
		CreatedAt: &domainVisitor.CreatedAt,
		UpdatedAt: &utils.NullTime{
			Time:  domainVisitor.UpdatedAt.Time,
			Valid: domainVisitor.UpdatedAt.Valid,
		},
		DeletedAt: &domainVisitor.DeletedAt,
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
		CreatedAt:       &domainEmployee.CreatedAt,
		UpdatedAt: &utils.NullTime{
			Time:  domainEmployee.UpdatedAt.Time,
			Valid: domainEmployee.UpdatedAt.Valid,
		},
		DeletedAt: &domainEmployee.DeletedAt,
	}
}

func ConvertToMemberResponse(domainMember *domain.Member) response.MemberResponse {
	return response.MemberResponse{
		ID:              domainMember.ID,
		NIN:             domainMember.NIN,
		Name:            domainMember.Name,
		BirthPlace:      domainMember.BirthPlace,
		BirthDate:       &domainMember.BirthDate,
		Address:         domainMember.Address,
		Gender:          domainMember.Gender,
		TelephoneNumber: domainMember.TelephoneNumber,
		CreatedAt:       &domainMember.CreatedAt,
		UpdatedAt: &utils.NullTime{
			Time:  domainMember.UpdatedAt.Time,
			Valid: domainMember.UpdatedAt.Valid,
		},
		DeletedAt: &domainMember.DeletedAt,
	}
}

func ConvertToUserResponse(domainUser *domain.User) response.UserResponse {
	return response.UserResponse{
		ID:       domainUser.ID,
		MemberID: domainUser.MemberID,
		Username: domainUser.Username,
		Email:    domainUser.Email,
		Password: domainUser.Password,
		EmailVerifiedAt: &utils.NullTime{
			Time:  domainUser.EmailVerifiedAt.Time,
			Valid: domainUser.EmailVerifiedAt.Valid,
		},
		CreatedAt: &domainUser.CreatedAt,
		UpdatedAt: &utils.NullTime{
			Time:  domainUser.UpdatedAt.Time,
			Valid: domainUser.UpdatedAt.Valid,
		},
		DeletedAt: &domainUser.DeletedAt,
	}
}

package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	custRep repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRep: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRep.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custReponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custReponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRep.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotfoundError("User not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custReponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custReponse, nil

}

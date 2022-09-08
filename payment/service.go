package payment

import (
	"errors"
	"fmt"
	"time"
)

// Middleware decorates a service.

type Service interface {
	Authorise(total float32) (Authorisation, error) // GET /paymentAuth
	Health() []Health                               // GET /health

type Authorisation struct {
	Authorised bool   `json:"authorised"`
	Message    string `json:"message"`
	IsFraud    string `json:"isFraud"`
}

type Middleware func(Service) Service

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

// NewFixedService returns a simple implementation of the Service interface,
// fixed over a predefined set of socks and tags. In a real service you'd
// probably construct this with a database handle to your socks DB, etc.
func NewAuthorisationService(declineOverAmount float32) Service {
	return &service{
		declineOverAmount: declineOverAmount,
	}
}

type service struct {
	declineOverAmount float32
}


func (s *service) Health() []Health {
	var health []Health
	app := Health{"payment", "OK", time.Now().String()}
	health = append(health, app)
	return health
}

func (s *service) Authorise(amount float32) (Authorisation, error) {
	if amount == 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	if amount < 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	authorised := false
	message := "Payment declined"
	isFraud := "no"
	if amount <= s.declineOverAmount {
		authorised = true
		message = "Payment authorised"
		isFraud = "yes"
	} else {
		message = fmt.Sprintf("Payment declined: amount exceeds %.2f", s.declineOverAmount)
		isFraud = "no"
	}
	authorisation := Authorisation{
                    		Authorised: authorised,
                    		Message:    message,
                    		IsFraud:    isFraud,
                    	}
    log.Println("Sending payment authorization response %v", authorisation)
	return authorisation, nil
}

func (s *service) Health() []Health {
	var health []Health
	app := Health{"payment", "OK", time.Now().String()}
	health = append(health, app)
	return health
}

var ErrInvalidPaymentAmount = errors.New("Invalid payment amount")

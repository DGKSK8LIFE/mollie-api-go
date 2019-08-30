package mollie

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"
)

// Resource constants
const (
	PaymentsResourceName string = "payments"
)

var (
	//PaymentsResourcePattern to construct URIs
	PaymentsResourcePattern = PaymentsResourceName + "/%v"
)

type Embedded struct {
	Payments []PaymentResponse `json:"payments"`
}

//Payment contains all the relevant information to call
//the payments resource endpoint.
type Payment struct {
	ID          string      `json:"id,omitempty"`
	Amount      Currency    `json:"amount"`
	Description string      `json:"description"`
	RedirectURL *HalURL     `json:"redirectUrl,omitempty"`
	WebhookURL  *HalURL     `json:"webhookUrl,omitempty"`
	Locale      string      `json:"locale,omitemtpy"`
	Method      string      `json:"method,omitempty"`
	Metadata    interface{} `json:"metadata,omitempty"`
	Sequence    string      `json:"sequenceType,omitemtpy"`
	CustomerID  string      `json:"customerId,omitemtpy"`
	MandateID   string      `json:"mandateId,omitempty"`
	ProfileID   string      `json:"profileId,omitempty"`
}

//PaymentResponse is a response descriptor for single
//result operations like Find and Update
type PaymentResponse struct {
	Payment
	hal              `json:"_links"`
	Status           string    `json:"status,omitempty"`
	Cancelable       bool      `json:"isCancelable,omitempty"`
	Resource         string    `json:"resource,omitempty"`
	ExpiresAt        time.Time `json:"expiresAt,omitempty"`
	ExpiredAt        time.Time `json:"expiredAt,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	AuthorizedAt     time.Time `json:"authorizedAt,omitempty"`
	PaidAt           time.Time `json:"paidAt,omitempty"`
	CanceledAt       time.Time `json:"canceledAt,omitempty"`
	FailedAt         time.Time `json:"failedAt,omitempty"`
	Mode             string    `json:"mode,omitempty"`
	AmountRefunded   Currency  `json:"amountRefunded,omitempty"`
	AmountRemaining  Currency  `json:"amountRemaining,omitempty"`
	AmountCaptured   Currency  `json:"amountCaptured,omitempty"`
	CountryCode      string    `json:"countryCode,omitempty"`
	SettlementAmount Currency  `json:"settlementAmount,omitempty"`
	SettlementID     string    `json:"settlementId,omitempty"`
	OrderID          string    `json:"orderId,omitempty"`
}

//PaymentsListResponse is a response descriptor for requests
//that expects multiple elements on the response, it contains
//other relevant information like Pagination metadata and
//HAL information.
type PaymentsListResponse struct {
	Count    int      `json:"count"`
	Embedded Embedded `json:"_embedded"`
	hal      `json:"_links"`
}

//PaymentsResource describes the actions for the service.
type PR interface {
	List(ctx context.Context) (PaymentsListResponse, error)
	Get(ctx context.Context, id string) (PaymentResponse, error)
	Create(ctx context.Context, p Payment) (PaymentResponse, error)
	Update(ctx context.Context, p Payment) (PaymentResponse, error)
	Cancel(id string) (PaymentResponse, error)
}

type PaymentResource struct {
	api *APIClient
}

//List will retrieve a paginated response containing a result count
//hal links, pagination information and also a list of entities
func (r *PaymentResource) List(ctx context.Context) (PaymentsListResponse, error) {
	var pl PaymentsListResponse
	req, err := r.api.NewAPIRequest("", PaymentsResourceName, nil, true)
	if err != nil {
		return pl, err
	}

	resp, err := r.api.HTTPClient.Do(req)
	if err != nil {
		return pl, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pl, err
	}
	err = json.Unmarshal(content, &pl)
	if err != nil {
		return pl, err
	}
	return pl, err
}

func (r *PaymentResource) Get(ctx context.Context, id string) (PaymentResponse, error) {
	return PaymentResponse{}, nil
}

func (r *PaymentResource) Create(ctx context.Context, p Payment) (PaymentResponse, error) {
	return PaymentResponse{}, nil
}

func (r *PaymentResource) Update(ctx context.Context, p Payment) (PaymentResponse, error) {
	return PaymentResponse{}, nil
}

func (r *PaymentResource) Cancel(id string) (PaymentResponse, error) {
	return PaymentResponse{}, nil
}

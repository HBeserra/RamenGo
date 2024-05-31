package orders

import (
	"encoding/json"
	"errors"
	"net/http"
)

type OrderResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var (
	ErrInvalidRequest  = errors.New("both brothId and proteinId are required")
	ErrExternalService = errors.New("external service error")
)

// Create an order of ramen
//
// Parameters:
// - brothId: The ID of the broth to order
// - proteinId: The ID of the protein to order
//
// Return type:
// - Order: The created order
func Create(brothId, proteinId string) (*OrderResponse, error) {

	var (
		order *OrderResponse
		err   error
	)

	// Validate inputs
	if brothId == "" || proteinId == "" {
		err = ErrInvalidRequest
		return nil, err
	}

	// Get order ID
	orderId, err := getOrderId()
	if err != nil {
		err = errors.Join(err, ErrExternalService)
		return nil, err
	}

	order = &OrderResponse{
		ID:          orderId,
		Description: "Salt and Chasu Ramen",
		Image:       "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	}

	return order, nil
}

// getOrderId returns a available order ID
func getOrderId() (string, error) {
	req, err := http.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("x-api-key", "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to generate order ID")
	}

	var orderIdResponse struct {
		OrderId string `json:"orderId"`
	}
	err = json.NewDecoder(resp.Body).Decode(&orderIdResponse)
	if err != nil {
		return "", err
	}

	return orderIdResponse.OrderId, nil
}

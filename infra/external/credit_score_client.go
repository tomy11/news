package external

import "news-api/business/entities"

type CreditScoreClient struct {
	apiURL string
	apiKey string
}

func NewCreditScoreClient(apiURL, apiKey string) *CreditScoreClient {
	return &CreditScoreClient{
		apiURL: apiURL,
		apiKey: apiKey,
	}
}

func (c *CreditScoreClient) GetCreditScore(customer *entities.Customer) (int, error) {
	// External API call to get credit score
	return 750, nil
}
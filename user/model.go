package user

import "context"

// Usecase represent the User's usecases.
type Usecase interface {
	SetToken(string) error
	GetList() ([]TitleList, error)
}

// Repository represent the User's repository contract.
type Repository interface {
	SetToken(ctx context.Context, bytes []byte)
	GetToken(ctx context.Context) *string
}

// HttpDelivery represent the User's transport
type HttpDelivery interface {
	GetSubscriptions(string) (*List, error)
}

type List struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}
type TitleList struct {
	Title string `json:"title"`
}

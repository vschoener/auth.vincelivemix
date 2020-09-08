package dto

// AuthenticatedResponse contains the authenticated response
type UserFetchResponse struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
}

package redash

// User is returned for all /api/user
type User struct {
	AuthType        string `json:"auth_type"`
	Name            string `json:"name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	ProfileImageURL string `json:"profile_image_url"`
	ID              int    `json:"id"`
	Groups          []int  `json:"groups"`
	Email           string `json:"email"`
}

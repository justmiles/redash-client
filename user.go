package redash

import "time"

// User is returned for all /api/user
type User struct {
	AuthType            string      `json:"auth_type,omitempty"`
	IsDisabled          bool        `json:"is_disabled,omitempty"`
	UpdatedAt           string      `json:"updated_at,omitempty"`
	ProfileImageURL     string      `json:"profile_image_url,omitempty"`
	IsInvitationPending bool        `json:"is_invitation_pending,omitempty"`
	Groups              []int       `json:"groups,omitempty"`
	ID                  int         `json:"id,omitempty"`
	Name                string      `json:"name,omitempty"`
	CreatedAt           string      `json:"created_at,omitempty"`
	DisabledAt          interface{} `json:"disabled_at,omitempty"`
	IsEmailVerified     bool        `json:"is_email_verified,omitempty"`
	ActiveAt            time.Time   `json:"active_at,omitempty"`
	Email               string      `json:"email,omitempty"`
}

type AutoGenerated struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Results  []struct {
		IsArchived        bool     `json:"is_archived"`
		RetrievedAt       string   `json:"retrieved_at"`
		UpdatedAt         string   `json:"updated_at"`
		IsFavorite        bool     `json:"is_favorite"`
		Query             string   `json:"query"`
		ID                int      `json:"id"`
		Description       string   `json:"description"`
		LastModifiedByID  int      `json:"last_modified_by_id"`
		Tags              []string `json:"tags"`
		Version           int      `json:"version"`
		QueryHash         string   `json:"query_hash"`
		APIKey            string   `json:"api_key"`
		DataSourceID      int      `json:"data_source_id"`
		IsSafe            bool     `json:"is_safe"`
		LatestQueryDataID int      `json:"latest_query_data_id"`
		Schedule          struct {
			Interval  int         `json:"interval"`
			Until     interface{} `json:"until"`
			DayOfWeek string      `json:"day_of_week"`
			Time      string      `json:"time"`
		} `json:"schedule"`
		IsDraft   bool    `json:"is_draft"`
		Name      string  `json:"name"`
		CreatedAt string  `json:"created_at"`
		Runtime   float64 `json:"runtime"`
		Options   struct {
			Parameters []interface{} `json:"parameters"`
		} `json:"options"`
	} `json:"results"`
}

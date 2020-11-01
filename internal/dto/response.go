package dto

type ApiResponse struct {
	Headers    map[string]string `json:"-"`
	StatusCode int               `json:"statusCode,omitempty"`
	Message    string            `json:"message,omitempty"`

	// in success case, statusCode = 2xx
	Data  interface{} `json:"data,omitempty"`
	Debug interface{} `json:"debug,omitempty"`

	Errors map[string][]string `json:"errors,omitempty"`
}

package handler

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type JsonMessage struct {
	// in: body
	Message string `json:"message"`
}

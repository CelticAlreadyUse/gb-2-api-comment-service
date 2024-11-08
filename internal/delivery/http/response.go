package httphandler

type Response struct {
	Data        any            `json:"data,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}
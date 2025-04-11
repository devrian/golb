package response

type (
	JSONResponse struct {
		Data        interface{}            `json:"data,omitempty"`
		Message     string                 `json:"message,omitempty"`
		Code        string                 `json:"code"`
		StatusCode  int                    `json:"statusCode"`
		ErrorString string                 `json:"error,omitempty"`
		Error       error                  `json:"-"`
		RealError   string                 `json:"-"`
		Latency     string                 `json:"latency"`
		Log         map[string]interface{} `json:"-"`
		HTMLPage    bool                   `json:"-"`
		Result      interface{}            `json:"result,omitempty"`
	}
)

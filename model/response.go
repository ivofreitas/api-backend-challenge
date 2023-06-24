package model

// Meta - Response's header
type Meta struct {
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
	RecordCount int `json:"record_count"`
}

// Response - general response object
type Response struct {
	Meta    Meta          `json:"meta"`
	Records []interface{} `json:"records"`
}

// NewResponse - Response's constructor
func NewResponse(offSet, limit, recordCount int, records []interface{}) *Response {
	meta := Meta{
		Offset:      offSet,
		Limit:       limit,
		RecordCount: recordCount,
	}
	body := &Response{
		Meta:    meta,
		Records: records,
	}

	return body
}

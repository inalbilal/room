package room

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// TestNewResponse tests the NewResponse function.
func TestNewResponse(t *testing.T) {
	// Mock HTTP response
	httpResp := &http.Response{
		StatusCode: http.StatusOK,
		Request: &http.Request{
			Method: http.MethodGet,
			URL: &url.URL{
				Scheme: "https",
				Host:   "example.com",
				Path:   "/test",
			},
		},
		Header: http.Header{},
		Body:   io.NopCloser(strings.NewReader("test body")),
	}

	// Test case for successful creation of Response
	resp, err := NewResponse(httpResp)
	if err != nil {
		t.Errorf("NewResponse() returned unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("NewResponse() returned response with unexpected status code: %d", resp.StatusCode)
	}

	// Test case for error in NewResponse
	httpResp.Body = nil
	resp, err = NewResponse(httpResp)
	if err == nil {
		t.Error("NewResponse() did not return expected error for HTTP client error")
	}
	var result any
	if resp.DTO(&result) == nil {
		t.Error("Response DTO() did not return expected error for HTTP client error")
	}

	if resp.DTOorFail(&result) == nil {
		t.Error("Response DTOorFail() did not return expected error for HTTP client error")
	}

	if resp.ResponseBody() != nil {
		t.Error("Response ResponseBody() did not return expected error for HTTP client error")
	}

	_, err = resp.ResponseBodyOrFail()
	if err != nil {
		return
	}
}

// TestResponse_OK tests the OK method of the Response struct.
func TestResponse_OK(t *testing.T) {
	// Test case for successful response (status code 200)
	response := Response{StatusCode: 200}
	if !response.OK() {
		t.Error("Response OK() returned false for status code 200")
	}

	// Test case for unsuccessful response (status code 404)
	response = Response{StatusCode: 404}
	if response.OK() {
		t.Error("Response OK() returned true for status code 404")
	}
}

// TestResponse_SetHeader tests the SetHeader method of the Response struct.
func TestResponse_SetHeader(t *testing.T) {
	// Test case for setting header
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	response := Response{}
	response = response.setHeader(header)
	if response.Header == nil {
		t.Error("Response SetHeader() did not set the header correctly")
	}
}

// TestResponse_SetRequestHeader tests the SetRequestHeader method of the Response struct.
func TestResponse_SetRequestHeader(t *testing.T) {
	// Test case for setting request header
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	response := Response{}
	response = response.setRequestHeader(header)
	if response.RequestHeader == nil {
		t.Error("Response SetRequestHeader() did not set the request header correctly")
	}
}

// TestResponse_SetRequestURI tests the SetRequestURI method of the Response struct.
func TestResponse_SetRequestURI(t *testing.T) {
	// Test case for setting request URI
	httpReq := &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "example.com",
			Path:   "/test",
		},
	}
	response := Response{}
	response = response.setRequestURI(httpReq)

	if response.RequestURI.Scheme() != "https" {
		t.Error("Response SetRequestURI() did not set the request URI correctly")
	}
}

// TestResponse_SetData tests the SetData method of the Response struct.
func TestResponse_SetData(t *testing.T) {
	// Test case for setting response data
	httpResp := &http.Response{
		Body: io.NopCloser(strings.NewReader("test data")),
	}
	response := Response{}
	response, err := response.setData(httpResp)
	if err != nil {
		t.Errorf("Response SetData() returned unexpected error: %v", err)
	}
	if len(response.Data) == 0 {
		t.Error("Response SetData() did not set the response data correctly")
	}
}

func TestResponse_SetRequestBody(t *testing.T) {
	// Test case for setting request body
	httpReq := &http.Request{
		Body: io.NopCloser(strings.NewReader(`{"key": "value"}`)),
	}

	response := Response{}
	response = response.setRequestBody(httpReq)
	if response.RequestBody == nil {
		t.Error("Response SetRequestBody() did not set the request body correctly")
	}

	// Test case for setting request body with invalid JSON
	defer func() { _ = recover() }()
	httpReq = &http.Request{
		Body: io.NopCloser(strings.NewReader("test body")),
	}

	response = Response{}
	response = response.setRequestBody(httpReq)
	if response.RequestBody == nil {
		t.Error("Response SetRequestBody() did not set the request body correctly")
	}

}

func TestNewDTOFactory(t *testing.T) {
	type Data struct {
		Key string `json:"key"`
	}

	// Test case for creating new DTOFactory with content type application/json
	factory := NewDTOFactory("application/json")
	if factory == nil {
		t.Error("NewDTOFactory() returned nil")
	}

	err := factory.marshall([]byte(`{"key": "value"}`), &Data{})
	if err != nil {
		t.Errorf("DTOFactory marshall() returned unexpected error: %v", err)
	}

	// Test case for creating new DTOFactory with content type text/xml
	factory = NewDTOFactory("text/xml")
	if factory == nil {
		t.Error("NewDTOFactory() returned nil")
	}

	err = factory.marshall([]byte(`<?xml version="1.0" encoding="UTF-8"?><key>value</key>`), &Data{})
	if err != nil {
		t.Errorf("DTOFactory marshall() returned unexpected error: %v", err)
	}

	// Test case for creating new DTOFactory with multiple content types
	factory = NewDTOFactory("application/json", "text/xml")
	if factory == nil {
		t.Error("NewDTOFactory() returned nil for multiple content types")
	}

	// Test case for creating new DTOFactory with empty content type
	factory = NewDTOFactory()
	if factory == nil {
		t.Error("NewDTOFactory() returned nil for empty content type")
	}
}

package room

import (
	"testing"
)

func TestNewConnector(t *testing.T) {
	c := NewConnector("http://example.com")
	if c.baseUrl != "http://example.com" {
		t.Errorf("NewConnector() returned unexpected base URL: %s", c.baseUrl)
	}

	c = NewConnector("http://example.com", WithHeaderConnector(NewHeader()))
	if c.Header == nil {
		t.Error("NewConnector() did not set the header")
	}

	c = NewConnector("http://example.com", WithHeaderContextBuilder(NewContextBuilder(30)))
	if c.contextBuilder == nil {
		t.Error("NewConnector() did not set the context builder")
	}

}

func TestConnector_Send(t *testing.T) {
	c := NewConnector("http://example.com")
	_, err := c.Send("/path/to/resource")
	if err != nil {
		t.Errorf("Send() returned an error: %v", err)
	}

}

func TestConnector_Do(t *testing.T) {
	c := NewConnector("http://example.com")
	resp, err := c.Do(NewRequest("/path/to/resource"))
	if err != nil {
		t.Errorf("Do() returned an error: %v", err)
	}

	c = NewConnector("http://example.com", WithHeaderConnector(NewHeader()))
	resp, err = c.Do(NewRequest("/path/to/resource"))

	if resp.Header == nil {
		t.Error("Do() did not set the header")
	}

	c = NewConnector("http://example.com", WithHeaderContextBuilder(NewContextBuilder(30)))
	resp, err = c.Do(NewRequest("/path/to/resource"))
	if c.contextBuilder == nil {
		t.Error("Do() did not set the context builder")
	}

}

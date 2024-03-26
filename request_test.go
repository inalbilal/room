package room

import "testing"

func TestNewRequest(t *testing.T) {
	data := map[string]interface{}{"key": "value"}
	bodyParser := NewJsonBodyParser(data)

	urlData := map[string]interface{}{"key1": "value1", "key2": "value2"}
	query := NewQuery(urlData)

	builder := NewContextBuilder(0)

	r := NewRequest("http://example.com", WithMethod("GET"), WithHeader(NewHeader()), WithBody(bodyParser), WithQuery(query), WithContextBuilder(builder))
	_, err := r.Send()
	if err != nil {
		t.Error("NewRequest() failed to send request")
	}

}

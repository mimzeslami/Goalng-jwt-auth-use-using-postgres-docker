package utils

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestReadJSON(t *testing.T) {
	// Create a sample JSON payload
	payload := []byte(`{"name":"John","age":30}`)

	// Create a request with the sample payload
	req, err := http.NewRequest("POST", "/example", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the ReadJSON function
	var data map[string]interface{}
	err = ReadJSON(rr, req, &data)
	if err != nil {
		t.Fatalf("ReadJSON failed: %s", err)
	}

	// Verify the decoded JSON data
	expectedData := map[string]interface{}{
		"name": "John",
		"age":  float64(30),
	}

	if data["name"] != expectedData["name"] {
		t.Errorf("ReadJSON name mismatch. Got %s, expected %s", data["name"], expectedData["name"])
	}

	if !reflect.DeepEqual(data["age"], expectedData["age"]) {
		t.Errorf("ReadJSON age mismatch. Got %s, expected %s", data["age"], expectedData["age"])
	}

}
func TestWriteJSON(t *testing.T) {
	// Create a sample data to be encoded as JSON
	data := map[string]interface{}{
		"message": "Hello, world!",
		"error":   false,
		"data":    nil,
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the WriteJSON function
	err := WriteJSON(rr, http.StatusOK, data)
	if err != nil {
		t.Fatalf("WriteJSON failed: %s", err)
	}

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("WriteJSON status code mismatch. Got %d, expected %d", rr.Code, http.StatusOK)
	}

	// Verify the response body
	expectedBody := `{"data":null,"error":false,"message":"Hello, world!"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("WriteJSON response body mismatch. Got %s, expected %s", rr.Body.String(), expectedBody)
	}
}

func TestErrorJSON(t *testing.T) {
	// Create a sample error
	err := errors.New("Some error")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the ErrorJSON function
	err = ErrorJSON(rr, err, http.StatusInternalServerError)
	if err != nil {
		t.Fatalf("ErrorJSON failed: %s", err)
	}

	// Verify the response status code
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("ErrorJSON status code mismatch. Got %d, expected %d", rr.Code, http.StatusInternalServerError)
	}

	// Verify the response body
	expectedBody := `{"error":true,"message":"Some error"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("ErrorJSON response body mismatch. Got %s, expected %s", rr.Body.String(), expectedBody)
	}
}

func TestSuccessJSON(t *testing.T) {
	// Create a sample data to be encoded as JSON
	data := map[string]interface{}{
		"message": "Success",
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the SuccessJSON function
	err := SuccessJSON(rr, "Success message", data, http.StatusOK)
	if err != nil {
		t.Fatalf("SuccessJSON failed: %s", err)
	}

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("SuccessJSON status code mismatch. Got %d, expected %d", rr.Code, http.StatusOK)
	}

	// Verify the response body
	expectedBody := `{"error":false,"message":"Success message","data":{"message":"Success"}}`
	if rr.Body.String() != expectedBody {
		t.Errorf("SuccessJSON response body mismatch. Got %s, expected %s", rr.Body.String(), expectedBody)
	}
}

func TestBase64Encode(t *testing.T) {
	str := "Hello, world!"
	expectedEncodedStr := "SGVsbG8sIHdvcmxkIQ=="

	encodedStr := base64Encode(str)
	if encodedStr != expectedEncodedStr {
		t.Errorf("Base64Encode result mismatch. Got %s, expected %s", encodedStr, expectedEncodedStr)
	}
}

func TestBase64Decode(t *testing.T) {
	encodedStr := "SGVsbG8sIHdvcmxkIQ=="
	expectedDecodedStr := "Hello, world!"

	decodedStr, err := base64Decode(encodedStr)
	if err {
		t.Errorf("Base64Decode failed: Error decoding base64 string")
	}

	if decodedStr != expectedDecodedStr {
		t.Errorf("Base64Decode result mismatch. Got %s, expected %s", decodedStr, expectedDecodedStr)
	}
}

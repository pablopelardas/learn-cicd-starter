package auth

import (
	"net/http"
	"reflect"
	"testing"
)


func TestGetAPIKey(t *testing.T) {
    header := http.Header{
        "Authorization": []string{"ApiKey a"},
    }
    got, err := GetAPIKey(header)
    if err != nil {
        t.Fatalf("expected: %v, got: %v", nil, err)
    }
    want := "a"
    if !reflect.DeepEqual(want, got) {
         t.Fatalf("expected: %v, got: %v", want, got)
    }
}
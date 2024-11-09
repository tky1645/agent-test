package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)


func TestTestHandler(t *testing.T){
	w:=httptest.NewRecorder()
	r:=httptest.NewRequest("GET", "/", nil)
	const expected ="head"
	r.Header.Set("Content-Type", expected)

	headerHandler(w,r)
	res :=w.Result()
	defer	 res.Body.Close()

	
	var response response
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.Header != expected{
		t.Fatalf("test fail")
	}
}


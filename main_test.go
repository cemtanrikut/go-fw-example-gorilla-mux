package main

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

//TODO: Add Gorilla Mux tests

// func TestMainHandler(t *testing.T) {
// 	t.Parallel()

// 	r, err := http.NewRequest(http.MethodGet, "/", nil)
// 	if err != nil{
// 		fmt.Println(err)
// 	}

// 	w := httptest.NewRecorder()
// 	handler := http.HandlerFunc(MainHandler)
// 	handler.ServeHTTP(w, r)
	
	
// 	if status := w.Code; status != http.StatusOK{
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	expected := "Welcome endpoint"
// 	// assert equality
// 	assert.Equal(t, expected, w.Body.String(), "they should be equal")

// 	expected = "Wrong endpoint message"
// 	// assert inequality
// 	assert.NotEqual(t, expected, w.Body.String(), "they should not be equal")
// }

func TestMainHandler(t *testing .T){
	type args struct {
		expected string
	}
	tests := []struct {
		name 	string
		want 	string
	}{
		{
			name: "Happy scenario",
			want: "Welcome endpoint",
		},
		{
			name: "Unhappy scenario",
			want: "wrong endpoint message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil{
				fmt.Println(err)
			}

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(MainHandler)
			handler.ServeHTTP(w, r)

			if status := w.Code; status != http.StatusOK{
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			if tt.want != w.Body.String(){
				assert.NotEqual(t, tt.want, w.Body.String(), "they should not be equal")
				return
			}

			assert.Equal(t, tt.want, w.Body.String(), "they should be equal")
		})

		
	}
}

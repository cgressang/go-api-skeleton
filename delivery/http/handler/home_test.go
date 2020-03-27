package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cgressang/go-api-skeleton/usecase/interactor/mocks"
)

func TestHandlers_Home(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "good",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "You have a good test!",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

			mockInteractor := &mocks.Home{}

			mockInteractor.On("GetMessage").Return(test.expectedBody)

			c := NewHomeHandler(nil, mockInteractor)
			c.Home(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected: %s\ngot: %s\n", test.expectedBody, body)
			}
		})
	}

}

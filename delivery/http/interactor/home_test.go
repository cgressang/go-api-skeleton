package interactor

import (
	"testing"
)

func TestInteractors_GetMessage(t *testing.T) {
	tests := []struct {
		name             string
		expectedResponse string
	}{
		{
			name:             "good",
			expectedResponse: "You are Home!",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

			i := NewHomeInteractor()
			response := i.GetMessage()
			if response != test.expectedResponse {
				t.Logf("expected: %s\ngot: %s\n", test.expectedResponse, response)
				t.Fail()
			}
		})
	}

}

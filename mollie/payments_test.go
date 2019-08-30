package mollie

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/mollie/testdata"
)

func TestPaymentsList(t *testing.T) {
	setup()
	defer teardown()
	paymentsAPIEndpoint := "/v2/payments"

	testMux.HandleFunc(paymentsAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Header[AuthHeader]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}
		raw := testdata.ListPaymentsExample
		fmt.Fprint(w, raw)
	})

	err := testClient.WithAPIKey("dummy-key")
	if err != nil {
		t.Fatal(err)
	}

	var ctx context.Context

	tests := []struct {
		name    string
		want    int
		wantErr bool
		err     error
	}{
		{
			name:    "test request is successful",
			wantErr: false,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testClient.Payments.List(ctx)
			if tt.wantErr && err != nil {
				if !reflect.DeepEqual(err, tt.err) {
					t.Fatal(err)
				}
			}

			if err != tt.err {
				t.Errorf("wanted nil, got %+v", err)
			}
		})
	}
}

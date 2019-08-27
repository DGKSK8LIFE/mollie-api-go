package mollie

import (
	"reflect"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	type args struct {
		code  string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Currency
		wantErr bool
		err     error
	}{
		{
			name: "test a currency is correctly instantiated",
			args: args{
				code:  "MXN",
				value: "1000.00",
			},
			want: Currency{
				CurrencyCode: "MXN",
				Value:        "1000.00",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "test invalid alpha3 code returns error",
			args: args{
				code:  "Mexican peso",
				value: "1000.00",
			},
			want:    Currency{},
			wantErr: true,
			err:     errInvalidISO4217Code,
		},
		{
			name: "test an empty currency code defaults to EUR",
			args: args{
				code:  "",
				value: "1000.00",
			},
			want: Currency{
				CurrencyCode: "EUR",
				Value:        "1000.00",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "test a string containing a space as currency code defaults to EUR",
			args: args{
				code:  " ",
				value: "1000.00",
			},
			want: Currency{
				CurrencyCode: "EUR",
				Value:        "1000.00",
			},
			wantErr: false,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCurrency(tt.args.code, tt.args.value)
			if tt.wantErr && err != nil {
				if tt.err.Error() != err.Error() {
					t.Errorf("error mismatch, got: %v | want %v", err, tt.err)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) || err != nil {
					t.Errorf("expectation mismatch: got %v | want: %v", got, tt.want)
				}
			}
		})
	}
}

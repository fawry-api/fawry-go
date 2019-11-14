package fawry

import "testing"

func TestFawryClient_getURL(t *testing.T) {
	type fields struct {
		IsSandbox      bool
		FawrySecureKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "returns production link if isSandbox false",
			fields: fields{
				IsSandbox:      false,
				FawrySecureKey: "test",
			},
			want: "https://www.atfawry.com/ECommerceWeb/Fawry/payments/",
		},
		{
			name: "returns sandbox link if isSandbox true",
			fields: fields{
				IsSandbox:      true,
				FawrySecureKey: "test",
			},
			want: "https://atfawry.fawrystaging.com/ECommerceWeb/Fawry/payments/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := FawryClient{
				IsSandbox:      tt.fields.IsSandbox,
				FawrySecureKey: tt.fields.FawrySecureKey,
			}
			if got := fc.getURL(); got != tt.want {
				t.Errorf("FawryClient.getURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

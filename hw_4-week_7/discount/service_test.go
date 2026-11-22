package discount

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestValidatePercent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		percent int
		wantErr error
	}{
		{name: "zero is valid", percent: 0},
		{name: "ten is valid", percent: 10},
		{name: "hundred is valid", percent: 100},
		{name: "negative is invalid", percent: -1, wantErr: ErrInvalidPercent},
		{name: "above hundred is invalid", percent: 101, wantErr: ErrInvalidPercent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidatePercent(tt.percent)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestApply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		price   int
		percent int
		want    int
		wantErr error
	}{
		{name: "regular discount", price: 200, percent: 10, want: 180},
		{name: "zero discount", price: 150, percent: 0, want: 150},
		{name: "full discount", price: 150, percent: 100, want: 0},
		{name: "negative price", price: -1, percent: 10, wantErr: ErrNegativePrice},
		{name: "percent below zero", price: 100, percent: -1, wantErr: ErrInvalidPercent},
		{name: "percent above hundred", price: 100, percent: 101, wantErr: ErrInvalidPercent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Apply(tt.price, tt.percent)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				require.Zero(t, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFinalPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		items   []Item
		want    int
		wantErr error
	}{
		{
			name: "empty list",
			want: 0,
		},
		{
			name: "several valid items",
			items: []Item{
				{Price: 200, DiscountPercent: 10},
				{Price: 100, DiscountPercent: 0},
				{Price: 50, DiscountPercent: 50},
			},
			want: 305,
		},
		{
			name: "invalid price",
			items: []Item{
				{Price: 100, DiscountPercent: 10},
				{Price: -5, DiscountPercent: 10},
			},
			wantErr: ErrNegativePrice,
		},
		{
			name: "invalid discount",
			items: []Item{
				{Price: 100, DiscountPercent: 10},
				{Price: 80, DiscountPercent: 150},
			},
			wantErr: ErrInvalidPercent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := FinalPrice(tt.items)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				require.Zero(t, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

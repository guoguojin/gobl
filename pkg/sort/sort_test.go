package sort

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	t.Run("Sort ints", func(t *testing.T) {
		vals := []int{5, 4, 3, 2, 1}
		Quick(vals, 0, len(vals)-1, func(i, j int) bool {
			return i < j
		})
		if vals[0] != 1 || vals[len(vals)-1] != 5 {
			t.Errorf("Expected [1, 5], got %v", vals)
		}
	})
	t.Run("Sort int8s", func(t *testing.T) {
		vals := []int8{5, 4, 3, 2, 1}
		Quick(vals, 0, len(vals)-1, func(i, j int8) bool {
			return i < j
		})
		if vals[0] != 1 || vals[len(vals)-1] != 5 {
			t.Errorf("Expected [1, 5], got %v", vals)
		}
	})
	t.Run("Sort float32", func(t *testing.T) {
		bals := []float32{5, 4, 3, 2, 1}
		Quick(bals, 0, len(bals)-1, func(i, j float32) bool {
			return i < j
		})
		if bals[0] != 1 || bals[len(bals)-1] != 5 {
			t.Errorf("Expected [1, 5], got %v", bals)
		}
	})
	t.Run("Sort float64", func(t *testing.T) {
		bals := []float64{5, 4, 3, 2, 1}
		Quick(bals, 0, len(bals)-1, func(i, j float64) bool {
			return i < j
		})
		if bals[0] != 1 || bals[len(bals)-1] != 5 {
			t.Errorf("Expected [1, 5], got %v", bals)
		}
	})
	t.Run("Sort string", func(t *testing.T) {
		bals := []string{"5", "4", "3", "2", "1"}
		Quick(bals, 0, len(bals)-1, func(i, j string) bool {
			return i < j
		})
		if bals[0] != "1" || bals[len(bals)-1] != "5" {
			t.Errorf("Expected [1, 5], got %v", bals)
		}
	})
	t.Run("Sort ints even number of elements", func(t *testing.T) {
		bals := []int{6, 5, 4, 3, 2, 1}
		Quick(bals, 0, len(bals)-1, func(i, j int) bool {
			return i < j
		})
		if bals[0] != 1 || bals[len(bals)-1] != 6 {
			t.Errorf("Expected [1, 6], got %v", bals)
		}
	})
}

func TestDecimal(t *testing.T) {
	type args struct {
		v  []decimal.Decimal
		hi int
		lo int
	}
	tests := []struct {
		name string
		args args
		want []decimal.Decimal
	}{
		{
			name: "sort values",
			args: args{
				v: []decimal.Decimal{
					decimal.NewFromInt(10),
					decimal.NewFromInt(7),
					decimal.NewFromInt(8),
					decimal.NewFromInt(9),
					decimal.NewFromInt(1),
					decimal.NewFromInt(5),
				},
				lo: 0,
				hi: 5,
			},
			want: []decimal.Decimal{
				decimal.NewFromInt(1),
				decimal.NewFromInt(5),
				decimal.NewFromInt(7),
				decimal.NewFromInt(8),
				decimal.NewFromInt(9),
				decimal.NewFromInt(10),
			},
		},
		{
			name: "empty list",
			args: args{
				v:  nil,
				lo: 0,
				hi: 0,
			},
			want: nil,
		},
		{
			name: "fractional values",
			args: args{
				v: []decimal.Decimal{
					decimal.NewFromFloat(1.0001),
					decimal.NewFromFloat(1.00001),
					decimal.NewFromFloat(1.0000000001),
					decimal.NewFromFloat(1.00000001),
					decimal.NewFromFloat(1.0000000000001),
					decimal.NewFromFloat(1.0000001),
				},
				lo: 0,
				hi: 5,
			},
			want: []decimal.Decimal{
				decimal.NewFromFloat(1.0000000000001),
				decimal.NewFromFloat(1.0000000001),
				decimal.NewFromFloat(1.00000001),
				decimal.NewFromFloat(1.0000001),
				decimal.NewFromFloat(1.00001),
				decimal.NewFromFloat(1.0001),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quick(tt.args.v, tt.args.lo, tt.args.hi, func(a, b decimal.Decimal) bool {
				return a.LessThan(b)
			})
			for i, v := range tt.args.v {
				assert.True(t, v.Equal(tt.want[i]), "Expected %v, got %v", tt.want[i], v)
			}
		})
	}
}

func TestDates(t *testing.T) {
	type args struct {
		v  []time.Time
		hi int
		lo int
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{
			name: "Dates",
			args: args{
				v: []time.Time{
					time.Unix(500, 0),
					time.Unix(400, 0),
					time.Unix(300, 0),
					time.Unix(200, 0),
					time.Unix(100, 0),
				},
				hi: 0,
				lo: 4,
			},
			want: []time.Time{
				time.Unix(100, 0),
				time.Unix(200, 0),
				time.Unix(300, 0),
				time.Unix(400, 0),
				time.Unix(500, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quick(tt.args.v, tt.args.lo, tt.args.hi, func(a, b time.Time) bool {
				return a.Before(b)
			})
			assert.ElementsMatch(t, tt.want, tt.args.v)
		})
	}
}

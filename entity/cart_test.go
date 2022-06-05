package entity

import (
	"reflect"
	"testing"
)

func TestCart_AddToCart(t *testing.T) {
	type fields struct {
		Items map[string]CartItem
	}
	type args struct {
		item CartItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]CartItem
	}{
		{
			"add to cart 1",
			fields{map[string]CartItem{}},
			args{
				CartItem{
					ProductSku: "abc",
					Price:      1000.00,
					Qty:        2,
				},
			},
			map[string]CartItem{
				"abc": {
					ProductSku: "abc",
					Price:      1000.00,
					Qty:        2,
					TotalPrice: 2000.00,
				},
			},
		},
		{
			"add to cart 2",
			fields{
				map[string]CartItem{
					"abc": {
						ProductSku: "abc",
						Price:      1000.00,
						Qty:        2,
						TotalPrice: 2000.00,
					},
				},
			},
			args{
				CartItem{
					ProductSku: "abc",
					Price:      1000.00,
					Qty:        2,
				},
			},
			map[string]CartItem{
				"abc": {
					ProductSku: "abc",
					Price:      1000.00,
					Qty:        4,
					TotalPrice: 4000.00,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCart("trx1")
			c.Items = tt.fields.Items
			c.AddToCart(tt.args.item)
			if !reflect.DeepEqual(c.Items, tt.want) {
				t.Errorf("Cart.Items = %v, want %v", c.Items, tt.want)
			}
		})
	}
}

func TestCart_AddToApplicablePromo(t *testing.T) {
	type fields struct {
		applicablePromo map[string]CartApplicablePromo
	}
	type args struct {
		applicablePromo CartApplicablePromo
	}
	type want struct {
		applicablePromo map[string]CartApplicablePromo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			"add 1",
			fields{map[string]CartApplicablePromo{}},
			args{CartApplicablePromo{
				PromoID:    1,
				ProductSku: "abc",
				Amount:     100,
				Qty:        1,
			}},
			want{
				map[string]CartApplicablePromo{
					"abc": {
						PromoID:    1,
						ProductSku: "abc",
						Amount:     100,
						Qty:        1,
					},
				},
			},
		},
		{
			"add 2",
			fields{
				map[string]CartApplicablePromo{
					"abc": {
						PromoID:    1,
						ProductSku: "abc",
						Amount:     100,
						Qty:        1,
					},
				}},
			args{CartApplicablePromo{
				PromoID:    1,
				ProductSku: "abc",
				Amount:     100,
				Qty:        1,
			}},
			want{
				map[string]CartApplicablePromo{
					"abc": {
						PromoID:    1,
						ProductSku: "abc",
						Amount:     100,
						Qty:        2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				ApplicablePromo: tt.fields.applicablePromo,
			}

			c.AddToApplicablePromo(tt.args.applicablePromo)
			if !reflect.DeepEqual(c.ApplicablePromo, tt.want.applicablePromo) {
				t.Errorf("Cart.ApplicablePromo = %v, want %v", c.ApplicablePromo, tt.want.applicablePromo)
			}
		})
	}
}

func TestCart_SubstractApplicablePromo(t *testing.T) {
	type fields struct {
		ApplicablePromo map[string]CartApplicablePromo
	}
	type args struct {
		productSku string
		qty        int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]CartApplicablePromo
	}{
		{
			"substract 1",
			fields{map[string]CartApplicablePromo{
				"abc": {Qty: 3},
			}},
			args{"abc", 1},
			map[string]CartApplicablePromo{
				"abc": {Qty: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				ApplicablePromo: tt.fields.ApplicablePromo,
			}
			c.SubstractApplicablePromo(tt.args.productSku, tt.args.qty)
			if !reflect.DeepEqual(c.ApplicablePromo, tt.want) {
				t.Errorf("Cart.ApplicablePromo = %v, want %v", c.ApplicablePromo, tt.want)
			}
		})
	}
}

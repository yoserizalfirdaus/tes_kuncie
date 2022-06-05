package checkout

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yoserizalfirdaus/tes_kuncie/entity"
	mock_checkout "github.com/yoserizalfirdaus/tes_kuncie/mocks/usecase/checkout"
)

func TestCheckoutUsecase_AddToCart(t *testing.T) {
	type fields struct {
		productRepo *mock_checkout.MockProductRepository
		cartRepo    *mock_checkout.MockCartRepository
		promoRepo   *mock_checkout.MockPromoRepository
	}
	type args struct {
		input AddToCartInput
	}
	tests := []struct {
		name    string
		fields  func(ctrl *gomock.Controller) fields
		args    args
		mock    func(f fields)
		want    *entity.Cart
		wantErr bool
	}{
		{
			"success",
			func(ctrl *gomock.Controller) fields {
				return fields{
					productRepo: mock_checkout.NewMockProductRepository(ctrl),
					cartRepo:    mock_checkout.NewMockCartRepository(ctrl),
					promoRepo:   mock_checkout.NewMockPromoRepository(ctrl),
				}
			},
			args{
				AddToCartInput{"111", "aaa", 1},
			},
			func(f fields) {
				f.cartRepo.EXPECT().GetByTransactionID("111").Return(entity.NewCart("111"), nil)
				f.productRepo.EXPECT().GetBySku("aaa").Return(entity.Product{
					Sku:   "aaa",
					Price: 100,
				}, nil)
				f.promoRepo.EXPECT().GetPromoByProductSku("aaa").Return(entity.Promotion{}, fmt.Errorf("not found"))
				f.cartRepo.EXPECT().SaveCart(gomock.Any()).Return(nil)
			},
			&entity.Cart{
				TransactionID: "111",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := tt.fields(ctrl)
			tt.mock(fields)
			u := CheckoutUsecase{
				productRepo: fields.productRepo,
				cartRepo:    fields.cartRepo,
				promoRepo:   fields.promoRepo,
			}
			got, err := u.AddToCart(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckoutUsecase.AddToCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.TransactionID, tt.want.TransactionID) {
				t.Errorf("CheckoutUsecase.AddToCart() = %v, want %v", got.TransactionID, tt.want.TransactionID)
			}
		})
	}
}

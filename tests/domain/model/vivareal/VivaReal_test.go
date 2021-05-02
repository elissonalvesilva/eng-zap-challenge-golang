package vivamodel

import (
	"reflect"
	"testing"

	mock "github.com/elissonalvesilva/eng-zap-challenge-golang/tests/mock"

	viva "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/model/vivareal"
	modelError "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/errors"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
)

func TestNewImovel(t *testing.T) {
	type args struct {
		imovel protocols.Imovel
	}
	tests := []struct {
		name      string
		args      args
		want      protocols.Imovel
		wantErr   bool
		typeError interface{}
	}{
		{
			"Should return a error if Total Rental Price Value is invalid",
			args{
				imovel: mock.MockImovelErrorInvalidRentalTotalPrice,
			},
			mock.MockImovelErrorInvalidRentalTotalPrice,
			true,
			modelError.InvalidTotalRentalPriceValue(mock.MockImovelErrorInvalidRentalTotalPrice.ID),
		},
		{
			"Should return a error if MontlyCondoFee is invalid",
			args{
				imovel: mock.MockImovelInvalidMonthlycondofeeValue,
			},
			mock.MockImovelInvalidMonthlycondofeeValue,
			true,
			modelError.InvalidMonthlycondofeeValue(mock.MockImovelInvalidMonthlycondofeeValue.ID),
		},
		{
			"Should return a created imovel and a error equals to nil if success",
			args{
				imovel: mock.MockImovelSuccess,
			},
			mock.MockImovelSuccess,
			false,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := viva.NewImovel(tt.args.imovel)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewImovel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImovel() = %v, want %v", got, tt.want)
			}
		})
	}
}

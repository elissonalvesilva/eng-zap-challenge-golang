package zapmodel

import (
	"reflect"
	"testing"

	zap "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/model/zap"
	mock "github.com/elissonalvesilva/eng-zap-challenge-golang/tests/mock"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	modelError "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/errors"
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
			"Should return a error if Usableareas is equal to 0",
			args{
				imovel: mock.MockImovelErrorUsableareas,
			},
			mock.MockImovelErrorUsableareas,
			true,
			modelError.InvalidUsableareasValue(mock.MockImovelErrorUsableareas.ID),
		},
		{
			"Should return a error if price is less or equals to 1",
			args{
				imovel: mock.MockImovelErrorInvalidPrice,
			},
			mock.MockImovelErrorInvalidPrice,
			true,
			modelError.InvalidPriceValue(mock.MockImovelErrorInvalidPrice.ID),
		},
		{
			"Should return a error if value per meter is less than 3500",
			args{
				imovel: mock.MockImovelErrorNotElegiblePricePerMeterValue,
			},
			mock.MockImovelErrorNotElegiblePricePerMeterValue,
			true,
			modelError.NotElegiblePricePerMeterValue(mock.MockImovelErrorNotElegiblePricePerMeterValue.ID),
		},
		{
			"Should return a error if min value price is greater than price if is bounding box",
			args{
				imovel: mock.MockImovelErrorNotElegibleMinSaleValue,
			},
			mock.MockImovelErrorNotElegibleMinSaleValue,
			true,
			modelError.NotElegibleMinSaleValue(mock.MockImovelErrorNotElegibleMinSaleValue.ID),
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
			got, err := zap.NewImovel(tt.args.imovel)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewImovel() error = %v, wantErr %v, modelError= %v", err, tt.wantErr, tt.typeError)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImovel() = %v, want %v", got, tt.want)
			}
		})
	}
}

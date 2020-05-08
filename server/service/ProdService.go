package service

import (
	"context"
)

type ProdService struct {

}

func(this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	if request.ProdArea == ProdAreas_A {
		stock = 30
	} else if request.ProdArea == ProdAreas_B {
		stock = 40
	} else {
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func(this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 28},
		&ProdResponse{ProdStock: 29},
		&ProdResponse{ProdStock: 30},
		&ProdResponse{ProdStock: 31},
	}
	return &ProdResponseList{Prodres: Prodres}, nil
}

func(this *ProdService) GetProdInfo(ctx context.Context, request *ProdRequest) (*ProdModel, error) {
	ret := ProdModel{
		ProdId:        101,
		ProdName:      "testing",
		ProdPrice:     20.7,
	}
	return &ret, nil
}
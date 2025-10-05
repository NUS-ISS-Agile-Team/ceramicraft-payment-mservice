package grpc

import (
	"context"
	"time"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/common/paymentpb"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/log"
)

type PaymentService struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *PaymentService) PayOrder(ctx context.Context, req *paymentpb.PayOrderRequest) (*paymentpb.PayOrderResponse, error) {
	log.Logger.Infof("Received: %v", req)
	return &paymentpb.PayOrderResponse{Code: int32(paymentpb.RespCode_SUCCESS),
		PayOrderInfo: &paymentpb.PayOrderInfo{
			PayOrderId:  "abced",
			Amount:      10,
			UserId:      0,
			CreatedTime: time.Now().Unix(),
		}}, nil
}

func (s *PaymentService) QueryPayOrder(ctx context.Context, req *paymentpb.PayOrderQueryRequest) (*paymentpb.PayOrderQueryResponse, error) {
	log.Logger.Infof("Received: %v", req)
	payOrder := &paymentpb.PayOrderInfo{
		PayOrderId:  "aceeee",
		Amount:      10,
		UserId:      0,
		CreatedTime: time.Now().Unix(),
	}
	return &paymentpb.PayOrderQueryResponse{
		Code:          int32(paymentpb.RespCode_SUCCESS),
		PayOrderInfos: []*paymentpb.PayOrderInfo{payOrder}}, nil
}

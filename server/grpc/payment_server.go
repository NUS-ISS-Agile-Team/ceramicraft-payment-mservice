package grpc

import (
	"context"
	"fmt"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/common/paymentpb"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/repository/model"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/service"
)

type PaymentService struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *PaymentService) PayOrder(ctx context.Context, req *paymentpb.PayOrderRequest) (*paymentpb.PayOrderResponse, error) {
	log.Logger.Infof("[grpc-svr] method=PayOrder, req=%v", req)
	changeLog, err := service.GetUserAccountService().PayOrder(ctx, int(req.UserId), req.BizId, int(req.Amount))
	if err != nil {
		return nil, fmt.Errorf("failed to pay order: %v", err)
	}
	log.Logger.Infof("Payment successful, change log: %+v", changeLog)
	return &paymentpb.PayOrderResponse{
		Code: int32(paymentpb.RespCode_SUCCESS),
		PayOrderInfo: &paymentpb.PayOrderInfo{
			PayOrderId:  genPayOrderId(changeLog),
			Amount:      int32(changeLog.Amount),
			UserId:      req.UserId,
			CreatedTime: changeLog.CreatedAt.Unix(),
		}}, nil
}

func (s *PaymentService) QueryPayOrder(ctx context.Context, req *paymentpb.PayOrderQueryRequest) (*paymentpb.PayOrderQueryResponse, error) {
	log.Logger.Infof("[grpc-svr] method=QueryPayOrder, req=%v", req)
	if req.UserId == nil && req.PayOrderId == nil {
		log.Logger.Warnf("Either UserId or PayOrderId must be provided")
		return &paymentpb.PayOrderQueryResponse{
			Code: -1, //todo
		}, nil
	}
	changeLogs, err := service.GetUserAccountService().GetUserPayHistory(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to query pay order: %v", err)
	}
	ret := make([]*paymentpb.PayOrderInfo, 0)
	for _, cLog := range changeLogs {
		ret = append(ret, &paymentpb.PayOrderInfo{
			PayOrderId:  genPayOrderId(cLog),
			Amount:      int32(cLog.Amount),
			UserId:      *req.UserId,
			CreatedTime: cLog.CreatedAt.Unix(),
		})
	}
	return &paymentpb.PayOrderQueryResponse{
		Code:          int32(paymentpb.RespCode_SUCCESS),
		PayOrderInfos: ret}, nil
}

func genPayOrderId(changeLog *model.UserAccountChangeLog) string {
	return fmt.Sprintf("%d_%s_%d", changeLog.AccountId, changeLog.IdempotentKey, changeLog.ID)
}

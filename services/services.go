package services

import (
	"context"
	"fmt"
	"time"

	cm "framework/common"
)

//SubscriberServices is service definition
type PaymentServices interface {
	OrderHandler(context.Context, cm.Order) cm.Message
	CustomerHandler(context.Context, cm.Customer) cm.Message
	ProductHandler(context.Context, cm.Product) cm.Message
	FaspayHandler(context.Context, cm.RequestFaspay) cm.ResponseFaspay
}

type PaymentService struct{}

type ServiceMiddleware func(PaymentServices) PaymentServices

func utc() string {
	return time.Now().Format("2006-01-02 15:04:05 +0700")
}

func panicRecovery() {
	if r := recover(); r != nil {
		fmt.Printf("Recovering from panic: %v \n", r)
	}
}

package config

import (
	"context"
	"github.com/google/uuid"
	"github.com/order-service/common/service"
	"github.com/order-service/db"
	"sync"
)

type PaymentService struct {
	service.UnimplementedPaymentServiceServer
}

func (ps *PaymentService) FindAll(_ context.Context, _ *service.Empty) (*service.PaymentsResponse, error) {
	var (
		res  []*service.PaymentRequest
		_db  db.Database
		rslt *service.PaymentsResponse
	)
	_db.Connect()

	row, err := _db.Conn.Query("SELECT * FROM payment")
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var temp service.PaymentRequest

		if err := row.Scan(
			&temp.Id, &temp.CustomerId, &temp.CustomerName,
			&temp.ProductID, &temp.ProductName, &temp.Price,
			&temp.DateOrder, &temp.ServiceShipmentName,
			&temp.DateShipment, &temp.ShipmentMethod,
		); err != nil {
			return nil, err
		}

		res = append(res, &temp)
	}

	rslt.Payments = res

	defer _db.Conn.Close()
	return rslt, nil
}

func (ps *PaymentService) FindId(_ context.Context, in *service.PaymentId) (*service.PaymentRequest, error) {
	var (
		res *service.PaymentRequest
		_db db.Database
	)
	_db.Connect()

	if err := _db.Conn.QueryRow("SELECT * FROM payment WHERE id = ?", in.GetId()).Scan(
		&res.Id, &res.CustomerId, &res.CustomerName,
		&res.ProductID, &res.ProductName, &res.Price,
		&res.DateOrder, &res.ServiceShipmentName,
		&res.DateShipment, &res.ShipmentMethod,
	); err != nil {
		return nil, err
	}

	defer _db.Conn.Close()
	return res, nil
}

func (ps *PaymentService) AddPayment(_ context.Context, in *service.PaymentRequest) (*service.PaymentResponse, error) {
	var (
		_db db.Database
	)
	_db.Connect()

	if _, err := _db.Conn.Exec("INSERT INTO payment VALUES (?,?,?,?,?,?,?,?,?,?)",
		uuid.New(), in.GetCustomerId(), in.GetCustomerName(),
		in.GetProductID(), in.GetProductName(), in.GetPrice(),
		in.GetDateOrder(), in.GetServiceShipmentName(),
		in.GetDateShipment(), in.GetShipmentMethod(),
	); err != nil {
		return nil, err
	}

	defer _db.Conn.Close()
	return &service.PaymentResponse{Status: "ok"}, nil
}

func (ps *PaymentService) AddPayments(ctx context.Context, in *service.PaymentsResponse) (*service.PaymentResponse, error) {
	var wg sync.WaitGroup

	errCh := make(chan error, len(in.GetPayments()))

	wg.Add(len(in.GetPayments()))

	for _, value := range in.GetPayments() {
		go func(payment *service.PaymentRequest, c context.Context) {
			defer wg.Done()
			if _, err := ps.AddPayment(c, payment); err != nil {
				errCh <- err
			}
		}(value, ctx)
	}
	wg.Wait()

	close(errCh)

	var allErrors []error
	for err := range errCh {
		allErrors = append(allErrors, err)
	}
	if allErrors[0] != nil {
		return nil, allErrors[0]
	}

	return &service.PaymentResponse{Status: "ok"}, nil
}

func (ps *PaymentService) DeletePayment(_ context.Context, in *service.PaymentId) (*service.PaymentResponse, error) {
	var _db db.Database
	_db.Connect()

	if _, err := _db.Conn.Exec("DELETE FROM payment WHERE id = ?", in.GetId()); err != nil {
		return nil, err
	}

	defer _db.Conn.Close()

	return &service.PaymentResponse{Status: "ok"}, nil
}

func (ps *PaymentService) DeletePayments(ctx context.Context, in *service.PaymentIds) (*service.PaymentResponse, error) {
	var wg sync.WaitGroup

	wg.Add(len(in.GetPaymentId()))
	errCh := make(chan error, len(in.GetPaymentId()))

	for _, value := range in.GetPaymentId() {
		go func(id *service.PaymentId, c context.Context) {
			defer wg.Done()
			if _, err := ps.DeletePayment(c, id); err != nil {
				errCh <- err
			}
		}(value, ctx)
	}

	wg.Wait()

	close(errCh)

	var allErrors []error
	for err := range errCh {
		allErrors = append(allErrors, err)
	}
	if allErrors[0] != nil {
		return nil, allErrors[0]
	}

	return &service.PaymentResponse{Status: "ok"}, nil
}

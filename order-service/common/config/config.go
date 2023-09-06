package config

import (
	"context"
	"github.com/order-service/common/service"
	"github.com/order-service/db"
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

	return res, nil
}

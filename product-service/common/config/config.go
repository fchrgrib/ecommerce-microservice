package config

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/product-service/common/service"
	"github.com/product-service/db"
	"time"
)

type ProductServiceServer struct {
	service.UnimplementedProductServiceServer
}

func (ps *ProductServiceServer) FindOneById(_ context.Context, in *service.ProductId) (*service.ProductRequest, error) {
	var (
		result service.ProductRequest
		_db    db.Database
	)
	_db.Connect()

	if err := _db.DB.QueryRow("SELECT * FROM product WHERE id = ?", in.GetId()).Scan(&result.Id, &result.Name, &result.Category, &result.Type, &result.Price, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return nil, err
	}

	return &result, nil
}

func (ps *ProductServiceServer) FindAll(_ context.Context, _ *service.Empty) (*service.ProductRequests, error) {
	var values []*service.ProductRequest
	var result service.ProductRequests
	var _db db.Database
	_db.Connect()

	row, err := _db.DB.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer func(row *sql.Rows) {
		if err := row.Close(); err != nil {
			return
		}
	}(row)

	for row.Next() {
		var temp service.ProductRequest

		if err := row.Scan(&temp.Id, &temp.Name, &temp.Category, &temp.Type, &temp.Price, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			return nil, err
		}

		values = append(values, &temp)
	}

	result.Product = values

	defer _db.DB.Close()

	return &result, nil
}

func (ps *ProductServiceServer) AddProduct(_ context.Context, in *service.ProductRequest) (*service.Response, error) {
	var (
		response service.Response
		_db      db.Database
	)
	_db.Connect()

	if _, err := _db.DB.Exec("INSERT INTO product(id, product_name, category, product_type, price, created_at, updated_at) VALUES (?,?,?,?,?,?)", uuid.New(), in.GetName(), in.GetCategory(), in.GetType(), in.GetPrice(), time.Now().String(), time.Now().String()); err != nil {
		return nil, err
	}

	response.Status = "ok"

	defer _db.DB.Close()

	return &response, nil
}

func (ps *ProductServiceServer) AddProducts(_ context.Context, in *service.ProductRequests) (*service.Response, error) {

	var (
		response service.Response
		_db      db.Database
	)
	_db.Connect()

	for _, ins := range in.GetProduct() {
		if _, err := _db.DB.Exec("INSERT INTO product(id, product_name, category, product_type, created_at, updated_at) VALUES (?,?,?,?,?,?)", uuid.New(), ins.GetName(), ins.GetCategory(), ins.GetType(), ins.GetPrice(), time.Now().String(), time.Now().String()); err != nil {
			return nil, err
		}
	}

	response.Status = "ok"
	defer _db.DB.Close()

	return &response, nil
}

func (ps *ProductServiceServer) DeleteById(_ context.Context, in *service.ProductId) (*service.Response, error) {
	var (
		response service.Response
		_db      db.Database
	)
	_db.Connect()

	if _, err := _db.DB.Exec("DELETE FROM product WHERE id = ?", in.GetId()); err != nil {
		return nil, err
	}

	response.Status = "ok"
	return &response, nil
}

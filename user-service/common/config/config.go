package config

import (
	"context"
	"github.com/user-service/common/service"
	"github.com/user-service/db"
	"time"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func (us *UserService) AddUser(ctx context.Context, in *service.UserRequest) (*service.UserResponse, error) {
	var (
		_db      db.Database
		response service.UserResponse
	)
	_db.Connect()

	if _, err := _db.DB.Exec("INSERT INTO users(id, user_id, user_name, phone_number, email, address, born, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)", in.GetId(), in.GetUserId(), in.GetUserName(), in.GetPhoneNumber(), in.GetEmail(), in.GetAddress(), in.GetBorn(), time.Now(), time.Now()); err != nil {
		return nil, err
	}

	response.Status = "ok"
	defer _db.DB.Close()

	return &response, nil
}

func (us *UserService) UpdateUser(_ context.Context, in *service.UserUpdate) (*service.UserResponse, error) {
	var (
		_db      db.Database
		response service.UserResponse
	)
	_db.Connect()

	if _, err := _db.DB.Exec("UPDATE users SET user_id=?, user_name=?, phone_number=?, email=?, address=?, born=?, updated_at=? WHERE id=?", in.GetUserRequest().GetUserId(), in.GetUserRequest().GetUserName(), in.GetUserRequest().GetPhoneNumber(), in.GetUserRequest().GetEmail(), in.GetUserRequest().GetAddress(), in.GetUserRequest().GetBorn(), time.Now(), in.GetUserId().GetId()); err != nil {
		return nil, err
	}

	response.Status = "ok"
	defer _db.DB.Close()

	return &response, nil
}

func (us *UserService) FindUserById(_ context.Context, in *service.Id) (*service.UserRequest, error) {
	var (
		_db      db.Database
		response service.UserRequest
	)
	_db.Connect()

	if err := _db.DB.QueryRow("SELECT * FROM users WHERE id = ?", in.GetId()).Scan(&response.Id, &response.UserId, &response.UserName, &response.PhoneNumber, &response.Email, &response.Address, &response.Born, &response.CreatedAt, &response.UpdatedAt); err != nil {
		return nil, err
	}

	defer _db.DB.Close()
	return &response, nil

}

func (us *UserService) DeleteUserById(_ context.Context, in *service.Id) (*service.UserResponse, error) {
	var (
		_db      db.Database
		response service.UserResponse
	)
	_db.Connect()

	if _, err := _db.DB.Exec("DELETE FROM users WHERE id = ?", in.GetId()); err != nil {
		return nil, err
	}
	defer _db.DB.Close()

	response.Status = "ok"

	return &response, nil
}

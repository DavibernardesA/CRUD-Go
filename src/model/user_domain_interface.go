package model

import rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string

	SetID(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.ApiErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

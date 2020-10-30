package dgraphland

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type User struct {
	Model
	Username string `json:"username,omitempty" dgland:"string @index(exact) @upsert"`
	Password string `json:"password,omitempty" dgland:"password"`
}

type Post struct {
	Model
	Title       string `json:"title,omitempty" dgland:"string @index(exact) @upsert"`
	Description string `json:"description,omitempty" dgland:"string"`
}

// WrongStruct1 มี Model Nested Structure แต่ไม่มี field ที่มี Json Tag และ Dglang Tag เลย
type WrongStruct1 struct {
	Model
}

// WrongStruct2 ไม่มี Model Nested Structure
type WrongStruct2 struct {
	Title       string `json:"title,omitempty" dgland:"string @index(exact) @upsert"`
	Description string `json:"description,omitempty" dgland:"string"`
}

// WrongStruct3 ไม่มี Field Structure ใดๆเลย
type WrongStruct3 struct{}

type UtilsTestSuite struct {
	suite.Suite
	SchemaResult string
}

func (suite *UtilsTestSuite) SetupTest() {
	suite.SchemaResult = `created_at: string .
updated_at: string .
deleted_at: string .
username: string @index(exact) @upsert .
password: password .
title: string @index(exact) @upsert .
description: string .

type User {
	username
	password
	created_at
	updated_at
	deleted_at
}

type Post {
	title
	description
	created_at
	updated_at
	deleted_at
}
`
}

func (suite *UtilsTestSuite) TestCtxFunc() {
	suite.Equal(context.Background(), ctx())
}

func (suite *UtilsTestSuite) TestSchemaFunc() {
	result, err := Schema(&User{}, &Post{})
	suite.Nil(err)
	suite.Equal(suite.SchemaResult, result)

	result, err = Schema(&WrongStruct1{})
	suite.Nil(err)
	suite.Equal(result, "")

	result, err = Schema(&WrongStruct2{})
	suite.NotNil(err)
	suite.Equal(result, "")

	result, err = Schema(&WrongStruct3{})
	suite.Nil(err)
	suite.Equal(result, "")
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

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
	suite.Equal(suite.SchemaResult, Schema(&User{}, &Post{}))
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

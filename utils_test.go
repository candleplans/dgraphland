package dgraphland

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type User struct {
	Uid       string    `json:"uid,omitempty"`
	DType     []string  `json:"dgraph.type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" dgland:"string"`
	UpdatedAt time.Time `json:"updated_at,omitempty" dgland:"string"`
	DeletedAt time.Time `json:"deleted_at,omitempty" dgland:"string"`
	Username  string    `json:"username,omitempty" dgland:"string @index(exact) @upsert"`
	Password  string    `json:"password,omitempty" dgland:"password"`
}

type UtilsTestSuite struct {
	suite.Suite
	UserModel    User
	SchemaResult string
}

func (suite *UtilsTestSuite) SetupTest() {
	suite.SchemaResult = `created_at: string .
updated_at: string .
deleted_at: string .
username: string @index(hash) @upsert .
password: password .

type User {
	created_at
	updated_at
	deleted_at
	username
	password
}`
}

func (suite *UtilsTestSuite) TestCtxFunc() {
	suite.Equal(context.Background(), ctx())
}

func (suite *UtilsTestSuite) TestSchemaFunc() {
	suite.Equal(suite.SchemaResult, Schema(suite.UserModel))
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

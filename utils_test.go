package dgraphland

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (suite *UtilsTestSuite) SetupTest() {}

func (suite *UtilsTestSuite) TestCtxFunc() {
	suite.Equal(ctx(), context.Background())
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

package dgraphland

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AlterTestSuite struct {
	suite.Suite
}

func (suite *AlterTestSuite) SetupTest() {

}

// ทดสอบ DropAll
func (suite *AlterTestSuite) TestDropAllFunc() {
	cancel := SetupClient(target)
	defer cancel()

	err := DropAll()
	suite.NoError(err)
}

// ทดสอบ AutoMigrate
func (suite *AlterTestSuite) TestAutoMigrateFunc() {
	cancel := SetupClient(target)
	defer cancel()

	err := AutoMigrate(&User{}, &Post{})
	suite.NoError(err)
}

func TestAlterTestSuite(t *testing.T) {
	suite.Run(t, new(AlterTestSuite))
}

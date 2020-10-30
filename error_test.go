package dgraphland

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ErrorTestSuite struct {
	suite.Suite
	TypeName string
}

func (suite *ErrorTestSuite) SetupTest() {
	suite.TypeName = "User"
}

func (suite *ErrorTestSuite) TestErrEmptyStructFunc() {
	suite.EqualError(errEmptyStruct(suite.TypeName),
		`DgError: `+suite.TypeName+
			` Type -> "เป็น Struct ที่ว่างเปล่า"`)
}

func (suite *ErrorTestSuite) TestErrNoModelStructFunc() {
	suite.EqualError(errNoModelStruct(suite.TypeName),
		`DgError: `+suite.TypeName+
			` Type -> "ไม่มีการเรียกใช้ Model Struct หลักของ Dgraphland"`)
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}

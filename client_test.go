package dgraphland

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

const target = "localhost:9080"

type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) SetupTest() {

}

// ทดสอบการเชื่อมต่อ
func (suite *ClientTestSuite) TestConnectFunc() {
	cancel := SetupClient(target)
	defer cancel()
	err := DropAll()
	suite.NoError(err)
}

// ทดสอบการยกเลิกการเชื่อมต่อ
func (suite *ClientTestSuite) TestDisconnectFunc() {
	cancel := SetupClient(target)
	cancel()
	err := DropAll()
	suite.EqualError(err, "rpc error: code = Canceled desc = grpc: the client connection is closing")
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

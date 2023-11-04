package smartbft

import (
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/internal/pkg/comm"
)

type MockOrderer struct {
	address     string
	ledgerArray []*common.Block
	logger      *flogging.FabricLogger
	grpcServer  *comm.GRPCServer
	channel     chan string
}

func (mo *MockOrderer) Broadcast(server orderer.AtomicBroadcast_BroadcastServer) error {
	mo.logger.Infof("Broadcast called")
	mo.channel <- "Called Broadcast"
	panic("implement me: Broadcast")
	return nil
}

func (mo *MockOrderer) Deliver(server orderer.AtomicBroadcast_DeliverServer) error {
	mo.logger.Infof("Deliver called")
	mo.channel <- "Called Deliver"
	panic("implement me: Deliver")
	return nil
}

func (mo *MockOrderer) deliverBlocks(server orderer.AtomicBroadcast_DeliverServer) error {
	mo.channel <- "Called DeliverBlocks"
	mo.logger.Infof("DeliverBlocks called")
	panic("implement me: DeliverBlocks")
	return nil
}

func NewMockOrderer(address string, ledgerArray []*common.Block, options comm.SecureOptions, ch chan string) (*MockOrderer, error) {
	sc := comm.ServerConfig{
		SecOpts: options,
	}

	logger := flogging.MustGetLogger("mockorderer")

	grpcServer, err := comm.NewGRPCServer(address, sc)
	if err != nil {
		logger.Errorf("Error creating GRPC server: %s", err)
	}

	go grpcServer.Start()

	mo := &MockOrderer{
		address:     address,
		ledgerArray: ledgerArray,
		logger:      flogging.MustGetLogger("mockorderer"),
		grpcServer:  grpcServer,
		channel:     ch,
	}

	return mo, nil
}

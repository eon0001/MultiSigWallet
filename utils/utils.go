package utils

import (
	"os"
)

func GetPorts() (portList map[string]string, err error) {

	peeringPort := os.Getenv("P2P_PORT")
	rpcPort := os.Getenv("RPC_PORT")
	abciAppPort := os.Getenv("ABCI_PORT")

	return map[string]string{"P2P_PORT": peeringPort, "RPC_PORT": rpcPort, "ABCI_PORT": abciAppPort}, nil
}

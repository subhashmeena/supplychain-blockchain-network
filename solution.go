package main

import (
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type LogisticsChaincode struct {
}

func (t *LogisticsChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initiated the chaincode");
	_, args := stub.GetFunctionAndParameters()
	

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	

	return shim.Success(nil)
}


func (t *LogisticsChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(LogisticsChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
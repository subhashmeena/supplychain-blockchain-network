package main

import (
	"fmt"
	//"strconv"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type LogisticsChaincode struct {
}

type Seller struct {
	Id string
	Name string
	Location string
}

type Buyer struct {
	Id string
	Name string
	Location string
}

type LogisticsProvider struct {
	Id string
	Name string
	Location string
}

var buyerStore map[string]Buyer
var sellerStore map[string]Seller
var logisticsProviderStore map[string]LogisticsProvider

 


func (t *LogisticsChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initiated the chaincode");
	_, args := stub.GetFunctionAndParameters()
	

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	return shim.Success(nil)
}


func (t *LogisticsChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function,_ := stub.GetFunctionAndParameters();

	fmt.Println("The function invoked is ",function);
	if(function == "registerSeller") {
		return t.registerSeller(stub);
	} else if( function == "registerLogisticsProvider") {
		return t.registerLogisticsProvider(stub);
	} else if( function == "registerBuyer") {
		return t.registerBuyer(stub);
	} else if(function == "getSeller") {
		return t.getSeller(stub);
	} else if( function == "getBuyer" ) {
		return t.getBuyer(stub);
	} else if( function == "getLogisticsProvider") {
		return t.getLogisticsProvider(stub);
	} else {
		return shim.Success(nil);
	}
}

func main() {
	err := shim.Start(new(LogisticsChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *LogisticsChaincode) registerSeller(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("In the buyer seller function")

	sellerStore := make(map[string]Seller);
	sellerStore["seller01"] = Seller{ "seller01", "Worldwide Seller", "Mumbai"};

	fmt.Println(sellerStore);
	bytearray, _ := json.Marshal(sellerStore);

	fmt.Println(string(bytearray));
	
	err := stub.PutState("sellerstore", bytearray);
	if( err != nil) {
		fmt.Println("While writing sellerstore to ledger, error encountered ",err);
		return shim.Error("Error occurrered while writing sellerstore to the ledger");
	}

	return shim.Success([]byte("Successfully written sellerstore to the ledger"));
}

func (t *LogisticsChaincode) registerBuyer(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("In the buyer buyer function");


	buyerStore := make(map[string]Buyer);
	buyerStore["buyer01"] = Buyer{ "buyer01", "Retail Expo", "Chennai"};

	fmt.Println(buyerStore);
	bytearray, _ := json.Marshal(buyerStore);

	fmt.Println(string(bytearray));
	
	err := stub.PutState("buyerstore", bytearray);
	if( err != nil) {
		fmt.Println("While writing buyerstore to ledger, error encountered ",err);
		return shim.Error("Error occurrered while writing buyerstore to the ledger");
	}

	return shim.Success([]byte("Successfully written buyerstore to the ledger"));

}

func (t *LogisticsChaincode) registerLogisticsProvider(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("In the buyer logistics provider function");
	
	logisticsProviderStore := make(map[string]LogisticsProvider);
	logisticsProviderStore["transporter01"] = LogisticsProvider{ "transporter01", "Indian Transporter", "Mumbai"};

	fmt.Println(logisticsProviderStore);
	bytearray, _ := json.Marshal(logisticsProviderStore);

	fmt.Println(string(bytearray));
	
	err := stub.PutState("logisticsproviderstore", bytearray);
	if( err != nil) {
		fmt.Println("While writing logisticsProviderstore to ledger, error encountered ",err);
		return shim.Error("Error occurrered while writing logisticsProviderstore to the ledger");
	}

	return shim.Success([]byte("Successfully written logisticsProviderstore to the ledger"));
}


func (t *LogisticsChaincode) getSeller(stub shim.ChaincodeStubInterface) pb.Response {
	_,parameters := stub.GetFunctionAndParameters();
	
	//put a check here
	var sellerId = parameters[0];

	sellerbytes,err := stub.GetState("sellerstore");
	if(err != nil ) {
		return shim.Error("Could not retrieve seller store from the ledger");
	}


	sellerStore = make(map[string]Seller);

	err = json.Unmarshal(sellerbytes,&sellerStore);
	if(err != nil ) {
		fmt.Println(string(sellerbytes));
		fmt.Println(err);
		return shim.Error("Error while unmarshaling data retrieved from the ledger");
	}

	fmt.Println(sellerStore[sellerId].Name);
	return shim.Success([]byte("Successfully retrieved the json data from the ledger"));
	
}

func (t *LogisticsChaincode) getBuyer(stub shim.ChaincodeStubInterface) pb.Response {
	_,parameters := stub.GetFunctionAndParameters();
	
	//put a check here
	var buyerId = parameters[0];

	buyerbytes,err := stub.GetState("buyerstore");
	if(err != nil ) {
		return shim.Error("Could not retrieve buyer store from the ledger");
	}


	buyerStore = make(map[string]Buyer);

	err = json.Unmarshal(buyerbytes,&buyerStore);
	if(err != nil ) {
		fmt.Println(string(buyerbytes));
		fmt.Println(err);
		return shim.Error("Error while unmarshaling data retrieved from the ledger");
	}

	fmt.Println(buyerStore[buyerId].Name);
	return shim.Success([]byte("Successfully retrieved the json data from the ledger"));
	
}


func (t *LogisticsChaincode) getLogisticsProvider(stub shim.ChaincodeStubInterface) pb.Response {
	_,parameters := stub.GetFunctionAndParameters();
	
	//put a check here
	var logisticsProviderId = parameters[0];

	logisticsproviderbytes,err := stub.GetState("logisticsproviderstore");
	if(err != nil ) {
		return shim.Error("Could not retrieve logistics provider store from the ledger");
	}


	logisticsProviderStore = make(map[string]LogisticsProvider);

	err = json.Unmarshal(logisticsproviderbytes,&logisticsProviderStore);
	if(err != nil ) {
		fmt.Println(string(logisticsproviderbytes));
		fmt.Println(err);
		return shim.Error("Error while unmarshaling data retrieved from the ledger");
	}

	fmt.Println(logisticsProviderStore[logisticsProviderId].Name);
	return shim.Success([]byte("Successfully retrieved the json data from the ledger"));
}
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"crypto/sha256"
	"encoding/hex"
)

type SmartContract struct{}

type Data struct {
	Message string `json:"message"`
	Hash    string `json:"hash"`
	ID      string `json:"id"`
}

// chaincode initiate
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// invoke functions
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	// setMessage, setHash, getData

	if function == "setData" {
		return s.setData(APIstub, args)
	} else if function == "setHash" {
		return s.setHash(APIstub, args)
	} else if function == "getData" {
		return s.getData(APIstub, args)
	} else {
		fmt.Println("Check function : ", function)
		return shim.Error("Unknown function")
	}
}

// enroll message
func (s *SmartContract) setData(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args : msg, id
	if len(args) != 2 {
		return shim.Error("arguments : message, id")
	}

	// only for RequestOrg
	creator, _ := APIstub.GetCreator()
	if string(creator[2:12]) != "RequestOrg" {
		return shim.Error("setData is only for requestOrg")
	}

	// check stored id
	IDbytes, _ := APIstub.GetState(args[1])
	if string(IDbytes[:]) != "" {
		return shim.Error("there is the id")
	}
	
	var data = Data{Message: args[0], Hash: "", ID: args[1]}

	// data to json
	DataJSONBytes, _ := json.Marshal(data)
	err := APIstub.PutState(data.ID, DataJSONBytes)

	if err != nil {
		return shim.Error("Fail to create data : " + data.ID)
	}

	return shim.Success(nil)
}

// set hash of message
func (s *SmartContract) setHash(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args : id
	if len(args) != 1 {
		return shim.Error("arguments : id")
	}

	// only for HashingOrg
	creator, _ := APIstub.GetCreator()
	if string(creator[2:12]) != "HashingOrg" {
		return shim.Error("setHash is only for HashingOrg")
	}

	dataBytes, err := APIstub.GetState(args[0])

	if err != nil {
		fmt.Println("Fail to load data : " + err.Error())
		return shim.Error(err.Error())
	}

	// check stored id
	if string(dataBytes[:]) == "" {
		return shim.Error("there isn't the id")
	}

	// json to data
	data := Data{}
	json.Unmarshal(dataBytes, &data)

	msg := data.Message

	// msg hashing to string
	msgBytes := []byte(msg)
	Hash := sha256.Sum256(msgBytes)
	strHash := hex.EncodeToString(Hash[:])

	data.Hash = strHash

	dataBytes, _ = json.Marshal(data)
	err2 := APIstub.PutState(args[0], dataBytes)

	if err2 != nil {
		return shim.Error("Fail to hashing : " + args[0])
	}

	return shim.Success(nil)
}

// get data : format = bytes
func (s *SmartContract) getData(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args : id
	if len(args) != 1 {
		return shim.Error("arguments : id")
	}

	// getData is for all Org

	dataBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println("Fail to load data : " + err.Error())
		return shim.Error(err.Error())
	}

	data := Data{}
	json.Unmarshal(dataBytes, &data)

	// bytes like json
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}

	buffer.WriteString("{\"Message\":")
	buffer.WriteString("\"")
	buffer.WriteString(data.Message)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Hash\":")
	buffer.WriteString("\"")
	buffer.WriteString(data.Hash)
	buffer.WriteString("\"")

	buffer.WriteString(", \"ID\":")
	buffer.WriteString("\"")
	buffer.WriteString(data.ID)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Println("Error starting chaincode : ", err.Error())
	}
}
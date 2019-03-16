import (
  "bytes"
  "encoding/json"
  "fmt"
  "time"
  "strconv"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/core/chaincode/lib/cid"
  sc "github.com/hyperledger/fabric/protos/peer"
)

type LicenseAsset struct{
}

type License struct{
    license_key string `json:"license_key"`
    CreatedAt time.Time `json:"created_at"`
    license_owner string `json:"license_owner"`
    license_user string `json:"license_user"`
    license_val string `json:"license_val"`
}

type License_Vendor struct{
    license_key string `json:"license_key"`
    license_type string `json:"license_type"`
    vendor_id string `json:"vendor_id"`
    service_provider_id string `json:"service_provider_id"`
    status string `json:"status"`
    timestamp string `json:"timestamp"`
    CreatedAt time.Time `json:"created_at"`
}

type License_SP struct{
    license_key string `json:"license_key"`
    license_type string `json:"license_type"` // l1,l2,l3
    vendor_id string `json:"vendor_id"` 
    service_provider_id string `json:"service_provider_id"`
    client_id string `json:"client_id"`
    service_type string `json:"service_type"` // s1,s2,s3
    service_id string `json:"service_id"`
    status string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

type License_Client struct{
    license_key string `json:"license_key"`
    service_provider_id string `json:"service_provider_id"`
    service_type string `json:"service_type"` // s1,s2,s3
    service_id string `json:"service_id"`
    vendor_id string `json:"vendor_id"`
    status string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

func main(){
    err := shim.Start(new(LicenseAsset))
    if err != nil {
        fmt.Printf("Error starting LicenseAsset: %s", err)
    }
}

func (s *LicenseAsset) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

//TODO : Invoke function is incomplete, Add functions to invoke.
/*
func (s *LicenseAsset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    if function == ""{
        return s.func_name(stub, args)
    } else if function == ""{
        return s.func_name(stub, args)
    }
    ...
}*/

func (s *LicenseAsset) initLicense(stub shim.ChaincodeStubInterface) pb.Response {
    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments")
    }

    License_key = args[0]
    license_type = args[1]

    licenseAsBytes, err := stub.GetState(License_type)
    if err != nil {
        return shim.Error("Failed to get License: " + err.Error())
    } else if licenseAsBytes != nil {
        fmt.Println("This License already exists: " + license_key)
        return shim.Error("This License already exists: " + license_key)
    }

    objectType := "License"
    License := &License{objectType, License_type, license_key}
    LicenseJSONasBytes, err := json.Marshal(License)
    if err != nil {
        return shim.Error(err.Error())
    }

    err = stub.PutState(License_type, )

}

func (s *LicenseAsset) ShowLicense(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var license_key, jsonResp string
    var err error

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting key of the License to query")
    }

    license_key = args[0]
    valAsbytes, err := stub.GetState(license_key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + license_key + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"Error\":\"License does not exist: " + license_key + "\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}

func (s *LicenseAsset) DeleteLicense(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var jsonResp string
    var LicenseJSON License
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments.")
    }
    License_key = args[0]

    valAsbytes, err := stub.GetState(License_key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + License_key + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"Error\":\"License does not exist: " + License_key + "\"}"
        return shim.Error(jsonResp)
    }

    err = json.Unmarshal([]byte(valAsbytes), &LicenseJSON)
    if err != nil {        
        jsonResp = "{\"Error\":\"Failed to decode JSON of: " + License_key + "\"}"
        return shim.Error(jsonResp)
    }

    err = stub.DelState(License_key) //remove the marble from chaincode state
    if err != nil {
        return shim.Error("Failed to delete state:" + err.Error())
    }

    indexName := "license_val"
    LicenseValIndex, err := stub.CreateCompositeKey(indexName, []string{LicenseJSON.license_val, LicenseJSON.license_key})
    if err != nil {
        return shim.Error(err.Error())
    }

    err = stub.DelState(LicenseValIndex)
    if err != nil {
        return shim.Error("Failed to delete state: " + err.Error())
    }
    return shim.Success(nil)
}

func (s *LicenseAsset) AssignLicense(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 4 {
        return shim.Error("Incorrect number of arguments.")
    } 

    License_key := args[0]
    license_type := args[1]
    vendor_id := args[2]
    service_provider_id := args[3]

    fmt.Println("- License Transfer initiate.", vendor_id, service_provider_id)

    licenseAsBytes, err := stub.GetState(License_key)
    if err != nil {
        return shim.Error("Failed license transfer." + err.Error())
    } else if licenseAsBytes == nil {
        return shim.Error("License does not exist")
    }

    licenseToTransfer := License{}
    err := json.Unmarshal(licenseAsBytes, &licenseToTransfer)
    if err != nil {
        return shim.Error(err.Error())
    }
    licenseToTransfer.license_user = service_provider_id

    LicenseJSONasBytes, _ := json.Marshal(licenseToTransfer)
    err = stub.PutState(License_key, LicenseJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("- end License transfer.")
    return shim.Success(nil)

}

func (s *LicenseAsset) initService(stub shim.ChaincodeStubInterface) pb.Response {

}

func (s *LicenseAsset) AssignService(stub shim.ChaincodeStubInterface) pb.Response {
    

    License_key = args[0]
    License_type = args[1]

    if License_type == 'L1'{
        service_type = 'S1'
    } else if License_type == 'L2'{
        service_type = 'S2'
    } else if License_type == 'L3'{
        service_type = 'S3'
    }

}



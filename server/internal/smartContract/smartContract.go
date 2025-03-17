// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartContract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SmartContractMetaData contains all meta data concerning the SmartContract contract.
var SmartContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceCurr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"targetCurr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"}],\"name\":\"TradeExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"addCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"currencies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceCurr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetCurr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"}],\"name\":\"executeTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currencyName\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currencyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f600255348015610013575f5ffd5b5061005c6040518060400160405280600481526020017f4d454d4500000000000000000000000000000000000000000000000000000000815250620f424061013960201b60201c565b6100a46040518060400160405280600381526020017f4554480000000000000000000000000000000000000000000000000000000000815250620f424061013960201b60201c565b6100ec6040518060400160405280600381526020017f4254430000000000000000000000000000000000000000000000000000000000815250620f424061013960201b60201c565b6101346040518060400160405280600481526020017f5553445400000000000000000000000000000000000000000000000000000000815250620f424061013960201b60201c565b610607565b5f826040516101489190610277565b90815260200160405180910390206002015f9054906101000a900460ff16156101a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019d906102e7565b60405180910390fd5b6040518060600160405280838152602001828152602001600115158152505f836040516101d39190610277565b90815260200160405180910390205f820151815f0190816101f49190610538565b50602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055509050505050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61025182610225565b61025b818561022f565b935061026b818560208601610239565b80840191505092915050565b5f6102828284610247565b915081905092915050565b5f82825260208201905092915050565b7f43757272656e637920616c7265616479206578697374730000000000000000005f82015250565b5f6102d160178361028d565b91506102dc8261029d565b602082019050919050565b5f6020820190508181035f8301526102fe816102c5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061037657607f821691505b60208210810361038957610388610332565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103eb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103b0565b6103f586836103b0565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61043961043461042f8461040d565b610416565b61040d565b9050919050565b5f819050919050565b6104528361041f565b61046661045e82610440565b8484546103bc565b825550505050565b5f5f905090565b61047d61046e565b610488818484610449565b505050565b5b818110156104ab576104a05f82610475565b60018101905061048e565b5050565b601f8211156104f0576104c18161038f565b6104ca846103a1565b810160208510156104d9578190505b6104ed6104e5856103a1565b83018261048d565b50505b505050565b5f82821c905092915050565b5f6105105f19846008026104f5565b1980831691505092915050565b5f6105288383610501565b9150826002028217905092915050565b61054182610225565b67ffffffffffffffff81111561055a57610559610305565b5b610564825461035f565b61056f8282856104af565b5f60209050601f8311600181146105a0575f841561058e578287015190505b610598858261051d565b8655506105ff565b601f1984166105ae8661038f565b5f5b828110156105d5578489015182556001820191506020850194506020810190506105b0565b868310156105f257848901516105ee601f891682610501565b8355505b6001600288020188555050505b505050505050565b61152a806106145f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c806305d01f50146100645780631dd7cecf146100945780635c49c825146100c45780638a9051db146100e0578063bdadad5214610110578063c8f5bc061461012c575b5f5ffd5b61007e60048036038101906100799190610add565b61015e565b60405161008b9190610b4f565b60405180910390f35b6100ae60048036038101906100a99190610add565b610196565b6040516100bb9190610b4f565b60405180910390f35b6100de60048036038101906100d99190610b92565b610265565b005b6100fa60048036038101906100f59190610bec565b610351565b6040516101079190610b4f565b60405180910390f35b61012a60048036038101906101259190610c9b565b610787565b005b61014660048036038101906101419190610d07565b610867565b60405161015593929190610dc8565b60405180910390f35b6001602052815f5260405f20818051602081018201805184825260208301602085012081835280955050505050505f91509150505481565b5f5f826040516101a69190610e3e565b90815260200160405180910390206002015f9054906101000a900460ff16610203576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fa90610e9e565b60405180910390fd5b60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208260405161024e9190610e3e565b908152602001604051809103902054905092915050565b5f826040516102749190610e3e565b90815260200160405180910390206002015f9054906101000a900460ff16156102d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c990610f06565b60405180910390fd5b6040518060600160405280838152602001828152602001600115158152505f836040516102ff9190610e3e565b90815260200160405180910390205f820151815f0190816103209190611121565b50602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055509050505050565b5f5f856040516103619190610e3e565b90815260200160405180910390206002015f9054906101000a900460ff166103be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b59061123a565b60405180910390fd5b5f846040516103cd9190610e3e565b90815260200160405180910390206002015f9054906101000a900460ff1661042a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610421906112a2565b60405180910390fd5b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20866040516104769190610e3e565b90815260200160405180910390205410156104c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bd9061130a565b60405180910390fd5b8160015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20856040516105129190610e3e565b9081526020016040518091039020541015610562576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055990611372565b60405180910390fd5b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20866040516105ae9190610e3e565b90815260200160405180910390205f8282546105ca91906113bd565b925050819055508160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208560405161061d9190610e3e565b90815260200160405180910390205f82825461063991906113f0565b925050819055508260015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208660405161068c9190610e3e565b90815260200160405180910390205f8282546106a891906113f0565b925050819055508160015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20856040516106fb9190610e3e565b90815260200160405180910390205f82825461071791906113bd565b925050819055505f60025f81548092919061073190611423565b9190505590507fb843df2016ffb5d0c50535c12fe486a11ee928b6dd2b0487c9a815f31c673e5a813389898989896040516107729796959493929190611479565b60405180910390a18091505095945050505050565b5f826040516107969190610e3e565b90815260200160405180910390206002015f9054906101000a900460ff166107f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ea90610e9e565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208360405161083f9190610e3e565b90815260200160405180910390205f82825461085b91906113f0565b92505081905550505050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461089d90610f51565b80601f01602080910402602001604051908101604052809291908181526020018280546108c990610f51565b80156109145780601f106108eb57610100808354040283529160200191610914565b820191905f5260205f20905b8154815290600101906020018083116108f757829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff16905083565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61097082610947565b9050919050565b61098081610966565b811461098a575f5ffd5b50565b5f8135905061099b81610977565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6109ef826109a9565b810181811067ffffffffffffffff82111715610a0e57610a0d6109b9565b5b80604052505050565b5f610a20610936565b9050610a2c82826109e6565b919050565b5f67ffffffffffffffff821115610a4b57610a4a6109b9565b5b610a54826109a9565b9050602081019050919050565b828183375f83830152505050565b5f610a81610a7c84610a31565b610a17565b905082815260208101848484011115610a9d57610a9c6109a5565b5b610aa8848285610a61565b509392505050565b5f82601f830112610ac457610ac36109a1565b5b8135610ad4848260208601610a6f565b91505092915050565b5f5f60408385031215610af357610af261093f565b5b5f610b008582860161098d565b925050602083013567ffffffffffffffff811115610b2157610b20610943565b5b610b2d85828601610ab0565b9150509250929050565b5f819050919050565b610b4981610b37565b82525050565b5f602082019050610b625f830184610b40565b92915050565b610b7181610b37565b8114610b7b575f5ffd5b50565b5f81359050610b8c81610b68565b92915050565b5f5f60408385031215610ba857610ba761093f565b5b5f83013567ffffffffffffffff811115610bc557610bc4610943565b5b610bd185828601610ab0565b9250506020610be285828601610b7e565b9150509250929050565b5f5f5f5f5f60a08688031215610c0557610c0461093f565b5b5f610c128882890161098d565b955050602086013567ffffffffffffffff811115610c3357610c32610943565b5b610c3f88828901610ab0565b945050604086013567ffffffffffffffff811115610c6057610c5f610943565b5b610c6c88828901610ab0565b9350506060610c7d88828901610b7e565b9250506080610c8e88828901610b7e565b9150509295509295909350565b5f5f5f60608486031215610cb257610cb161093f565b5b5f610cbf8682870161098d565b935050602084013567ffffffffffffffff811115610ce057610cdf610943565b5b610cec86828701610ab0565b9250506040610cfd86828701610b7e565b9150509250925092565b5f60208284031215610d1c57610d1b61093f565b5b5f82013567ffffffffffffffff811115610d3957610d38610943565b5b610d4584828501610ab0565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610d8082610d4e565b610d8a8185610d58565b9350610d9a818560208601610d68565b610da3816109a9565b840191505092915050565b5f8115159050919050565b610dc281610dae565b82525050565b5f6060820190508181035f830152610de08186610d76565b9050610def6020830185610b40565b610dfc6040830184610db9565b949350505050565b5f81905092915050565b5f610e1882610d4e565b610e228185610e04565b9350610e32818560208601610d68565b80840191505092915050565b5f610e498284610e0e565b915081905092915050565b7f43757272656e637920646f6573206e6f742065786973740000000000000000005f82015250565b5f610e88601783610d58565b9150610e9382610e54565b602082019050919050565b5f6020820190508181035f830152610eb581610e7c565b9050919050565b7f43757272656e637920616c7265616479206578697374730000000000000000005f82015250565b5f610ef0601783610d58565b9150610efb82610ebc565b602082019050919050565b5f6020820190508181035f830152610f1d81610ee4565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f6857607f821691505b602082108103610f7b57610f7a610f24565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610fdd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fa2565b610fe78683610fa2565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61102261101d61101884610b37565b610fff565b610b37565b9050919050565b5f819050919050565b61103b83611008565b61104f61104782611029565b848454610fae565b825550505050565b5f5f905090565b611066611057565b611071818484611032565b505050565b5b81811015611094576110895f8261105e565b600181019050611077565b5050565b601f8211156110d9576110aa81610f81565b6110b384610f93565b810160208510156110c2578190505b6110d66110ce85610f93565b830182611076565b50505b505050565b5f82821c905092915050565b5f6110f95f19846008026110de565b1980831691505092915050565b5f61111183836110ea565b9150826002028217905092915050565b61112a82610d4e565b67ffffffffffffffff811115611143576111426109b9565b5b61114d8254610f51565b611158828285611098565b5f60209050601f831160018114611189575f8415611177578287015190505b6111818582611106565b8655506111e8565b601f19841661119786610f81565b5f5b828110156111be57848901518255600182019150602085019450602081019050611199565b868310156111db57848901516111d7601f8916826110ea565b8355505b6001600288020188555050505b505050505050565b7f536f757263652063757272656e637920646f6573206e6f7420657869737400005f82015250565b5f611224601e83610d58565b915061122f826111f0565b602082019050919050565b5f6020820190508181035f83015261125181611218565b9050919050565b7f5461726765742063757272656e637920646f6573206e6f7420657869737400005f82015250565b5f61128c601e83610d58565b915061129782611258565b602082019050919050565b5f6020820190508181035f8301526112b981611280565b9050919050565b7f496e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f6112f4601483610d58565b91506112ff826112c0565b602082019050919050565b5f6020820190508181035f830152611321816112e8565b9050919050565b7f5461726765742068617320696e73756666696369656e742062616c616e6365005f82015250565b5f61135c601f83610d58565b915061136782611328565b602082019050919050565b5f6020820190508181035f83015261138981611350565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6113c782610b37565b91506113d283610b37565b92508282039050818111156113ea576113e9611390565b5b92915050565b5f6113fa82610b37565b915061140583610b37565b925082820190508082111561141d5761141c611390565b5b92915050565b5f61142d82610b37565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361145f5761145e611390565b5b600182019050919050565b61147381610966565b82525050565b5f60e08201905061148c5f83018a610b40565b611499602083018961146a565b6114a6604083018861146a565b81810360608301526114b88187610d76565b905081810360808301526114cc8186610d76565b90506114db60a0830185610b40565b6114e860c0830184610b40565b9897505050505050505056fea26469706673582212205e7a9469b1000d227a1df3d86e35f253c025964b611bdf0e3c49b2898e60908c64736f6c634300081d0033",
}

// SmartContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartContractMetaData.ABI instead.
var SmartContractABI = SmartContractMetaData.ABI

// SmartContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SmartContractMetaData.Bin instead.
var SmartContractBin = SmartContractMetaData.Bin

// DeploySmartContract deploys a new Ethereum contract, binding an instance of SmartContract to it.
func DeploySmartContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SmartContract, error) {
	parsed, err := SmartContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SmartContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartContract{SmartContractCaller: SmartContractCaller{contract: contract}, SmartContractTransactor: SmartContractTransactor{contract: contract}, SmartContractFilterer: SmartContractFilterer{contract: contract}}, nil
}

// SmartContract is an auto generated Go binding around an Ethereum contract.
type SmartContract struct {
	SmartContractCaller     // Read-only binding to the contract
	SmartContractTransactor // Write-only binding to the contract
	SmartContractFilterer   // Log filterer for contract events
}

// SmartContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartContractSession struct {
	Contract     *SmartContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartContractCallerSession struct {
	Contract *SmartContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SmartContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartContractTransactorSession struct {
	Contract     *SmartContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SmartContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartContractRaw struct {
	Contract *SmartContract // Generic contract binding to access the raw methods on
}

// SmartContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartContractCallerRaw struct {
	Contract *SmartContractCaller // Generic read-only contract binding to access the raw methods on
}

// SmartContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartContractTransactorRaw struct {
	Contract *SmartContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartContract creates a new instance of SmartContract, bound to a specific deployed contract.
func NewSmartContract(address common.Address, backend bind.ContractBackend) (*SmartContract, error) {
	contract, err := bindSmartContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartContract{SmartContractCaller: SmartContractCaller{contract: contract}, SmartContractTransactor: SmartContractTransactor{contract: contract}, SmartContractFilterer: SmartContractFilterer{contract: contract}}, nil
}

// NewSmartContractCaller creates a new read-only instance of SmartContract, bound to a specific deployed contract.
func NewSmartContractCaller(address common.Address, caller bind.ContractCaller) (*SmartContractCaller, error) {
	contract, err := bindSmartContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartContractCaller{contract: contract}, nil
}

// NewSmartContractTransactor creates a new write-only instance of SmartContract, bound to a specific deployed contract.
func NewSmartContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartContractTransactor, error) {
	contract, err := bindSmartContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartContractTransactor{contract: contract}, nil
}

// NewSmartContractFilterer creates a new log filterer instance of SmartContract, bound to a specific deployed contract.
func NewSmartContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartContractFilterer, error) {
	contract, err := bindSmartContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartContractFilterer{contract: contract}, nil
}

// bindSmartContract binds a generic wrapper to an already deployed contract.
func bindSmartContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartContract *SmartContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartContract.Contract.SmartContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartContract *SmartContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartContract.Contract.SmartContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartContract *SmartContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartContract.Contract.SmartContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartContract *SmartContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartContract *SmartContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartContract *SmartContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartContract.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_SmartContract *SmartContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _SmartContract.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_SmartContract *SmartContractSession) Balances(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _SmartContract.Contract.Balances(&_SmartContract.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_SmartContract *SmartContractCallerSession) Balances(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _SmartContract.Contract.Balances(&_SmartContract.CallOpts, arg0, arg1)
}

// Currencies is a free data retrieval call binding the contract method 0xc8f5bc06.
//
// Solidity: function currencies(string ) view returns(string name, uint256 totalSupply, bool exists)
func (_SmartContract *SmartContractCaller) Currencies(opts *bind.CallOpts, arg0 string) (struct {
	Name        string
	TotalSupply *big.Int
	Exists      bool
}, error) {
	var out []interface{}
	err := _SmartContract.contract.Call(opts, &out, "currencies", arg0)

	outstruct := new(struct {
		Name        string
		TotalSupply *big.Int
		Exists      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TotalSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Currencies is a free data retrieval call binding the contract method 0xc8f5bc06.
//
// Solidity: function currencies(string ) view returns(string name, uint256 totalSupply, bool exists)
func (_SmartContract *SmartContractSession) Currencies(arg0 string) (struct {
	Name        string
	TotalSupply *big.Int
	Exists      bool
}, error) {
	return _SmartContract.Contract.Currencies(&_SmartContract.CallOpts, arg0)
}

// Currencies is a free data retrieval call binding the contract method 0xc8f5bc06.
//
// Solidity: function currencies(string ) view returns(string name, uint256 totalSupply, bool exists)
func (_SmartContract *SmartContractCallerSession) Currencies(arg0 string) (struct {
	Name        string
	TotalSupply *big.Int
	Exists      bool
}, error) {
	return _SmartContract.Contract.Currencies(&_SmartContract.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x1dd7cecf.
//
// Solidity: function getBalance(address account, string currencyName) view returns(uint256)
func (_SmartContract *SmartContractCaller) GetBalance(opts *bind.CallOpts, account common.Address, currencyName string) (*big.Int, error) {
	var out []interface{}
	err := _SmartContract.contract.Call(opts, &out, "getBalance", account, currencyName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x1dd7cecf.
//
// Solidity: function getBalance(address account, string currencyName) view returns(uint256)
func (_SmartContract *SmartContractSession) GetBalance(account common.Address, currencyName string) (*big.Int, error) {
	return _SmartContract.Contract.GetBalance(&_SmartContract.CallOpts, account, currencyName)
}

// GetBalance is a free data retrieval call binding the contract method 0x1dd7cecf.
//
// Solidity: function getBalance(address account, string currencyName) view returns(uint256)
func (_SmartContract *SmartContractCallerSession) GetBalance(account common.Address, currencyName string) (*big.Int, error) {
	return _SmartContract.Contract.GetBalance(&_SmartContract.CallOpts, account, currencyName)
}

// AddCurrency is a paid mutator transaction binding the contract method 0x5c49c825.
//
// Solidity: function addCurrency(string name, uint256 initialSupply) returns()
func (_SmartContract *SmartContractTransactor) AddCurrency(opts *bind.TransactOpts, name string, initialSupply *big.Int) (*types.Transaction, error) {
	return _SmartContract.contract.Transact(opts, "addCurrency", name, initialSupply)
}

// AddCurrency is a paid mutator transaction binding the contract method 0x5c49c825.
//
// Solidity: function addCurrency(string name, uint256 initialSupply) returns()
func (_SmartContract *SmartContractSession) AddCurrency(name string, initialSupply *big.Int) (*types.Transaction, error) {
	return _SmartContract.Contract.AddCurrency(&_SmartContract.TransactOpts, name, initialSupply)
}

// AddCurrency is a paid mutator transaction binding the contract method 0x5c49c825.
//
// Solidity: function addCurrency(string name, uint256 initialSupply) returns()
func (_SmartContract *SmartContractTransactorSession) AddCurrency(name string, initialSupply *big.Int) (*types.Transaction, error) {
	return _SmartContract.Contract.AddCurrency(&_SmartContract.TransactOpts, name, initialSupply)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x8a9051db.
//
// Solidity: function executeTrade(address targetAddress, string sourceCurr, string targetCurr, uint256 sourceAmount, uint256 targetAmount) returns(uint256)
func (_SmartContract *SmartContractTransactor) ExecuteTrade(opts *bind.TransactOpts, targetAddress common.Address, sourceCurr string, targetCurr string, sourceAmount *big.Int, targetAmount *big.Int) (*types.Transaction, error) {
	return _SmartContract.contract.Transact(opts, "executeTrade", targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x8a9051db.
//
// Solidity: function executeTrade(address targetAddress, string sourceCurr, string targetCurr, uint256 sourceAmount, uint256 targetAmount) returns(uint256)
func (_SmartContract *SmartContractSession) ExecuteTrade(targetAddress common.Address, sourceCurr string, targetCurr string, sourceAmount *big.Int, targetAmount *big.Int) (*types.Transaction, error) {
	return _SmartContract.Contract.ExecuteTrade(&_SmartContract.TransactOpts, targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0x8a9051db.
//
// Solidity: function executeTrade(address targetAddress, string sourceCurr, string targetCurr, uint256 sourceAmount, uint256 targetAmount) returns(uint256)
func (_SmartContract *SmartContractTransactorSession) ExecuteTrade(targetAddress common.Address, sourceCurr string, targetCurr string, sourceAmount *big.Int, targetAmount *big.Int) (*types.Transaction, error) {
	return _SmartContract.Contract.ExecuteTrade(&_SmartContract.TransactOpts, targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xbdadad52.
//
// Solidity: function mintTokens(address to, string currencyName, uint256 amount) returns()
func (_SmartContract *SmartContractTransactor) MintTokens(opts *bind.TransactOpts, to common.Address, currencyName string, amount *big.Int) (*types.Transaction, error) {
	return _SmartContract.contract.Transact(opts, "mintTokens", to, currencyName, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xbdadad52.
//
// Solidity: function mintTokens(address to, string currencyName, uint256 amount) returns()
func (_SmartContract *SmartContractSession) MintTokens(to common.Address, currencyName string, amount *big.Int) (*types.Transaction, error) {
	return _SmartContract.Contract.MintTokens(&_SmartContract.TransactOpts, to, currencyName, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0xbdadad52.
//
// Solidity: function mintTokens(address to, string currencyName, uint256 amount) returns()
func (_SmartContract *SmartContractTransactorSession) MintTokens(to common.Address, currencyName string, amount *big.Int) (*types.Transaction, error) {
	return _SmartContract.Contract.MintTokens(&_SmartContract.TransactOpts, to, currencyName, amount)
}

// SmartContractTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the SmartContract contract.
type SmartContractTradeExecutedIterator struct {
	Event *SmartContractTradeExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartContractTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartContractTradeExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartContractTradeExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartContractTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartContractTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartContractTradeExecuted represents a TradeExecuted event raised by the SmartContract contract.
type SmartContractTradeExecuted struct {
	Id            *big.Int
	SourceAddress common.Address
	TargetAddress common.Address
	SourceCurr    string
	TargetCurr    string
	SourceAmount  *big.Int
	TargetAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTradeExecuted is a free log retrieval operation binding the contract event 0xb843df2016ffb5d0c50535c12fe486a11ee928b6dd2b0487c9a815f31c673e5a.
//
// Solidity: event TradeExecuted(uint256 id, address sourceAddress, address targetAddress, string sourceCurr, string targetCurr, uint256 sourceAmount, uint256 targetAmount)
func (_SmartContract *SmartContractFilterer) FilterTradeExecuted(opts *bind.FilterOpts) (*SmartContractTradeExecutedIterator, error) {

	logs, sub, err := _SmartContract.contract.FilterLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return &SmartContractTradeExecutedIterator{contract: _SmartContract.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0xb843df2016ffb5d0c50535c12fe486a11ee928b6dd2b0487c9a815f31c673e5a.
//
// Solidity: event TradeExecuted(uint256 id, address sourceAddress, address targetAddress, string sourceCurr, string targetCurr, uint256 sourceAmount, uint256 targetAmount)
func (_SmartContract *SmartContractFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *SmartContractTradeExecuted) (event.Subscription, error) {

	logs, sub, err := _SmartContract.contract.WatchLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartContractTradeExecuted)
				if err := _SmartContract.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradeExecuted is a log parse operation binding the contract event 0xb843df2016ffb5d0c50535c12fe486a11ee928b6dd2b0487c9a815f31c673e5a.
//
// Solidity: event TradeExecuted(uint256 id, address sourceAddress, address targetAddress, string sourceCurr, string targetCurr, uint256 sourceAmount, uint256 targetAmount)
func (_SmartContract *SmartContractFilterer) ParseTradeExecuted(log types.Log) (*SmartContractTradeExecuted, error) {
	event := new(SmartContractTradeExecuted)
	if err := _SmartContract.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceCurr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"targetCurr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"}],\"name\":\"TradeExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"addCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"currencies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceCurr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"targetCurr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"}],\"name\":\"executeTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currencyName\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"currencyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f6002556008600355348015610018575f5ffd5b5061007a6040518060400160405280600381526020017f4254430000000000000000000000000000000000000000000000000000000000815250600354600a6100619190610454565b620f424061006f919061049e565b61020360201b60201c565b6100db6040518060400160405280600481526020017f444f474500000000000000000000000000000000000000000000000000000000815250600354600a6100c29190610454565b620f42406100d0919061049e565b61020360201b60201c565b61013c6040518060400160405280600481526020017f5348494200000000000000000000000000000000000000000000000000000000815250600354600a6101239190610454565b620f4240610131919061049e565b61020360201b60201c565b61019d6040518060400160405280600481526020017f424f4e4b00000000000000000000000000000000000000000000000000000000815250600354600a6101849190610454565b620f4240610192919061049e565b61020360201b60201c565b6101fe6040518060400160405280600481526020017f5045504500000000000000000000000000000000000000000000000000000000815250600354600a6101e59190610454565b620f42406101f3919061049e565b61020360201b60201c565b6108b8565b5f826040516102129190610531565b90815260200160405180910390206002015f9054906101000a900460ff1615610270576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610267906105a1565b60405180910390fd5b6040518060600160405280838152602001828152602001600115158152505f8360405161029d9190610531565b90815260200160405180910390205f820151815f0190816102be91906107e9565b50602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055509050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156103715780860481111561034d5761034c6102ef565b5b600185161561035c5780820291505b808102905061036a8561031c565b9450610331565b94509492505050565b5f826103895760019050610444565b81610396575f9050610444565b81600181146103ac57600281146103b6576103e5565b6001915050610444565b60ff8411156103c8576103c76102ef565b5b8360020a9150848211156103df576103de6102ef565b5b50610444565b5060208310610133831016604e8410600b841016171561041a5782820a905083811115610415576104146102ef565b5b610444565b6104278484846001610328565b9250905081840481111561043e5761043d6102ef565b5b81810290505b9392505050565b5f819050919050565b5f61045e8261044b565b91506104698361044b565b92506104967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461037a565b905092915050565b5f6104a88261044b565b91506104b38361044b565b92508282026104c18161044b565b915082820484148315176104d8576104d76102ef565b5b5092915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61050b826104df565b61051581856104e9565b93506105258185602086016104f3565b80840191505092915050565b5f61053c8284610501565b915081905092915050565b5f82825260208201905092915050565b7f43757272656e637920616c7265616479206578697374730000000000000000005f82015250565b5f61058b601783610547565b915061059682610557565b602082019050919050565b5f6020820190508181035f8301526105b88161057f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061063057607f821691505b602082108103610643576106426105ec565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261066a565b6106af868361066a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6106ea6106e56106e08461044b565b6106c7565b61044b565b9050919050565b5f819050919050565b610703836106d0565b61071761070f826106f1565b848454610676565b825550505050565b5f5f905090565b61072e61071f565b6107398184846106fa565b505050565b5b8181101561075c576107515f82610726565b60018101905061073f565b5050565b601f8211156107a15761077281610649565b61077b8461065b565b8101602085101561078a578190505b61079e6107968561065b565b83018261073e565b50505b505050565b5f82821c905092915050565b5f6107c15f19846008026107a6565b1980831691505092915050565b5f6107d983836107b2565b9150826002028217905092915050565b6107f2826104df565b67ffffffffffffffff81111561080b5761080a6105bf565b5b6108158254610619565b610820828285610760565b5f60209050601f831160018114610851575f841561083f578287015190505b61084985826107ce565b8655506108b0565b601f19841661085f86610649565b5f5b8281101561088657848901518255600182019150602085019450602081019050610861565b868310156108a3578489015161089f601f8916826107b2565b8355505b6001600288020188555050505b505050505050565b611569806108c55f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c80635c49c825116100595780635c49c825146100fd5780638a9051db14610119578063bdadad5214610149578063c8f5bc06146101655761007b565b806305d01f501461007f5780631dd7cecf146100af578063313ce567146100df575b5f5ffd5b61009960048036038101906100949190610b1c565b610197565b6040516100a69190610b8e565b60405180910390f35b6100c960048036038101906100c49190610b1c565b6101cf565b6040516100d69190610b8e565b60405180910390f35b6100e761029e565b6040516100f49190610b8e565b60405180910390f35b61011760048036038101906101129190610bd1565b6102a4565b005b610133600480360381019061012e9190610c2b565b610390565b6040516101409190610b8e565b60405180910390f35b610163600480360381019061015e9190610cda565b6107c6565b005b61017f600480360381019061017a9190610d46565b6108a6565b60405161018e93929190610e07565b60405180910390f35b6001602052815f5260405f20818051602081018201805184825260208301602085012081835280955050505050505f91509150505481565b5f5f826040516101df9190610e7d565b90815260200160405180910390206002015f9054906101000a900460ff1661023c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023390610edd565b60405180910390fd5b60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20826040516102879190610e7d565b908152602001604051809103902054905092915050565b60035481565b5f826040516102b39190610e7d565b90815260200160405180910390206002015f9054906101000a900460ff1615610311576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030890610f45565b60405180910390fd5b6040518060600160405280838152602001828152602001600115158152505f8360405161033e9190610e7d565b90815260200160405180910390205f820151815f01908161035f9190611160565b50602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055509050505050565b5f5f856040516103a09190610e7d565b90815260200160405180910390206002015f9054906101000a900460ff166103fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f490611279565b60405180910390fd5b5f8460405161040c9190610e7d565b90815260200160405180910390206002015f9054906101000a900460ff16610469576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610460906112e1565b60405180910390fd5b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20866040516104b59190610e7d565b9081526020016040518091039020541015610505576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fc90611349565b60405180910390fd5b8160015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20856040516105519190610e7d565b90815260200160405180910390205410156105a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610598906113b1565b60405180910390fd5b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20866040516105ed9190610e7d565b90815260200160405180910390205f82825461060991906113fc565b925050819055508160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208560405161065c9190610e7d565b90815260200160405180910390205f828254610678919061142f565b925050819055508260015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20866040516106cb9190610e7d565b90815260200160405180910390205f8282546106e7919061142f565b925050819055508160015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208560405161073a9190610e7d565b90815260200160405180910390205f82825461075691906113fc565b925050819055505f60025f81548092919061077090611462565b9190505590507fb843df2016ffb5d0c50535c12fe486a11ee928b6dd2b0487c9a815f31c673e5a813389898989896040516107b197969594939291906114b8565b60405180910390a18091505095945050505050565b5f826040516107d59190610e7d565b90815260200160405180910390206002015f9054906101000a900460ff16610832576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082990610edd565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208360405161087e9190610e7d565b90815260200160405180910390205f82825461089a919061142f565b92505081905550505050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0180546108dc90610f90565b80601f016020809104026020016040519081016040528092919081815260200182805461090890610f90565b80156109535780601f1061092a57610100808354040283529160200191610953565b820191905f5260205f20905b81548152906001019060200180831161093657829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff16905083565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109af82610986565b9050919050565b6109bf816109a5565b81146109c9575f5ffd5b50565b5f813590506109da816109b6565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a2e826109e8565b810181811067ffffffffffffffff82111715610a4d57610a4c6109f8565b5b80604052505050565b5f610a5f610975565b9050610a6b8282610a25565b919050565b5f67ffffffffffffffff821115610a8a57610a896109f8565b5b610a93826109e8565b9050602081019050919050565b828183375f83830152505050565b5f610ac0610abb84610a70565b610a56565b905082815260208101848484011115610adc57610adb6109e4565b5b610ae7848285610aa0565b509392505050565b5f82601f830112610b0357610b026109e0565b5b8135610b13848260208601610aae565b91505092915050565b5f5f60408385031215610b3257610b3161097e565b5b5f610b3f858286016109cc565b925050602083013567ffffffffffffffff811115610b6057610b5f610982565b5b610b6c85828601610aef565b9150509250929050565b5f819050919050565b610b8881610b76565b82525050565b5f602082019050610ba15f830184610b7f565b92915050565b610bb081610b76565b8114610bba575f5ffd5b50565b5f81359050610bcb81610ba7565b92915050565b5f5f60408385031215610be757610be661097e565b5b5f83013567ffffffffffffffff811115610c0457610c03610982565b5b610c1085828601610aef565b9250506020610c2185828601610bbd565b9150509250929050565b5f5f5f5f5f60a08688031215610c4457610c4361097e565b5b5f610c51888289016109cc565b955050602086013567ffffffffffffffff811115610c7257610c71610982565b5b610c7e88828901610aef565b945050604086013567ffffffffffffffff811115610c9f57610c9e610982565b5b610cab88828901610aef565b9350506060610cbc88828901610bbd565b9250506080610ccd88828901610bbd565b9150509295509295909350565b5f5f5f60608486031215610cf157610cf061097e565b5b5f610cfe868287016109cc565b935050602084013567ffffffffffffffff811115610d1f57610d1e610982565b5b610d2b86828701610aef565b9250506040610d3c86828701610bbd565b9150509250925092565b5f60208284031215610d5b57610d5a61097e565b5b5f82013567ffffffffffffffff811115610d7857610d77610982565b5b610d8484828501610aef565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610dbf82610d8d565b610dc98185610d97565b9350610dd9818560208601610da7565b610de2816109e8565b840191505092915050565b5f8115159050919050565b610e0181610ded565b82525050565b5f6060820190508181035f830152610e1f8186610db5565b9050610e2e6020830185610b7f565b610e3b6040830184610df8565b949350505050565b5f81905092915050565b5f610e5782610d8d565b610e618185610e43565b9350610e71818560208601610da7565b80840191505092915050565b5f610e888284610e4d565b915081905092915050565b7f43757272656e637920646f6573206e6f742065786973740000000000000000005f82015250565b5f610ec7601783610d97565b9150610ed282610e93565b602082019050919050565b5f6020820190508181035f830152610ef481610ebb565b9050919050565b7f43757272656e637920616c7265616479206578697374730000000000000000005f82015250565b5f610f2f601783610d97565b9150610f3a82610efb565b602082019050919050565b5f6020820190508181035f830152610f5c81610f23565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610fa757607f821691505b602082108103610fba57610fb9610f63565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261101c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fe1565b6110268683610fe1565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61106161105c61105784610b76565b61103e565b610b76565b9050919050565b5f819050919050565b61107a83611047565b61108e61108682611068565b848454610fed565b825550505050565b5f5f905090565b6110a5611096565b6110b0818484611071565b505050565b5b818110156110d3576110c85f8261109d565b6001810190506110b6565b5050565b601f821115611118576110e981610fc0565b6110f284610fd2565b81016020851015611101578190505b61111561110d85610fd2565b8301826110b5565b50505b505050565b5f82821c905092915050565b5f6111385f198460080261111d565b1980831691505092915050565b5f6111508383611129565b9150826002028217905092915050565b61116982610d8d565b67ffffffffffffffff811115611182576111816109f8565b5b61118c8254610f90565b6111978282856110d7565b5f60209050601f8311600181146111c8575f84156111b6578287015190505b6111c08582611145565b865550611227565b601f1984166111d686610fc0565b5f5b828110156111fd578489015182556001820191506020850194506020810190506111d8565b8683101561121a5784890151611216601f891682611129565b8355505b6001600288020188555050505b505050505050565b7f536f757263652063757272656e637920646f6573206e6f7420657869737400005f82015250565b5f611263601e83610d97565b915061126e8261122f565b602082019050919050565b5f6020820190508181035f83015261129081611257565b9050919050565b7f5461726765742063757272656e637920646f6573206e6f7420657869737400005f82015250565b5f6112cb601e83610d97565b91506112d682611297565b602082019050919050565b5f6020820190508181035f8301526112f8816112bf565b9050919050565b7f496e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f611333601483610d97565b915061133e826112ff565b602082019050919050565b5f6020820190508181035f83015261136081611327565b9050919050565b7f5461726765742068617320696e73756666696369656e742062616c616e6365005f82015250565b5f61139b601f83610d97565b91506113a682611367565b602082019050919050565b5f6020820190508181035f8301526113c88161138f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61140682610b76565b915061141183610b76565b9250828203905081811115611429576114286113cf565b5b92915050565b5f61143982610b76565b915061144483610b76565b925082820190508082111561145c5761145b6113cf565b5b92915050565b5f61146c82610b76565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361149e5761149d6113cf565b5b600182019050919050565b6114b2816109a5565b82525050565b5f60e0820190506114cb5f83018a610b7f565b6114d860208301896114a9565b6114e560408301886114a9565b81810360608301526114f78187610db5565b9050818103608083015261150b8186610db5565b905061151a60a0830185610b7f565b61152760c0830184610b7f565b9897505050505050505056fea26469706673582212209f8f55f34c2810ec8b95987cbca55036a4c70f3c273a14e5f5cfe304eb870fc464736f6c634300081d0033",
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

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_SmartContract *SmartContractCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_SmartContract *SmartContractSession) Decimals() (*big.Int, error) {
	return _SmartContract.Contract.Decimals(&_SmartContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_SmartContract *SmartContractCallerSession) Decimals() (*big.Int, error) {
	return _SmartContract.Contract.Decimals(&_SmartContract.CallOpts)
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

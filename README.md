# Ethereum Go Client

This is a Ethereum compatible Go Client
which implements the 
[Eth JSON RPC Module](https://github.com/ethereum/wiki/wiki/JSON-RPC),
[Personal JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-personal-module) and
[NET JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-net-module#net_version).

## Status

This package is currently under active development. It is not already stable and the infrastructure is not complete and there are still several RPCs left to implement and the API is not stable yet.

## Contribute!

### In Progress = ![](https://placehold.it/15/FFFF00/000000?text=+)

TODO List

- [x] web3_clientVersion                      
- [x] web3_sha3                               
- [x] net_version                             
- [x] net_peerCount                           
- [x] net_listening                           
- [x] eth_protocolVersion                     
- [x] eth_syncing                             
- [x] eth_coinbase                            
- [x] eth_mining                              
- [x] eth_hashrate                            
- [x] eth_gasPrice                            
- [x] eth_accounts                            
- [x] eth_blockNumber                         
- [x] eth_getBalance                          
- [x] eth_getStorageAt (deprecated)
- [ ] eth_getTransactionCount                 
- [ ] eth_getBlockTransactionCountByHash      
- [ ] eth_getBlockTransactionCountByNumber    
- [ ] eth_getUncleCountByBlockHash            
- [ ] eth_getUncleCountByBlockNumber          
- [ ] eth_getCode                             
- [ ] eth_sign                                
- [x] eth_sendTransaction                     
- [ ] eth_sendRawTransaction                  
- [ ] eth_call                                
- [x] eth_estimateGas                         
- [ ] eth_getBlockByHash                      
- [ ] ![](https://placehold.it/15/FFFF00/000000?text=+) eth_getBlockByNumber                    
- [x] eth_getTransactionByHash                
- [ ] eth_getTransactionByBlockHashAndIndex   
- [ ] eth_getTransactionByBlockNumberAndIndex 
- [x] eth_getTransactionReceipt               
- [ ] eth_getUncleByBlockHashAndIndex         
- [ ] eth_getUncleByBlockNumberAndIndex       
- [ ] eth_getCompilers                        
- [ ] eth_compileLLL                          
- [x] eth_compileSolidity (deprecated)                    
- [ ] eth_compileSerpent                      
- [ ] eth_newFilter                           
- [ ] eth_newBlockFilter                      
- [ ] eth_newPendingTransactionFilter         
- [ ] eth_uninstallFilter                     
- [ ] eth_getFilterChanges                    
- [ ] eth_getFilterLogs                       
- [ ] eth_getLogs                             
- [ ] eth_getWork                             
- [ ] eth_submitWork                          
- [ ] eth_submitHashrate                      
- [ ] db_putString                            
- [ ] db_getString                            
- [ ] db_putHex                               
- [ ] db_getHex                               
- [ ] shh_post                                
- [ ] shh_version                             
- [ ] shh_newIdentity                         
- [ ] shh_hasIdentity                         
- [ ] shh_newGroup                            
- [ ] shh_addToGroup                          
- [ ] shh_newFilter                           
- [ ] shh_uninstallFilter                     
- [ ] shh_getFilterChanges                    
- [ ] shh_getMessages                         
- [x] personal_listAccounts                   
- [x] personal_newAccount                     
- [x] personal_sendTransaction                
- [x] personal_unlockAccount                  

## Installation

### go get

```bash
go get -u github.com/regcostajr/go-web3
```

### glide

```bash
glide get github.com/regcostajr/go-web3
```

## Usage

```go
import (
	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
	"github.com/regcostajr/go-web3/providers"
)

func test() {

	web3Client := web3.NewWeb3(providers.NewHTTPProvider("http://127.0.0.1:8545", 10))
	balance, err := web3Client.Eth.GetBalance("0x00035DB1C858Fe4C2772a779C6fEF0FdB850dE42", block.LATEST)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(balance)

}
```

More samples in the 'examples' directory.

### Requirements

* go ^1.8.3

[Go installation instructions.](https://golang.org/doc/install)

## Testing
```bash
go test -v test/*.go
```

## License

Package go-web3 is licensed under the [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html) License.

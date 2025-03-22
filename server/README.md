<!-- @author Dinh Le Hoang Anh - 105508318 -->
<!-- @author Pham Vu Minh - 105110564 -->
# Backend operations
We assume your working directory is `server/`, which is the same directory as this file.

# Configuration
The backend first reads the `.env` file at the root `server/` folder (or wherever your binary is) for config envvar. First copy the provided `.env.example` to `.env`:

```
$ cp .env.example .env
```

Then you can start opening `.env` file with your favorite text editor and follow the comments there

# Starting the backend
First, start the localnet:

```sudo kurtosis run github.com/ethpandaops/ethereum-package --args-file ../testnet/network_params.yaml --image-download always --enclave testnet```

Then, start the backend:

```go run .```

Now you can start interact with the backend at `localhost:8080`


## Building the backend
You can build the backend first with
```go build```
Then you can run the server with 
```./server```
Note that you still need to spin up the localnet for the server to work correctly at all

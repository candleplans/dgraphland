# dgraphland
Dgraphland is a Dgraph framework build on top of [Dgo](https://github.com/dgraph-io/dgo), the official Dgraph Go client which communicates with the server using gRPC.

# DGRAPH
- Download [dgraph-windows-amd64.zip](https://github.com/dgraph-io/dgraph/releases) (Last Version)
- Extract to ~/dgraph
- cd ~/dgraph && ./dgraph.exe zero
- cd ~/dgraph && ./dgraph.exe alpha --lru_mb 2048
- cd ~/dgraph && ./dgraph-ratel.exe
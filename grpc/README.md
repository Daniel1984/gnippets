### INSTRUCTION
install dependencies:
1. `go get -u google.golang.org/grpc`
2. `go get -u github.com/golang/protobuf/protoc-gen-go`

3. setup protoc:
3.1 MacOSX: `brew install protobuf` 
3.2 Ubuntu (Linux)

- Find the correct protocol buffers version based on your Linux Distro: https://github.com/google/protobuf/releases
- Example with x64:
```
  # Make sure you grab the latest version
  curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
  # Unzip
  unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
  # Move protoc to /usr/local/bin/
  sudo mv protoc3/bin/* /usr/local/bin/
  # Move protoc3/include to /usr/local/include/
  sudo mv protoc3/include/* /usr/local/include/
  # Optional: change owner
  sudo chown [user] /usr/local/bin/protoc
```

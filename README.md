# mini-docker-container-for-go-app

### Description
inspired by https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

tls-ca-bundle.pem is from: https://github.com/golang/go/blob/0ff40a76ad81e2b02c24e83eee8aa93352498f6f/src/crypto/x509/root_linux.go#L10

### Build Image
```
chmod +x ./build.sh
./build.sh
```
### Run
```
docker run -ti bluealert/hello
```

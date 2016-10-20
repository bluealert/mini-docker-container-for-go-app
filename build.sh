EXE=hello
CGO_ENABLED=0 GOOS=linux go build -a -o $EXE
docker build -t bluealert/$EXE .
rm -rf $EXE

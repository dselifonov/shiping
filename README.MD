### Build
```shell script
docker build -t consignment .
docker build -t cli .
docker build -t vessel .
```

### Run
```shell script
docker run -p 50051:50051 \
      -e MICRO_SERVER_ADDRESS=:50051 \
      consignment

docker run cli
```
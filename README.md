# transaction-app




## 🚀 Introduction

This is a project aiming to construct an microservice architecture, exposing REST endpoints through transaction-service and connecting to exchange-service usingo gRPC endpoints.

The first step to get everything running its to run the docker-compose.yaml in tools/environment. In case of any port having conflitc problems, change at will in the .env file next to the docker-compose file.

This compose contains the necessary tools and technologies to get everything running. In addition, there are a few extra tools to facilitate the testing and exploration of the project. Some of them beeing:
 - grpcox - tool to connect and send requests in grpc given the port and protofile
    - access it through localhost:6969
    - add the port, such as localhost:32001, which is the exchange-service grpc port, and use the exchange.proto file in https://github.com/walbety/transaction-app/tree/main/exchange-service/tools/protos
 - MongoExpress - tool to visualize and manipulate data in the mongodb.
     - access it through localhost:19001 


## 🛠️ Technologies

- Golang
- MongoDB
- python(behave)


## 📌 Services and components

## transaction-service

Check its own readme for more details on the endpoints.

Run using the Makefile:
```
make run
```



Test coverage command and svg (a lot of tests are yet to come...):
```
make cover-tree
```

![image](https://raw.githubusercontent.com/walbety/transaction-app/9d881e0b268276bf720977dab1dd259d866bd9b0/transaction-service/cover.svg)

## transaction-test

This is the automated functional tests for transaction-service.
To run it, first you need to set up the python environment:
```
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

After that, just run the command below to start the test:
```
behave
```





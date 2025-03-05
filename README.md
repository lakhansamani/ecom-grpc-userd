# userd

User service for ecom-grpc example

```
docker run --name postgres-cluster -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

```

```sh
grpcurl -plaintext -d '{ "name": "John Doe", "email": "john@example.com", "password": "securepass" }' -proto=apis/user/v1/user.proto localhost:50051 user.v1.UserService/Register

grpcurl -plaintext -d '{ "email": "john@example.com", "password": "securepass" }' -proto=apis/user/v1/user.proto localhost:50051 user.v1.UserService/Login

grpcurl -plaintext -H "authorization: bearer JWT_TOKEN" -proto=apis/user/v1/user.proto localhost:50051 user.v1.UserService/Me
```
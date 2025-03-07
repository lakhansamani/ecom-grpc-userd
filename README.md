# userd

User service for ecom-grpc example

```
docker run --name postgres-cluster -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
docker exec -it postgres-cluster psql -U postgres -c "CREATE DATABASE userdb;"

docker run -d --name jaeger \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  jaegertracing/all-in-one:latest

```

```sh
grpcurl -plaintext -d '{ "name": "John Doe", "email": "john@example.com", "password": "securepass" }' -proto=apis/user/v1/user.proto localhost:50051 user.v1.UserService/Register

grpcurl -plaintext -d '{ "email": "john@example.com", "password": "securepass" }' -proto=apis/user/v1/user.proto localhost:50051 user.v1.UserService/Login

grpcurl -plaintext -H "authorization: bearer JWT_TOKEN" -proto=apis/user/v1/user.proto localhost:50051 user.v1.UserService/Me
```
# Desafio 3 da pós gradução FullCycle Go Expert

## Enunciado do desafio

Olá devs!
Agora é a hora de botar a mão na massa. 

Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Instruções de desenvolvimento

Para regerar fontes da injeção de dependências (Wire):
```bash
./wire_gen.sh
```

Para regerar fontes GraphQL:
```bash
./gql_gen.sh
```
Para regerar fontes gRPC:
```bash
./grpc_proto_gen.sh
```
Para rodar migrations:
```
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
````


## Instruções de uso

Inicializar dependências da aplicação (Banco de dados MySQL e Broker de mensageria RabbitMQ)

```bash
docker-compose up -d
```

Executar aplicação

```bash
./run.sh
```

Portas padrão:
- Servidor HTTP: 8000
- Servidor gRPC: 50051
- Servidor GraphQL: 8080

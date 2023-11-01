# nellcorp-chalenge-pedro-domingos
Para Executar:
    docker-compose up


para efectuar o deposito:
    curl -X POST -H "Content-Type: application/json" -d '{"idconta": 1, "montante": 30}' localhost:8080/depositar



para efectuar o levantamento

    curl -X POST -H "Content-Type: application/json" -d '{"idconta": 1, "montante": 10}' localhost:8080/levantamento




para testar pode se usar postman,insomia ou mesmo curl

para executar a transferencia
    curl -X POST -H "Content-Type: application/json" -d '{"idcontaemissora": 1,"idcontareceptora": 2, "montante": 10}' localhost:8080/transferir


para listar as transações de uma determinada conta(da conta 1 por exemplo):
    curl -X GET localhost:8080/transacoes/1

 

para consultar saldo de uma  conta(da conta 1 por exemplo):
    curl -X GET localhost:8080/consultar/1


para reembolsar:(transacao 1 exemplo)
    curl -X GET localhost:8080/reembolsar/1



para listar todas as contas:
    curl -X GET localhost:8080/contas

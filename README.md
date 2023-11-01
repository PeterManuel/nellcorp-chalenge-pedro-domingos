# nellcorp-chalenge-pedro-domingos
Para Executar:
    docker-compose up


para efectuar o deposito

    POST
    
    localhost:8080/depositar
    {
        "idconta":1,
        "montante":30
    }


para efectuar o levantamento

    POST
    
    localhost:8080/levantamento
    {
        "idconta":1,
        "montante":10
    }




para testar pode se usar postman,insomia ou mesmo curl

para executar a transferencia

    POST

    localhost:8080/transferir
    
    {
	"idcontaemissora":1,
	"idcontareceptpra":2,
	"montante":10
    }

para listar as transações de uma determinada conta(da conta 1 por exemplo):
GET
    localhost:8080/transacoes/1
 

para consultar saldo de uma  conta(da conta 1 por exemplo):
GET
    localhost:8080/consutar/1
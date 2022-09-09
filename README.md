POC HEXAGONAL COM GO

>COMANDO PARA GERAR O MOCK:
mockgen -destination=application/mocks/application.go -source=application/product.go application

>CONECTAR NO BANCO SQLITE VIA TERMINAL:
sqlite3 sqlite.db

>PARA ADICIONAR O CLI DO COBRA VIA TERMINAL:
cobra-cli add cli

>PARA ADICIONAR O HTTP DO COBRA VIA TERMINAL:
cobra-cli add http

>PARA CRIAR UM PRODUTO VIA CLI NO TERMINAL :
go run main.go cli -a=create -n=TV -p=2000
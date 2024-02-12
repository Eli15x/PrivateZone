# ZCOM

rodar mongodb

instalar dependencias mongod e mongosh
caso acontença algum problema 

utilize o comando na raiz: 
sudo rm -rf /tmp/mongodb-27017.sock

fluxo natural:
service mongod start
mongosh


# Inicializando Todo Fluxo

Para inicializar as instancias com o kafka local e o mongodb:
    Na pasta /PrivateZone/docker rode no terminal o seguinte comando:
    docker-compose up
    (Tenha o docker baixado corretamente em sua maquina)

Na pasta raiz /PrivateZone
no terminal rode o seguinte comando
    go run main.go

Na pasta PrivateZone/consumer, onde iniciará o Consumidor das mensagens kafka
no terminal rode o seguinte comando:
    go run main.go



### dados tecnicos 

banco: 
mongo

mensageria:
kafka

lambdas:
1. alteração de conteudos de arquivo 
2. edições de imagens de usuario.

cache:
redis


redirecionamento 

kafka -> trativa no proprio.
kafka -> lamda

kafka/ou outro -> serviço de autenticação celular.
kafka/ou outro -> serviço de autenticação email

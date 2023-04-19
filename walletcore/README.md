### To access the database do the following steps:

1: access the container of mysql:
`docker-compose exec mysql bash`
2: sign in with the following user:
`mysql -uroot -p wallet`
3: then type the password: root


### To execute the project do the following steps:
1: access the container of the main app:
`docker-compose exec goapp bash`
2: open the dir with the main.go file (cd cmd/walletcore)
`cd cmd/walletcore `
3: execute the following:
`go run main.go `


### To see the topics on confluence just use the following link:
`localhost:9021`
 

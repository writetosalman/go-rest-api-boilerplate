## GoLang Rest API Boilerplate

This is the rest api boiler plate built in GoLang. It will help you kick start your next REST API project quickly.
It supports authorization and JWT tokens.

### Setup

* Checkout project in `$GOPATH/src/github.com/writetosalman/go-rest-api-boilerplate` folder.

* It requires Mysql database.

* It requires following commands to run to setup dependencies

```
go get github.com/gorilla/context
go get github.com/gorilla/mux
go get github.com/subosito/gotenv
go get github.com/go-sql-driver/mysql
go get github.com/dgrijalva/jwt-go
go get golang.org/x/crypto/bcrypt
go get golang.org/x/net/context
```

* Create MySQL database and run SQL files in the `__SQL` folder to create necessary tables.

* Copy `.env.example` as `.env`

* Edit `.env` and enter port and MySQL connection string based on your system information.


### Deployment

To run in dev, run command

```
go run main.go
```

### Production

To run in production, run command

```
go build main.go
chmod +x main
nohup ./main &> main.log &
```

To kill it, just run `kill -9 <pID>`

### Future Plans

* Write tests. Should have been written with dev but I was in hurry to put things together.

* Refactor utilities so that different functions are in their separate package. Right now it is bit hotchpoch.

* Built support for 2FA.

* More SQL databases support.

-


# SSO サンプル

# Run SignIn Service
```bash
$ mkdir -p $GOPATH/src/github/m0cchi
$ cd $GOPATH/src/github/m0cchi
$ git clone git@github.com:m0cchi/gfalcon-signin-service.git
$ cd gfalcon-signin-service
$ dep ensure
$ cd vendor/github.com/m0cchi/gfalcon/metatools/sql
$ mysql -u $USER -p < create_database.sql
$ mysql -u $USER -p -D gfalcon < create_database.sql
$ mysql -u $USER -p -D gfalcon < create_table.sql
$ mysql -u $USER -p -D gfalcon < create_defaultdata.sql
$ cd ../../../../../../
$ bower install
$ go run test/init_data.go --dbhost 'user:password@unix(/tmp/mysql.sock)/gfalcon?parseTime=true'
$ go run app/server.go --dbhost 'user:password@unix(/tmp/mysql.sock)/gfalcon?parseTime=true' --allowed-host 'saas.m0cchi.net'
```

# Run this sample
```bash
$ cd $GOPATH/src/github/m0cchi
$ git clone git@github.com:m0cchi/gfalcon-sso-sample.git
$ cd gfalcon-sso-sample
$ dep ensure
$ go run app/server.go --dbhost 'user:password@unix(/tmp/mysql.sock)/gfalcon?parseTime=true' --port 3000
```

# View
![](https://i.gyazo.com/1cde44d51b4356e8cedbc8029b9be131.gif)

Creating the migrations: 

```
 goose -dir "db/migrations" create create_user_table sql

```

Apply the migrations:

```

goose -dir "db/migrations" mysql "root:Jyoti@123@tcp(127.0.0.1:3306)/auth_dev" up

```

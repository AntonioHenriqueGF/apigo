# Instruções para as Migrations

Para criar uma migration, basta rodar o comando:

```bash
migrate create -ext sql -dir db/migrations -seq descricao_da_operacao
```

As descrições devem seguir a convenção `create_<entidade>_table` para a criação de tabelas e `add_<campo>_to_<entidade>` para adição de campos.

As migrations serão guardadas na pasta `db/migrations`.

## Rodar as Migrations

Para rodar as migrations, inicie o container do banco de dados e utilize o comando:

```bash
migrate -path db/migrations -database "mysql://<usuario>:<senha>@tcp(<host>:3306)/<database>" up
```

No nosso caso, a string de conexão seria `mysql://<usuario>:<senha>@tcp(127.0.0.1:3306)/api_go_database`. Desta forma, o `migrate` conseguirá se conectar ao banco de dados dentro do container corretamente e aplicar as migrations.
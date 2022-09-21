# pismo-transactions

## Requirements
- To execute program you need have Golang 1.17 or newer and Docker instaled


### MakeFile
- Use Make commands to start apps easily:

```yaml
generate-docs:
	Use to generate swagger documentation

run-migration:
	Use this to create database and insert default value. I recomend use only in the frist execution

start-app:
	Use to start server

database-up:
	Use to start database server in docker-compose

database-down:
	Use to turn off database server in docker-compose

```

### Reminders
- Remember to change your personal volumes in  docker-compose file.
- Remeber to change  SQL_SCRIPTS_DIR for your personal project dir in .env files.
- Check auth basic credencials in .env file
- By Default server is running in http://localhost:1323
- Docs are avalible in http://localhost:1323/swagger/index.html

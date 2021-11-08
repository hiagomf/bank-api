## Para buildar o Dockerfile
docker build -t db_bank_api:latest .

## Para rodar
docker run -p 5432:5432 -e POSTGRES_PASSWORD=dev1234 -d db_bank_api
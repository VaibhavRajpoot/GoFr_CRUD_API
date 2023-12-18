
# CRUD API Using GoFr Framework

This is an API of the Library Management System made using GoLang using the GoFr framework. It contains a single [main.go](https://github.com/VaibhavRajpoot/GoFr_CRUD_API/blob/main/main.go) file which consists of all the route handles for performing the CRUD operations like Read, Create, Update, Delete operations and all the MySQL queries. It uses the Docker Mysql for storing the records. 


## Deployment

1. Clone this git repository

2. Make sure that Go is installed in your system & run the below command for installing the GoFr framework

```bash
go get gofr.dev

```
3. Now open up the terminal and enter the below-mentioned command to run the MySql docker container and create the table in it
```bash
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30

sudo docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE customers (id INT AUTO_INCREMENT PRIMARY KEY, bookname VARCHAR(255) NOT NULL,publication VARCHAR(255) NOT NULL,issuername VARCHAR(255) NOT NULL, fine INT NOT NULL);"
```
4. After creating the Database run the main.go file 
```bash
go run main.go
```
5. Now open up the postman and check the working of all the CRUD operations.

## UML & Sequence Diagram

![SequenceDiagram](https://github.com/VaibhavRajpoot/GoFr_CRUD_API/assets/50841702/c94d90a2-0cbe-405a-b603-36cf673ae23a)

![UML Diagram](https://github.com/VaibhavRajpoot/GoFr_CRUD_API/assets/50841702/0537d58c-5e2d-4f9b-964a-17e314346486)
## Screenshots

![Postman With get request](https://github.com/VaibhavRajpoot/GoFr_CRUD_API/assets/50841702/7137ca7f-a4de-44a4-93d4-998afc608557)
![Postman with Post Request](https://github.com/VaibhavRajpoot/GoFr_CRUD_API/assets/50841702/863f133b-19c2-428e-804c-69337ba222ea)

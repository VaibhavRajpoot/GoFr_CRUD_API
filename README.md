# CRUD API Using GoFr Framework

This is an API of the Library Management System made using GoLang using the GoFr framework. It contains a single main.go file which consists of all the route handles for performing the CRUD operations like Read, Create, Update, Delete operations and all the MySQL queries. It uses the Docker Mysql for storing the records. 

## Installation
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

## UML and Sequence Diagrams
SEQUENCE DIAGRAM
<img src="https://github.com/VaibhavRajpoot/GoFr_CRUD_API/tree/main/UML%20%26%20Sequence%20Diagrams" alt="Sequence Diagram" title="Sequence Diagram">


## Postman Images of working Project





create database BookStore;
show databases;
use BookStore;

create table users(
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO users (id,username, email,password) VALUES (1,'Priyanka', 'priyanka18@gmail.com','bachhav');

create table books(
Id INT NOT NULL PRIMARY KEY AUTO_INCREMENT ,
Title varchar(40) NOT NULL,
Author varchar(40) NOT NULL,
bookQuantity INT NOT NULL
);
INSERT INTO books (id, title, author) values (1, "The Alchmist", "Paulo Coelho");

create table cart(
	bookName varchar(50),
    bookId INT NOT NULL,
    cartId INT NOT NULL AUTO_INCREMENT,
    quantity INT NOT NULL,
    PRIMARY KEY(cartId)
);
INSERT INTO cart (bookName, bookId, cartId, quantity) values ("The Alchmist", 1, 1, 5);
select * from cart;

create table orders(
	orderID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    userId INT NOT NULL,
    bookId INT NOT NULL,
    quantity INT NOT NULL,
    orderDate varchar(50),
    price INT NOT NULL,
    orderStatus varchar(50)
);
INSERT INTO orders (orderID, userId, bookId, quantity, orderDate, price, orderStatus) values (1, 1, 22, 10,"01-01-2023", 100, 1);
select * from orders;


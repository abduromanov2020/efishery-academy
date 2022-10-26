create database order_database;

create table if not exists customers (
	id serial not null,
	customer_name char(50) not null
);

insert into customers (customer_name) values ('Rizki');
insert into customers (customer_name) values ('Billar');

update customers set customer_name = 'Febian' where customer_name  = 'Billar';

delete from customers where customer_name = 'Rizki';
select * from customers;

create table if not exists products (
	id serial not null,
	name char(50) not null
);

insert into products (name) values ('Keyboard');
insert into products (name) values ('Mouse');

update products  set name = 'Cable' where id  = 1;

delete from products where id = 2;

select * from products;

create table if not exists orders (
	order_id serial not null,
	customer_id int not null,
	product_id int not null,
	order_date date not null,
	total double precision not null
);

insert into orders (
customer_id, product_id, order_date, total
) values (
	1,4,'2002-09-02',70000
);

update orders set total = 69000 where order_id = 1;

delete from orders where customer_id = 1;

select * from orders;


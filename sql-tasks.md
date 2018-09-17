# Tasks

1. Install / setup postgres.

2. Create a database (createdb DatabaseName) and connect to it.

3. Create our first table. 
   ```sql
   create table pizza (
   id serial primary key,
   name varchar(50) not null, 
   ingredients varchar(200) not null, 
   price money not null,
   size varchar(2) not null,
   comments varchar(300) );
   ```

4. Add data to our table (second statement causes error).
   ```sql
   insert into pizza (id, name, ingredients, price, size, comments) values 
   (1, 'pepperoni', 'cheese, peperoni', 10, 'S', '') ; 
   insert into pizza (id, name, ingredients, price, size, comments) values 
   (1, 'pepperoni', 'cheese, peperoni', 15, 'L', '') ; 
   ```

5. Add more.
   ```sql
   insert into pizza (id, name, ingredients, price, size, comments) values 
   (2, 'pepperoni', 'cheese, peperoni', 15, 'L', '') , 
   (3, 'margherita', 'cheese, peperoni', 10, 'S', '') , 
   (4, 'margherita', 'cheese, peperoni', 15, 'L', '') ;
   ```

6. View our data.
   ``` select * from pizza; ```
   ``` select name, price  from pizza ; ```
   ``` select * from pizza where price < 12 ; ```

   
   



   

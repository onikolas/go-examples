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
   (1, 'pepperoni', 'cheese, pepperoni', 10, 'S', '') ; 
   insert into pizza (id, name, ingredients, price, size, comments) values 
   (1, 'pepperoni', 'cheese, pepperoni', 15, 'L', '') ; 
   ```

5. Add more.
   ```sql
   insert into pizza (id, name, ingredients, price, size, comments) values 
   (2, 'pepperoni', 'cheese, pepperoni', 15, 'L', '') , 
   (3, 'margherita', 'cheese, pepperoni', 10, 'S', '') , 
   (4, 'margherita', 'cheese, pepperoni', 15, 'L', '') ;
   ```

6. View our data.
   ```sql
   select * from pizza; 
   select name, price  from pizza ; 
   select * from pizza where price < '$12' ; 
   ```
7. Update margheritas.
   ```sql
   update pizza set ingredients = 'cheese, tomato' where name = 'margherita'; ```
   ```

8. Some problems with our table:
   - Duplication.
   - How would we search for pizzas without pepperoni?

9. Lets make a table to hold size/price. Each pizza can be priced differently.
	```sql
	create table prices(
	id serial primary key,
	pizza serial references pizza(id),
	size varchar(2) not null,
	price decimal not null
	);
	```

10. Remove the price and size columns from pizza.
   ```sql
   alter table pizza drop column price ;
   alter table pizza drop column size ;
   ```

11. Add prices.
	Pepperoni:	
	```sql
	insert into prices values 
	(1, 1, 'S', 10.50), (2, 1, 'M', 12.50), (3, 1, 'L', 15.50) ; 
	```
	Margherita:
	```sql
	insert into prices values 
	(4, 2, 'S', 10.00), (5, 2, 'M', 12.00), (6, 2, 'L', 15.00) ; 
	```

12. Delete duplicate pizzas. This will violate a constraint.
	```	delete from pizza where id in (2,4) ;```
	
13. Update foreign keys and then delete.
	``` update prices set pizza = 3 where pizza = 2; ```

14. Get the complete picture.
	Cartesian product (usually not what we want):	
	``` select * from pizza, prices ; ```
	
	Filter with a where clause:
	```  select * from pizza, prices where pizza.id = prices.pizza ```
	
	Or with a join clause (faster, usually):
	```select * from pizza join prices on (pizza.id = prices.pizza) ```
	
15. Practice:
	- Make a menu. Query to get  the name, ingredients, size and price of each pizza.
	- What if we also want the pizza id.

16. Lets also fix the ingredients problem. While we are at it, lets add nutritional info about 
	each ingredient. This will require a different solution than price/size.

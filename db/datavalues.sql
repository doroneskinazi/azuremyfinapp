--INSERT INTO myfpuser(name) VALUES ('deskinazi')




--INSERT INTO trans_category(name, type_id, user_id) values ('Dividend',2,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Salary',2,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Cloths',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Taxes',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Household Items',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Utilities',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Rent',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Auto Insurance',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Groceries',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Auto Gas',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Medical Insurance',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Auto Registration',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Auto Repair',1,1)
--INSERT INTO trans_category(name, type_id, user_id) values ('Entertainment',1,1)
--
--
--insert into trans_type (name, user_id) values('Expense',1)
--insert into trans_type (name, user_id) values('Income',1)
--
--
insert into trans_money (user_id, date, category_id, amount, description)
select 1, '1/15/18', id,  20.25, 'Jons market'
from trans_category
where id = 9

insert into trans_money (user_id, date, category_id, amount, description)
select 1, '1/17/18', id,  230.25, 'Amazon'
from trans_category
where id = 5


select * from trans_category
select * from trans_type
select * from trans_money
select * from myfpuser
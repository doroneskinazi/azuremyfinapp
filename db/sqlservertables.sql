/*
use myfp
go
drop table myfp.dbo.trans_money
drop table myfp.dbo.trans_category
drop table myfp.dbo.trans_type
DROP TABLE myfp.dbo.sysusers
 */
create table trans_money(
    id numeric(18,0) identity,
	user_id numeric(18,0),
	date datetime, 
    category_id numeric(4,0),
    amount money,
    description varchar(100)
)

create table trans_type(
	id numeric(4,0) identity,
	user_id numeric(18,0),
	name varchar(80)
)
	
create table trans_category(
	id	numeric(4,0) identity,
	user_id numeric(18,0),
	type_id numeric(4,0),
	name varchar(80)
)
create table myfpuser (
	id numeric(18,0) identity,
	name varchar(80)
)

select * from trans_category

create table users
(
    user_id text UNIQUE,
    password text,
    insert_date timestamp with time zone,
    update_date timestamp with time zone
);
create table user_names
(
    user_id text UNIQUE,
    user_name text
)
--update users set user_id='aaa', insert_date='2020/04/11 11:11:11';--
--insert into users values('bbb','bbb','2020/04/11 11:11:11','2020/04/11 11:11:11')
create table users
(
    user_id text UNIQUE,
    password text,
    insert_date timestamp with time zone,
    update_date timestamp with time zone
);

--update users set user_id='aaa', insert_date='2020/04/11 11:11:11';--
--insert into users values('bbb','bbb','2020/04/11 11:11:11','2020/04/11 11:11:11')
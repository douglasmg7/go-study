-- drop table if exists entrance;
-- enable foreign keys
-- not working, reset to off when back to db
pragma foreign_keys = on;

create table entrance (
  id integer primary key autoincrement,
  user_id integer not null,
  school varchar(64) not null,
  created date not null
);

create table user (
  id integer primary key autoincrement,
  name varchar(64) null,
  -- carteira de identidade
  rg varchar(64) null,
  -- mobile number
  mobile varchar(64) null,
  email varchar(64) null,
  created date null
);


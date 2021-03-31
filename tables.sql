drop table if exists authors cascade;
drop table if exists series cascade;
drop table if exists genres cascade;
drop table if exists books cascade;
drop table if exists series_books cascade;
drop table if exists reviews cascade;

create table authors (
	id varchar(21) primary key not null,
	name varchar(255) not null unique
);

create table series (
	id varchar(21) primary key not null,
	name text not null unique
);

create table genres (
	id varchar(21) primary key not null,
	name text not null unique
);

create table books (
	id varchar(21) primary key not null,
	author_id varchar(21) not null references authors(id) ON DELETE CASCADE ON UPDATE CASCADE,
	genre_id varchar(21) not null references genres(id) ON DELETE set null ON UPDATE CASCADE,
	title text not null,
	age int not null default 0,
	cover_url varchar(2000) not null
);

create table series_books (
	series_id varchar(21) not null references series(id) ON DELETE CASCADE ON UPDATE CASCADE,
	book_id varchar(21) not null references books(id) ON DELETE CASCADE ON UPDATE CASCADE,
	number smallint not null,
	primary key(series_id, book_id)
);

create table reviews (
	id varchar(21) not null primary key,
	book_id varchar(21) not null references books(id) ON DELETE CASCADE ON UPDATE CASCADE,
	rating smallint not null,
	review text not null,
	created_at timestamp not null default (now() at time zone 'utc')
);
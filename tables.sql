drop table if exists authors;
drop table if exists series;
drop table if exists genres;
drop table if exists books;
drop table if exists series_books;
drop table if exists reviews;

create table authors (
	id varchar(21) primary key not null,
	name varchar(255) not null
);

create table series (
	id varchar(21) primary key not null,
	name text not null
);

create table genres (
	id varchar(21) primary key not null,
	name text not null
);

create table books (
	id varchar(21) primary key not null,
	author_id varchar(21) not null references authors(id),
	genre_id varchar(21) not null references genres(id),
	title text not null,
	age int not null default 0,
	cover_url varchar(2000) not null
);

create table series_books (
	series_id varchar(21) not null references series(id),
	book_id varchar(21) not null references books(id),
	number smallint not null,
	primary key(series_id, book_id)
);

create table reviews (
	id varchar(21) not null primary key,
	book_id varchar(21) not null references books(id),
	rating smallint not null,
	review text not null
);
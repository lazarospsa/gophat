create table todos (
    id serial not null,
    todo varchar(128),
    done boolean,
    primary key(id)
);

insert into todos(todo, done) values('Hello!', false);

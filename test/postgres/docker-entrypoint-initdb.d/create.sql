
CREATE USER user_todo;
ALTER USER user_todo WITH ENCRYPTED PASSWORD '1122';


CREATE DATABASE api_todo;
\connect api_todo;

CREATE TABLE todos (
    id serial primary key,
    title varchar,
    description text,
    done bool default FALSE
);



insert into dbo.people
    (id, name, age)
values
    ('1', 'Alessandra', '2'),
    ('2', 'Kopec', '10'),
    ('3', 'Zacchi', '20');

-- update dbo.people SET name = "Aleh" WHERE id = 10;
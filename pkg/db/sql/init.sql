-- create users table
create table users (
    id serial primary key,
    username varchar(15) not null unique,
    password varchar(255) not null,
    email varchar(255) not null unique,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp  
);

-- create posts table
create table posts (
    id serial primary key,
    content varchar(1000) not null,
    user_id int not null references users(id),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- create likes
create table likes (
    id serial primary key,
    user_id int not null references users(id),
    post_id int not null references posts(id),
    created_at timestamp not null default current_timestamp
);
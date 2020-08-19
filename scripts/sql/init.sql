create table if not exists users (
    id serial primary key,
    username varchar(255) not null unique,
    created_at timestamp not null default now()
);

create table if not exists chat (
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp not null default now(),
    recent_msg_at timestamp not null default now()
);

create table if not exists message (
    id serial primary key,
    chat_id int not null,
    user_id int not null,
    body text not null,
    created_at timestamp not null default now(),
    foreign key(chat_id) references chat(id),
    foreign key(user_id) references users(id)
);

create table if not exists users_chat (
    chat_id int not null,
    user_id int not null,
    foreign key(chat_id) references chat(id),
    foreign key(user_id) references users(id)
);


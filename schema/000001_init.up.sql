CREATE TABLE users_balances
(
    id      bigserial primary key,
    user_id bigint unique               not null,
    balance bigint CHECK (balance >= 0) not null
);

CREATE TABLE reservations
(
    id         bigserial primary key,
    user_id    bigint                              not null,
    order_id   bigint                              not null,
    service_id bigint                              not null,
    value      bigint CHECK (value >= 0)           not null,
    created_at timestamp default current_timestamp not null,
    CONSTRAINT user_id FOREIGN KEY (user_id) references users_balances (user_id)
);

CREATE UNIQUE INDEX reservation ON reservations (user_id, order_id, service_id);

CREATE TABLE purchases
(
    id         bigserial primary key,
    user_id    bigint                              not null,
    order_id   bigint                              not null,
    service_id bigint                              not null,
    value      bigint CHECK (value >= 0)           not null,
    created_at timestamp default current_timestamp not null,
    CONSTRAINT user_id FOREIGN KEY (user_id) references users_balances (user_id)
);

CREATE TABLE transfers
(
    id           bigserial primary key,
    from_user_id bigint                              not null,
    to_user_id   bigint                              not null,
    value        bigint CHECK (value >= 0)           not null,
    comment      varchar(255),
    created_at   timestamp default current_timestamp not null,
    CONSTRAINT from_user_id FOREIGN KEY (from_user_id) references users_balances (user_id),
    CONSTRAINT to_user_id FOREIGN KEY (to_user_id) references users_balances (user_id)
);
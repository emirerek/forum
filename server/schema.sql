DROP TABLE IF EXISTS account CASCADE;
DROP TABLE IF EXISTS subforum CASCADE;
DROP TABLE IF EXISTS thread CASCADE;
DROP TABLE IF EXISTS reply CASCADE;

CREATE TABLE account (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(256) NOT NULL UNIQUE,
    email VARCHAR(256) NOT NULL UNIQUE,
    password_hash VARCHAR(512) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE subforum (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE thread (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    account_id INT NOT NULL,
    subforum_id INT NOT NULL,
    title TEXT NOT NULL UNIQUE,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (subforum_id) REFERENCES subforum(id)
);

CREATE TABLE reply (
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    account_id INT NOT NULL,
    thread_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (thread_id) REFERENCES thread(id)
);

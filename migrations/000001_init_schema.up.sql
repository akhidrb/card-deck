BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS decks;

create table decks
(
    id         uuid default uuid_generate_v4() not null primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    shuffled   bool                            not null,
    cards      varchar[]                       not null
);


COMMIT;
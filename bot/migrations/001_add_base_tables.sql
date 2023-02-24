CREATE TABLE IF NOT EXISTS owners (
    id serial primary key,
    gh_id integer,
    login varchar(128),
    url varchar(160),
    avatar_url varchar(160),
    type varchar(20),
    twitter_username varchar(128)
);

CREATE TABLE IF NOT EXISTS bounties (
    id serial primary key,
    owner_id integer references owners(gh_id),
    title text,
    url varchar(160),
    reward bigint
);
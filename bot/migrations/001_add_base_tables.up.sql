CREATE TABLE IF NOT EXISTS owners (
    gh_id integer primary key,
    login varchar(128),
    url varchar(160),
    avatar_url varchar(160),
    type varchar(20),
    twitter_username varchar(128),
    wallet_address varchar(64) default '',

    UNIQUE (gh_id)
);

CREATE TABLE IF NOT EXISTS bounties (
    gh_id integer primary key,
    owner_gh_id integer references owners(gh_id),
    title text,
    url varchar(160),
    reward bigint
);
-- +goose Up

-- Extensions
create
extension if not exists "uuid-ossp";
create
extension if not exists "pg_trgm";

-- Accounts
create table if not exists "accounts"
(
    id varchar primary key
);

-- Settlements
create table if not exists "settlements"
(
    "id"       uuid primary key default uuid_generate_v4
(
),
    "name"     varchar not null,
    "location" point   not null
    );

create index "settlements_name_gin_idx"
    on "settlements"
    using gin ("name" gin_trgm_ops);

-- Quests
create table if not exists "quests"
(
    "id"            uuid primary key default uuid_generate_v4
(
),
    "settlement_id" uuid    not null references "settlements"
(
    "id"
),
    "name"          varchar not null,
    "description"   varchar not null,
    "type"          varchar not null,
    "avg_duration"  bigint  not null,
    "reward"        decimal not null default 0
    );

create table if not exists "quests_steps"
(
    "id"         uuid primary key default uuid_generate_v4
(
),
    "quest_id"   uuid    not null references "quests"
(
    "id"
),
    "order"      int     not null,
    "name"       varchar not null,
    "place_type" varchar not null,
    "address"    varchar not null,
    "phone"      varchar,
    "email"      varchar,
    "website"    varchar,
    "schedule"   jsonb,
    "location"   point   not null
    );

-- Encyclopedia
create table if not exists "encyclopedia_items"
(
    "id"            uuid primary key default uuid_generate_v4
(
),
    "settlement_id" uuid    not null references "settlements"
(
    "id"
),
    "name"          varchar not null,
    "type"          varchar not null,
    "article"       text
    );

-- Achievements
create table if not exists "achievements"
(
    "id"          uuid primary key default uuid_generate_v4
(
),
    "name"        varchar not null,
    "icon"        varchar,
    "steps"       int     not null,
    "description" varchar not null default ''
    );

create table if not exists "accounts_achievements"
(
    "account_id"     varchar not null references "accounts"
(
    "id"
),
    "achievement_id" uuid    not null references "achievements"
(
    "id"
),
    "passed"         int     not null default 0,
    "is_completed"   bool    not null default false,
    primary key
(
    "account_id",
    "achievement_id"
)
    );

-- +goose Down
drop table if exists "accounts_achievements";
drop table if exists "achievements";
drop table if exists "encyclopedia_items";
drop table if exists "quests_steps";
drop table if exists "quests";
drop index if exists "settlements_name_gin_idx";
drop table if exists "settlements";
drop table if exists "accounts";
drop
extension if exists "pg_trgm";
drop
extension if exists "uuid-ossp";

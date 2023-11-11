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
    "id"       uuid primary key default uuid_generate_v4(),
    "name"     varchar not null,
    "location" point   not null
);

create index if not exists "settlements_name_gin_idx"
    on "settlements"
        using gin ("name" gin_trgm_ops);

-- Quests
create table if not exists "quests"
(
    "id"            uuid primary key default uuid_generate_v4(),
    "settlement_id" uuid    not null references "settlements"("id"),
    "name"          varchar not null,
    "description"   varchar not null,
    "type"          varchar not null,
    "avg_duration"  bigint  not null,
    "reward"        decimal not null default 0
);

create table if not exists "quests_steps"
(
    "id"         uuid primary key default uuid_generate_v4(),
    "quest_id"   uuid    not null references "quests" ("id"),
    "order"      int     not null,
    "name"       varchar not null,
    "place_type" varchar not null,
    "address"    varchar not null,
    "phone"      varchar,
    "email"      varchar,
    "website"    varchar,
    "schedule"   jsonb,
    "location"   point   not null,
    "status"     varchar not null default ''
);

create table if not exists "account_quests"
(
    "account_id" varchar not null references "accounts" ("id"),
    "quest_id"   uuid    not null references "quests" ("id"),
    "is_active"  bool    not null default false,
    primary key ("account_id","quest_id")
);

create table if not exists "account_quest_steps"
(
    "quest_id" uuid    not null references "quests" ("id"),
    "quest_step_id" uuid    not null references "quests_steps" ("id"),
    "account_id" varchar not null references "accounts" ("id"),
    "status"   varchar not null,
    primary key ("account_id","quest_step_id")
);

-- Encyclopedia
create table if not exists "encyclopedia_items"
(
    "id" uuid primary key default uuid_generate_v4(),
    "settlement_id" uuid    not null references "settlements"("id"),
    "name"          varchar not null,
    "type"          varchar not null,
    "article"       text
);

-- Achievements
create table if not exists "achievements"
(
    "id"          uuid primary key default uuid_generate_v4(),
    "name"        varchar not null,
    "icon"        varchar,
    "steps"       int     not null,
    "description" varchar not null default ''
);

create table if not exists "accounts_achievements"
(
    "account_id"     varchar not null references "accounts" ("id"),
    "achievement_id" uuid    not null references "achievements" ("id"),
    "passed"         int     not null default 0,
    "is_completed"   bool    not null default false,
    primary key ("account_id", "achievement_id")
);

--Riddles
create table if not exists "riddles"
(
    "id" uuid primary key default uuid_generate_v4(),
    "quest_step_id" uuid not null default uuid_generate_v4(),
    "name" varchar not null,
    "description" varchar not null,
    "status" varchar not null default 'not_passed',
    "letter" varchar not null
 );

create table if not exists "account_riddles"
(
    "account_id" varchar not null references "accounts"("id"),
    "riddle_id" uuid not null references "riddles"("id"),
    "riddle_status" varchar not null default 'not_passed',
    "riddle_letter" varchar not null,
    primary key ("account_id", "riddle_id")
);

-- +goose Down
drop table if exists "accounts_achievements";
drop table if exists "account_quests";
drop table if exists "account_quest_steps";
drop table if exists "achievements";
drop table if exists "encyclopedia_items";
drop table if exists "quests_steps";
drop table if exists "quests";
drop index if exists "settlements_name_gin_idx";
drop table if exists "settlements";
drop table if exists "accounts";
drop table if exists "riddles";
drop table if exists "account_riddles";
drop
    extension if exists "pg_trgm";
drop
    extension if exists "uuid-ossp";

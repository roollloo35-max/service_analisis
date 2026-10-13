create table campaigns  (
    id bigserial primary key,
    names text not null,
    statuse text not null,
    created_at timestamptz not null default NOW(),
    budget numeric(12,2) not null default 0,
    cost_per_click numeric(12,2) not null default 0,
    target_url text not null,
    start_dates timestamptz not null,
    end_date timestamptz,
    last_update timestamptz
);

create table daily_stat(
    id bigserial primary key,
    campaign_id bigint not null,
    dates timestamptz not null,
    impressions bigint not null,
    clicks bigint not null,
    cost numeric(12,2) not null,
    conversion int,
    revenue numeric(12, 2),

    unique (campaign_id, dates)
);

create table conversion_event (
    id bigserial primary key,
    campaign_id bigint not null,
    occurred_at timestamptz not null,
    amount numeric(12, 2)
); 

create table users (
    id bigserial primary key,
    email text unique not null,
    password_hash text  not null,
    created_at timestamp not null default NOW()
);


alter table daily_stat add foreign key ("campaign_id") references "campaigns" ("id") on delete cascade;

alter table conversion_event add foreign key("campaign_id") references "campaigns" ("id") on delete cascade;

create index on "campaigns" ("names");
 
create index on "daily_stat" ("campaign_id");

create index on "conversion_event" ("campaign_id");


comment on column "conversion_event"."amount" is 'can be negative or positive'
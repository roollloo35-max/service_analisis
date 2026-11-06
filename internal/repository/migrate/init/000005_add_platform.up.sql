create table ad_platform (
    id bigserial primary key,
    code text not null unique,
    name_platform text not null,
    created_at timestamptz not null default now()
);

alter table campaigns 
    add column external_id bigint,
    add column delete_atr bool not null default false,
    add column version_atr bigint not null default 0; 

alter table users 
    add column delete_atr bool default false, 
    add column version_atr bigint not null default 0; 


alter table campaigns
    add constraint fk_campaigns_ad_platform
    foreign key ("external_id") references "ad_platform" ("id") on delete cascade;


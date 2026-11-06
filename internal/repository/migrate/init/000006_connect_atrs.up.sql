-- поправить сущность daily stat 
-- поменять тип поля conversion на bigint 
-- сделать из поля conversion внешний ключ на таблицу conversion_event

alter table daily_stat
    alter column conversion type bigint;

alter table daily_stat
    add constraint fk_daily_stat_conversion_event
    foreign key ("conversion") references "conversion_event" ("id") on delete cascade;
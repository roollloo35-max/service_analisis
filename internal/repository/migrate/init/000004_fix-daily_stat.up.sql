alter table daily_stat
    rename column dates to date_daily;
    
alter table daily_stat
    alter column date_daily type date using date_daily::date;

alter table daily_stat
    add column reach bigint not null default 0;
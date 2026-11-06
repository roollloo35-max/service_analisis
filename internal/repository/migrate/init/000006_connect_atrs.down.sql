alter table daily_stat
    alter column conversion type int;

alter table daily_stat
    drop constraint fk_daily_stat_conversion_event;
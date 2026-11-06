alter table daily_stat
    rename column date_daily to dates;
    
alter table daily_stat
    alter column dates type timestamptz using dates::timestamptz;

alter table daily_stat
    drop column reach;
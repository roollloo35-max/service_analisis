alter table campaigns
    rename column start_date to start_dates;

alter table campaigns
    rename column name to names;

alter table campaigns
    rename column status to statuses;
    
alter table campaigns
    alter column start_dates type timestamptz using start_dates::timestamptz,
    alter column end_date type timestamptz using end_date::timestamptz;
    

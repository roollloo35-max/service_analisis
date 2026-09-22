alter table campaigns
    rename column start_dates to start_date;

alter table campaigns
    alter column start_date type date using start_date::date,
    alter column end_date type date using end_date::date;


    
    
alter table campaigns
    rename column names to name;

alter table campaigns
    rename column statuses to status;
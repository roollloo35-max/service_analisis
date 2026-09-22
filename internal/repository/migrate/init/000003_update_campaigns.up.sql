alter table campaigns
    rename column start_dates to start_date;

alter table campaigns
    alter column start_date type date using start_date::date;



    
    
alter table campaigns
    rename column names to name;

alter table campaigns
    rename column statuse to status;
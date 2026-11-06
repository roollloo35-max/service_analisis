alter table campaigns 
    drop constraint fk_campaigns_ad_platform;
    
alter table campaigns 
    drop column external_id,
    drop column delete_atr,
    drop column version_atr;

alter table users
    drop column delete_atr,
    drop column version_atr;


drop table ad_platform;
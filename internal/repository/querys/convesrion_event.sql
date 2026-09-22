-- name: CreateConversionEvent :one

insert into conversion_event
(

    campaign_id,
    occurred_at,
    amount

) values ($1, $2, $3) returning id, campaign_id, occurred_at, amount;


-- name: GetConversion :one 

select id, campaign_id, occurred_at, amount from  conversion_event
    where id = $1;


-- name: UpDataConversion :one

update conversion_event
    set 
        campaign_id =   coalesce(sqlc.narg('new_campaign_id'),campaign_id),
        occurred_at =   coalesce(sqlc.narg('new_occurred_at'),occurred_at),
        amount      =   coalesce(sqlc.narg('new_amount'),     amount)
    where id = sqlc.arg('id') returning id, campaign_id, occurred_at, amount;


-- name: CreateDailyStat :one

insert into daily_stat(
    campaign_id,
    date_daily,
    impressions,
    clicks,
    cost,
    conversion,
    revenue, 
    reach
) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id, campaign_id, date_daily, impressions, clicks,  cost, conversion, revenue, reach;

-- name: GetDailys :many

select  id,campaign_id, date_daily, impressions,
    clicks,
    cost,
    conversion,
    revenue, 
    reach from daily_stat
    order by id
    limit $1 offset $2;

-- name: GetDaily :one 
    select  id,campaign_id, date_daily, impressions,
    clicks,
    cost,
    conversion,
    revenue, 
    reach from daily_stat
    where id = $1;


-- name: UpdateDaily :one 
    update daily_stat
    set 
        campaign_id =   coalesce(sqlc.narg('new_campaign_id'), campaign_id),
        date_daily  =   coalesce(sqlc.narg('date_daily'),      date_daily),
        impressions =   coalesce(sqlc.narg('impressions'),     impressions),
        clicks      =   coalesce(sqlc.narg('clicks'),          clicks),
        cost        =   coalesce(sqlc.narg('cost'),            cost),
        conversion  =   coalesce(sqlc.narg('conversion'),      conversion),
        revenue     =   coalesce(sqlc.narg('revenue'),         revenue),
        reach       =   coalesce(sqlc.narg('reach'),           reach)
where 
    id = sqlc.arg('id') returning id, campaign_id, date_daily, impressions, clicks,  cost, conversion, revenue, reach;

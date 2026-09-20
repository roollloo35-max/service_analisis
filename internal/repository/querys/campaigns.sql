-- name: CreateCampaign :one
insert into
    campaigns (
        name,
        status,
        budget,
        target_url,
        start_date,
        end_date,
        last_update
    )
values
    ($1, $2, $3, $4::text, $5, $6, now()) 
    returning 
    id,
    name,
    status,
    budget,
    target_url,
    start_date,
    end_date,
    last_update;

-- name: GetCampaign :one
select
    id,
    name,
    status,
    budget,
    target_url,
    start_date,
    end_date,
    last_update
from
    campaigns
where
    id = $1;

-- name: GetAllCampaigns :many
select
    id,
    name,
    status,
    budget,
    target_url,
    start_date,
    end_date,
    last_update
from
    campaigns
order by id;

-- name: UpdateCampaign :one
update
    campaigns
set
    name = coalesce(sqlc.narg('new_name'), name),
    status = coalesce(sqlc.narg('new_status'), status),
    budget = coalesce(sqlc.narg('new_budget'), budget),
    target_url = coalesce(sqlc.narg('new_target_url')::text, target_url),
    start_date = coalesce(sqlc.narg('new_start_date'), start_date),
    end_date = coalesce(sqlc.narg('new_end_date'), end_date),
    last_update = now()
where
    id = sqlc.arg('id') returning 
    id,
    name,
    status,
    budget,
    target_url,
    start_date,
    end_date,
    last_update;
-- name: CreateCampaign :one
insert into
    campaigns (
        name,
        status,
        budget,
        target_url,
        start_date,
        end_date,
        last_update,
        external_id,
        delete_atr,
        version_atr
    )
values
    ($1, $2, $3, $4, $5, $6, now(), $7, $8, $9) 
    returning 
    id,
    name,
    status,
    budget,
    created_at,
    target_url,
    start_date,
    end_date,
    last_update,
    external_id,
    delete_atr,
    version_atr;

-- name: GetCampaign :one
select
    id,
    name,
    status,
    budget,
    target_url,
    start_date, 
    created_at,
    end_date,
    last_update,
    external_id,
    delete_atr,
    version_atr
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
    created_at,
    end_date,
    last_update,
     external_id,
    delete_atr,
    version_atr
from
    campaigns
order by id offset $1 limit $2;

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
    last_update = now(),
    external_id = coalesce(sqlc.narg('new_external_id'), external_id),
    delete_atr = coalesce(sqlc.narg('new_delete_atr'), delete_atr),
    version_atr = coalesce(sqlc.narg('new_version_atr'), version_atr)
where
    id = sqlc.arg('id') returning 
    id,
    name,
    status,
    budget,
    target_url,
    start_date,
    end_date,
    last_update,
    external_id,
    delete_atr,
    version_atr;


-- name: CreatePlatform :one
insert into ad_platform 
    (code, 
    name_platform, 
    created_at) 
    values ($1, $2, now()) 
    returning id,code, name_platform, created_at;


-- name: GetAdPlatform :one

select 
    id, 
    code, 
    name_platform, 
    created_at 
    from ad_platform where id = $1;

-- name: GetListAdPlatform :many
select 
    id, 
    code, 
    name_platform, 
    created_at 
    from ad_platform order by id offset $1 limit $2;


-- name: UpdateAdPlatform :one
update 
    ad_platform 
set
    code = coalesce(sqlc.narg('new_code'), code),
    name_platform = coalesce(sqlc.narg('new_name_platform'), name_platform)
where
    id = sqlc.arg('id') returning
    id,
    code,
    name_platform,
    created_at;


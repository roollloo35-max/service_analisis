-- name: CreateUser :one
insert into users (
    email,
    password_hash
) values (
    $1, $2
) returning *; 


-- name: GetUserByEmail :one
select * from users
where email = $1 limit 1;

-- name: UpвateUsers :one

update users
set 
    password_hash = coalesce(sqlc.narg('new_password_hash'), password_hash),
    email = coalesce(sqlc.narg('new_email'), email)
where 
    email = sqlc.arg('target_email')
returning *;

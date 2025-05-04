-- name: LoginUserWithUsername :one
SELECT username, passwdhash FROM users WHERE username = $1;  

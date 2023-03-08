module go_manager_db

go 1.18

require (
	github.com/jmoiron/sqlx v1.3.5
	github.com/mattn/go-sqlite3 v1.14.16
	go_manager_utils v0.0.0-00010101000000-000000000000
)

replace go_manager_utils => ../utils

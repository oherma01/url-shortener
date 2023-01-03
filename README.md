# url-shortener

A simple url shortener written in Go, using PostgreSQL as a database, and Svelte with JavaScript for the frontend.

GORM is used for the database, and the database is automatically created and migrated on startup.
Fiber is used for the webserver, and the frontend is served from the same server.
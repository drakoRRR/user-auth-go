# User Register Application
## Start Application
* To start application
`make local-up`

Backend reach by: `http://localhost:8080`

## Migrations

* To apply migrations
`make migrate-up`

* To down migrations
`make migrate-down`

* To create migration
`make migration name_of_migration`

## Tests
* To run tests
`make test`

## Docs
![](media_readme/swagger.png)
Reach by: `http://localhost:8080/swagger/index.html`

* To update docs
`make update-docs`
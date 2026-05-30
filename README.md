# mimer

Mimer is a _well of knowledge_, a single _source of thruth_,
which can be consumed by applications.

Right now, it provides information about
the **committees** and **trustees** of the IT chapter.

## Developing

To set the project up, you will need:

- [Golang](https://go.dev/) for coding
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations
- [sqlc](https://docs.sqlc.dev/) for generating structs and functions from SQL

### Frontend

- [templ](https://templ.guide/) for server-side rendering
- [Tailwind CSS](https://tailwindcss.com/) for styling
  - We are using the [Tailwind CLI](https://tailwindcss.com/docs/installation/tailwind-cli), so if you want to make changes to the Tailwind configuration you need to have it installed.

### Yaak

If you are using [Yaak](https://yaak.app/) for local API testing you can find workspace files in `yaak/` so you do not have to write requests for all endpoints manually.

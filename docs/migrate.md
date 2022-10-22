# Introduction to Database Migration Tools

If you already understand Flyway's design philosophy or its ideas and using experience,
you can skip introduction section and go directly to the [Getting Started](gettingstarted) section.
This article only covers a simple start. It is recommended to read the official documentation (15 minutes to get
started)

## What tools are available for Golang Developer?

- Migrate: [GitHub Repo](https://github.com/golang-migrate/migrate)
- Atlas: [Atlas](https://atlasgo.io)

Other:

- Flyway: [Flyway](https://flywaydb.org/) CLI supported

Reference standards for use of suitable tools:

1. Function: Refer to the functions of Flyway, the data migration tool in Java next door
2. Based on its GitHub and community
3. Easy to use

## Migrate and Atlas introduction

- Migrate is the most used library, and many other libraries are based on his ideas or secondary development, such as
  Atlas.
- The advantage of Migrate is that it is simple, whether it is a CLI or want to integrate it in a project and use code
  to control it through a Package.
- Atlas has very rich documentation, incorporates a lot of its own ideas, and provides community support. The
  disadvantage is that it has a higher learning cost than Migrate. If you need to use version control, you still have to
  borrow other tools.

## Why use database migration tools?

1. https://flywaydb.org/documentation/getstarted/why

Other advantages (personal opinion):

1. Great for unit and integration testing without polluting the stage/dev/test database
2. Occupy less resources, you needn't to prepare multiple databases for different environments and different versions to
   occupy resources
3. No confusion
4. Applicable to various situations, it can well avoid database version conflicts between different function and
   different project groups

# How to use the migration tools

Use Docker to start MySql and PostgreSql as test database

```bash
# --platform the parameters are for macos M1/M2 chips
# Other parameters can also be set by themselves without any effect

# MySQL 8
docker run --platform linux/amd64 --name mysqltmp -e MYSQL_ROOT_PASSWORD=pwd -p 3306:3306 -v `pwd`/mysql8:/var/lib/mysql -d mysql:latest

# PostgreSQL
docker run --name postgresql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pwd -p 5432:5432 -v `pwd`/pstgres:/var/lib/postgresql/data -d postgres
```

## Migrate CLI Usage

Before starting, you need to prepare an empty database in the container

### CLI

1. Installation

```bash
# Mac Homebrew
brew install golang-migrate

# Windows Scoop
scoop install migrate

# Other
# See Doc: Plz read https://github.com/golang-migrate/migrate/tree/master/cmd/migrate 
```

2. Command: -help
   View available commands with help

```bash
# Help
migrate -help

# Output:
Usage: migrate OPTIONS COMMAND [arg...]
       migrate [ -version | -help ]

Options:
  -source          Location of the migrations (driver://url)
  -path            Shorthand for -source=file://path
  -database        Run migrations against this database (driver://url)
  -prefetch N      Number of migrations to load in advance before executing (default 10)
  -lock-timeout N  Allow N seconds to acquire database lock (default 15)
  -verbose         Print verbose logging
  -version         Print version
  -help            Print usage

Commands:
  create [-ext E] [-dir D] [-seq] [-digits N] [-format] [-tz] NAME
	   Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
	   Use -seq option to generate sequential up/down migrations with N digits.
	   Use -format option to specify a Go time format string. Note: migrations with the same time cause "duplicate migration version" error.
           Use -tz option to specify the timezone that will be used when generating non-sequential migrations (defaults: UTC).

  goto V       Migrate to version V
  up [N]       Apply all or N up migrations
  down [N] [-all]    Apply all or N down migrations
	Use -all to apply all down migrations
  drop [-f]    Drop everything inside database
	Use -f to bypass confirmation
  force V      Set version V but don not run migration (ignores dirty state)
  version      Print current migration version

Source drivers: bitbucket, github-ee, godoc-vfs, file, gcs, s3, github, gitlab, go-bindata
Database drivers: firebirdsql, mongodb+srv, postgres, sqlserver, stub, mongodb, neo4j, cassandra, mysql, postgresql, redshift, spanner, clickhouse, cockroach, cockroachdb, crdb-postgres, firebird, pgx
```

3. Command: create

```bash
# -ext file extension
# -dir directory
# -seq migrate to the N order up or down
migrate create -ext sql -dir db/migration -seq sample

# Output:
# The generated file will be two files in the format of [sequence number]_[file name].[up/down].[extension], up is the upgrade file, down is the rollback file
# It is similar to Flyway
/{path}/000001_sample.up.sql
/{path}/db/migration/000001_sample.down.sql
```

Since we chose to use the sql extension in the previous steps, we can edit the file with SQL.

4. Start to migrate

```bash
migrate -path db/migration -database "{database}://{user}:{pwd}@{database url}?{parameters}" -verbose {command}

# postgresql example
# sslmode=disable
migrate -path db/migration -database "postgresql://root:root@localhost:33062/sample?sslmode=disable" -verbose up

# mysql example
# You may need to add tcp connection
migrate -path db/migration -database "mysql://root:root@tcp(localhost:33062)/sample?" -verbose goto 1
```

5. Give an error example
   The local database uses Mysql 8 as this example
   The sql in the 000001_sample.up.sql is:

```sql
CREATE TABLE users
(
    id         BIGSERIAL primary key,
    first_name TEXT not null,
    last_name  TEXT,
    created_at TIMESTAMP default now()
);
```

After executing the migration
statement `migrate -path db/migration -database "mysql://root:root@tcp(localhost:33062)/sample?" -verbose goto 1` will
get the following error message:

```bash
2022/09/26 23:11:14 Start buffering 1/u sample
2022/09/26 23:11:14 Read and execute 1/u sample
2022/09/26 23:11:14 error: migration failed in line 0: CREATE TABLE users (
   id BIGSERIAL primary key,
   first_name TEXT not null,
   last_name TEXT,
   created_at TIMESTAMP default now()
);
 (details: Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'BIGSERIAL primary key,
   first_name TEXT not null,
   last_name TEXT,
   create' at line 2)
```

Fix SQL in `000001_sample.up.sql`:

```sql
CREATE TABLE users
(
    id         BIGINT primary key,
    first_name TEXT not null,
    last_name  TEXT,
    created_at TIMESTAMP default now()
);
```

Then executing it again with
command `migrate -path db/migration -database "mysql://root:root@tcp(localhost:33062)/sample?" -verbose goto 1`,
you will get error message like: `2022/09/26 23:17:00 error: Dirty database version 1. Fix and force version.`.
This is similar to Flyway that each version can only be executed once, and a record will exist for each execution.
Every operation you executed with the CLI is recorded in a table called `schema_migrations`. The CLI tool will create it
the first time the command is executed if it does not exist.

So in order to solve the above error, there is only one way to upgrade the version.
Our .down file was originally used for downgrade, but since version 1 was not created successfully, the .down file will
also fail.
We need the force command to solve it.

5. Upgrade migration
   Create a new migration file, the new migration file will add a new serial number to the file name in order, and
   perform the migration as before
   All migration files are executed in order, and all new migration files only need to add the changed parts, not the
   complete sql of the entire database

## Atlas

- Atlas CLI is an open source tool that helps developers manage their database schema by applying modern DevOps
  principles
- Atlas DDL is a declarative, Terraform-like configuration language designed to capture an organization's data topology.
  Currently, it supports defining schemas for SQL databases such as MySQL, Postgres, SQLite, and MariaDB

> DDL - Data Definition Language (DDL) is a subset of SQL. It is a language used to describe data in a database and its
> relationships.
> You can generate DDL for database objects in scripts to: Keep a snapshot of the database structure

It provides migration in two ways: declarative, versioned

### CLI

Before starting, you need to prepare a database and a Users table. The table structure is:

```sql
CREATE table users
(
    id   BIGINT PRIMARY KEY,
    name varchar(50)
);
```

1. Installation
   MacOS:
   `brew install ariga/tap/atlas`

Other OS: `Plz read https://atlasgo.io/getting-started Installation section`

2. Create a new database

```bash
atlas schema inspect -u "mysql://{usr}:{pwd}@{ip}:{port}/{dbname}" > {name}.hcl

# Output: no return structure
# If the output file path is not specified, the specified file will be stored in the path when the command is executed, and the file extension is .hcl
# The content of the file is similar to the following struct format
 table "users" {
  schema = schema.example
  column "id" {
    null = false
    type = int
  }
  column "name" {
    null = true
    type = varchar(100)
  }
  primary_key {
    columns = [column.id]
  }
}

```

3. Add a new table named `blog_posts`
   The table format is as follows
   ![blog_post](https://atlasgo.io/uploads/images/blog-erd.png)

Change the content of the .hcl file and add the following:

```
table "blog_posts" {
  schema = schema.example
  column "id" {
    null = false
    type = int
  }
  column "title" {
    null = true
    type = varchar(100)
  }
  column "body" {
    null = true
    type = text
  }
  column "author_id" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "author_fk" {
    columns     = [column.author_id]
    ref_columns = [table.users.column.id]
  }
  index "author_id" {
    unique  = false
    columns = [column.author_id]
  }
}
```

4. Migration
   a. Declarative migration

```bash
atlas schema apply \
  -u "mysql://root:pass@localhost:3306/example" \
  -f schema.hcl

# Output: Select Apply or Abort. 
# To solve the error iF there is an error (the above example buried the error content, the error message is friendly, and it is easy to solve). 
# Check the database and find that the table has been created successfully  
```

b. Version migration
Change the respective .hcl files to change the column named `name` in the Users table to `varchar(150)

```bash
# use command `atlas migrate diff` to check difference between files
atlas migrate diff create_blog_posts \
  --dir="file://migrations" \
  --to="file://schema.hcl" \
  --dev-url="mysql://root:pass@:3306/test"

# Output: This control file.sql file, which contains the modified content, can be used by tools such as golang-migration
```

Migrate using your favorite migration tool

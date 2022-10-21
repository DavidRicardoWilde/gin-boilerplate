# gin-boilerplate
A starter project with Golang and Gin framework.

Quickstart web service integrates different tools you may need.
Do not worry about import too many redundant tools into the project that make project too bloated,
because the tools used are all in different git branches,
you can only merge/rebase the needed ones into the development branch.

## Tools
- [x] [Gin](https://gin-gonic.com/docs/) -- web server framework
- [x] [Native Database Client](https://pkg.go.dev/database/sql) -- database connection
- [x] [GORM](https://gorm.io/docs/index.html) -- database connection with ORM
- [x] [Viper](https://github.com/spf13/viper) -- configuration solution
- [x] [Logrus](https://github.com/sirupsen/logrus) -- logging solution
- [x] [Testify](https://github.com/stretchr/testify) -- unit testing

## How to use this boilerplate
The core branch is mater, which is the basic use of Gin, including routing and middleware. 
If other tools are needed, you can be merged into master from other branches. 
The tools currently include: database connection, logging, configuration solution, unit testing, see the 'Tools' list for full details.

# API CRUD Point-of-Sales (POS) System 

## Business Requirement
You would able to create invoice, edit, get detail & get list of Invoices. Each of Invoice have a customer data and items that being purchased.

You could also see a Demo Video to give you better understanding how to use these APIs.

## Database Migration
The system could migrate SQL scripts which will be triggered [in the system startup](https://github.com/akbar12/wzly/blob/main/implementation-test-1-crud-api/src/util/database.go#L34).

The new SQL script must be placed in the [migration folder](https://github.com/akbar12/wzly/tree/main/implementation-test-1-crud-api/src/util/migrations).

## Authentication
The system use JWT token for Authentication. You could see it's implementation in the [http request's middleware](https://github.com/akbar12/wzly/blob/main/implementation-test-1-crud-api/src/util/middleware.go).

## Logging
The system prints log (error & info) for any catched error & optional info. It will be formatted as JSON. You could see it's implementation in the [logger](https://github.com/akbar12/wzly/blob/main/implementation-test-1-crud-api/src/util/logger.go).

## Unit Test
The system equiped with some unit tests to prevent breaking changes. Unit test's files have the `_test.go` prefixes. You could see it's implementaion in the common [util](https://github.com/akbar12/wzly/tree/main/implementation-test-1-crud-api/src/util), [business process logics](https://github.com/akbar12/wzly/tree/main/implementation-test-1-crud-api/src/model) and [controller](https://github.com/akbar12/wzly/tree/main/implementation-test-1-crud-api/src/controller). 

## Design Patter
The system also implements some of the well known design pattern, such as Singleton and Facade. 

### Singleton
You could see Singleton implementation in the [database initialization](https://github.com/akbar12/wzly/blob/main/implementation-test-1-crud-api/src/util/database.go#L22) in [main.go](https://github.com/akbar12/wzly/blob/main/implementation-test-1-crud-api/main.go#L23). 

### Facade
The could see Facade implementation in the [logger](https://github.com/akbar12/wzly/blob/main/implementation-test-1-crud-api/src/util/logger.go) util. The logger will simplify the utilization of logging library.

## System Installation

### System Requirement
- Golang +1.16
- MySQL 5.7.xx

### Step by Step Installation
- [Download & Install](https://go.dev/doc/install) Golang +1.16 
- [Download & Install](https://dev.mysql.com/downloads/windows/installer/5.7.html) MySQL 5.7.xx
- Create an empty database & import a database dump that has been attached along with this system. Please adjust the API's database configuration in the config.json based on your database config (host, port, username, password, dbname, etc)
- Install this system by build & run the API source code (in Golang). Go to /api root folder and run `go mod vendor` to download all the dependencies needed. 
- Run the API source code by running `go build && /pos-api.exe` (for windows) or `go build && ./pos-api` (for linux). You will see message `listening on port :9999` in terminal that indicates the API is running.
- Test / Try the API by using a Postman Collection that has been attached along with this system. Dont forget to replace Postman Env Variable {{api-pos}} with `http://localhost/9999`

### Try it
To use or try these APIs, you could use a [Postman](https://www.postman.com/downloads/) Collection that has been attached along with this system.

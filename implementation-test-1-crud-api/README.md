
# Point of Sales 

This system is APIs of simple Invoice Management System.

To use / try these APIs, you could use a [Postman](https://www.postman.com/downloads/) Collection that has been attached along with this system.

You could also see a Demo Video to give you better understanding how to use these APIs.

# System Installation

## Software
- Golang +1.16
- MySQL 5.7.xx

## Step by Step
- [Download & Install](https://go.dev/doc/install) Golang +1.16 
- [Download & Install](https://dev.mysql.com/downloads/windows/installer/5.7.html) MySQL 5.7.xx
- Create an empty database & import a database dump that has been attached along with this system. Please adjust the API's database configuration in the config.json based on your database config (host, port, username, password, dbname, etc)
- Install this system by build & run the API source code (in Golang). Go to /api root folder and run `go mod vendor` to download all the dependencies needed. 
- Run the API source code by running `go build && /pos-api.exe` (for windows) or `go build && ./pos-api` (for linux). You will see message `listening on port :9999` in terminal that indicates the API is running.
- Test / Try the API by using a Postman Collection that has been attached along with this system. Dont forget to replace Postman Env Variable {{api-pos}} with `http://localhost/9999`

## Notes
You may see a lot of weakness / holes in this source code, because I done it mostly in Working days. I am also running out of time to create any unit tests. So don't hesitate to tell me for any suggestion / improvement. Thank you
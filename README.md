# Library Management System v2.0
 
 This is a simple library management system implemented in <span style="color:skyblue">Go</span>. It uses <span style="color:yellow">mysql</span> as the database. It has same features as the previous Library Management System implemented in <span style="color:green">node.js</span>. It is based on MVC architecture.

 ## Setup 
 ### Manual
 - Clone the repo. From root directory:
    ```zsh
    go mod vendor
    go mod tidy
    cp .env.example .env
    ```
- Update the `.env` file with your database credentials.
- Run `chmod +x ./cmd/startup.sh` to make the script executable.
- Run `./cmd/startup.sh` to start the process of hosting the server and also add control for migrations.
- It uses the `Makefile` to complete the database. (**Note:** The databse is created completely after migration version `4` and `5` just adds some dummy data for testing. Version above `5` is for testing purposes only.) 
- Makefile after startup.sh has following utilities:
    ```zsh
    make build
    make run
    make clean
    make host
    make migration_create
    make migration_up
    make migration_down
    make migration_fix
    ```
- The `startup.sh` will make up the database upto version `5`.
- It will also allow you to set default admin credentials.
- For running the server on localhost, make sure that in `/pkg/controller/login-register.go` the `Domain` in `SetCookie` function is set to empty string or else the cookies will not be set. And for virtual host, set the `Domain` to the domain name (Default: `mvc.libmansys.com`).
- Use `make host` to host the server.
### Docker
- Run the following command.
    ```zsh
        cp .env.example .env
    ```
- Change the variables as per your use.
- Compose docker.
    ```zsh
        docker compose up --build
    ```

## Unit Test
- The unit test is available for `InsertUser` function in `models` package.
To run the test use the following command from root directory:
    ```zsh
    go test ./pkg/models
    ```
- Another unit test is available for `HashPassword` function in `utils` package.
To run the test use the following command from root directory:
    ```zsh
    go test ./pkg/utils
    ```
## Migration
- For migration testing, use the Makefile. The Makefile has the following commands:
    ```zsh
    make migration_create
    make migration_up
    make migration_down
    make migration_fix
    ```
    
## Features

- Users can search for a book using title, filter to view books of a particular genre, written by a particular author. 
- Users can have only 1 copy of a particular book. They have to request admin for checkin/checkout of a book.
- Users can see their borrowing history.
- Users are notified on visiting the site regarding their status of their requests.
- Admin can add a new book, update quantity of a book, delete a book, view all books, review requests from users.
- The entire system is based on MVC architecture.
- Database is formed using migrations and hence can be easily updated to a new version.
- The system is secure and uses JWT for authentication.
- Run `make` to see the list of commands available.
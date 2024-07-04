# Library Management System v1.3
 
 This is a simple library management system implemented in <span style="color:skyblue">Go</span>. It uses <span style="color:yellow">mysql</span> as the database. It has same features as the previous Library Management System implemented in <span style="color:green">node.js</span>. It is based on MVC architecture.

 ## Setup
 - Clone the repo. From root directory:
    ```zsh
    go mod vendor
    go mod tidy
    cp ./config/.env.example ./config/.env
    ```


- __Database Setup__
    1. Run `mysql -u root -p < database/db.sql` from root directory.
    2. Update the `config/.env` file with your database credentials.
    3. Use the Makefile to create the database and tables. (Database is setup in form of migrations. So, to jump to the most stable model, enter `5` when asked for version)

- __Running the server__
    ```zsh
    go run cmd/main.go
    ```
- __Unit Test__
    
    The unit test is available for `InsertUser` function in `models` package.
    To run the test use the following command from root directory:
    ```zsh
    go test ./pkg/models
    ```
- __Migration__

    For migration testing, use the Makefile. The Makefile has the following commands:
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

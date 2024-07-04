# Library Management System v1.0
 
 This is a simple library management system implemented in <span style="color:skyblue">Go</span>. It uses <span style="color:yellow">mysql</span> as the database. It has same features as the previous Library Management System implemented in <span style="color:green">node.js</span>. It is based on MVC architecture.

 ## Setup
 - Clone the repo. From root directory:
    ```bash
    go mod vendor
    go mod tidy
    cp ./config/.env.example ./config/.env
    ```

- __Database Setup__
    1. Run `mysql -u root -p < config/db.sql` from root directory.

- Running the server
    ```bash
    go run cmd/main.go
    ```
- __Unit Test__
    
    The unit test is available for `InsertUser` function in `models` package.
    To run the test use the following command from root directory:
    ```zsh
    go test ./pkg/models
    ```
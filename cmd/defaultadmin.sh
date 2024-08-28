#!/bin/bash
defaultadmin(){
echo -e "Please enter the default admin credentials(username will be admin)\n"
echo "Password: "
read -r -s pswd
echo "Confirm Password: "
read -r -s confirm_password
if [ -z "$pswd" ] || [ -z "$confirm_password" ]
then
	echo "Password cannot be empty"
	exit 1
fi
if [ "$pswd" != "$confirm_password" ]
then
	echo "Passwords do not match"
	exit 1
fi
go build  -o admincred ./config/admincred.go
hashedPassword=$(./admincred  "$pswd")

if [ $? -ne 0 ]
then
	echo "Error in hashing password"
	exit 1
fi
sqlquery="INSERT INTO users (username, password, role) VALUES ('admin', '$hashedPassword', 'admin');"
mysql -u root -p$password -e "$sqlquery" lib_db

if [ $? -ne 0 ]
then
	echo "Error in inserting admin credentials"
	exit 1
fi
echo -e "Admin credentials inserted successfully\n"
rm admincred
}

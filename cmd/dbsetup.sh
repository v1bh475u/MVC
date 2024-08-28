#!/bin/bash
dbsetup(){
echo -e "Please enter your mysql root user password(necessary for migration)\n"
	echo "Password: "
	read -r -s password
	if [ -z "$password" ]
	then
		echo "Password cannot be empty"
		exit 1
	fi
	mysql -u root -p$password -e "DROP DATABASE IF EXISTS lib_db; CREATE DATABASE lib_db;"
	echo -e "Database created successfully\n"
	curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
	sudo apt-get install -y migrate
	if [ -f "Makefile" ]
	then
		echo -e "Makefile already exists\n"
	else
		mv Makefile.example Makefile
		echo "
	migration_up:
			@read -p \"Enter amount by which you want to up the db: \" v; \\
			migrate -path database/migration/ -database \"mysql://root:$password@tcp(localhost:3306)/lib_db?\" -verbose up	\$\$v
	migration_down:
			@read -p \"Enter amount by which you want to down the db: \" v; \\
			migrate -path database/migration/ -database \"mysql://root:$password@tcp(localhost:3306)/lib_db?\" -verbose down \$\$v
	migration_fix:
			@read -p \"Enter version: \" v; \\
			migrate -path database/migration/ -database \"mysql://root:$password@tcp(localhost:3306)/lib_db?\" force \$\$v
	" >> Makefile
		echo -e "Makefile created successfully\n"
	fi
	echo 5 | make migration_up
}
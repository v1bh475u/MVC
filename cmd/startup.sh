#!/bin/bash
path=$(pwd)
echo -e "This script will install and configure apache2, create a virtual host and help in creation of database.\n"
sudo apt install apache2 -y
sudo a2enmod proxy proxy_http 
cd /etc/apache2/sites-available
if [ -f "mvc.libmansys.local.conf" ]
then
	echo -e "Virtual host already exists\n"
else
	echo -e "Please enter your emailid\n"
	read -r emailid

	if [ -z "$emailid" ]
	then
		echo "Emailid cannot be empty"
		exit 1
	fi

	sudo tee -a mvc.libmansys.local.conf > /dev/null <<EOL
	<VirtualHost *:80>
		ServerName mvc.libmansys.local
		ServerAdmin $emailid
		ProxyPreserveHost On
		ProxyPass / http://127.0.0.1:8080/
		ProxyPassReverse / http://127.0.0.1:8080/
		TransferLog /var/log/apache2/mvc_access.log
		ErrorLog /var/log/apache2/mvc_error.log
	</VirtualHost>
EOL

	sudo a2ensite mvc.libmansys.local.conf
	ip=$(hostname -I | awk '{print $1}')

	echo "$ip	mvc.libmansys.local" | sudo tee -a /etc/hosts > /dev/null

	sudo a2dissite 000-default.conf
	sudo apache2ctl configtest
	sudo systemctl restart apache2
	# sudo systemctl status apache2
	echo -e "Apache2 installed and configured successfully\n"
fi
cd $path
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
	echo 5 | make migration_up
fi
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
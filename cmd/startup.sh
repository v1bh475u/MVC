#!/bin/zsh
path=$(pwd)
sudo apt install apache2 -y
sudo a2enmod proxy proxy_http 
cd /etc/apache2/sites-available

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
	ProxyPass / http://127.0.0.1:8000/
	ProxyPassReverse / http://127.0.0.1:8000/
	TransferLog /var/log/apache2/mvc_access.log
	ErrorLog /var/log/apache2/mvc_error.log
</VirtualHost>
EOL

sudo a2ensite mvc.libmansys.local.conf
echo " 127.0.0.1	mvc.libmansys.local" | sudo tee -a /etc/hosts > /dev/null

sudo a2dissite 000-default.conf
sudo apache2ctl configtest
sudo systemctl restart apache2
# sudo systemctl status apache2
echo -e "Apache2 installed and configured successfully\n"
echo -e "Please visit http://mvc.libmansys.local to view the application\n"

echo -e "Please enter your mysql username(necessary for migration)\n"
echo "Username: "
read -r username
if [ -z "$username" ]
then
	echo "Username cannot be empty"
	exit 1
fi
echo "Password: "
read -r -s password
if [ -z "$password" ]
then
	echo "Password cannot be empty"
	exit 1
fi
cd $path
# mv Makefile.example Makefile
echo "
migration_up:
		migrate -path database/migration/ -database \"mysql://$username:$password@tcp(localhost:3306)/lib_db?\" -verbose up	
migration_down:
		migrate -path database/migration/ -database \"mysql://$username:$password@tcp(localhost:3306)/lib_db?\" -verbose down
migration_fix:
		@read -p \"Enter version: \" v; \\
		migrate -path database/migration/ -database \"mysql://$username:$password@tcp(localhost:3306)/lib_db?\" force \$\$v
" >> Makefile
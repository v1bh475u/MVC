#!/bin/bash
apachesetup(){
echo -e "Would you like to install and configure apache2?(y/n)\n"
read -r reply
if [ "$reply" != "y" ]
then
	echo -e "Skipping apache2 installation and configuration\n"
else
	sudo apt update
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
fi
}
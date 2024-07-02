#!/usr/bin/env zsh

sudo apt install apache2 -y
sudo a2enmod proxy proxy_http 
cd /etc/apache2/sites-available

echo -n "Please enter your emailid"
read emailid
echo emailid
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
sudo systemctl status apache2
echo -n "Apache2 installed and configured successfully"
echo -n "Please visit http://mvc.libmansys.local to view the application"

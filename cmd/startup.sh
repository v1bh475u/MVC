#!/bin/bash
path=$(pwd)
source dbsetup.sh
source virtualhost.sh
source defaultadmin.sh
echo -e "This script will install and configure apache2, create a virtual host and help in creation of database.\n"
apachesetup
echo -e "Would you like to setup mysql database?(y/n)\n"
read -r reply
if [ "$reply" != "y" ]
then
	echo -e "Skipping mysql database setup\n"
	exit 0
else
	cd $path
	dbsetup
fi
defaultadmin
./mvc
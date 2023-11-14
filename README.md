# contentServer

sudo apt-get install postgresq

https://ubuntu.com/tutorials/install-and-configure-nginx#2-installing-nginx
sudo apt install nginx


https://www.postgresqltutorial.com/install-postgresql-linux/
sudo -i -u postgres
psql
CREATE DATABASE content_server_development;
ALTER ROLE postgres WITH PASSWORD 'password';
\q
sudo systemctl restart postgresql

sudo nano /etc/nginx/sites-available/jukedec

upstream buffalo_app {
    server 127.0.0.1:3000;
}

server {
    listen 80;
    server_name jukedec.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        proxy_pass       http://buffalo_app;
    }
}

sudo service nginx restart

sudo apt-get install golang

https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-mac-linux/#slide-10

https://gofi.sh/#install
gofish init
gofish install buffalo

mkdir ~/go/src/github.com/
mkdir ~/go/src/github.com/jukedec

sudo apt install nodejs
sudo apt install npm

buffalo new contentServer

nano database.yaml

buffalo build

./bin/contentServer

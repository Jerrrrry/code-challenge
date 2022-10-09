# code challenge

## Introduction
cd to project folder 

dev local prerequistes:
install docker

sh start.sh (docker image build for web and gin-api)

Once ready 

web is listening at localhost(127.0.0.1:80)

gin-api is listening at localhost:8080

## GIN API
cd /gin-api

sh rebuild.sh

build image and run a container 


## WEB

cd /web 

test purpose : npm run dev 

or 

sh rebuild.sh 

build image and run a container 



## test cases:

*** 1. domain:google record: A ***

*** 2. domain:google record: AAAA ***

*** 3. domain: google.com record: ANY ***

*** 4. domain : google.com record: CAA ***

*** 5. domain : google.com record: CNAME ***

*** 6. domain : google.com record: MX ***

*** 7. domain : google.com record: NS ***

*** 8. domain : 8.8.8.8 record: PTR***

*** 9. domain : google.com record: SOA ***

*** 10. domain : _xmpp-server._tcp.google.com record: SRV ***

*** 11. domain : google.com record: TXT ***


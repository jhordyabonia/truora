# truora

1) Install CockroachDB.
    - wget -qO- https://binaries.cockroachdb.com/cockroach-v19.1.3.linux-amd64.tgz | tar  xvz
    - cp -i cockroach-v19.1.3.linux-amd64/cockroach /usr/local/bin

2) Start up Cockroach a secure or insecure local cluster:
    - cockroach start --insecure --background --advertise-host= <[host]>
    - cockroach sql --insecure
    - >CREATE USER IF NOT EXISTS jhordy;
    - >CREATE DATABASE truora;
    - >GRANT ALL ON DATABASE truora TO jhordy;
    - >\q
3) Install Whois
    - sudo apt-get install whois

4) Install library go - Cockroach
    - go get -u github.com/lib/pq

5) Install library to go - Router CHI
    - go get -u github.com/go-chi/chi

6) Build
    - go build

7) Run
    - ./truora
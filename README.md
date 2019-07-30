# truora

1) Install CockroachDB.
- wget -qO- https://binaries.cockroachdb.com/cockroach-v19.1.3.linux-amd64.tgz | tar  xvz
- cp -i cockroach-v19.1.3.linux-amd64/cockroach /usr/local/bin

2) Start up Cockroach a secure or insecure local cluster:
-cockroach start --insecure --background --advertise-host= [host]>

3) Install library to go-Cockroach
-go get -u github.com/lib/pq
#!/bin/bash

export VAULT_ADDR=http://127.0.0.1:8200
# shellcheck disable=SC2125
export VAULT_TOKEN=nD^MNCD@pPff*^2Q
#export ALLOWED_ROLES=bardia
#export PORT_POSTGRES=5432
#export CONTAINER_POSTGRES=psg_demo
#export DATABASE_NAME=bardiadb
#export DATABASE_USER_NAME=bardia
#export DATABASE_PASSWORD=bardiapw
#export DEFAULT_TLL=1h
#export MAX_TTL=24h

vault login nD^MNCD@pPff*^2Q

sleep 1
echo '******** secrets enable database ********'
vault secrets enable database

sleep 1
echo '******** Configure the postgresql plugin ********'
vault write database/config/bardiadb \
    plugin_name=postgresql-database-plugin \
    allowed_roles="bardia" \
    connection_url="postgresql://{{username}}:{{password}}@psg_demo:5432/bardiadb?sslmode=disable" \
    username="bardia" \
    password="bardiapw"

sleep 1
echo '******** Configure a role to be used ********'
 vault write database/roles/bardia \
    db_name="bardiadb" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

sleep 1
echo '******** Generate credentials on the DB from the role ********'
vault read database/creds/bardia

#Renew the lease
#vault lease renew -increment=3600 database/creds/admin/O4RaKi1yWJCl1Ak0xnqb4wd9
#vault lease renew -increment=96400 database/creds/admin/O4RaKi1yWJCl1Ak0xnqb4wd9

#Revoke the lease
#vault lease revoke database/creds/admin/O4RaKi1yWJCl1Ak0xnqb4wd9

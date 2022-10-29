sleep 1
echo '\e[0;34m ✅️******** secrets enable database ******** \e[0m'
vault secrets enable database

sleep 2
echo '\e[0;34m ✅️******** Configure the postgresql plugin ******** \e[0m'
vault write database/config/bardiadb \
    plugin_name=postgresql-database-plugin \
    allowed_roles="bardia" \
    connection_url="postgresql://{{username}}:{{password}}@localhost:5432/bardiadb?sslmode=disable" \
     username="bardia" \
     password="bardiapw"

sleep 2
echo '\e[0;34m ✅️******** Configure a role to be used ******** \e[0m'
 vault write database/roles/bardia \
    db_name="bardiadb" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

sleep 2
echo '\e[0;34m ✅️******** Generate credentials on the DB from the role ******** \e[0m'
vault read database/creds/bardia

#Renew the lease
#vault lease renew -increment=3600 database/creds/admin/O4RaKi1yWJCl1Ak0xnqb4wd9
#vault lease renew -increment=96400 database/creds/admin/O4RaKi1yWJCl1Ak0xnqb4wd9

#Revoke the lease
#vault lease revoke database/creds/admin/O4RaKi1yWJCl1Ak0xnqb4wd9

# Dynamic PostgreSQL credentials using HashiCorp Vault (with Go examples)

### Dynamic Secrets
As mentioned, Vault supports a number of dynamic secret backends for generating secrets dynamically when needed. For example, with the AWS and Google Cloud backends, you can create access credentials based on IAM policies. The Databases backend, meanwhile, generates database credentials based on configured roles.

### Dynamic Secrets:
* are generated on demand
* have limited access based on role
* are leased for a period of time
* can be revoked
* come with an audit trail
```shell
vault server -dev
```

#### Set env variable
```shell
export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=AddYourVaultTokenHere
vault login
```

## PostgresSQL Initial Setup
#### first run command ```make up-pgsql``` for docker compose postgresql
```yaml
username: bardia
role: bardia
dbname: bardiadb
```
#### Run the following commands in order:
* ```docker ps```
* ```docker exec -it <name-container> bash```
* ```psql -U <user-name> -W <db-name>```
* ```CREATE TABLE authors(author_id SERIAL PRIMARY KEY, first_name VARCHAR(100) NOT NULL, last_name VARCHAR(100) NOT NULL);```
* ```INSERT INTO authors (first_name, last_name) VALUES ('Tamsyn', 'Muir'), ('Ann', 'Leckie');```

## Configuring Vault to use our PostgresSQL database
Vault can manage secrets using its Secrets Engines which range from AWS, GCP, Key Value, LDAP, SSH, databases... and so on. See the complete list on their docs.
Secrets Engines are Vault components that store, generate & encrypt secrets. The one that we are interested in is the database engine.
The database engine supports a wide varity of database flavors including but not limited to PostgreSQL, MySQL, Redshift and Elasticsearch.
Let's enable the engine and configure it to use our Postgrs database.

#### Enable the database secrets engine
```shell
vault secrets enable database
```
#### Configure the postgresql plugin
```shell
vault write database/config/bardiadb \
    plugin_name=postgresql-database-plugin \
    allowed_roles="bardia" \
    connection_url="postgresql://{{username}}:{{password}}@localhost:5432/bardiadb?sslmode=disable" \
     username="bardia" \
     password="bardiapw"
```
#### Configure a role to be used
Now, let's create a Vault role that will manage the credential creation in both Vault & Postgres.

When creating a role, we supply:

* Creation statements: Vault will use this to know how to create the user in Postgres whenever we ask for new credentials.
* Revocation statements: Vault will execute these commands against Postgres whenever the credentials have expired (TTL reached).
* TTL (Time-To-Live) of the credentials. Once the credentials expire Vault will execute the Revokation statements and will remove the credentials from its storage.
```shell
 vault write database/roles/bardia \
    db_name="bardiadb" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"
```
#### Generate credentials on the DB from the role
```shell
vault read database/creds/bardia
```
#### output
```shell
Key                Value
---                -----
lease_id           database/creds/bardia/NeFeB5oqoiAzguF68xUCmhD6
lease_duration     1h
lease_renewable    true
password           tiGyUy-it2pSZSLoLxt3
username           v-root-bardia-lMJOPHdemg2mr3QPWXnt-1665731759
```
## Web ui Vault
### secrets engines
<img src="screenshot/1.png">

### secrets database
<img src="screenshot/2.png">

### secrets engines database
<img src="screenshot/3.png">

## go run main.go
```
Connected to PostgreSQL db using user <v-root-bardia-OQZIsxFo71ytoXb3aIKo-1660475644> and password <4u-A74pocYWqzUaSqW5L>                                                                     │
Listening on 127.0.0.1:4030
```
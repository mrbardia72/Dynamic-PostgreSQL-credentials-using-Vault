# Dynamic PostgreSQL credentials using HashiCorp Vault (withGo examples)

Basically, have a web application connect to a database (PostgreSQL) using dynamically generated credentials (username & password), that you can rotate whenever you want and it'll all be transparent to your app.

Vault handles the credentials generation (and thus creating a corresponding username & password in PostgreSQL) and expiration (and thus, removing the username from the DB).

## Vault Initial Setup
In the output, you'll get the three important information:

The API endpoint (which is the same as the UI URL). If you ran the dev server without any arguments this is probably http://127.0.0.1:8200.
The unseal token. This is used to unseal and Vault from its sealed state. Whenever Vault is rebooted and/or initialized, it starts in a sealed state so you'll need to unseal it first. We don't have to worry about this because when using dev server, Vault is already initilized and unsealed.
The root token. The token we'll use to authenticate our requests to the API. This is only a good idea when running a dev server and trying out some stuff, but in the real world the root token is only there for emergencies and for initial setup of users/policies... etc.
Now that we have a Vault server running, leave that Terminal open and open a new one (or a new Tmux pane or whatever).

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
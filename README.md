# go-rest-api-course

### Cloud dependency

#### CockroachDB Serverless database cluster
* PostgreSQL Compatibility
* A generous free-forever tier
* Automatic redundancy
* No database administration
* Explicit spending limits

https://www.cockroachlabs.com/docs/cockroachcloud/quickstart.html
sign in with github account... choose the Serverless plan.
created a cluster in Google's cloud... "bad-mojo" (video game subtitled "The Roach Game")

1. Download the CRDB Client
```bash
curl https://binaries.cockroachdb.com/cockroach-v21.2.2.darwin-10.9-amd64.tgz \
  | tar -xz; sudo cp -i cockroach-v21.2.2.darwin-10.9-amd64/cockroach /usr/local/bin/
```

2. Download the CA certificate
```bash
curl --create-dirs -o $HOME/.postgresql/root.crt \
  -O https://cockroachlabs.cloud/clusters/c32d1222-5275-4205-9bab-8215bd5659e8/cert
```

3. Run this command to connect to your database
```bash
cockroach sql --url "postgresql://postgres:PASSWORD@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/
  defaultdb?sslmode=verify-full&sslrootcert=$HOME/.postgresql/root.crt&options=--cluster%3Dbad-mojo-5100"
```

Add your connection string into your application code.
```bash
postgresql://postgres:<ENTER-PASSWORD>@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/
  defaultdb?sslmode=verify-full&sslrootcert=$HOME/.postgresql/root.crt&options=--cluster%3Dbad-mojo-5100
```

### JWT token creation

* jwt.io: enter missionimpossible within "verify signature" [your-256-bit-secret]
* paw: add https://paw.cloud/extensions/JsonWebTokenDynamicValue
 HS256, payload can immitate jwt.io sub "1234567890", name "John Doe", ist 1516239022
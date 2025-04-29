- pre-init the database:

```bash
sudo -u postgres psql
> CREATE DATABASE contacts;
> \c contacts
> CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    number VARCHAR(16),
    created_at TIMESTAMP DEFAULT
    CURRENT_TIMESTAMP
    );
```

- get contact with specific id:

```bash
curl http://localhost:8080/contacts/<id>
```

- get all contacts:

```bash
curl http://localhost:8080/contacts
```

- add a contact:

```bash
curl -X POST http://localhost:8080/contacts -H "Content-Type: application/json" -d '{"name":"<name>","number":"<number>"}'
```

- update a contact:

```bash
curl -X PUT http://localhost:8080/contacts/<id> -H "Content-Type: application/json" -d '{"name":"<name>","number":"<number>"}'
```

- delete a contact:

```bash
curl -X DELETE http://localhost:8080/contacts/<id>
```

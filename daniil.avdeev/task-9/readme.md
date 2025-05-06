Get all contacts
```
curl http://localhost:8080/contacts
```

Get contact with {id}
```
curl http://localhost:8080/contacts/{id}
```

Add new contact
```
curl -X POST -d '{"name": "bla", "phone": "+7(123)456-78-90"}' http://localhost:8080/contacts
```

Update contact with {id}
```
curl -X PUT -d '{"name": "foo", "phone": "+7(098)765-43-21"}' http://localhost:8080/contacts/{id} 
```

Delete contact with {id}
```
curl -X DELETE http://localhost:8080/contacts/{id}
```


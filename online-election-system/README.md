
# Online election system

An online simulation of real world elections.




## API Reference

#### Add new user

```http
  POST localhost:8080/api/user/add
```

```curl
curl --location --request POST 'localhost:8080/api/user/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "user",
    "name": "Lorem3",
    "email": "Lorem3",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem3",
        "father_name": "Lorem3",
        "dob": "2016-04-08",
        "age": 21,
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "C:/Users/Vidhi/Downloads/529_3_334359_1666357514_Databricks - Generic.pdf"
    }
}'
```

#### Verify user

```http
  POST localhost:8080/api/user/verify/{id}
```
```curl
curl --location --request PUT 'localhost:8080/api/user/verify/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "admin",
    "name": "Lorem1",
    "email": "Lorem1",
    "is_verified": true
}'
```

#### Update user

```http
  POST localhost:8080/api/user/update
```
```curl
curl --location --request POST 'localhost:8080/api/user/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": "2016-04-08",
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": "6c6c2d84de0f09d5007bdef8",
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        "ccebdc4611f903debe1b15ac"
    ]
}'
```

#### Search one user

```http
  POST localhost:8080/api/user/search/{id}
```
```curl
curl --location --request GET 'localhost:8080/api/user/search/{id}' \
--header 'Content-Type: application/json'
```
#### Search multiple users

```http
  POST localhost:8080/api/user/search-by-filter
```
```curl
curl --location --request POST 'localhost:8080/api/user/search-by-filter' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "user",
    "city": "Lorem",
    "state": "Lorem",
    "zip_code": "Lorem",
    "country": "Lorem",
    "is_verified": true
}'
```

#### Delete user

```http
  DELETE localhost:8080/api/user/delete/{id}
```
```curl
curl --location --request DELETE 'localhost:8080/api/user/delete/{id}' \
--header 'Content-Type: application/json'
```

#### Deactivate user

```http
  POST localhost:8080/api/user/deactivate
```
```curl
curl --location --request POST 'localhost:8080/api/user/deactivate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "is_verified": false
}'
```

### Add new election
```http
POST localhost:8080/api/election/add
```

```curl
curl --location --request POST 'localhost:8080/api/election/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "location": "Lorem1",
    "election_date": "2023-04-08",
    "result_date": "2023-04-08",
    "result": "",
    "election_status": "voting"
}'
```

### Add new candidate
```http
PUT localhost:8080/api/candidate/add
```

```curl
curl --location --request PUT 'localhost:8080/api/candidate/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "6368b8691503ded80405a7a8",
    "name": "Lorem1",
    "user_id": "63635d75a68e40fe497eac67",
    "vote_count": "",
    "vote_sign": "C:/Users/Vidhi/Downloads/529_3_334359_1666357514_Databricks - Generic.pdf"
}'
```

### Verify candidate
```http
PUT localhost:8080/api/candidate/verify/{id}
```

```curl
curl --location --request PUT 'localhost:8080/api/candidate/verify/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "admin",
    "name": "Lorem1update",
    "email": "Lorem1update",
    "user_id": "63635d75a68e40fe497eac67",
    "nomination_status": "verified"
}'
```

### Update election
```http
PUT localhost:8080/api/election/update/{id}
```

```curl
curl --location --request PUT 'localhost:8080/api/election/update/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_date": "2022-04-08",
    "result_date": "2022-04-08",
    "election_staus": "Lorem"
}'
```

### Search election
```http
GET localhost:8080/api/election/search/{id}
```

```curl
curl --location --request GET 'localhost:8080/api/election/search/{id}' \
--header 'Content-Type: application/json'
```

### Search multiple election
```http
POST localhost:8080/api/election/search-by-filter
```

```curl
curl --location --request POST 'localhost:8080/api/election/search-by-filter' \
--header 'Content-Type: application/json' \
--data-raw '{
            "location": "Lorem",
            "election_date": "2022-04-08",
            "result_date": "2022-04-08",
            "result": "Lorem",
            "election_status": "Lorem"
}'
```

### Deactivate election
```http
PUT localhost:8080/api/election/deactivate/{id}
```

```curl
curl --location --request PUT 'localhost:8080/api/election/deactivate/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_status": "deactivated"
}'
```

### Cast a vote
```http 
POST localhost:8080/cast-vote
```

```curl
curl --location --request POST 'localhost:8080/cast-vote' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "",
    "candidate_id": ""
}'
```

### Search election result
```http
POST localhost:8080/election-result
```

```curl
curl --location --request POST 'localhost:8080/election-result' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "",
    "location": "",
    "election_date": "",
    "result_date": "",
    "election_status": ""
}'
```

### Search candidate profile

```http
POST localhost:8080/candidates-profile
```

```curl
curl --location --request POST 'localhost:8080/candidates-profile' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "",
    "location": "",
    "election_date": "",
    "election_status": ""
}'
```
## Authors

- [Vidhi Goel](https://www.github.com/gic-vidhi)


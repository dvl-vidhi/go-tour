
# Online election system

An online simulation of real world elections.




## API Reference

#### Add new user

```http
  POST localhost:8080/api/user/add
```

```curl
curl --location --request POST 'localhost:8080/api/user/add \
--header 'Content-Type: application/json' \
--data-raw '{
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
        "document_path": "Lorem"
    }
}'
```

#### Verify user

```http
  POST localhost:8080/api/user/verify
```
```curl
curl --location --request POST 'localhost:8080/api/user/verify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
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
  POST localhost:8080/api/user/search
```
```curl
  curl --location --request POST 'localhost:8080/api/user/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem"
}'
```
#### Search multiple users

```http
  POST localhost:8080/api/user/search-by-filter
```
```curl
curl --location --request POST 'localhost:8080/api/user/search-by-filter' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "Lorem",
    "city": "Lorem",
    "state": "Lorem",
    "zip_code": "Lorem",
    "country": "Lorem",
    "is_verified": true,
    "voted": [
        "ccebdc4611f903debe1b15ac"
    ]
}'
```

#### Delete user

```http
  POST localhost:8080/api/user/delete
```
```curl
curl --location --request POST 'localhost:8080/api/user/delete' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem"
}'
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
    "location": "Lorem",
    "election_date": "Lorem",
    "result_date": "Lorem",
    "result": "Lorem",
    "election_status": "Lorem"
}'
```

### Add new candidate
```http
POST localhost:8080/api/candidate/add
```

```curl
curl --location --request POST 'localhost:8080/api/candidate/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "candidate": [
        {
            "election_id": "",
            "candidate_name": " ",
            "vote_count": " ",
            "vote_sign": " ",
            "nomination_status": " ",
            "nomination_verified_by": " "
        }
    ]
}'
```

### Verify candidate
```http
POST localhost:8080/api/candidate/verify
```

```curl
curl --location --request POST 'localhost:8080/api/candidate/verify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": " ",
    "nomination_status": " "
}'
```

### Update election
```http
POST localhost:8080/api/election/update
```

```curl
curl --location -g --request POST 'localhost:8080/api/election/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_date": " ",
    "result_date": " ",
    "election_staus": " ",
    "location": " "
}'
```

### Search election
```http
POST localhost:8080/api/election/search
```

```curl
curl --location --request POST 'localhost:8080/api/election/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "election_date": " ",
    "result_date": " ",
    "election_staus": " ",
    "candidate_name": " ",
    "location": " "
}'
```

### Deactivate election
```http
DELETE localhost:8080/api/election/deactivate
```

```curl
curl --location -g --request DELETE 'localhost:8080/api/election/deactivate'
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


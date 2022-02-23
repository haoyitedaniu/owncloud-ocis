---
title: Users
weight: 30
geekdocRepo: https://github.com/owncloud/ocis
geekdocEditPath: edit/master/docs/extensions/graph
geekdocFilePath: users.md
---

{{< toc >}}

## Users API

The Users API is implementing a subset of the functionality of the
[MS Graph User resource](https://docs.microsoft.com/en-us/graph/api/resources/user?view=graph-rest-1.0)
The JSON representation of a User handled by the Users API looks like this:

```
{
    "displayName": "Albert Einstein",
    "id": "4c510ada-c86b-4815-8820-42cdf82c3d51",
    "mail": "einstein@example.org",
    "onPremisesSamAccountName": "einstein"
}
```

Our implemenation currently supports only a limited set of Attributes of Users:

| Attribute	| Description |
|---------------|-------------|
| displayName	| The full name of the user, usually a combination for givenname and lastname|
| mail		| The user's email address |
| onPremisesSamAccountName | The loginname/account name of the user|
| id		| An unique, stable readonly identifier for the user that stays the same for the whole lifetime of the User, usually a UUID|
| passwordProfile | Contains the password of the users. This is only present when updating or creating users. It is never returned by the API|


### Reading users

#### `GET /me`

Returns the user object of the currently signed-in user

Example:
```
curl -k 'https://localhost:9200/graph/v1.0/me' -u user:password
```

Response:
```
{
    "displayName": "Albert Einstein",
    "id": "4c510ada-c86b-4815-8820-42cdf82c3d51",
    "mail": "einstein@example.org",
    "onPremisesSamAccountName": "einstein"
}
```

#### `GET /users`

Returns a list of all users 

Example:

```
curl -k 'https://localhost:9200/graph/v1.0/users' -u user:password

```

Response:

```
{
    "value": [
        {
            "displayName": "Albert Einstein",
            "id": "4c510ada-c86b-4815-8820-42cdf82c3d51",
            "mail": "einstein@example.org",
            "onPremisesSamAccountName": "einstein"
        },
        {
            "displayName": "Maurice Moss",
            "id": "058bff95-6708-4fe5-91e4-9ea3d377588b",
            "mail": "moss@example.org",
            "onPremisesSamAccountName": "moss"
        }
    ]
}
```

#### `GET /users/{userid or accountname}`

Example:

```
curl -k 'https://localhost:9200/graph/v1.0/users/058bff95-6708-4fe5-91e4-9ea3d377588b' -u user:password
```

Response:

```
{
    "displayName": "Maurice Moss",
    "id": "058bff95-6708-4fe5-91e4-9ea3d377588b",
    "mail": "moss@example.org",
    "onPremisesSamAccountName": "moss"
}
```

### Creating / Updating Users

#### `POST /users`

Use this to create a new user.

##### Request Body

Note the missing `"id"` Attribute. It will be generated by the server:

```
{
    "displayName": "Example User",
    "mail": "example@example.org",
    "onPremisesSamAccountName": "example",
    "passwordProfile": {
    	"password": "ThePassword"
    }
}
```

##### Response

When successful, the Reponse will return the new user, without the password, but including the newly allocated `"id"`:

```
{
    "displayName":"Example User",
    "id":"c067b139-c91c-4e47-8be6-669156a0587b",
    "mail":"example@example.org",
    "onPremisesSamAccountName":"example"
}
```

#### `DELETE /users/{id}`

Example:

```
curl -k --request DELETE 'https://localhost:9200/graph/v1.0/users/c067b139-c91c-4e47-8be6-669156a0587b' -u user:password
```

When successful the API returns no Response Body and the HTTP status code 204 (No Content)


#### `PATCH /users/{id}`

Updating attributes of a single user can be done with a patch request. The Request Body contains the new values of the attributes
to be updated. E.g. to update the `displayName` Attribute:

```
 curl -k --header "Content-Type: application/json" \
         --request PATCH  --data '{"displayName": "Test User" }' \
	 'https://localhost:9200/graph/v1.0/users/c54b0588-7157-4521-bb52-c1c8ca84ea71' -u user:password
```

Similar to creating a user via `POST`,  the `PATCH` request will return the user object containing the new attribute values.
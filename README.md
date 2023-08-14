
# Identities microsservice

Microsservice API responsible about all identities core data
  
## Informations

### Version

0.1.0

## Content negotiation

### URI Schemes

* http
* https

### Consumes

* application/json

### Produces

* application/json

## All endpoints

### operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/chall-recovery | [post chall recovery](#post-chall-recovery) | Checks if who request recovery ticket are the account's owner |
| POST | /api/login | [post login](#post-login) | Sign in user |
| POST | /api/recovery | [post recovery](#post-recovery) | Open recovery account process |
| POST | /api/register | [post register](#post-register) | Sign up user in service |
  
## Paths

### <span id="post-chall-recovery"></span> Checks if who request recovery ticket are the account's owner (*PostChallRecovery*)

```
POST /api/chall-recovery
```

#### Consumes

* application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| password | `formData` | string | `string` |  | ✓ |  | New account password |
| ticket | `formData` | string | `string` |  | ✓ |  | The ticket identity field |
| validation | `formData` | integer | `int64` |  | ✓ |  | A random validation integer. The number get's a range between 100000 ~ 999999 |

#### All responses

| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#post-chall-recovery-204) | No Content | Recovery account with success |  | [schema](#post-chall-recovery-204-schema) |
| [400](#post-chall-recovery-400) | Bad Request | requested body reaches malformed |  | [schema](#post-chall-recovery-400-schema) |
| [403](#post-chall-recovery-403) | Forbidden | incorrect validation number |  | [schema](#post-chall-recovery-403-schema) |
| [410](#post-chall-recovery-410) | Gone | Recovery ticket has been expired |  | [schema](#post-chall-recovery-410-schema) |
| [500](#post-chall-recovery-500) | Internal Server Error | Occurs unexpected internal error |  | [schema](#post-chall-recovery-500-schema) |

#### Responses

##### <span id="post-chall-recovery-204"></span> 204 - Recovery account with success

Status: No Content

###### <span id="post-chall-recovery-204-schema"></span> Schema

##### <span id="post-chall-recovery-400"></span> 400 - requested body reaches malformed

Status: Bad Request

###### <span id="post-chall-recovery-400-schema"></span> Schema

##### <span id="post-chall-recovery-403"></span> 403 - incorrect validation number

Status: Forbidden

###### <span id="post-chall-recovery-403-schema"></span> Schema

##### <span id="post-chall-recovery-410"></span> 410 - Recovery ticket has been expired

Status: Gone

###### <span id="post-chall-recovery-410-schema"></span> Schema

##### <span id="post-chall-recovery-500"></span> 500 - Occurs unexpected internal error

Status: Internal Server Error

###### <span id="post-chall-recovery-500-schema"></span> Schema

### <span id="post-login"></span> Sign in user (*PostLogin*)

```
POST /api/login
```

#### Consumes

* application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| password | `formData` | string | `string` |  | ✓ |  | The users password. Used to check if who is attemped to sign in is really the account owner |
| user | `formData` | string | `string` |  | ✓ |  | The user identity field. Can be informed email ou username |

#### All responses

| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-login-200) | OK | Logged in with success | ✓ | [schema](#post-login-200-schema) |
| [400](#post-login-400) | Bad Request | requested body reaches malformed |  | [schema](#post-login-400-schema) |
| [403](#post-login-403) | Forbidden | users credential are incorrect |  | [schema](#post-login-403-schema) |
| [500](#post-login-500) | Internal Server Error | Occurs unexpected internal error |  | [schema](#post-login-500-schema) |

#### Responses

##### <span id="post-login-200"></span> 200 - Logged in with success

Status: OK

###### <span id="post-login-200-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| X-JWT | string | `string` |  |  | Token to grant access to user in restrict services |

##### <span id="post-login-400"></span> 400 - requested body reaches malformed

Status: Bad Request

###### <span id="post-login-400-schema"></span> Schema

##### <span id="post-login-403"></span> 403 - users credential are incorrect

Status: Forbidden

###### <span id="post-login-403-schema"></span> Schema

##### <span id="post-login-500"></span> 500 - Occurs unexpected internal error

Status: Internal Server Error

###### <span id="post-login-500-schema"></span> Schema

### <span id="post-recovery"></span> Open recovery account process (*PostRecovery*)

```
POST /api/recovery
```

#### Consumes

* application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `formData` | string | `string` |  | ✓ |  | The user identity field. Can be informed email, username or phonenumber |

#### All responses

| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-recovery-201) | Created | Created ticket recovery with success |  | [schema](#post-recovery-201-schema) |
| [400](#post-recovery-400) | Bad Request | requested body reaches malformed |  | [schema](#post-recovery-400-schema) |
| [404](#post-recovery-404) | Not Found | user not found |  | [schema](#post-recovery-404-schema) |
| [500](#post-recovery-500) | Internal Server Error | Occurs unexpected internal error |  | [schema](#post-recovery-500-schema) |
| [503](#post-recovery-503) | Service Unavailable | Email service is unavailable and fails in sent recovery message. |  | [schema](#post-recovery-503-schema) |

#### Responses

##### <span id="post-recovery-201"></span> 201 - Created ticket recovery with success

Status: Created

###### <span id="post-recovery-201-schema"></span> Schema

[PostRecoveryCreatedBody](#post-recovery-created-body)

##### <span id="post-recovery-400"></span> 400 - requested body reaches malformed

Status: Bad Request

###### <span id="post-recovery-400-schema"></span> Schema

##### <span id="post-recovery-404"></span> 404 - user not found

Status: Not Found

###### <span id="post-recovery-404-schema"></span> Schema

##### <span id="post-recovery-500"></span> 500 - Occurs unexpected internal error

Status: Internal Server Error

###### <span id="post-recovery-500-schema"></span> Schema

##### <span id="post-recovery-503"></span> 503 - Email service is unavailable and fails in sent recovery message

Status: Service Unavailable

###### <span id="post-recovery-503-schema"></span> Schema

###### Inlined models

**<span id="post-recovery-created-body"></span> PostRecoveryCreatedBody**

**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | string| `string` |  | | The ticket recovery ID |  |

### <span id="post-register"></span> Sign up user in service (*PostRegister*)

```
POST /api/register
```

#### Consumes

* application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| address | `formData` | string | `string` |  |  |  |  |
| avatar | `formData` | string | `string` |  |  |  | Image or link that will be used to represents avatar users account |
| birthday | `formData` | string | `string` |  | ✓ |  |  |
| email | `formData` | string | `string` |  | ✓ |  |  |
| name | `formData` | string | `string` |  | ✓ |  | The users complete name |
| password | `formData` | string | `string` |  | ✓ |  | The password is used to log in |
| phoneNumber | `formData` | integer | `int64` |  | ✓ |  |  |
| username | `formData` | string | `string` |  | ✓ |  | The username to be used in project |

#### All responses

| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-register-201) | Created | Created user with success |  | [schema](#post-register-201-schema) |
| [400](#post-register-400) | Bad Request | The user sent some bad data in form |  | [schema](#post-register-400-schema) |
| [409](#post-register-409) | Conflict | users with respective email, username or phone number already exists |  | [schema](#post-register-409-schema) |
| [500](#post-register-500) | Internal Server Error | Unexpected error occurs in internal server |  | [schema](#post-register-500-schema) |

#### Responses

##### <span id="post-register-201"></span> 201 - Created user with success

Status: Created

###### <span id="post-register-201-schema"></span> Schema

##### <span id="post-register-400"></span> 400 - The user sent some bad data in form

Status: Bad Request

###### <span id="post-register-400-schema"></span> Schema

##### <span id="post-register-409"></span> 409 - users with respective email, username or phone number already exists

Status: Conflict

###### <span id="post-register-409-schema"></span> Schema

##### <span id="post-register-500"></span> 500 - Unexpected error occurs in internal server

Status: Internal Server Error

###### <span id="post-register-500-schema"></span> Schema

## Models

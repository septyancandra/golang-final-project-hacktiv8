# MyGram - Aplikasi Media Sosial untuk Berbagi Foto dan Komentar

## About this Project => https://github.com/septyancandra/golang-final-project-hacktiv8

See Postman Collection --->> https://github.com/septyancandra/golang-final-project-hacktiv8/blob/main/final-project-golang-hacktiv8.postman_collection.json
* [Photo](#photo)
* [User](#user)
* [Comment](#comment)
* [Social Media](#social-media)

# Photo

## POST Create Photo

POST /photos

> Body Parameters

```yaml
title: Someday 2
photo_url: https://upload.wikimedia.org/wikipedia/commons/8/82/Pemandangan_Gunung_Kerinci_Dari_Swarga_Village_01.jpg
caption: Pray Until Something Happend

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» title|body|string| yes |none|
|» photo_url|body|string| yes |none|
|» caption|body|string| yes |none|

> Response Examples

> Create First Photo

```json
{
    "id": 3,
    "title": "Someday 2",
    "caption": "Pray Until Something Happend",
    "photo_url": "https://www.instagram.com/septyancandra/",
    "user_id": 1,
    "created_at": "2024-02-16T19:41:23.890961+07:00"
}
```

## GET Get All Photo

GET /photos/add

> Response Examples

> Success Get All Photo

```json
[
    {
        "id": 1,
        "title": "Someday in Ciloto",
        "caption": "Pray Until Something Happend (PUSH)",
        "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
        "user_id": 1,
        "created_at": "2024-02-16T15:28:12.889112+07:00",
        "updated_at": "2024-02-16T19:40:50.922624+07:00",
        "User": {
            "email": "septyancandra27@gmail.com",
            "username": "septyancandra"
        }
    },
    {
        "id": 3,
        "title": "Someday 2",
        "caption": "Pray Until Something Happend",
        "photo_url": "https://www.instagram.com/septyancandra/",
        "user_id": 1,
        "created_at": "2024-02-16T19:41:23.890961+07:00",
        "updated_at": "2024-02-16T19:41:23.890961+07:00",
        "User": {
            "email": "septyancandra27@gmail.com",
            "username": "septyancandra"
        }
    },
    {
        "id": 4,
        "title": "Someday ",
        "caption": "Pray Until Something Happend",
        "photo_url": "https://www.instagram.com/septyancandra/",
        "user_id": 1,
        "created_at": "2024-02-16T20:40:53.35654+07:00",
        "updated_at": "2024-02-16T20:40:53.35654+07:00",
        "User": {
            "email": "septyancandra27@gmail.com",
            "username": "septyancandra"
        }
    }
]
```

> Failed not login

```json
{
  "error_message": "sign in to proceed",
  "error_status": "Unauthorized"
}
```

## GET Get One Photo

GET /photos/1

> Response Examples

> Get One Photo

```json
{
    "id": 1,
    "created_at": "2024-02-16T15:28:12.889112+07:00",
    "updated_at": "2024-02-16T19:40:50.922624+07:00",
    "title": "Someday in Ciloto",
    "caption": "Pray Until Something Happend (PUSH)",
    "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
    "user_id": 1
}
```

## PUT Update Photo

PUT /photos/update/1

> Body Parameters

```yaml
title: Someday in Ciloto
caption: Pray Until Something Happend (PUSH)
photo_url: https://www.instagram.com/septyancandra/suka-suka

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» title|body|string| yes |none|
|» caption|body|string| yes |none|
|» photo_url|body|string| yes |none|

> Response Examples

> Update Photo

```json
{
    "id": 1,
    "title": "Someday in Ciloto",
    "caption": "Pray Until Something Happend (PUSH)",
    "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
    "user_id": 1,
    "updated_at": "2024-02-16T19:40:50.922624+07:00"
}
```

## DELETE Delete Photo

DELETE /photos/delete/3

> Response Examples

> Delete Photo

```json
{
    "message": "Your photo has been successfully deleted"
}
```

# User

## POST Register MyGram

POST /users/register

> Body Parameters

```yaml
username: septyanpyan
email: septyanpyan9@gmail.com
password: "12345678"
age: "25"

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» username|body|string| yes |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|
|» age|body|string| yes |none|

> Response Examples

> Register MyGram

```json
{
    "id": 3,
    "age": 25,
    "email": "septyanpyan9@gmail.com",
    "username": "septyanpyan"
}
```

> Failed Register MyGram

```json
{
  "error": "Bad Request",
  "message": "Age must be greater than 8"
}
```

## POST Login MyGram

POST /users/login

> Body Parameters

```yaml
email: septyanpyan9@gmail.com
password: "12345678"

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|

> Response Examples

> Login MyGram

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNlcHR5YW5weWFuOUBnbWFpbC5jb20iLCJpZCI6M30.0B23n5BwSyA1LWzsIDBlPb4xuNdVUQG3zE6zhF3uGAw"
}
```

# Comment

## POST Post Comment

POST /comments/add

> Body Parameters

```yaml
photo_id: 1
message: congrats

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» message|body|string| yes |none|

> Response Examples

> Post Comment

```json
{
    "id": 8,
    "message": "congrats",
    "user_id": 1,
    "photo_id": 1,
    "updated_at": "2024-02-16T19:41:58.598511+07:00"
}
```


## GET Get All Comment

GET /comments

> Body Parameters

```yaml
{}

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Get All Comment

```json
[
    {
        "id": 1,
        "message": "wah happy ya",
        "user_id": 1,
        "photo_id": 1,
        "created_at": "2024-02-16T18:26:03.95114+07:00",
        "updated_at": "2024-02-16T18:26:03.95114+07:00",
        "User": {
            "id": 1,
            "email": "septyancandra27@gmail.com",
            "username": "septyancandra"
        },
        "Photo": {
            "id": 1,
            "title": "Someday in Ciloto",
            "caption": "Pray Until Something Happend (PUSH)",
            "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
            "user_id": 1
        }
    },
    {
        "id": 4,
        "message": "wah happy ya",
        "user_id": 1,
        "photo_id": 1,
        "created_at": "2024-02-16T18:38:32.452145+07:00",
        "updated_at": "2024-02-16T18:38:32.452145+07:00",
        "User": {
            "id": 1,
            "email": "septyancandra27@gmail.com",
            "username": "septyancandra"
        },
        "Photo": {
            "id": 1,
            "title": "Someday in Ciloto",
            "caption": "Pray Until Something Happend (PUSH)",
            "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
            "user_id": 1
        }
    },
    {
        "id": 5,
        "message": "wah happy ya",
        "user_id": 1,
        "photo_id": 1,
        "created_at": "2024-02-16T18:40:12.478702+07:00",
        "updated_at": "2024-02-16T18:40:12.478702+07:00",
        "User": {
            "id": 1,
            "email": "septyancandra27@gmail.com",
            "username": "septyancandra"
        },
        "Photo": {
            "id": 1,
            "title": "Someday in Ciloto",
            "caption": "Pray Until Something Happend (PUSH)",
            "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
            "user_id": 1
        }
    }
]
```

## DELETE Delete Comment

DELETE /comments/delete/5

> Response Examples

> Delete Comment

```json
{
    "message": "Your comment has been successfully deleted"
}
```

## PUT Update Comment

PUT /comments/update/5

> Body Parameters

```yaml
message: Update comment dulu

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» message|body|string| yes |none|

> Response Examples

> Update Comment

```json
{
    "id": 5,
    "message": "hahahaa",
    "user_id": 1,
    "photo_id": 1,
    "updated_at": "2024-02-16T19:42:05.394582+07:00"
}
```

## GET Get One Comment

GET /comments/5

> Body Parameters

```yaml
{}

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Get One Comment

```json
{
    "comment": {
        "id": 5,
        "created_at": "2024-02-16T18:40:12.478702+07:00",
        "updated_at": "2024-02-16T18:40:12.478702+07:00",
        "user_id": 1,
        "User": {
            "id": 1,
            "created_at": "2024-02-13T16:18:08.762246+07:00",
            "updated_at": "2024-02-13T16:18:08.762246+07:00",
            "username": "septyancandra",
            "email": "septyancandra27@gmail.com",
            "password": "$2a$08$69tlRRP4pPLQxFg3luVfQOdeAoYhdI0GFInHSl3GIMecaF4Gtc7oS",
            "age": 9,
            "profile_image_url": "",
            "photos": null,
            "comments": null,
            "social_medias": null
        },
        "photo_id": 1,
        "Photo": {
            "id": 1,
            "created_at": "2024-02-16T15:28:12.889112+07:00",
            "updated_at": "2024-02-16T19:40:50.922624+07:00",
            "title": "Someday in Ciloto",
            "caption": "Pray Until Something Happend (PUSH)",
            "photo_url": "https://www.instagram.com/septyancandra/suka-suka",
            "user_id": 1
        },
        "message": "wah happy ya"
    }
}
```

# Social Media

## POST Create Social Media

POST social-media/add

> Body Parameters

```yaml
name: septyan candra
social_media_url: https://www.instagram.com/septyancandra/

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» social_media_url|body|string| yes |none|

> Response Examples

> Create Social Media

```json
{
    "id": 3,
    "name": "septyan candra",
    "social_media_url": "https://www.instagram.com/septyancandra/",
    "user_id": 1,
    "created_at": "2024-02-16T19:39:41.792043+07:00"
}
```

## GET Get All Social Media

GET /social-media/

> Response Examples

> Get All Social Media

```json
[
    {
        "id": 3,
        "name": "septyan candra",
        "social_media_url": "https://www.instagram.com/septyancandra/",
        "user_id": 1,
        "updated_at": "2024-02-16T19:39:41.792043+07:00",
        "created_at": "2024-02-16T19:39:41.792043+07:00",
        "User": {
            "id": 1,
            "username": "septyancandra",
            "profile_image_url": ""
        }
    },
    {
        "id": 4,
        "name": "pyan jr",
        "social_media_url": "https://www.instagram.com/pyanjr/",
        "user_id": 1,
        "updated_at": "2024-02-16T20:55:03.485312+07:00",
        "created_at": "2024-02-16T20:55:03.485312+07:00",
        "User": {
            "id": 1,
            "username": "septyancandra",
            "profile_image_url": ""
        }
    }
]
```

## GET Get One Social Media

GET /social-media/3

> Response Examples

> Get One Social Media

```json
{
    "id": 3,
    "created_at": "2024-02-16T19:39:41.792043+07:00",
    "updated_at": "2024-02-16T19:39:41.792043+07:00",
    "name": "septyan candra",
    "social_media_url": "https://www.instagram.com/septyancandra/",
    "user_id": 1,
    "User": null
}
```

## PUT Update Social Media

PUT /social-media/update/3

> Body Parameters

```yaml
name: Pyan
social_media_url: https://www.instagram.com/septyancandra/

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» social_media_url|body|string| yes |none|

> Response Examples

> Update Social Media

```json
{
    "id": 2,
    "name": "Pyan",
    "social_media_url": "https://www.instagram.com/septyancandra/",
    "user_id": 1,
    "updated_at": "2024-02-16T19:39:19.1383+07:00"
}
```

## DELETE Delete Social Media

DELETE /socialmedias/delete/2

> Response Examples

> Delete Social Media

```json
{
    "message": "Your social media has been successfully deleted"
}
```

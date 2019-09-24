
# GO-Blog



## Indices

* [Author](#author)

  * [Get All Authors](#1-get-all-authors)
  * [Get Author By ID](#2-get-author-by-id)
  * [Delete Author By ID](#3-delete-author-by-id)
  * [Create Author](#4-create-author)

* [Comment](#comment)

  * [Get All Comments](#1-get-all-comments)
  * [Get Comments Of Post](#2-get-comments-of-post)
  * [Get Comments Of Author](#3-get-comments-of-author)
  * [Get Comment By ID](#4-get-comment-by-id)
  * [Create Comment](#5-create-comment)
  * [Delete Comment](#6-delete-comment)

* [Post](#post)

  * [Get All Posts](#1-get-all-posts)
  * [Get Posts Of Author](#2-get-posts-of-author)
  * [Get Post By ID](#3-get-post-by-id)
  * [Like Post](#4-like-post)
  * [Create Post](#5-create-post)
  * [Delete Post](#6-delete-post)

* [Tag](#tag)

  * [Create Tag](#1-create-tag)
  * [Get Tag By ID](#2-get-tag-by-id)
  * [Get Tags Of Post](#3-get-tags-of-post)
  * [Get All Tags](#4-get-all-tags)
  * [Delete Tag By ID](#5-delete-tag-by-id)


--------


## Author



### 1. Get All Authors



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/authors
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/authors |  |  |



### 2. Get Author By ID



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/authors/2
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/authors/2 |  |  |



### 3. Delete Author By ID



***Endpoint:***

```bash
Method: DELETE
Type: URLENCODED
URL: http://localhost:8080/authors/1
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/authors/1 |  |  |



### 4. Create Author



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/authors
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
	"firstname": "Lara",
	"lastname": "Gerhard",
	"email": "lara.gerhard@email.com",
	"avatarUrl": null
}
```



## Comment



### 1. Get All Comments



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/comments
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/comments |  |  |



### 2. Get Comments Of Post



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/comments/post/1
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/comments/post/1 |  |  |



### 3. Get Comments Of Author



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/comments/author/2
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/comments/author/1 |  |  |



### 4. Get Comment By ID



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/comments/2
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/comments/1 |  |  |



### 5. Create Comment



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/comments
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
	"content": "Amazing Post!",
	"author_id": 2,
	"post_id": 1
}
```



### 6. Delete Comment



***Endpoint:***

```bash
Method: DELETE
Type: 
URL: http://localhost:8080/comments/2
```



## Post



### 1. Get All Posts



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/posts
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/posts |  |  |



### 2. Get Posts Of Author



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/posts/author/1
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/posts/author/1 |  |  |



### 3. Get Post By ID



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/posts/1
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/posts/1 |  |  |



### 4. Like Post



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/posts/1/like
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/posts/1/like |  |  |



### 5. Create Post



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/posts
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
	"title": "How to cook Curry",
	"content": "Mmmmhhh",
	"likes": 0,
	"author_id": 3
}
```



### 6. Delete Post



***Endpoint:***

```bash
Method: DELETE
Type: URLENCODED
URL: http://localhost:8080/posts/2
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/posts/2 |  |  |



## Tag



### 1. Create Tag



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: http://localhost:8080/tags/post/1
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
	content: Tech
}
```



### 2. Get Tag By ID



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/tags/7
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/tags/7 |  |  |



### 3. Get Tags Of Post



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/tags/post/1
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/tags/post/1 |  |  |



### 4. Get All Tags



***Endpoint:***

```bash
Method: GET
Type: URLENCODED
URL: http://localhost:8080/tags
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/tags |  |  |



### 5. Delete Tag By ID



***Endpoint:***

```bash
Method: DELETE
Type: URLENCODED
URL: http://localhost:8080/tags/8
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| http://localhost:8080/tags/8 |  |  |



---
[Back to top](#go-blog)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2019-09-24 16:17:45 by [docgen](https://github.com/thedevsaddam/docgen)

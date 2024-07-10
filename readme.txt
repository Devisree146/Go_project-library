*** LRU Eviction policy by using In-memory and Redis

*** Outline:

This project implements an LRU (Least Recently Used) policy in in-memory and Redis, including TTL (Time-to-Live) functionality.


The cache operations available are set, get, getall, delete, deleteall, accessed by POST, GET and DELETE HTTP methods.

It includes testing, benchmarking and API handlers implemented by using Gin Framework.

Postman is used for testing with API.

Installations:

** Go
** Redis
** Postman  

Go installation:
** Install the Go latest version

Redis installation:
** search redis download.
** Go to the Redis downloaded directory: cd C:\Users\devis\Downloads\Redis-x64-3.2.100--This is my directory
** Start the server:  .\redis-server.exe  

Run this application.
** we can use go run and the application name.
** If you want to run in_memory then use go run in_memory
** If you want to run redis_cache then use go run redis
** If you want to run multicache then use go run multicache

URL:
** http://localhost:8080/cache

*** Endpoints

*** Set Key-Value Pair

*   **URL:** `/cache`
*   **Method:** `POST`
*   **Request Body:** `{ key": "your-key", "value": "your-value", "ttl": 60 }`  
>   TTL in seconds
*   **Response:** `{ "message": "Key-Value pair set successfully" }`

*** Get Value by Key


*  **URL:** `cache/get?key="your-key"
*  **Method:** `GET`
*  **Response:** `{ "key": "your-key", "value": "your-value" }`

*** Get All Keys
*   **URL:** `/cache/all`
*   **Method:** `GET`
*   **Response:** `{ "Key-Value Pairs": { key1: value1, key2:value2 } }`

*** Delete Key

*   **URL:** `cache/delete?key="your-key"`
*   **Method:** `DELETE`
*   **Response:** `Key deleted successfully` 

*** Delete All Keys
*   **URL:** `/cache/all`
*   **Method:** `DELETE`
*   **Response:** : `All keys deleted successfully` 

**These are the same operations performed by in_memory,redis and multicache

** Benchmarking
To benchmark the performance of the LRU cache:
1.  Run the benchmark tests:
    Go to the exact directory where benchmark test file is located
    `go test -bench=. -benchtime=5s -v'


** Configuration

Configuration for TTL and max cache size is done via POST.  
The default configuration is:
*   `REDIS_ADDR`: Address of the Redis server (default: `localhost:6379`).
*   `REDIS_PASSWORD`: Password for the Redis server (default: `""`).
*   `REDIS_DB`: Redis database number (default: `0`).
*   `SIZE`: Default size is `3`.
*   `TTL`: Default TTL is `60` seconds.

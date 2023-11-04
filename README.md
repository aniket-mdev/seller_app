# seller_app


```
[#f03c15] SERVICE ENDPOINTS
    - "/adspace":
            To create a new Ad Space
            Method : POST
            Request Body :
                {
                    "description":"Banner Ad",
                    "base_price":123,
                    "status":"open",
                    "end_time":"2023-11-04 16:04"
                }

    - "/adspace":
            To fetch all Ad Space which has a status OPEN
            Method : GET

    - "/adspace/{id}":
            To Fetch a Specific Ad Space By Space ID
            Method : GET
            Path Param : url/1

    - "/adspace/{id}":
            To update a Ad Space by Space Id
            Method : PUT
            Path Param : url/1
            Request Body:
                {
                    "description":"Banner Ad",
                    "base_price":123,
                    "status":"open",
                    "end_time":"2023-11-04 16:04"
                }

    - "/adspace/{id}:
            To delete a Ad Space by space id
            Method : DELETE
            Path Param : url/1

    - "/check_adspace_endtime":
            This endpoint will call from scheduler after every 5 min or we can define a any time for that.This will ensure that space status is open and endtime is expired for that ad space , if it find a any space in the query then it will update that space status as a close.

            Method : GET                                

```                

## Docker Installation
```
1. system should have a docker install
2. build a docker image by following a command
    * docker build -t demmand_service . *
3. to run a image
    * docker run -p 8080:8080 demmand_service*    

```
Data flow:
    1) Web request (see sample request) >
    2) "Ingestion Agent" (php) >
    3) "Delivery Queue" (redis)
    4) "Delivery Agent" (go) >
    5) Web response (see sample response)

App Operation - Ingestion Agent (php):
    1) Accept incoming http request
    2) Push a "postback" object to Redis for each "data" object contained in accepted request.

App Operation - Delivery Agent (go):
    1) Continuously pull "postback" objects from Redis
    2) Deliver each postback object to http endpoint:
        Endpoint method: request.endpoint.method.
        Endpoint url: request.endpoint.url, with {xxx} replaced with values from each request.endpoint.data.xxx element.
    3) Log delivery time, response code, response time, and response body.

Sample Request:
    (POST) http://{server_ip}/ingest.php
    (RAW POST DATA) 
    {  
      "endpoint":{  
        "method":"GET",
        "url":"http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}"
      },
      "data":[  
        {  
          "mascot":"Gopher",
          "location":"https://blog.golang.org/gopher/gopher.png"
        }
      ]
    }

Sample Response (Postback):
    GET http://sample_domain_endpoint.com/data?title=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo=
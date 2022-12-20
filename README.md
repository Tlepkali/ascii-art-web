# ASCII-ART-WEB-STYLIZE

ascii-art-web is a webpage, that allows create ascii representation of input string

## Usage/Examples

Clone the repository and run docker container

  docker build .
  docker images
  ____________________________________________________________________________________________
  |                  |                                                                        |
  | You will see     |  REPOSITORY             TAG       IMAGE ID       CREATED         SIZE  |
  | table like this  | <none>                 <none>    cb45c286cf70   9 seconds ago   1GB    |
  | and you need     | asciiartwebdockerize   latest    077f36111734   5 hours ago     1GB    |
  | to copy IMAGE ID | golang                 latest    f7d4f5578ce3   9 days ago      992MB  |
  |__________________|________________________________________________________________________|

  docker run -p 8080:8080 {IMAGE ID} ## you can use any port before ':', paste the copied IMAGE ID instead {IMAGE ID}

go to http://127.0.0.1:8080/

1.Write your input string

2.Choose one of the ascii-template

3.Press SUBMIT button
### Implementation details

main.go Creates multiplexer and starts server on 8080 port. When the request reaches the server, a multiplexer will inspect the URL being requested and redirect the request to the correct handler fucntion

handlers.go

Handlers fucntions process requests. When the processing is complete, the handler passes the data to the template engine, which will use templates to generate HTML to be returned to the client.

#### Authors

@Tlep - https://github.com/Tlepkali
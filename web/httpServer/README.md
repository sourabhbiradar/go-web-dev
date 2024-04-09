# HTTP Server

ListenAndServe() function from net/http is used to create HTTP server.

http.Handler handles the HTTP requests from Server. 

Read from request before writing response.

http.ResponseWriter  :
Header() method returns Header map that is set to response.
Write()  method write data to connection as a part of HTTP reply.
WriteHeader() method is used to send provided Status code.

http.Request : 
HTTP request recieved from server (OR sent by a clint)
HTTP method GET,POST,PUT etc..
URL *url.URL : URI being requested or URL to access


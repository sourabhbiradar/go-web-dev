# io.Reader & io.Writer

a)io.Reader : used to READ BYTES from various input sources.

io.Reader has core method Read(name) which reads bytes from source and stores in 'name' 
& returns number of bytes read & error (if any).

- bufio.Reader         : read for improved performance
- ioutil.ReadAll(file) : reads from 'file'
- http.Request.Body    : allows you to supply request body via reader.

b)io.Writer : used to WRITE BYTES.

io.Writer has core method Write(data) which writes bytes to destination from 'data'.
& returns number of bytes written & error (if any).

- os.Stdout           : writes on console.
- http.ResponseWriter : writes response body of HTTP request.
- ioutil.WriteAll     : writes to 'file'




# serveMux

Upgraded in Go 1.22

Capable of replacing third party routers. 

serveMux is router multipexer match different rquest handlers to different URLs served by server.

Ensure right METHOD is used.[GET,POST etc]

Wildcards :
- `{path_param}` matches a single path segment till next forward slash. `PathValue()` to get this value.

- `...` matches rest of URL. Allowed only at the end of url.

- `{$}` matches path "/"

NOTE : be specific with patterns to avoid Path conflict.

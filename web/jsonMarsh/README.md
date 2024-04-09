# marshalling & unmarshalling

a.k.a. serialization & deserialization

marshalling : converting structs,maps & slices to JSON format. 

Unmarshalling : reverse marshalling

Provided by package 'encoding/json'

JSON tags (within BACKTICKS ` `) are used to map struct fields to JSON data.

JSON data is in '[]byte'.

# encoding & decoding

json.NewEncoder().Encode() is used to write the JSON data directly to an io.Writer like a file, network connection, or HTTP response.

json.NewDecoder().Decode() is used for decoding JSON data from an input source (like an io.Reader such as a file, network connection, or HTTP response) into Go data structures.


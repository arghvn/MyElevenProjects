Marshalling and Unmarshalling in Golang :

In Golang, struct data is converted into JSON and JSON data to string with Marshal() and Unmarshal() method. The methods returned data in byte format and we need to change returned data into JSON or String.
What is Marshalling in Golang
Converting Go objects into JSON is called marshalling in Golang. As the JSON is a language-independent data format, the Golang has inbuilt encoding/json package to perform json related operations.

The encoding/json package has json.Marshal() method. This method takes object as parameter and returned byte code.
What is Unmarshalling in Golang ?
Unmarshalling just opposite of Marshalling. The Golang encoding/json package has json.Unmarshal() method that used to convert json(Byte data) into Struct. As we have covered Marshalling of struct into JSON, now we will take that JSON string and Unmarshal that JSON into a Struct.
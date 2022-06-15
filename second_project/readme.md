 next we will build s api using GO 
 we will have something called as a movie server which will give us the list of all movies and we can create update and delete ...
 its name is crud and crud stands for this server.
 we dont have database here and we use strucs and slices in localhost  8000 

 we will have five different routes : get all , get by id , create , update and delete
 Each path will have a function that uses methods to display a movie
 
 functions for each routes :

get all , getmovies
get by id , getmovie
create , createmovie
update , updatemovie
delete , deletemovie

there are two types of endpoints slash /movies and /movies/id
we have to use depending on what we are doing in the sense ,if we are creating
 then obviously /m if we are updating then obviously id is required so it has to be /m/i right then. 

 and there are methods and there are getmethods...get fo get all and get by id function and we will use post (another method)
 for create function and we will use put for update  function and we will use delete for delete function.

 and we wil using postman to test all these endpoints.

 what is the rest Api ?
 REST is an "architectural style" or "design template" for APIs.
  Various applications and web pages are used to obtain information about various sources.
  Where does this data come from? This data is received from servers and applications are connected to web servers to receive and send data.
A REST API is a way to connect two computer systems, such as a web browser and servers, via HTTP.
API‌ is the title for the "Application Programming Interface".
  An API is an intermediary that allows communication between two applications.

postman :
Postman makes it very easy for us to test, execute, and test written WebAPI code and methods.
In the programming environment, when we implement a Web API to test it in the same environment, we have to write client-side code that we can check its output, that is, we have to simulate the original code that we are going to write, but always do This is not cost-effective for us in terms of time, and we like to see the output of our program very quickly, and when we change our code, we can easily see the output and the result of our changes. Also, sometimes testing the output by the main program is a difficult and challenging task. So having an interface software for this is very efficient and using it will help us.

in postman/collection we create a new folder by name rough and in this folder we create a folder by name GO-Movies
in folder GO-Movies we create a request by name GetAll and another request named Get By ID another request named create another request named update and another request named delete.
 in section url we enter http://localhost:8000/movies and send it to server
 showing our code

 lets get a movie by enter http://localhost:8000/movies/1
 the firs movie is here

 to create a movie enter http://localhost:8000/movies
 in body section and part json 
 showing the code and to create a new movie furthermore two movies that we have in past , we can change the Id and first name Aknil and lastname Sharrea
 it has added

 to update a movie enter http://localhost:8000/movies/ID 
 in body section and part json 
 we can change each parameters we would like
 for example change Aknil to Arnold and  Sharrea to S and change Isbn
 the movie has been updated

 to delete a movie enter http://localhost:8000/movies/ID
 the movie was deleted

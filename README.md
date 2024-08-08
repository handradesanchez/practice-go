# practice-go

In order to better understand what's the purpose of interfaces and how to use them, we should do some practice that involves them.

Task steps:
- Create a very simple go application that just has a main() func that calls another func called GetUsers() and hardcode some users on that func to be returned. Main func should just print out those users (you should create a user struct and make GetUsers() return a slice of those. Just give them some username and age to start with).

- Create an echo http server and has only 1  "hello world" endpoint. It doesn't need to have any fancy file structure, you can just leave everything dumped into main. Endpoint should just return a hardcoded text saying "hello world!" along with a 200 status code. Get it working and tested (manually with curl, no unit tests involved).

- Create a 2nd endpoint called /users and make it return 200 along with the GetUsers() result. Give it a test run.

- Create an empty struct, called: "RedisDB". Now and implement a METHOD called GetUsers() on that new struct. Just make it return a slice of hardcoded users with different username and ages. Replace the implementation of your /users endpoint to now use that METHOD and delete your FUNCTION. Make your HTTP endpoint return redis.GetUsers() along with a 200 status code. Give it a test run and make sure it works.

- Add a new DELETE /users endpoint and make the RedisDB struct implement a DeleteUser() method that takes a username as an argument (string type). That argument value must come from the http call either with c.Bind() or c.QueryParams() and be passed along to the reids.DeleteUser(string) method. Again, the DeleteUser() method implementation can be hardcoded, you can just return an empty user struct, but you need to make sure that you run the echo server, hit the endpoint, and it works.

- Let's say a new DB type comes in place and now we also want to retrieve users from a MySQLDB. You need to add that as an empty struct to know it exists now. Don't worry about adding methods to it, just the empty struct will suffice for now.

- Now, we got a new requirement for our application that states that we don't want the Get /users endpoint to retrieve users just from redis anymore. Let's pretend that we have a fully working implementation of the MySQLDB by now and we want to return a combination of the redis + mysql users on that endpoint. We need to create a function called GetAllUsers() that takes a slice of dbs as an argument, iterates through all of them and retrieves their users, and returns a slice containing all of the users retrieved from every source… You can start defining the GetAllUsers() func, but, what would be the type of that slice argument you need to iterate on it? that's where interfaces come in! Go ahead and create a DB interface. Take a look at your current api implementation and think of what methods you need to inject on that interface so that you can interpret ANY kind of struct as DB. What is the least amount of methods that a struct should have in order to be called a DB and suffice your api needs? Go ahead and add them to that new interface you just created. Then, use that new interface type as a slice argument on your GetAllUsers([]<interfaceName>) func and fully implement it. Again, you should iterate through each one of them, retrieve the users, and return a concatenation of all of them regardless of the source. Get to a point where you can hit your /users endpoint again and be able to see both redis + mysql users on the response.

- New requirement: the DELETE /user endpoint should also look for users on both DB and delete the user if it is found on any of them. You don't have to fully implement the DeleteUser method on both structs, just make sure they both have a method with the same signature and that they return some hardcoded stuff so that again you can iterate through them and execute that method. Get it working and tested by curl.

- Last new requirement: another DB comes in! add a new MongoDB struct and make sure that you pass it on to both endpoints so that we now look into 3 different DBs instead of 2.


Definition of Done:
Have a fully working example that accomplishes all of the tasks defined above.
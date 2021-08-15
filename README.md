**Go Auth**

__Objective:__ 
_To make a secure web app using go._

__TODO:__
- Session based login
- JWT implementation
- Setting server side cookies
- Implementing CSRF

__Secondary objective:__
- Clean code
- Writing tests
- Graceful shutdown

__Contains:__
- API:
    - Home
    - Login(Get: To load html and inject csrf token)
    - Login(Post: To make requst and consume api with csrf as header)
    - UserHome(Get: To load userhome after successful login)
- Data:
    - In memory (for simplicity) 
- Html Templates:
    - home
    - login
    - userhome

__Note:__
    - Dummy Data for successful login
        - Name:     "Dummy User",
	    - Email:    "dummy@email.com",
	    - Password: "dummyPassword",
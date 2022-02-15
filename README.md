# Go Auth

## Objective: 
_To make a secure web app using go._

## TODO:
* Session based login
* JWT implementation
* Setting server side cookies
* Implementing CSRF
* Rate Limiting to Avoid DOS

## Secondary objective:
* Clean code
* Writing tests
* Graceful shutdown

## Contains:
* API:
    * Home
    * Login(Get: To load html and inject csrf token)
    * Login(Post: To make requst and consume api with csrf as header)
    * UserHome(Get: To load userhome after successful login)
* Data:
    * In memory (for simplicity) 
* Html Templates:
    * home
    * login
    * userhome

## Note:
    * Dummy Data for successful login
_       * Name:     "Dummy User",
_	    * Email:    "dummy@email.com",
_	    * Password: "dummyPassword",

Read Me file for **Go Swagger**

when you first time use go swagger then first of all do this:-

1.create a temporary  directory:-
- set dir=%TEMP%\go-swagger
- mkdir %dir%

2.Clone the go-swagger repository into the temporary directory
 - git clone htp://github.com/go-swagger/go-swagger%dir%

3.change to the cloned repository
- cd%dir%

4.Install go-swagger using the "go.install" command
- go install ./cmd/swagger

5.Run this command for generate yml file
- swagger init spec

and then run this command for swagger validation 
- swagger validation ./swagger.yml

and then write yml file according to your requirement in generate swagger.yml file

6.For code generation write some command
- (a).Code generate for server side you Run this command:-
    swagger generate server -f <yml file name> 
- (b).Code generate for client side you Run this command:-
    swagger generate client -f <yml file name>

7.For see a lot of command Run this command:-
- swagger generate server --help

8.For generate JSON file Run this command:-
- swagger generate spec -o ./swagger.json

9.Then mapping auto generated handler with your logic
for Example:-
If you want some data insert,read,update and delete in database then first of all create
a separate directory name is dbhelper and write your query inside any 
function this function created by you,and after then call this query function in configure_movie.go file
in configureAPI function.

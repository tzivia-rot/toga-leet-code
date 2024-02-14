# toga-leet-code
## Instructions:
## What needs to be done to activate the system:
1. open cmd
2. run  ``` git clone git@github.com:tzivia-rot/toga-leet-code.git ```
3. #### to run server:
     1. run ```cd server```
     2. run ```docker-compose up```
4. #### to run client - cmd:
   1. run ```cd command-line```
   2. rum ```go run main.go```
> [!WARNING]
> the server must be runing befor you run this
### How to run the actions from the command line:
- First, you get all the existing exercises, you can choose which exercise you want to do
  or if you want to add a new exercise, you can choose another action and go to add an exercise.
  - Add an exercise: When you add an exercise you need to enter these parameters:
     Name, description and 3 examples. for this exercise each example includes output and how many parameters the function should accept.

     Apart from that, you need to add basisOperationNodeJS and basisOperationGO      
     These parameters are used to test the exercise for node.js/go

     This should include:
    
     1. get the parms from env
     2. sending from action function
     3. defintion from function that get right params
    
     This need to be all what that the user need to run code

    >for example new exercise:
    >
         name     
         node return input
         description     the function need return the input that he get
    

> [!NOTE]
> that the base name of the function must be action and contain contiguous blank brackets for applying the function value in its place

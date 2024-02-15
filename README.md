# Toga-leet-code
## Instructions:
## What needs to be done to activate the system:
1. open cmd
2. run  ``` git clone git@github.com:tzivia-rot/toga-leet-code.git ```
3. #### to run server and the mongoDB:
     1. run ```docker-compose up```
> [!WARNING]
>  When i try run the server on docker compose i get problom to connect to DB,
> I tried to search about this for long time without success,
> if is this not work you can run the DB in docker compose and the server local - ןt works
4. #### to run client - cmd:
   1. run ```cd command-line```
   2. rum ```go run main.go```
> [!WARNING]
> The server must be runing befor you run this
### How to run the actions from the command line:
- First, you get all the existing exercises, you can choose which exercise you want to do
  or if you want to add a new exercise, you can choose another action and go to add an exercise.
 - Add an exercise: When you add an exercise you need to enter these parameters:
     Name, description and 3 examples. for this exercise each example includes output and how many parameters the function should accept.

     Apart from that, you need to add basisOperationNodeJS and basisOperationGO      
     These parameters are used to test the exercise for node.js/go

     This should include:
    
     1. get the parms from env, MY_VARIABLE_INPUT_1 / MY_VARIABLE_INPUT_2 ... and MY_VARIABLE_OUTPUT
     2. convert env for the the type that the function need to get
     3. sending from action function
     4. defintion from function that get right params
     5. convert the output of action function to string because the ouput get string
     6. Checking whether the value returned from an action function is equal to the output value received in an environment variable, if is ```exit(0)``` else ```exit (1)```
    
     This need to be all what that the user need to run code

    >**For example new exercise:**
    >
    name
    
         return input
    
    description:
    
         the function need return the input that he get
    
    example:
    
         [input:["1"] , output:"1"]
         [input:["100"] , output:"100"]
         [input:["hellow"] , output:"hellow"]
    
   basisOperationNodeJS:
    
         function action(input1) {}
         function main(){
             var input1 = process.env.MY_VARIABLE_INPUT_1;
             var input2 = process.env.MY_VARIABLE_OUTPUT; 
             const output=action(input1)
             if(output===input2){
                 process.exit(0)
             }
             else
             {
                 process.exit(1)
             }
         }
          main();


   basisOperationGO:
    
          package main
          import (
              "fmt"
              "os"
          )
          
          func action(input1 string) string {}
          
          func main() {
              input1 := os.Getenv("MY_VARIABLE_INPUT_1")
              input2 := os.Getenv("MY_VARIABLE_OUTPUT")
          
              output := action(input1)
              if output == input2 {
                  fmt.Println("Output matches input2.")
                  os.Exit(0)
              } else {
                  fmt.Println("Output does not match input2.")
                  os.Exit(1)
              }
          }

          
  > [!NOTE]
  > That the base name of the function must be action and contain contiguous blank brackets for applying the function value in its place
   
 - delete, update
       you no need sent params
 - check exercise
    
  You need choose hew languge you went to write the function

  and write **anly** content the function

>**For example for the previous example**

         return input1;

You get if this work or not
- finally, you can choose if you went to continue or not

  Good luck :)
  Thank you for the challenge!
    

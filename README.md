# Logrus to stack driver logger package.




## Introduction
You can use it inside other server, clone it under gitlab.com/dekape/dekalogs

this contains several, level of logging 

 - Info()
 - Debug()
 - Warn()
 - Error()
 
Also default logging middleware contains here! The default middleware can be recover your panic error into error 500 internal server error and log it in command line interface or inside stack driver.
This package did not recover panic outside main goroutine.
feel free to add and change

 <br>

#### For newest go-boilerplate
follow these step to and setup standard boilerplate :
1. Setup the env, like these :
 ```env
  APP_NAME="Your_App_Name"
  APP_PORT=Your_Port
  ....
  LOG_HOOKS=none
  LOG_TYPES=JSON
  PROJECT_ID=dekape-dev-server
 ```
>>>
  You can change following items:
  - LOG_HOOKS with none or STACK
  - LOG_TYPES with none or JSON
  - PROJECT_ID = _your-gcp-project-id_
  
  > **Note**:You only need fill `PROJECT_ID` if `LOG_HOOKS` is equal to `STACK`.
  {: .note}
>>>


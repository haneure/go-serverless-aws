# Golang Serverless API
Golang Serverless stack using DynamoDB and deployed to AWS Lambda

## How to use
### Build go code
1. Make sure you are inside the correct directory, where the main.go is.
    ```
    cd cmd
    ```
2. Build the project code, i am outputting it into the build directory, you can put it anywhere you want.
    ```
    go build -o ../build/main.exe
    ```
3. Make the zip file of the main.exe. I use the ```zip``` package in linux to make the .zip file.
    you can do this code if you are on the root directory
    ```
    go-serverless-aws$ zip ./build/main.zip ./build/main.exe
    ```
4. Create a Lambda Function
5. Create DynamoDB 
# Introduction 
This is a sample application that demonstrates how to use the Temporal Go SDK to perform database operations.The main agenda of this POC is to identify the ideal performance we can achieve with the Temporal Go SDK
by tweaking the parameters in the config file.
About the folders:
1. config: contains the config.yaml file.
2. activity: contains the activity code.
3. workflow: contains the workflow code.
4. starter: contains the workflow initialization code.
5. worker: contains the worker code.
6. start_workers.sh: script to start the workers.

# Getting Started
 1.Clone this repository
 2.Update the config.yaml file with your database and temporal configuration.
 3. Run the application by executing the following commands:
    a. navigate to the root directory of the project.
    b. go build -o worker/main.go // will generate a worker executable.
    c. go run starter/main.go //workflow initialization code, this will start the specified number of workflows.
    d. ./start_workers.sh worker_count //worker code, this will start the specified number of workers.

  

# Build and Test
TODO: Describe and show how to build your code and run the tests. 

# Contribute
TODO: Explain how other users and developers can contribute to make your code better. 

If you want to learn more about creating good readme files then refer the following [guidelines](https://docs.microsoft.com/en-us/azure/devops/repos/git/create-a-readme?view=azure-devops). You can also seek inspiration from the below readme files:
- [ASP.NET Core](https://github.com/aspnet/Home)
- [Visual Studio Code](https://github.com/Microsoft/vscode)
- [Chakra Core](https://github.com/Microsoft/ChakraCore)



    

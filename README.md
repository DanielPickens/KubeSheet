# Kubewrap

kubewrap is an kubernetes command line utility with a focus to explore the kubernetes REST API's leveraging the go libraries available along with golang and cobra package.

If you still haven’t installed the cobra library, you can install it using the below command.
```
 go get -u github.com/spf13/cobra/cobra
 ```
##Next, create your CLI skeleton simply by executing below command


It will initialize the “kubewrap” project with cobra the library. You can observe that it created a few files in the project.

##Folder structure after cobra
Lets quickly understand the files created and its use here
root.go → is the file created for the root command. If you know that for a CLI utility, there would be a root command followed by some set of sub-commands and flags. Therefore, root.go will take care of all logic required for the main/root command.
main.go → as the name says, its the main entry point for your CLI utility. Usually, there would not be any code inside this file but just a call to a function present in the root.go
So to recall the flow main.go -> root.go
At this point, you could test it by executing below commands inside kubewrap folder. Note Minikube must be install as it checks for default server connect to in process.
```
 go install
 
 kubewrap
```


kubewrap is a root command and dig is a sub command to that
In the future, you could have a set of subcommands such as below
```
 kubewrap dig1
 ```
will list out all the pods
```
 kubewrap dig2
 ```
will list out all the deployments…vice versa.
So, to extend our existing skeleton to have subcommand, just execute the below

For you, it should be "kubewrap” folder. Again ignore here testkube
Now your flow would be → main.go -> root.go → dig.go
Now you could also see the dig.go file created in your existing file structure.
That's it. Done. Jump to code now…

Step 3:

a. Connect to your cluster using “kubeconfig” REST API call → root.go
The changes required in “step a” needs to be implemented in the root.go wherein you just need to provide the path of your own config file 

Refer to complete root.go code on a git repository and do change the config file path above
In root.go you just need to have above function which takes care of connecting to the minkube cluster using a REST call. Now since this is a snip of code, you still need to import packages. Refer to the complete code of root.go on git.
b. Once connected, you would simply make a call to rest API to
```
“list pods available” → dig.go
```
Simple, go to dig.go and follow the cha changes required to make a call to the REST API to list all the pods available in the cluster.

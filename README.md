# Kubewrap
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-overtime.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-responsibility.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-reason.svg)](https://forthebadge.com)

kubewrap is an kubernetes cli wrapper with a focus to explore the kubernetes REST API's leveraging the go libraries available along with golang and cobra package.

# Install  the cobra library:
```
 go get -u github.com/spf13/cobra/cobra
 ```
# Create the CLI skeleton 

```
cobra init --package-name kubewrap
```

It will initialize the “kubewrap” project with cobra as the library. 


# Set up minikube

 Execute below commands inside kubewrap folder. Minikube must be installed as it'll need to check for default server to connect to.
```
 go install
```

```

In powershell as admin:
New-Item -Path 'c:\' -Name 'minikube' -ItemType Directory -Force
Invoke-WebRequest -OutFile 'c:\minikube\minikube.exe' -Uri 'https://github.com/kubernetes/minikube/releases/latest/download/minikube-windows-amd64.exe' -UseBasicParsing


Next add your binary to the path:

$oldPath = [Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::Machine)
if ($oldPath.Split(';') -inotcontains 'C:\minikube'){ `
  [Environment]::SetEnvironmentVariable('Path', $('{0};C:\minikube' -f $oldPath), [EnvironmentVariableTarget]::Machine) `
}
```

# Docker

```
docker build -o kubewrap main.go 
```

```
docker pull kubewrap:latest
```
```
docker run kubewrap
```

# Install package dependencies
```
go get https://github.com/DanielPickens/kubewrap.git
```

# Cloning
```
 git clone https://github.com/DanielPickens/kubewrap.git 
 ```
 
# Root and Sub Command
kubewrap is a root command and dig is a sub command to that.
For this case, we could have a set of subcommands such as below:
```
 kubewrap dig1
 ```
will list out all the pods
```
 kubewrap dig2
 ```
This will list out all the deployments…vice versa.
So, to extend our existing skeleton to have subcommand, just execute the below:
```
cobra add dig
```

# Connect to cluster and config root.go:

a. Connect to a cluster using “kubeconfig” REST API call
'''
 → root.go
'''
The changes required in “step a” needs to be implemented in the root.go wherein you just need to provide the path of your own config file.

Locate to the config in root.go and change the config file path above when required.
In root.go you'll just need to have above function which takes care of connecting to the minkube cluster using a REST call.  
b. Once connected, you would simply need to make a call to rest API to:
```
“list pods available” → dig.go
```
Go to dig.go and follow the changes required to make a call to the REST API to list all the pods available in the cluster.

# Testing
```
go install
```
```
kubewrap dig
```

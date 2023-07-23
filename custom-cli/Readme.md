# Repo for Custom Golang CLI with cobra & viper


### mctl - small utility for common math operation

* Basic understanding of the utility program
    * Command represents the action:
    * Args are things
    * Flags are modifier for those action

```
APPNAME Command Args --Flags
#For example: 
mctl operation add --float
```

* Common steps
```
mkdir mctl
cd mctl
go init mod mctl
cobra-cli init 


# add a operation 
cobra-cli add operation

cobra-cli add add -p 'operationCmd'

go fmt cmd/...
go build .
```


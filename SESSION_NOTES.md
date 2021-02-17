# 2021-02 Golang training

# Performed by

Miki Tebeka

using go since 2010

prof dev since 25 years

50 50 go python

mostly consulting / writing code

organizer of several conferences

[https://www.353solutions.com/c/rh21/](https://www.353solutions.com/c/rh21/)


# 2021-02-17

exportable variables and functions start with a capital letter

these are accessible outside the package

[https://github.com/go-yaml/yaml](https://github.com/go-yaml/yaml)

consider using an alternative domain for your golang exports besides github.com

regarding go get

go get gopkg.in/yaml.v2

go doc .

write your example code into file

example_test.go

(break)

[hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))

“look at kubernetes, see what they are doing regarding structure” is usually a bad idea

[https://github.com/etcd-io/bbolt](https://github.com/etcd-io/bbolt)

golang internal packages

just a directory where you put your internal code, will not be accessible from the outside

golang http package since pre 1.0

where Context was introduced in 1.7

aschuett: as an aside for those that like to model go code after kubernetes project theres a talk about how this code is based of of Java and how it's all wrong. For those interested[ https://archive.fosdem.org/2019/schedule/event/kubernetesclusterfuck/](https://archive.fosdem.org/2019/schedule/event/kubernetesclusterfuck/)


## http

[wrk tool](https://github.com/wg/wrk)

[Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)

postman tool to test apis

[go-chi](go-chi)

github.com/gorilla

[https://www.postman.com/product/api-client/](https://www.postman.com/product/api-client/)

go tool pprof


## conditional building

// +build prof

go run -tag prof

go test also supports tags, so you can conditionally include code/tools for testing


## Dependencies

we want to use chi

dependencies are dangerous, rather use golang packages

malicious code

we fix the versions to protect against malicious code

use go mod vendor

problem about putting vendor folder into github

i.e. kubevirt has copy of kubernetes in vendor directory which is 65MB

[golangci-lint](https://github.com/golangci/golangci-lint)

 

## Process by commit and manual operation
1. [first commit](https://github.com/tarotaro0/exp-mm-repo/commit/3960c08ba51b7816142cfcd8384906655405dd72)
- add a new go program
2. [Init root go.mod](https://github.com/tarotaro0/exp-mm-repo/commit/7375c264d5f4b3214dbaa32fb34c3d8facd87616)
- ```$ go mod init``` in root directory
3. [Add foo service](https://github.com/tarotaro0/exp-mm-repo/commit/7aef6d7224b6a6e1c2be3c27451f4be2860f9e69)
- add a new service
- ```$ go mod init github.com/tarotaro0/exp-mm-repo/services/foo``` in `/foo` directory
- ```$ go mod edit -require github.com/tarotaro0/exp-mm-repo/services/foo@v1.0.0` in `root```
- ```$ go mod edit -replace github.com/tarotaro0/exp-mm-repo/services/foo@v1.0.0=./services/foo```
  - use local foo-service because there is no tag for `foo` in public yet

4. [Tag for foo-servicev1.0.0](https://github.com/tarotaro0/exp-mm-repo/releases/tag/services%2Ffoo%2Fv1.0.0)
- ```$ git tag services/foo/v1.0.0```
- ```$ git push origin master services/foo/v1.0.0```
  - or operate in github release for tagging
5. [Use foo service v1.0.0](https://github.com/tarotaro0/exp-mm-repo/commit/4302e96d496dcb90dc224bc1fb35fc08322eea41)
- use foo-service
```
$ go run main.go
Hello, Foo v1.0.0
```
6. [Drop replace](https://github.com/tarotaro0/exp-mm-repo/commit/9edef0f2a346f79998b714a6619556487f3707e6)
- ```$ go mod edit -dropreplace github.com/tarotaro0/exp-mm-repo/services/foo@v1.0.0```
  - because foo-service v1.0.0 is tagged in github and we can use it
7. [Update foo-service](https://github.com/tarotaro0/exp-mm-repo/commit/46bdc01c008461ed69349fda663cfb3f7009092b)
- update foo-service
8. [Tag for foo-service v1.1.0](https://github.com/tarotaro0/exp-mm-repo/releases/tag/services%2Ffoo%2Fv1.1.0)
- ```$ git tag services/foo/v1.1.0```
- and push it
9. [Require foo-service v1.1.0](https://github.com/tarotaro0/exp-mm-repo/commit/8efd7acf325a481bf7b7ba203eb585f19b6ff3fd)
- change the version of foo-service
```
$ go run main.go
Hello, Foo v1.1.0
```
10. [Add a new bar-service](https://github.com/tarotaro0/exp-mm-repo/commit/de2d5bd70326bc29c1a79285739a4aa97d1fb5de)
- ```$ go mod init github.com/tarotaro0/exp-mm-repo/services/bar```
- for checking in local
  - ```$ go  mod edit -require github.com/tarotaro0/exp-mm-repo/services/bar@v1.0.0``` in root
  - ```$ go mod edit -replace github.com/tarotaro0/exp-mm-repo/services/bar@v1.0.0=./services/bar/```
11. [Tag for bar-servicev1.0.0](https://github.com/tarotaro0/exp-mm-repo/releases/tag/services%2Fbar%2Fv1.0.0)
- same as 4. but use github release in this process
12. [Require bar-service v1.0.0](https://github.com/tarotaro0/exp-mm-repo/commit/5d1239d3f2ae756e0c7ff97075b86a5fd5b3eaab)
- same as 5.
```
$ go run main.go
Hello, Foo v1.1.0
Hello, Bar v1.0.0
```
13. [Add super log](https://github.com/tarotaro0/exp-mm-repo/commit/239a0b6fe31560d1560b6b95afa595622f9ca1c1)
- add the super log system on /conf to test the bi-directional dependency between sub module and main module
14. [Import super log system to foo-service](https://github.com/tarotaro0/exp-mm-repo/commit/3f4b6f550f5153e1a79ca05f92911f95306a051b)
- Add the dependency to root module from foo-service
15. [Tag for foo-service v1.1.1](https://github.com/tarotaro0/exp-mm-repo/releases/tag/services%2Ffoo%2Fv1.1.1)
- same as 11.
16. [Require foo-service v1.1.1](https://github.com/tarotaro0/exp-mm-repo/commit/0dbb69351d1d426e7390d7876fc834aacd22e0c7)
- same as 12.
```
$ go run main.go
Hello, Foo v1.1.0
2019/11/12 11:43:38 This is super log: Foo
Hello, Bar v1.0.0
```

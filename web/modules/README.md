# modules

modules help in depenedency management.

1. go modules : manage dependencies in Go projects. 

         cmd  : go mod init  github.com/username/repo
                [go mod init project_name (if you are practicing)]

                Dependencies are specified in the go.mod file, 
                and their versions are recorded in the go.sum

         cmd  : go mod tidy
                (removes unused dependencies or adds needed dependencies)
                 
        If you want to do it manually use below cmds.

         cmd  : go get site.com/dependency
                (by default downloads latest version)
                go get site.com/dependency@1.12.2
                (to get specified version)

         cmd  : go list -m all
                (lists all modulus used in project)



                
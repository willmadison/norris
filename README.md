norris
======
chucknorris.io client

Getting Set Up
--------------
    1. Download and install Go from https://golang.org/dl/

    2. Download and install Go dependency management tool "dep"

        go get -u github.com/golang/dep/cmd/dep
    
    3. Create new directory in your Go workspace's source directory called "norris"

    4. Run dep init
    
    5. Add testify assertion library for testing purposes
        
        dep ensure -add github.com/stretchr/testify
 




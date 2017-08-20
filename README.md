# Go-For-It #

### USE CASE ###

* Restful APIs providing Advertisement details
* Result contains Ad, RequestPath and Execution time, please see respective request detail below

### Which APIs are exposed ? ###

#### Advertisement Resources

 **[<code>GET</code> RandomAd](https://github.com/ankitggits/go-for-it/blob/master/endpoints/random_ad.md)**
 
 **[<code>GET</code> RandomCategoryAd](https://github.com/ankitggits/go-for-it/blob/master/endpoints/random_category_ad.md)**
 
 **[<code>GET</code> SearchTextCategoryAd](https://github.com/ankitggits/go-for-it/blob/master/endpoints/searchtext_category_ad.md)**
         
### Modules Used ###

   * net/http
   * log
   * time
   * io/ioutil
   * os
   * encoding/json
   * math/rand
   * strings
   * regexp
   * testing
	

### Project Structure ###

    go-for-it
        ├───advertisement
        │   ├───handler
        │   │       ├───adHandler.go
        │   │       └───regexHandler.go
        │   ├───model
        │   │       ├───ad.go
        │   │       ├───adCategory.go
        │   │       ├───adStore.go
        │   │       └───responseEntity.go
        │   ├───repo
        │   │      └───adRepository.go
        │   └───util
        │          └───adUtility.go
        ├───advertisement_test
        │   └───repo_test
        │               └───adRepository_test.go
        ├───endpoints
        │           ├───random_ad.md
        │           ├───random_category_ad.md
        │           └───searchtext_category_ad.md
        ├───.gitignore
        ├───ads.json
        ├───app.go
        └───README.md
        
### Handler-Chain ### 

    requestValidatorHandler->traceableHandler->headersHandler->regexpHandler->adHandler

*   *requestValidatorHandler*: Validates the HTTP request method, for this example only GET Method is allowed
*   *traceableHandler*: Traces all incoming HTTP requests with request path, method and execution timing
*   *headersHandler*: Adds http response header, In this example Content-Type header is added as application/json
*   *regexpHandler*: A pattern matcher for redirecting http requests to respective ad handler 

### How to run ? ### 

Application runs on port 8090 defined in main.go

# Swaggering the REST API

## Go-Swagger
https://github.com/go-swagger/go-swagger

## References for Dockerizing the App
https://docs.microsoft.com/en-us/dotnet/core/docker/building-net-docker-images

# Quick snippets

## Error while running *go get*
To address this problem, check your GOBIN variable by running `go env`.

If it is empty, then set this variable to your Go Bin directory, which uses to be `C:\Go\bin` using:

`> set GOBIN=C:\Go\bin`

## Generate spec from commented code
Take care! You'll see nothing in your `swagger.json` unless you comment out your code using one of the [*Spec from source*](https://goswagger.io/generate/spec.html) snippets.

After commenting it all, use the following command to generate a full `swagger.json` file:

`..\..\bin\swagger.exe generate spec -o .\swagger.json`

Note that I pointed the swagger.exe to a different directory, as my `bin` directory was outsite the working dir I was storing this source code. Take care to find the right path of your swagger.exe.

## Run Swagger UI

You can choose between Swagger UI or ReDoc to view an user-friendly documentation of your API, based on your swagger.json file. I chose to use Swagger as you can take advantage of its API Console. To do so, you just need to run the following command:

`..\..\bin\swagger.exe serve swagger.json -F swagger -p 8081`

Where:

* `-F swagger` means I want to use the Swagger 'flavor' instead of ReDoc;
* `-p 8081` is the port I want Swagger to expose my OpenAPI doc (swagger.json)

# Enable CORS on your Web API
Swagger exposes the documentation of your API through http://petstore.swagger.io and uses Cross Origin Resource Sharing (CORS) to make it possible. Given that, your Web API needs to enable it by adding `Access-Control-Allow-Origin` to the response header.

You can do this using one of these two options:

## On each method/operation

Add the response header using the Response Writer of the API Operation:

`w.Header().Set("Access-Control-Allow-Origin", "*")`

Where `*` can be the exact host, for security purposes: petstore.swagger.io

## For the Web Server you start:

Choose this option to enable CORS for every request of your API. 

Replace this: 

`log.Fatal(http.ListenAndServe(":8080", router))`

By this: 

`handler := cors.Default().Handler(router)`

`log.Fatal(http.ListenAndServe(":8080", handler))`

## DISCLAIMER: **DO NOT** use Microsoft Edge to navigate on the Swagger UI generated by this procedure

I still didn't find the root cause of it but, if you try out your API using Swagger on Microsoft Edge, you're going to get an error like this:

`Invariant Violation: Minified React error #31; visit http://facebook.github.io/react/docs/error-decoder.html?invariant=31&args[]=%5Bobject%20Blob%5D&args[]= for the full message or use the non-minified dev environment for full errors and additional helpful warnings.`

If you try it using Google Chrome, you're going to succeed.

# Dockerizing it

This sample does not run Swagger UI into the container. Instead, it uses the UI in http://petstore.swagger.io/ (and that's why you need to enable CORS). This is how Go-Swagger works, actually.

To make it possible, I added a few more lines to the Dockerfile to update the swagger.json file after compilation and then I copied it to the scratch image.

At last, I added a new function `SwaggerDoc` - associated to `/doc/{fileName}` route to read this file and retrieve its content so the petstore Swagger UI is able to render it.
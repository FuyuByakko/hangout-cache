#  MemoGo

MemoGo - is a simple in-memory key-value store that acts a6s a cache for your services.

It is possible to use the base verision, to store the and retrieve data as JSON objects.

An optional approach is to use Structure-based storage.

For more info on Structure based storage, please see the example inside `cacheWithStructs`.<br />
It uses the examples from a previous group project - Eventrip. More details can be found at the ![EvenTrip project page](https://github.com/cc-project-hangout/project-hangout)

The version of Eventrip that uses the cache is available at the ![Cached EvenTrip project page](https://github.com/FuyuByakko/project-hangout).

The acessible APIs and further info can be found in the *Notes* section

## Available Scripts

In the project directory, you can run:

### `go run </path/>server.go`

runs the server that acts as the chaching server.
use the link to the server as .

### `go build </path/>server.go`

Builds the app as an executable file in same folder.

### `go install </path/>server.go`

Builds the go package into an executable file in the bin folder in your root GoLang folder.<br />
Hot-reloading enabled.
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

Proxy is setup to connect to the back-end server running with the below command.

### Note

* When starting the server, if no PORT enviromental variable is available, the server port will default to `:9000`.

* Once the cache server is started, you can get more info at the base host page.

* Available APIs are:

> GET `~/api/v1/contents/?key=<...>` - where key is the key for the data you want to retrieve.

> POST `~/api/v1/contents/?key=<...>` - will post the Body of the request into the store, under the provided key.

In case of POST requests with the same key, the contents will be overwritten, so take care!

## Future Features

Add the business and API-call logic to the same GoLang-based server, to reduce the time.
Add concurrency to cache reading and writing.
Add a load balancer and serveral instances of the cache to speed up multiple request handling.
Add a backup Database

## Deployment

Please use this link access our deployed version to see the example: http://eventrip-ch.herokuapp.com/

Comparison of different languages in building a key-value pair store rest API with support for generic models, using a DynamoDB database

<b>Run the database locally</b>

`start-mock-infrastructure.sh`

<b>Run the API server for the chosen language e.g.</b>

`go/run.sh`

<b>Run a test client sending PUT and GET requests to benchmark the currently running server</b>

`python3 benchmark.py 10000`

### Sample Results:

#### C Sharp
```
PUT average time per request:
6.13 ms

GET average time per request:
4.15 ms
```

#### Go
```
PUT average time per request:
6.12 ms

GET average time per request:
4.38 ms
```

#### Node
```
PUT average time per request:
6.47 ms

GET average time per request:
4.77 ms
```

#### Python
```
PUT average time per request:
6.52 ms

GET average time per request:
4.82 ms
```


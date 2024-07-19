# Multithreading Zip code fech
Poject to resolve the Full Cycle Go Expert challenge. 

This project aims to control fetch zip code info from 2 or more different APIs, and Return the faster one before 1 second.

To test this implementation, make sure that your server is runnig. navigate to the cmd/server folder and run the following command:

```
go run main.go
```

Now you can send a GET reqeust to:

```
http://localhost:8000/{your_zipcode}
```

The results will be displayed in the command line terminal and as JSON in the API response 
This is for load testing DB, ultimately for testing Driver: bad connection error
1. If the number of concurrent connection is small (e.g. 100) no driver connection issue found
2. If the number of concurrent connection is big (1000), we will get an error ` Error 1040: Too many connections`
3. When we set the number of max connection to DB server in the client side, we can eliminate such error, but still, we got some timeout issue 
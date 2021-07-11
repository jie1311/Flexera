# Flexera

This is the Flexera code test for a job interview.

To run this code please make sure your local computeror your server has go environment, see more: https://golang.org/ 
This code was witten in go 1.16.5

After download/clone the code to your targeted environment, you can run the source code directly by 
```
  go run. <csvFilePath><applicationID>
```
![alt text](https://github.com/jie1311/Flexera/blob/main/img/go_run.png?raw=true)

Otherwise you can build the executable file and run it by
```
  go build
  .\flexera.exe <csvFilePath><applicationID>
```
![alt text](https://github.com/jie1311/Flexera/blob/main/img/go_build.png?raw=true)

Also, you can run the test by (-v gives you the details of each test)
```
  go test -v
```
![alt text](https://github.com/jie1311/Flexera/blob/main/img/go_test.png?raw=true)

Please note that the code was tested in Windows powershell.

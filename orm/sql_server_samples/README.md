See [sql server samples](https://github.com/microsoft/sql-server-samples/tree/master/samples/tutorials/go).

We can launch a docker container for sql server locally.
```bash
docker pull mcr.microsoft.com/mssql/server:2019-CU5-ubuntu-18.04
```

Then run this container
``bash
docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=<MyPassword>" ^
   -p 1433:1433 --name sql1 ^
   -d mcr.microsoft.com/mssql/server:2019-CU5-ubuntu-18.04
```
Note that the password includes angled brackets '<MyPassword>'. The password needs to be strong or the container won't launch.

Connect via SQL Server Management Studio (SSMS) using the public IP address, followed by comma separator and then the port (xxx.xx.xx.xxx,port)
Create a new database called 'TESTDB' from SSMS.

Then on with the samples...

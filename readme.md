# Reading SOVA files

A script that reads SOVA files in XML formt placed on the SFTP folder and stores the data in Oracle database

- Building in GO version go1.18.2 windows/amd64
- Database is Oracle 
- Uses the [pkg/sftp](https://github.com/pkg/sftp) and [golang.org/x/crypto/ssh](https://golang.org/x/crypto/ssh) for access the SFTP folder
- Uses [godror Oracle driver](https://github.com/godror/godror)
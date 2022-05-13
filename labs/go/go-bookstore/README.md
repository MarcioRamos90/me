
Build the image:
```s
sudo docker build -t simplerest-db .
```

Create database:
```s
sudo docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=#10293848FgT -e MYSQL_DATABASE=simplerest simplerest-db
```

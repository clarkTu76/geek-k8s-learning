FROM ubuntu

#COPY的 source 路径写相对路径
#在module2目录下运行 docker build -t gohttpserver:v1 .
#docker run -p 80:80 gohttpserver:v1
#sudo  ps aux|grep 0a267681c980 查询pid
#sudo nsenter -t 11019 -n ip addr 查看容器的网络设置
COPY httpServer  /dockerStudy/

EXPOSE 80
ENTRYPOINT  /dockerStudy/httpServer


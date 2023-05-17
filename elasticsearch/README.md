## 使用docker安装elasticsearch、kibana

```sh
# 当前目录下所有文件赋予权限(读、写、执行)
chmod -R 755 ./elasticsearch
# 创建 elasticsearch.yml 位置./config/elasticsearch.yml

# 创建 kibana.yml 位置./kibana/config/kibana.yml

# 运行
docker-compose up -d
# 运行后，再次给新创建的文件赋予权限
chmod -R 755 ./elasticsearch
```


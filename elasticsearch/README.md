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
# 设置密码
docker exec -it elasticsearch /bin/bash
elasticsearch-setup-passwords auto #随机生成
elasticsearch-setup-passwords interactive #自己设置
# 修改kibana.yml文件

```

## 安装ik[分词](https://so.csdn.net/so/search?q=分词&spm=1001.2101.3001.7020)


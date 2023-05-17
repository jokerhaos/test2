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

### 第一张安装方式解压压缩包到plugins下

```

IK 分词器 https://github.com/medcl/elasticsearch-analysis-ik/releases
拼音分词器 https://github.com/medcl/elasticsearch-analysis-pinyin/releases
```

```
#拷贝到 ./plugins 目录下

#解压
unzip -d ./elasticsearch-analysis-pinyin-7.6.2/ ./elasticsearch-analysis-pinyin-7.6.2.zip
unzip -d ./elasticsearch-analysis-ik-7.6.2/ ./elasticsearch-analysis-ik-7.6.2.zip

#删除压缩包
rm elasticsearch-analysis-pinyin-7.6.2
rm elasticsearch-analysis-ik-7.6.2

#重启容器
docker restart elasticsearch

```

### 第二种安装方式使用命令行安装

```
#进入容器
sudo docker exec -it elasticsearch /bin/bash

#进入bin目录
cd /usr/share/elasticsearch/bin

#执行命令
./elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-pinyin/releases/download/v7.6.2/elasticsearch-analysis-pinyin-7.6.2.zip
./elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.6.2/elasticsearch-analysis-ik-7.6.2.zip

#退出容器
exit

#重启容器
docker restart elasticsearch
```


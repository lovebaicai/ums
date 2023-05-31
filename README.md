### ums用户管理平台说明
- 本平台可对接内部ladp系统，通过web页面对ldap用户进行管理

### 一. 依赖环境
- 开发语言：Golang 1.18.5 + Node.js v10.22.1
- 后端框架：Gin 1.8.1
- 前端框架：Vue2 + [vue-element-admin](https://github.com/PanJiaChen/vue-element-admin) + [Element-UI](https://element.eleme.cn/#/zh-CN)
- 其他组件：Nginx、MySQL、Supervisord

### 二. 部署说明（默认路径：/opt/devops_platform/ums）
#### 1. 安装相关组件

- 安装前端依赖库
```
cd oms_frontend && npm install --unsafe-perm
```

- 安装 Nginx、MySQL、Supervisord、Node.js
```
自行安装
```

- 启动服务
```
一、启动后端服务(debug用)
go run main.go

2. supervisord后台启动, 生产使用
supervisorctl restart all

三、前端工程
1. 开发debug启动
npm run dev 

2. 生产打包
npm run build:prod
```

#### 2. 组件配置

- nginx配置：
```
worker_processes auto;

events {
    worker_connections  1024;
}
http {
    gzip on;
    gzip_min_length 1k;
    gzip_buffers 4 8k;
    gzip_comp_level 4;
    gzip_types text/plain application/x-javascript text/css text/xml application/xml text/javascript application/json application/javascript;

    include       /etc/nginx/mime.types;
    server {
        listen     80;
        server_name    ums.domain.cc;
        rewrite ^/(.*) https://ums.domain.cc/$1 permanent;
    }

    server {
        listen 443 ssl;
        server_name ums.domain.cc;
        ssl_certificate  "/etc/nginx/ssl/domain.cc.pem";
        ssl_certificate_key "/etc/nginx/ssl/domain.cc.key";
        charset     utf-8;

        # logs
        access_log  /opt/devops_platform/logs/ng_access.log;
        error_log   /opt/devops_platform/logs/ng_error.log;

        # max upload size
        client_max_body_size 1024M;   # adjust to taste

       location / {
            root   /opt/devops_platform/ums/ums_frontend/dist;
            index  index.html index.htm;
            try_files $uri $uri/ /index.html;
        }


        location /api {
            proxy_pass http://ums-server;
        }
    }
}
```


- supervisord配置
```
[unix_http_server]
file=/var/run/supervisor.sock
chmod=0700

[supervisord]
logfile=/var/log/supervisord.log ; main log file; default $CWD/supervisord.log
logfile_maxbytes=50MB        ; max main logfile bytes b4 rotation; default 50MB
logfile_backups=10           ; # of main logfile backups; 0 means none, default 10
loglevel=info                ; log level; default info; others: debug,warn,trace
pidfile=/var/run/supervisord.pid ; supervisord pidfile; default supervisord.pid
nodaemon=false               ; start in foreground if true; default false
silent=false                 ; no logs to stdout if true; default false
minfds=1024                  ; min. avail startup file descriptors; default 1024
minprocs=200                 ; min. avail process descriptors;default 200
;umask=022                   ; process file creation umask; default 022
;user=supervisord            ; setuid to this UNIX account at startup; recommended if root
;identifier=supervisor       ; supervisord identifier, default is 'supervisor'
;directory=/tmp              ; default is not to cd during start
;nocleanup=true              ; don't clean up tempfiles at start; default false
;childlogdir=/tmp            ; 'AUTO' child log dir, default $TEMP
;environment=KEY="value"     ; key value pairs to add to environment
;strip_ansi=false            ; strip ansi escape codes in logs; def. false

[rpcinterface:supervisor]
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface

[supervisorctl]
serverurl=unix:///var/run/supervisor.sock

[program:ums-server]
command=/opt/devops_platform/ums/ums_backend/ums
directory=/opt/devops_platform/ums/ums_backend
stdout_logfile=/opt/devops_platform/ums/logs/ums.log
autostart=true
autorestart=true
startsecs=0
stopwaitsecs=600
redirect_stderr=true
stopasgroup=true
killasgroup=true

[include]
files = /etc/supervisord.d/*.conf
```

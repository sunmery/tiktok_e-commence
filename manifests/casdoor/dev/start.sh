#!/usr/bin/env bash
# 启用 POSIX 模式并设置严格的错误处理机制
set -o posix errexit -o pipefail

#!/bin/bash

declare driverName
declare user
declare pass
declare db_port

# https://casdoor.org/zh/docs/basic/try-with-docker/

user="postgres"
pass="postgres"
host=postgres17
db_port="5432"
driverName="postgres"
dbname=tiktok
sslmode=disable
schema=users
dataSourceName="user=$user password=$pass host=$host port=$db_port sslmode=$sslmode dbname=$dbname search_path=$schema"

# https://github.com/casdoor/casdoor/blob/master/conf/app.conf
cat > app.conf <<EOF
appname = casdoor
httpport = 8000
runmode = dev
copyrequestbody = true
driverName = $driverName
dataSourceName = $dataSourceName
dbName = $dbname
tableNamePrefix =
showSql = false
redisEndpoint =
defaultStorageProvider =
isCloudIntranet = false
authState = "casdoor"
socks5Proxy = "127.0.0.1:10808"
verificationCodeTimeout = 10
initScore = 0
logPostOnly = true
isUsernameLowered = false
origin =
originFrontend =
staticBaseUrl = "https://cdn.casbin.org"
isDemoMode = false
batchSize = 100
enableErrorMask = false
enableGzip = true
inactiveTimeoutMinutes =
ldapServerPort = 389
radiusServerPort = 1812
radiusSecret = "secret"
quota = {"organization": -1, "user": -1, "application": -1, "provider": -1}
logConfig = {"filename": "logs/casdoor.log", "maxdays":99999, "perm":"0770"}
initDataNewOnly = false
initDataFile = "./init_data.json"
frontendBaseDir = "../cc_0"
EOF

# 运行Casdoor
#docker run \
#-itd \
#-p 8000:8000 \
#-v /home/casdoor/:/conf \
#casbin/casdoor:latest


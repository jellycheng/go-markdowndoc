
## xxx接口名称接口
### 功能描述
接口功能描述及注意事项todo。

### 请求说明
> 请求方式：POST<br>
请求URL ：/ <br>
请求格式：application/json <br>
开发者： 商城

### 请求头
参数       | 类型    | 必须   |说明
------------|-----------|-----------|-----------
app-type   |string|是 | app应用类型:cp-测评系统
app-platform   |string |是  |app所在平台：i ios，a Android
app-v|string|是|app客户端当前版本 exp: 1.0.0

> 其它请求头见接口规约《公共请求头》
### 请求参数
参数         |类型       | 必须       |说明
------------|-----------|-----------|-----------
phone |string   | 是 |手机号
sms_code |string   | 是 | 验证码

### 返回参数
参数         | 类型      | 必须 | 说明
------------|-----------|-----------|-----------
code |  Integer | 是| 0:成功，非0:失败 
msg |   String | 是 |消息
data |   Object |  是 |业务数据
↳user_id |   int|   是 | 如果登录成功，返回用户ID  
↳token |   String|   是 | 如果登录成功，登录态  
trace_id |   String | 是   |追踪id

### 错误状态码
状态码       |说明
------------|-----------
0           |正常
非0         |发生错误

### 请求示例
```
curl --location --request POST 'https://api.xxx.com/v1/user/login/loginBySms' \
--header 'x-from-service: manage-api' \
--header 'system-code: mall' \
--header 'app-type: mall' \
--header 'app-platform: m' \
--header 'app-v: 10.0.1' \
--header 'Content-Type: application/json' \
-d '{
    "phone": "18221891080",
    "sms_code": 740378
}'
```

### 返回结果示例
```
{
    "code": 0,
    "data": {
        "token": "7d823c70404a15e6",
        "user_id": 1
    },
    "msg": "success",
    "trace_id": "e5c3259362a7fb022d3e0d4a37ccfd06"
}
```
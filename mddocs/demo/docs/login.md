## 账号登录

### 功能描述
B端账号密码登录接口。

### 请求说明
> 请求方式：POST<br>
请求URL : /asst/user/v1/account/login <br>
请求格式：application/json <br>
开发者： 张三 <zhang.san@xxx.com.cn>

### 请求参数
| 参数 | 类型 | 必须 | 说明 |
| ---- | ---- | ---- | ---- |
| account_name | string | Yes | 账号 |
| password | string | Yes | 密码 |
| session_id | string | Yes | 会话id |
| captcha | string | Yes | 验证码 |

### 返回参数
| 参数 | 类型 | 必须 | 说明 |
| ---- | ---- | ---- | ---- |
| code | integer | Yes | 0:成功，非0:失败 |
| msg | string | Yes | 消息 |
| data | object | Yes | 业务数据 |
|↳token | string | Yes |登录token|
|↳servicer_id | integer | Yes |服务者id|
| trace_id | string | Yes | 追踪id |

### 错误状态码
| 状态码 | 说明 |
| ----- | ---- |
| 0     | 正常 |
| 非0   | 发生错误 |

### 请求示例
```bash

curl --location --request POST 'http://dev-api-开发者.xxx.com/asst/user/v1/account/login' \
--header 'x-from-service: mobile-api' \
--header 'x-account-type: B1' \
--header 'branchname: feature_asst_20210324' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_name": "test1",
    "password": "123456",
    "session_id": "1a673fc668ab07cb4d10166c623fc7ed",
    "captcha": "OnOr"
}'

```

### 返回结果示例
```json
{
    "code": 0,
    "data": {
        "token": "47177b4af7195cb75c54db5dfcaaaadd",
        "servicer_id": 1
    },
    "msg": "success",
    "trace_id": "ea3383af80b79c4a82c94922fc273def"
}
```


## 公共请求头
| 参数 |  类型 | 必须  | 说明 |
| :-------- | :--------:| :---: | :---: |
| token |  string | 否  | 登录token、登录令牌 |
| system-code |  string | 是  | 后台接入系统代号，应用场景仅用于控制菜单权限，manage-api：运营后台调用、supplier-api：商家后台调用 |
| enterprise-id| int | 是 | 入驻企业ID、默认1000表主公司 |
| app-type | string | 是  | 产品类型，mall：电商系统 |
| x-from-service |  string | 是  | 来源系统 mobile-api: 移动API（适用小程序、h5）、pc-api：pc网址调用、manage-api：运营后台调用、supplier-api：商家后台调用|
| app-platform | string | 是  | 终端类型：a：andorid、i：ios、h5：h5站、mp：小程序、<br> m：运营后台、s: 商家后台、pc：c端pc页面 |
| app-v |  string | 否  | app版本号，示例值：1.0.0 |
| app-channel |  string | 否  | 渠道编号，app_store：iOS渠道、site：Android主站渠道、<br> yyb：Android应用宝渠道 |
| app-os-version |  string | 否 | 系统版本号，示例值：12  |
| app-device-model |  string | 否 | 设备型号，示例值：iPhone，iPad，小米，华为 |
| app-device-id |  string | 否 | 设备ID，iOS使用openudid |
| app-idfv |  string | 否  | 供应商标识符（iOS端），android空 |
| app-idfa |  string | 否  | 广告标识符（iOS端） |
| share-id |  string | 否  | 分享id，h5和小程序才可能有值 |
| | | | &nbsp; | 


## 公共响应参数
| 参数 |  类型 | 必须  | 说明 |
| :-------- | :---:| :---:| :---: |
| code |  Integer | 是 | 错误代号，0:成功,非0:失败 |
| msg |   String | 是 | 错误消息 |
| data |   Object | 是 |业务数据 |
| trace_id | String | 否 | 追踪ID |
| elapsed_time | Double | 否 | 接口耗时,单位毫秒 |

## 返回结果示例
```json
{
  "code": 0,
  "msg": "success",
  "data": {
  },
  "trace_id": "0171050f4f63fa6780a119fe062e0000"
}
```

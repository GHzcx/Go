# 1. 获取VR价格
## 请求说明
> URL: ```/house_vr/Api_get_price```
> 
> Method: ```GET```<br>

## 请求参数

 字段       | 是否必填       | 字段说明
:------------:|:-----------:|:-----------:
 totalCount       |是        |套餐数量
 packageType       |是        |套餐类型 

> 套餐类型传枚举值 2为基础版 3为托管版

## 返回结果
```json  
{
    "code": 0,
    "message": "success",
    "data": {
        "totalPrice": 29,
        "unitPrice": 29,
        "isCoupon": false
    }
}
```

## 返回参数
字段     |字段说明
:------------:|:-----------:
totalPrice  |总价格
unitPrice | 单位价格
isCoupon | 是否为优惠价格

----

# 2. 获取活动文案配置
## 请求说明
> URL: ```/house_vr/Api_activity_config```
> 
> Method: ```GET```

## 请求参数

无

## 返回结果
```json  
{
    "code": 0,
    "message": "success",
    "data": {
        "activity_six": true,
        "content": {
            "packageMake": [
                "福利1：新用户首次自主拍摄超5套，免费赠5次升级3D服务。",
                "福利2：升级3d服务，低至5折起。单笔满2享7折，满5享5折。"
            ],
            "packageShoot": [
                "福利1：新用户首次自主拍摄超5套，免费赠5次升级3D服务。",
                "福利2：托管服务，低至4.5折。单笔满3享7折，满5享6折，满10享5.2折，满20享4.5折。"
            ]
        }
    }
}
```

## 返回参数
字段     |字段说明
:------------:|:-----------:
activity_six  |活动是否开始
content | 文案内容
packageMake |基础版文案内容
packageShoot | 托管版文案内容


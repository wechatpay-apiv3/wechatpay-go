# FavorAvailableTime

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**AvailableBeginTime** | **time.Time** | 批次开始时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 注意：商家券有效期最长为1年。 | 
**AvailableEndTime** | **time.Time** | 批次结束时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 注意：商家券有效期最长为1年。 | 
**AvailableDayAfterReceive** | **int64** | 日期区间内，券生效后x天内有效。例如生效当天内有效填1，生效后2天内有效填2，以此类推……注意，用户在有效期开始前领取商家券，则从有效期第1天开始计算天数，用户在有效期内领取商家券，则从领取当天开始计算天数，无论用户何时领取商家券，商家券在活动有效期结束后均不可用。可配合wait_days_after_receive一同填写，也可单独填写。单独填写时，有效期内领券后立即生效，生效后x天内有效。 | [可选] 
**AvailableWeek** | [**AvailableWeek**](AvailableWeek.md) | 可以设置多个星期下的多个可用时间段，比如每周二10点到18点 | [可选] 
**IrregularyAvaliableTime** | [**[]IrregularAvailableTime**](IrregularAvailableTime.md) | 无规律的有效时间，多个无规律时间段 | [可选] 
**WaitDaysAfterReceive** | **int64** | 日期区间内，用户领券后需等待x天开始生效。例如领券后当天开始生效则无需填写，领券后第2天开始生效填1，以此类推……用户在有效期开始前领取商家券，则从有效期第1天开始计算天数，用户在有效期内领取商家券，则从领取当天开始计算天数。无论用户何时领取商家券，商家券在活动有效期结束后均不可用。需配合available_day_after_receive一同填写，不可单独填写。注：最大不能超过30天 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)



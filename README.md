# 本项目主要用来满足一些日常造数需求,目前支持的数据有:
- [x] 地址相关: 省市,区域,街道和外国国名
- [x] 公司名称
- [x] 工作名称
- [x] 男/女姓名
- [x] 手机号码
- [x] 生成一段数字之间随机整数(最高支持18位)
- [x] 随机整数(最高支持18位)
- [x] 生成n位随机小数
- [x] 生成时间段之间的日期(yyyy-mm-ss,yyyy-mm-ss hh:mm:ss)
- [ ] 邮箱
- [ ] 身份证号
- [ ] ...
# 使用方式
## 直接使用
1. 获取街道小区信息
2. 获取省市名称
3. 获取详细地址
> 使用详情参考测试类
## 函数注册
1. init 方法中注册需要使用的函数

```
func init() {
	util.Seed(0)
	AddAddressFunc()
}
```

2. 根据函数名获取需要的信息

```
value, _ := util.GetFuncLookup("detailAddress").Call()
t.Log(value)
```
## 代码coding(建议)
使用方式:
```
value3, _ := CodeInfo([]string{"districts", "streets"}, "address")
t.Log(value3)
value4, err := CodeInfo([]string{"cityInfo"}, "city")
t.Log(value4,err)
value5, err := CodeInfo([]string{"prefix","middle"}, "company")
t.Log(value5,err)
value6, err := CodeInfo([]string{"middle"}, "phone")
t.Log(value6,err)
```
优势:可以自由组合自己想要的数据,比如只需要手机号的后四位,地址只需要街道等
# 注意事项
如果想生成随机数据,则每次运行时需要指定`util.Seed(0)`,可以将此函数写在init方法中
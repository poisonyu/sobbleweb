

## 12.5 reflect.Value设置值
通过reflect.ValueOf(x)返回reflect.Value的都是不可寻址的
```
d := reflect.ValueOf(&x).Elem() // 可寻址
d.CanAddr() // true
```

~~通过一个指针来间接获取一个可寻址的reflect.Value~~

从一个可寻址的reflect.Value()获取变量
```
x := 2
d := reflect.ValueOf(&x).Elem()
px := d.Addr().Interface().(*int)
*px = 3
fmt.Println(x) // 3
```

直接通过可寻址reflect.Value的来更新变量
```
x := 2
d := reflect.ValueOf(&x).Elem()
d.Set(reflect.ValueOf(4))
d.SetInt(3)
fmt.Println(x) // 4
```

确定一个reflect.Value是否可以寻址并且更改
```
fd.CanAddr()
fd.CanSet()
```

1. 基于反射的代码是很脆弱的  
对每一个reflect.Value都需要仔细注意它的类型、是否可寻址、是否可设置  
确保反射的使用完整封装在包里，在每个危险操作前都做额外的动态检查
2. 反射降低了自动重构和分析工具的安全性和准确度  
参数类型是interface{},reflect.Value，一定要写清楚参数类型和不变量
3. 基于反射的函数会比特定类型优化的函数慢一两个数量级  
与整体性能无关，可以使用反射，**测试**就很适合使用反射，因为大部分测试都使用小数据集，关键路径上的函数，最好避免使用反射
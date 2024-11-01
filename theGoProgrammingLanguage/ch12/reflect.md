

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
## 

# 10. 6 包及命名

保持可读性和和无歧义，辅助工具包可以使用imageutil或ioutil命名

包名通常使用统一的形式，标准包bytes,errors,strings使用**复数**来避免覆盖响应的预声明类型

temp作为温度转换的包名有歧义，(temporary temperature)可以使用**tempconv**，类似**strconv**
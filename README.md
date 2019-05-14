# golang
golang Learning 
## basic1

1.go语言变量类型化写在变量名之后  

2.编译器可以推测变量类型  

3.没有char 只有rune 并且其是32位的  

4.变量的定义可以简写成 := 但是包外不能这样定义  
5.只有使用Printf()使用格式化输出 其可以在进行欧拉公式计算的时候直接进行去.3f精度 以便于计算出0  
6.python中的i 写的是j 因为创造者认为 i 经常作为循环遍历的变量不适合做复数的表示形式 i在golang语句中 使用的是1i  
7.使用const定义常量  

## basic2  
在go中使用ioutil读取文件 返回两个值分别是byte[] 以及err  
if的条件中可以进行赋值  
if的条件中的作用域就在这个if的语句块中  

## loop 
for if 后面的条件没有括号  
if语句中也可以定义变量  
没有while  
switch不需要break 也可以直接switch多个条件
## fun  
返回值类型写在最后面  
可以返回多个值  
函数作为参数  
没有默认参数 可选参数  
## array  
可以使用for循环进行遍历  
可以使用range进行遍历  
数组中是值传递  
可以使用指针进行引用传递  
定义数组使用 arr := [...]int{...}  
但是上面的方式只能使用固定大小的数组 可以使用切片灵活操作  
使用range的优点 简洁 含义明确
## splice  
注意 splice是对数组的一种view 其不同于值传递和引用传递  
切片是左开右闭的一种形式  
切片只能向后扩展不能向前进行扩展  
向其中添加元素超过cap时 系统会重新分配底层数组的空间  
由于值传递的关系 必须接受append 的返回值  
s = append(s,val)

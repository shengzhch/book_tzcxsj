package main
/*
散列法：根据各元素的值来确定存储的位置，然后将位置保管在散列表中，从而实现数据的高速搜索

insert(data)  T(h(data.key))=data  [T为hash table，h为hash函数，data.key 为数据的关键字]

search(data) return T(h(data.key))

函数的返回值要求在0~m-1之间;m 为 哈希表T的大小

例如根据k值求下标  h(k) = k mod m

冲突： 不同k值对应同一个hash值。

开放地址法是解决hash冲突的常用手段之一。双散列结构中使用的开放地址法

H(k) = h(k,i) = (h1(k)+i*h2(k)) mod m

h(k,i) 先调用h1(K)计算下标，发生冲突后向后移动h2(k)个单位，直到找到空地址。

注意 ： h2(k)要与m互质
方法 ： 设置m为质数，h2(k)为小于m的值

例如

h1(k) = k mod m;
h2(k) = 1 + (k mod (m-1))



 */

 func


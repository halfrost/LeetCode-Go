---
title: 1.3 时间复杂度
type: docs
weight: 3
---

# 时间复杂度和空间复杂度


## 一. 时间复杂度数据规模

1s 内能解决问题的数据规模：10^6 ~ 10^7

- O(n^2) 算法可以处理 10^4 级别的数据规模(保守估计，处理 1000 级别的问题肯定没问题)
- O(n) 算法可以处理 10^8 级别的数据规模(保守估计，处理 10^7 级别的问题肯定没问题)
- O(nlog n) 算法可以处理 10^7 级别的数据规模(保守估计，处理 10^6 级别的问题肯定没问题)

| | 数据规模|时间复杂度 | 算法举例|
|:------:|:------:|:------:|:------:|
|1|10|O(n!)|permutation 排列|
|2|20~30|O(2^n)|combination 组合|
|3|50|O(n^4)|DFS 搜索、DP 动态规划|
|4|100|O(n^3)|任意两点最短路径、DP 动态规划|
|5|1000|O(n^2)|稠密图、DP 动态规划|
|6|10^6|O(nlog n)|排序，堆，递归与分治|
|7|10^7|O(n)|DP 动态规划、图遍历、拓扑排序、树遍历|
|8|10^9|O(sqrt(n))|筛素数、求平方根|
|9|10^10|O(log n)|二分搜索|
|10|+∞|O(1)|数学相关算法|
|------------------------------|------------------------------|------------------------------------------------------------------|------------------------------------------------------------------|


一些具有迷惑性的例子：

```c
void hello (int n){
    for( int sz = 1 ; sz < n ; sz += sz)
        for( int i = 1 ; i < n ; i ++)
            cout << "Hello" << endl;
}
```

上面这段代码的时间复杂度是 O(nlog n) 而不是 O(n^2)

```c
bool isPrime (int n){
    for( int x = 2 ; x * x <= n ; x ++ )
        if( n % x == 0)
            return false;
    return true;
}
```

上面这段代码的时间复杂度是 O(sqrt(n)) 而不是 O(n)。

再举一个例子，有一个字符串数组，将数组中的每一个字符串按照字母序排序，之后再降整个字符串数组按照字典序排序。两步操作的整体时间复杂度是多少呢？

如果回答是 O(n*nlog n + nlog n) = O(n^2log n)，这个答案是错误的。字符串的长度和数组的长度是没有关系的，所以这两个变量应该单独计算。假设最长的字符串长度为 s，数组中有 n 个字符串。对每个字符串排序的时间复杂度是 O(slog s)，将数组中每个字符串都按照字母序排序的时间复杂度是 O(n * slog s)。

将整个字符串数组按照字典序排序的时间复杂度是 O(s * nlog n)。排序算法中的 O(nlog n) 是比较的次数，由于比较的是整型数字，所以每次比较是 O(1)。但是字符串按照字典序比较，时间复杂度是 O(s)。所以字符串数组按照字典序排序的时间复杂度是 O(s * nlog n)。所以整体复杂度是 O(n * slog s) + O(s * nlog n) = O(n\*slog s + s\*nlogn) = O(n\*s\*(log s + log n))

## 二. 空间复杂度

递归调用是有空间代价的，递归算法需要保存递归栈信息，所以花费的空间复杂度会比非递归算法要高。

```c
int sum( int n ){
    assert( n >= 0 )
    int ret = 0;
    for ( int i = 0 ; i <= n ; i++)
        ret += i;
    return ret;
}
```

上面算法的时间复杂度为 O(n)，空间复杂度 O(1)。

```c
int sum( int n ){
    assert( n >= 0 )
    if ( n == 0 )
        return 0;
    return n + sum( n - 1);
}
```

上面算法的时间复杂度为 O(n)，空间复杂度 O(n)。

## 三. 递归的时间复杂度

### 1. 只有一次递归调用

如果递归函数中，只进行了一次递归调用，且递归深度为 depth，在每个递归函数中，时间复杂度为 T，那么总体的时间复杂度为 O(T * depth)

举个例子：

```c
int binarySearch(int arr[], int l, int r, int target){
	if( l > r)
	    return -1;
    int mid = l + (r-l)/2;//防溢出
    if(arr[mid] == target)
        return mid;
    else if (arr[mid]>target)
        return binarySearch(arr,l,mid-1,target);
    eles 
        return binarySearch(arr,mid+1,r,target);
}

```

在二分查找的递归实现中，只递归调用了自身。递归深度是 log n ，每次递归里面的复杂度是 O(1) 的，所以二分查找的递归实现的时间复杂度为 O(log n) 的。


### 2. 只有多次递归调用

针对多次递归调用的情况，就需要看它的计算调用的次数了。通常可以画一颗递归树来看。举例：

```c
int f(int n){
    assert( n >= 0 );
    if( n ==0 )
        return 1;
    return f( n - 1 ) + f ( n - 1 );

```

上述这次递归调用的次数为 2^0^ + 2^1^ + 2^2^ + …… + 2^n^ = 2^n+1^ - 1 = O(2^n)


> 关于更加复杂的递归的复杂度分析，请参考，主定理。主定理中针对各种复杂情况都给出了正确的结论。


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterOne/Algorithm/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/">下一章➡️</a></p>
</div>

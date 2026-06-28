---
title: 1.3 Time Complexity
type: docs
weight: 3
---

# Time Complexity and Space Complexity


## I. Data Scale for Time Complexity

Data scale of problems that can be solved within 1s: 10^6 ~ 10^7

- An O(n^2) algorithm can handle data at the 10^4 level (conservative estimate; handling problems at the 1000 level is definitely fine)
- An O(n) algorithm can handle data at the 10^8 level (conservative estimate; handling problems at the 10^7 level is definitely fine)
- An O(nlog n) algorithm can handle data at the 10^7 level (conservative estimate; handling problems at the 10^6 level is definitely fine)

| | Data Scale|Time Complexity | Algorithm Examples|
|:------:|:------:|:------:|:------:|
|1|10|O(n!)|permutation|
|2|20~30|O(2^n)|combination|
|3|50|O(n^4)|DFS search, DP dynamic programming|
|4|100|O(n^3)|Shortest paths between any two points, DP dynamic programming|
|5|1000|O(n^2)|Dense graphs, DP dynamic programming|
|6|10^6|O(nlog n)|Sorting, heaps, recursion and divide and conquer|
|7|10^7|O(n)|DP dynamic programming, graph traversal, topological sorting, tree traversal|
|8|10^9|O(sqrt(n))|Prime sieving, square root calculation|
|9|10^10|O(log n)|Binary search|
|10|+∞|O(1)|Math-related algorithms|


Some misleading examples:

```c
void hello (int n){
    for( int sz = 1 ; sz < n ; sz += sz )
        for( int i = 1 ; i < n ; i ++ )
            cout << "Hello" << endl;
}
```

The time complexity of the above code is O(nlog n), not O(n^2).

```c
bool isPrime (int n){
    if (num <= 1) return false;
    for( int x = 2 ; x * x <= n ; x ++ )
        if( n % x == 0 )
            return false;
    return true;
}
```

The time complexity of the above code is O(sqrt(n)), not O(n).

Here is another example: given an array of strings, sort each string in the array alphabetically, and then sort the entire string array lexicographically. What is the overall time complexity of these two operations?

If the answer is O(n*nlog n + nlog n) = O(n^2log n), this answer is wrong. The length of the strings has no relation to the length of the array, so these two variables should be calculated separately. Suppose the maximum string length is s, and there are n strings in the array. The time complexity of sorting each string is O(slog s), and the time complexity of sorting every string in the array alphabetically is O(n * slog s).

The time complexity of sorting the entire string array lexicographically is O(s * nlog n). The O(nlog n) in a sorting algorithm is the number of comparisons. Since what is compared is integer numbers, each comparison is O(1). But when strings are compared lexicographically, the time complexity is O(s). Therefore, the time complexity of sorting a string array lexicographically is O(s * nlog n). So the overall complexity is O(n * slog s) + O(s * nlog n) = O(n\*slog s + s\*nlogn) = O(n\*s\*(log s + log n)) = O(n\*s\*log(n\*s)).

## II. Space Complexity

Recursive calls have a space cost. Recursive algorithms need to save recursion stack information, so the space complexity consumed will be higher than that of non-recursive algorithms.

```c
int sum( int n ){
    assert( n >= 0 )
    int ret = 0;
    for ( int i = 0 ; i <= n ; i ++ )
        ret += i;
    return ret;
}
```

The time complexity of the above algorithm is O(n), and the space complexity is O(1).

```c
int sum( int n ){
    assert( n >= 0 )
    if ( n == 0 )
        return 0;
    return n + sum( n - 1 );
}
```

The time complexity of the above algorithm is O(n), and the space complexity is O(n).

## III. Time Complexity of Recursion

### 1. Only One Recursive Call

If, in a recursive function, only one recursive call is made, and the recursion depth is depth, and in each recursive function the time complexity is T, then the overall time complexity is O(T * depth).

For example:

```c
int binarySearch(int arr[], int l, int r, int target){
	if( l > r )
	    return -1;
    int mid = l + ( r - l ) / 2; // Prevent overflow
    if(arr[mid] == target)
        return mid;
    else if (arr[mid] > target)
        return binarySearch(arr,l,mid-1,target);
    else
        return binarySearch(arr,mid+1,r,target);
}

```

In the recursive implementation of binary search, it only recursively calls itself. The recursion depth is log n, and the complexity inside each recursion is O(1), so the time complexity of the recursive implementation of binary search is O(log n).


### 2. Multiple Recursive Calls

For the case of multiple recursive calls, you need to look at the number of computational calls. Usually, you can draw a recursion tree to examine it. Example:

```c
int f(int n){
    assert( n >= 0 );
    if( n == 0 )
        return 1;
    return f( n - 1 ) + f ( n - 1 );
}
```

The number of recursive calls above is 2^0^ + 2^1^ + 2^2^ + …… + 2^n^ = 2^n+1^ - 1 = O(2^n)


> For complexity analysis of more complicated recursion, please refer to the Master Theorem. The Master Theorem gives correct conclusions for various complex cases.

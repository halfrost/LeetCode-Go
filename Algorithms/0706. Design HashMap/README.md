# [706. Design HashMap](https://leetcode.com/problems/design-hashmap/)


## 题目:

Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

- `put(key, value)` : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
- `get(key)`: Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
- `remove(key)` : Remove the mapping for the value key if this map contains the mapping for the key.

**Example:**

    MyHashMap hashMap = new MyHashMap();
    hashMap.put(1, 1);          
    hashMap.put(2, 2);         
    hashMap.get(1);            // returns 1
    hashMap.get(3);            // returns -1 (not found)
    hashMap.put(2, 1);          // update the existing value
    hashMap.get(2);            // returns 1 
    hashMap.remove(2);          // remove the mapping for 2
    hashMap.get(2);            // returns -1 (not found)

**Note:**

- All keys and values will be in the range of `[0, 1000000]`.
- The number of operations will be in the range of `[1, 10000]`.
- Please do not use the built-in HashMap library.


## 题目大意

不使用任何内建的哈希表库设计一个哈希映射具体地说，你的设计应该包含以下的功能：

- put(key, value)：向哈希映射中插入(键,值)的数值对。如果键对应的值已经存在，更新这个值。
- get(key)：返回给定的键所对应的值，如果映射中不包含这个键，返回 -1。
- remove(key)：如果映射中存在这个键，删除这个数值对。

注意：

- 所有的值都在 [1, 1000000] 的范围内。
- 操作的总数目在 [1, 10000] 范围内。
- 不要使用内建的哈希库。


## 解题思路


- 简单题，设计一个 hashmap 的数据结构，要求有 `put(key, value)`，`get(key)`，`remove(key)`，这 3 个方法。设计一个 map 主要需要处理哈希冲突，一般都是链表法解决冲突。

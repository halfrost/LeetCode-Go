# [767. Reorganize String](https://leetcode.com/problems/reorganize-string/)

## 题目

Given a string S, check if the letters can be rearranged so that two characters that are adjacent to each other are not the same.

If possible, output any possible result.  If not possible, return the empty string.

Example 1:

```c
Input: S = "aab"
Output: "aba"
```

Example 2:

```c
Input: S = "aaab"
Output: ""
```

Note:   

S will consist of lowercase letters and have length in range [1, 500].


## 题目大意

给定一个字符串，要求重新排列字符串，让字符串两两字符不相同，如果可以实现，即输出最终的字符串，如果不能让两两不相同，则输出空字符串。

## 解题思路

这道题有 2 种做法。第一种做法是先统计每个字符的出现频率次数，按照频率次数从高往低排序。具体做法就是第 451 题了。如果有一个字母的频次次数超过了 (len(string)+1)/2 那么就返回空字符串。否则输出最终满足题意的字符串。按照频次排序以后，用 2 个指针，一个从 0 开始，另外一个从中间位置开始，依次取出一个字符拼接起来。

第二种做法是用优先队列，结点是一个结构体，结构体有 2 个字段，一个字段记录是哪个字符，另一个字段记录是这个字符的频次。按照频次的多作为优先级高，用大根堆建立优先队列。注意，这样建立成功的优先队列，重复字母只有一个结点，频次记录在结构体的频次字段中。额外还需要一个辅助队列。优先队列每次都出队一个优先级最高的，然后频次减一，最终结果加上这个字符。然后将这个结点入队。入队的意义是检测这个结点的频次有没有减到 0，如果还不为 0 ，再插入优先队列中。

```c
    string reorganizeString(string S) {
        vector<int> mp(26);
        int n = S.size();
        for (char c: S)
            ++mp[c-'a'];
        priority_queue<pair<int, char>> pq;
        for (int i = 0; i < 26; ++i) {
            if (mp[i] > (n+1)/2) return "";
            if (mp[i]) pq.push({mp[i], i+'a'});
        }
        queue<pair<int, char>> myq;
        string ans;
        while (!pq.empty() || myq.size() > 1) {
            if (myq.size() > 1) { // 注意这里要大于 1，如果是等于 1 的话，频次大的元素一直在输出了，答案就不对了。
                auto cur = myq.front();
                myq.pop();
                if (cur.first != 0) pq.push(cur);
            }
            if (!pq.empty()) {
                auto cur = pq.top();
                pq.pop();
                ans += cur.second;
                cur.first--;
                myq.push(cur);
            }
        }
        return ans;
    }
```















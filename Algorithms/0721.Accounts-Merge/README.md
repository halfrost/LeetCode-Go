# [721. Accounts Merge](https://leetcode.com/problems/accounts-merge/)


## 题目:

Given a list `accounts`, each element `accounts[i]` is a list of strings, where the first element `accounts[i][0]` is a name, and the rest of the elements are emailsrepresenting emails of the account.

Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there is some email that is common to both accounts. Note that even if two accounts have the same name, they may belong to different people as people could have the same name. A person can have any number of accounts initially, but all of their accounts definitely have the same name.

After merging the accounts, return the accounts in the following format: the first element of each account is the name, and the rest of the elements are emails **in sorted order**. The accounts themselves can be returned in any order.

**Example 1:**

    Input: 
    accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
    Output: [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
    Explanation: 
    The first and third John's are the same person as they have the common email "johnsmith@mail.com".
    The second John and Mary are different people as none of their email addresses are used by other accounts.
    We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'], 
    ['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.

**Note:**

- The length of `accounts` will be in the range `[1, 1000]`.
- The length of `accounts[i]` will be in the range `[1, 10]`.
- The length of `accounts[i][j]` will be in the range `[1, 30]`.


## 题目大意


给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该帐户的邮箱地址。现在，我们想合并这些帐户。如果两个帐户都有一些共同的邮件地址，则两个帐户必定属于同一个人。请注意，即使两个帐户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的帐户，但其所有帐户都具有相同的名称。合并帐户后，按以下格式返回帐户：每个帐户的第一个元素是名称，其余元素是按顺序排列的邮箱地址。accounts 本身可以以任意顺序返回。


注意：

- accounts 的长度将在 [1，1000] 的范围内。
- accounts[i] 的长度将在 [1，10] 的范围内。
- accounts[i][j] 的长度将在 [1，30] 的范围内。



## 解题思路


- 给出一堆账户和对应的邮箱。要求合并同一个人的多个邮箱账户。如果判断是同一个人呢？如果这个人名和所属的其中之一的邮箱是相同的，就判定这是同一个人的邮箱，那么就合并这些邮箱。
- 这题的解题思路是并查集。不过如果用暴力合并的方法，时间复杂度非常差。优化方法是先把每组数据都进行编号，人编号，每个邮箱都进行编号。这个映射关系用 `map` 记录起来。如果利用并查集的 `union()` 操作，把这些编号都进行合并。最后把人的编号和对应邮箱的编号拼接起来。
- 这一题有 2 处比较“坑”的是，不需要合并的用户的邮箱列表也是需要排序和去重的，同一个人的所有邮箱集合都要合并到一起。具体见测试用例。不过题目中也提到了这些点，也不能算题目坑，只能归自己没注意这些边界情况。

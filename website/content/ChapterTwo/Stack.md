---
title: Stack
type: docs
---

# Stack

![](https://img.halfrost.com/Leetcode/Stack.png)

- 括号匹配问题及类似问题。第 20 题，第 921 题，第 1021 题。
- 栈的基本 pop 和 push 操作。第 71 题，第 150 题，第 155 题，第 224 题，第 225 题，第 232 题，第 946 题，第 1047 题。
- 利用栈进行编码问题。第 394 题，第 682 题，第 856 题，第 880 题。
- **单调栈**。**利用栈维护一个单调递增或者递减的下标数组**。第 84 题，第 456 题，第 496 题，第 503 题，第 739 题，第 901 题，第 907 题，第 1019 题。



| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0020|Valid Parentheses|[Go]({{< relref "/ChapterFour/0020.Valid-Parentheses.md" >}})|Easy| O(log n)| O(1)||39.8%|
|0042|Trapping Rain Water|[Go]({{< relref "/ChapterFour/0042.Trapping-Rain-Water.md" >}})|Hard| O(n)| O(1)|❤️|50.8%|
|0071|Simplify Path|[Go]({{< relref "/ChapterFour/0071.Simplify-Path.md" >}})|Medium| O(n)| O(n)|❤️|33.6%|
|0084|Largest Rectangle in Histogram|[Go]({{< relref "/ChapterFour/0084.Largest-Rectangle-in-Histogram.md" >}})|Hard| O(n)| O(n)|❤️|36.9%|
|0094|Binary Tree Inorder Traversal|[Go]({{< relref "/ChapterFour/0094.Binary-Tree-Inorder-Traversal.md" >}})|Medium| O(n)| O(1)||65.5%|
|0103|Binary Tree Zigzag Level Order Traversal|[Go]({{< relref "/ChapterFour/0103.Binary-Tree-Zigzag-Level-Order-Traversal.md" >}})|Medium| O(n)| O(n)||49.8%|
|0144|Binary Tree Preorder Traversal|[Go]({{< relref "/ChapterFour/0144.Binary-Tree-Preorder-Traversal.md" >}})|Medium| O(n)| O(1)||57.2%|
|0145|Binary Tree Postorder Traversal|[Go]({{< relref "/ChapterFour/0145.Binary-Tree-Postorder-Traversal.md" >}})|Medium| O(n)| O(1)||57.2%|
|0150|Evaluate Reverse Polish Notation|[Go]({{< relref "/ChapterFour/0150.Evaluate-Reverse-Polish-Notation.md" >}})|Medium| O(n)| O(1)||37.6%|
|0155|Min Stack|[Go]({{< relref "/ChapterFour/0155.Min-Stack.md" >}})|Easy| O(n)| O(n)||46.1%|
|0173|Binary Search Tree Iterator|[Go]({{< relref "/ChapterFour/0173.Binary-Search-Tree-Iterator.md" >}})|Medium| O(n)| O(1)||59.7%|
|0224|Basic Calculator|[Go]({{< relref "/ChapterFour/0224.Basic-Calculator.md" >}})|Hard| O(n)| O(n)||38.0%|
|0225|Implement Stack using Queues|[Go]({{< relref "/ChapterFour/0225.Implement-Stack-using-Queues.md" >}})|Easy| O(n)| O(n)||47.0%|
|0232|Implement Queue using Stacks|[Go]({{< relref "/ChapterFour/0232.Implement-Queue-using-Stacks.md" >}})|Easy| O(n)| O(n)||51.7%|
|0331|Verify Preorder Serialization of a Binary Tree|[Go]({{< relref "/ChapterFour/0331.Verify-Preorder-Serialization-of-a-Binary-Tree.md" >}})|Medium| O(n)| O(1)||40.9%|
|0385|Mini Parser|[Go]({{< relref "/ChapterFour/0385.Mini-Parser.md" >}})|Medium||||34.3%|
|0394|Decode String|[Go]({{< relref "/ChapterFour/0394.Decode-String.md" >}})|Medium| O(n)| O(n)||52.4%|
|0402|Remove K Digits|[Go]({{< relref "/ChapterFour/0402.Remove-K-Digits.md" >}})|Medium| O(n)| O(1)||28.6%|
|0456|132 Pattern|[Go]({{< relref "/ChapterFour/0456.132-Pattern.md" >}})|Medium| O(n)| O(n)||30.6%|
|0496|Next Greater Element I|[Go]({{< relref "/ChapterFour/0496.Next-Greater-Element-I.md" >}})|Easy| O(n)| O(n)||65.2%|
|0503|Next Greater Element II|[Go]({{< relref "/ChapterFour/0503.Next-Greater-Element-II.md" >}})|Medium| O(n)| O(n)||58.2%|
|0636|Exclusive Time of Functions|[Go]({{< relref "/ChapterFour/0636.Exclusive-Time-of-Functions.md" >}})|Medium| O(n)| O(n)||54.0%|
|0682|Baseball Game|[Go]({{< relref "/ChapterFour/0682.Baseball-Game.md" >}})|Easy| O(n)| O(n)||66.2%|
|0726|Number of Atoms|[Go]({{< relref "/ChapterFour/0726.Number-of-Atoms.md" >}})|Hard| O(n)| O(n) |❤️|51.0%|
|0735|Asteroid Collision|[Go]({{< relref "/ChapterFour/0735.Asteroid-Collision.md" >}})|Medium| O(n)| O(n) ||43.2%|
|0739|Daily Temperatures|[Go]({{< relref "/ChapterFour/0739.Daily-Temperatures.md" >}})|Medium| O(n)| O(n) ||64.4%|
|0844|Backspace String Compare|[Go]({{< relref "/ChapterFour/0844.Backspace-String-Compare.md" >}})|Easy| O(n)| O(n) ||46.8%|
|0856|Score of Parentheses|[Go]({{< relref "/ChapterFour/0856.Score-of-Parentheses.md" >}})|Medium| O(n)| O(n)||62.3%|
|0880|Decoded String at Index|[Go]({{< relref "/ChapterFour/0880.Decoded-String-at-Index.md" >}})|Medium| O(n)| O(n)||28.3%|
|0895|Maximum Frequency Stack|[Go]({{< relref "/ChapterFour/0895.Maximum-Frequency-Stack.md" >}})|Hard| O(n)| O(n)  ||62.2%|
|0901|Online Stock Span|[Go]({{< relref "/ChapterFour/0901.Online-Stock-Span.md" >}})|Medium| O(n)| O(n)  ||61.2%|
|0907|Sum of Subarray Minimums|[Go]({{< relref "/ChapterFour/0907.Sum-of-Subarray-Minimums.md" >}})|Medium| O(n)| O(n)|❤️|33.2%|
|0921|Minimum Add to Make Parentheses Valid|[Go]({{< relref "/ChapterFour/0921.Minimum-Add-to-Make-Parentheses-Valid.md" >}})|Medium| O(n)| O(n)||74.6%|
|0946|Validate Stack Sequences|[Go]({{< relref "/ChapterFour/0946.Validate-Stack-Sequences.md" >}})|Medium| O(n)| O(n)||63.5%|
|1003|Check If Word Is Valid After Substitutions|[Go]({{< relref "/ChapterFour/1003.Check-If-Word-Is-Valid-After-Substitutions.md" >}})|Medium| O(n)| O(1)||56.2%|
|1019|Next Greater Node In Linked List|[Go]({{< relref "/ChapterFour/1019.Next-Greater-Node-In-Linked-List.md" >}})|Medium| O(n)| O(1)||58.2%|
|1021|Remove Outermost Parentheses|[Go]({{< relref "/ChapterFour/1021.Remove-Outermost-Parentheses.md" >}})|Easy| O(n)| O(1)||78.7%|
|1047|Remove All Adjacent Duplicates In String|[Go]({{< relref "/ChapterFour/1047.Remove-All-Adjacent-Duplicates-In-String.md" >}})|Easy| O(n)| O(1)||70.3%|
|1673|Find the Most Competitive Subsequence|[Go]({{< relref "/ChapterFour/1673.Find-the-Most-Competitive-Subsequence.md" >}})|Medium||||45.1%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Tree/">下一页➡️</a></p>
</div>


# LeetCode by Go
[LeetCode Online Judge] (https://leetcode.com/) is a website containing many **algorithm questions**. Most of them are real interview questions of **Google, Facebook, LinkedIn, Apple**, etc. This repo shows my solutions by Go with the code style strictly follows the [Google Golang Style Guide](https://github.com/golang/go/wiki/CodeReviewComments). Please feel free to reference and **STAR** to support this repo, thank you!


<p align='center'>
<img src='./logo.png'>
</p>




## 数据结构

| 数据结构 | 变种 | 相关题目 |
|:-------:|:-------|:------|
|顺序线性表：向量|||
|单链表|1.双链表<br>2.静态链表<br>3.对称矩阵<br>4.稀疏矩阵||
|栈|广义栈||
|队列|1.链表实现<br>2.循环数组实现<br>3.双端队列||
|字符串|1.KMP算法<br>2.有限状态自动机<br>3.模式匹配有限状态自动机<br>4.BM模式匹配算法<br>5.BM-KMP算法||
|树|1.二叉树<br>2.并查集<br>3.Huffman数||
|数组实现的堆|1.极大堆和极小堆<br>2.极大极小堆<br>3.双端堆<br>4.d叉堆||
|树实现的堆|1.左堆<br>2.扁堆<br>3.二项式堆<br>4.斐波那契堆<br>5.配对堆||
|查找|1.哈希表<br>2.跳跃表<br>3.排序二叉树<br>4.AVL树<br>5.B树<br>6.AA树<br>7.红黑树<br>8.排序二叉堆<br>9.Splay树<br>10.双链树<br>11.Trie树||


## Data Structures
* [Array](#array)
* [String](#string)
* [Linked List](#linked-list)
* [Stack](#stack)
* [Tree](#tree)
* [Dynamic programming](#dynamic-programming)
* [Depth-first search](#depth-first-search)
* [Math](#math)
* [Search](#search)
* [Sort](#sort)
* [Union Find](#union-find)

## Companies
* [Google](#google)
* [Facebook](#facebook)
* [Snapchat](#snapchat)
* [Uber](#uber)
* [Airbnb](#airbnb)
* [LinkedIn](#linkedin)
* [Amazon](#amazon)
* [Microsoft](#microsoft)


## 算法


## 一. 目录

|题号|题目|通过率|难度|收藏|  
|:--:|:--|:--: | :--: | :--: |  
|1|[Two Sum](./Algorithms/1.two-sum)|38%|Easy||  
|2|[Add Two Numbers](./Algorithms/2.add-two-numbers)|28%|Medium||
|3|[Longest Substring Without Repeating Characters](./Algorithms/3.longest-substring-without-repeating-characters)|24%|Medium||
|4|[Median of Two Sorted Arrays](./Algorithms/4.median-of-two-sorted-arrays)|23%|Hard||
|5|[Longest Palindromic Substring](./Algorithms/5.longest-palindromic-substring)|25%|Medium||
|6|[ZigZag Conversion](./Algorithms/6.zigzag-conversion)|27%|Medium||
|7|[Reverse Integer](./Algorithms/7.reverse-integer)|24%|Easy||
|8|[String to Integer (atoi)](./Algorithms/8.string-to-integer-atoi)|14%|Medium||
|9|[Palindrome Number](./Algorithms/9.palindrome-number)|35%|Easy||
|10|[Regular Expression Matching](./Algorithms/10.regular-expression-matching)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|11|[Container With Most Water](./Algorithms/11.container-with-most-water)|37%|Medium||
|12|[Integer to Roman](./Algorithms/12.integer-to-roman)|46%|Medium||
|13|[Roman to Integer](./Algorithms/13.roman-to-integer)|48%|Easy||
|14|[Longest Common Prefix](./Algorithms/14.longest-common-prefix)|31%|Easy||
|15|[3Sum](./Algorithms/15.3sum)|21%|Medium||
|16|[3Sum Closest](./Algorithms/16.3sum-closest)|31%|Medium||
|17|[Letter Combinations of a Phone Number](./Algorithms/17.letter-combinations-of-a-phone-number)|36%|Medium||
|18|[4Sum](./Algorithms/18.4sum)|27%|Medium||
|19|[Remove Nth Node From End of List](./Algorithms/19.remove-nth-node-from-end-of-list)|33%|Medium||
|20|[Valid Parentheses](./Algorithms/20.Valid-Parentheses)|33%|Easy||
|21|[Merge Two Sorted Lists](./Algorithms/21.merge-two-sorted-lists)|41%|Easy||
|22|[Generate Parentheses](./Algorithms/22.generate-parentheses)|48%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|23|[Merge k Sorted Lists](./Algorithms/23.merge-k-sorted-lists)|28%|Hard||
|24|[Swap Nodes in Pairs](./Algorithms/24.Swap-Nodes-In-Pairs)|39%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|25|[Reverse Nodes in k-Group](./Algorithms/25.Reverse-Nodes-In-k-Group)|31%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|26|[Remove Duplicates from Sorted Array](./Algorithms/26.remove-duplicates-from-sorted-array)|36%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|27|[Remove Element](./Algorithms/27.remove-element)|40%|Easy||
|28|[Implement strStr()](./Algorithms/28.implement-strstr)|28%|Easy||
|29|[Divide Two Integers](./Algorithms/29.divide-two-integers)|15%|Medium||
|30|[Substring with Concatenation of All Words](./Algorithms/30.substring-with-concatenation-of-all-words)|22%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|31|[Next Permutation](./Algorithms/31.next-permutation)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|32|[Longest Valid Parentheses](./Algorithms/32.longest-valid-parentheses)|23%|Hard||
|33|[Search in Rotated Sorted Array](./Algorithms/33.search-in-rotated-sorted-array)|31%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|34|[Search for a Range](./Algorithms/34.search-for-a-range)|31%|Medium||
|35|[Search Insert Position](./Algorithms/35.search-insert-position)|40%|Easy||
|36|[Valid Sudoku](./Algorithms/36.valid-sudoku)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|37|[Sudoku Solver](./Algorithms/37.sudoku-solver)|32%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|38|[Count and Say](./Algorithms/38.count-and-say)|36%|Easy||
|39|[Combination Sum](./Algorithms/39.combination-sum)|41%|Medium||
|40|[Combination Sum II](./Algorithms/40.combination-sum-ii)|35%|Medium||
|41|[First Missing Positive](./Algorithms/41.First-Missing-Positive)|25%|Hard||
|42|[Trapping Rain Water](./Algorithms/42.trapping-rain-water)|37%|Hard||
|43|[Multiply Strings](./Algorithms/43.multiply-strings)|28%|Medium||
|44|[Wildcard Matching](./Algorithms/44.wildcard-matching)|21%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|45|[Jump Game II](./Algorithms/45.jump-game-ii)|26%|Hard||
|46|[Permutations](./Algorithms/46.permutations)|47%|Medium||
|47|[Permutations II](./Algorithms/47.permutations-ii)|35%|Medium||
|48|[Rotate Image](./Algorithms/48.rotate-image)|41%|Medium||
|49|[Group Anagrams](./Algorithms/49.group-anagrams)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|50|[Pow(x, n)](./Algorithms/50.powx-n)|26%|Medium||
|51|[N-Queens](./Algorithms/51.n-queens)|33%|Hard||
|52|[N-Queens II](./Algorithms/52.n-queens-ii)|46%|Hard||
|53|[Maximum Subarray](./Algorithms/53.maximum-subarray)|40%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|54|[Spiral Matrix](./Algorithms/54.spiral-matrix)|27%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|55|[Jump Game](./Algorithms/55.jump-game)|29%|Medium||
|56|[Merge Intervals](./Algorithms/56.merge-intervals)|31%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|57|[Insert Interval](./Algorithms/57.insert-interval)|28%|Hard||
|58|[Length of Last Word](./Algorithms/58.length-of-last-word)|32%|Easy||
|59|[Spiral Matrix II](./Algorithms/59.spiral-matrix-ii)|41%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|60|[Permutation Sequence](./Algorithms/60.permutation-sequence)|29%|Medium||
|61|[Rotate List](./Algorithms/61.rotate-list)|24%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|62|[Unique Paths](./Algorithms/62.unique-paths)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|63|[Unique Paths II](./Algorithms/63.unique-paths-ii)|32%|Medium||
|64|[Minimum Path Sum](./Algorithms/64.minimum-path-sum)|40%|Medium||
|65|[Valid Number](./Algorithms/65.valid-number)|12%|Hard||
|66|[Plus One](./Algorithms/66.plus-one)|39%|Easy||
|67|[Add Binary](./Algorithms/67.add-binary)|34%|Easy||
|68|[Text Justification](./Algorithms/68.text-justification)|20%|Hard||
|69|[Sqrt(x)](./Algorithms/69.sqrtx)|28%|Easy||
|70|[Climbing Stairs](./Algorithms/70.climbing-stairs)|41%|Easy||
|71|[Simplify Path](./Algorithms/71.simplify-path)|26%|Medium||
|72|[Edit Distance](./Algorithms/72.edit-distance)|32%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|73|[Set Matrix Zeroes](./Algorithms/73.set-matrix-zeroes)|36%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|74|[Search a 2D Matrix](./Algorithms/74.search-a-2d-matrix)|34%|Medium||
|75|[Sort Colors](./Algorithms/75.sort-colors)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|76|[Minimum Window Substring](./Algorithms/76.minimum-window-substring)|26%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|77|[Combinations](./Algorithms/77.combinations)|41%|Medium||
|78|[Subsets](./Algorithms/78.subsets)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|79|[Word Search](./Algorithms/79.word-search)|28%|Medium||
|80|[Remove Duplicates from Sorted Array II](./Algorithms/80.remove-duplicates-from-sorted-array-ii)|36%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|81|[Search in Rotated Sorted Array II](./Algorithms/81.search-in-rotated-sorted-array-ii)|32%|Medium||
|82|[Remove Duplicates from Sorted List II](./Algorithms/82.remove-duplicates-from-sorted-list-ii)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|83|[Remove Duplicates from Sorted List](./Algorithms/83.remove-duplicates-from-sorted-list)|40%|Easy||
|84|[Largest Rectangle in Histogram](./Algorithms/84.largest-rectangle-in-histogram)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|85|[Maximal Rectangle](./Algorithms/85.maximal-rectangle)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|86|[Partition List](./Algorithms/86.partition-list)|33%|Medium||
|87|[Scramble String](./Algorithms/87.scramble-string)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|88|[Merge Sorted Array](./Algorithms/88.Merge-Sorted-Array)|32%|Easy||
|89|[Gray Code](./Algorithms/89.gray-code)|42%|Medium||
|90|[Subsets II](./Algorithms/90.subsets-ii)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|91|[Decode Ways](./Algorithms/91.decode-ways)|20%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|92|[Reverse Linked List II](./Algorithms/92.reverse-linked-list-ii)|31%|Medium||
|93|[Restore IP Addresses](./Algorithms/93.restore-ip-addresses)|28%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|94|[Binary Tree Inorder Traversal](./Algorithms/94.binary-tree-inorder-traversal)|49%|Medium||
|95|[Unique Binary Search Trees II](./Algorithms/95.unique-binary-search-trees-ii)|32%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|96|[Unique Binary Search Trees](./Algorithms/96.unique-binary-search-trees)|41%|Medium||
|97|[Interleaving String](./Algorithms/97.interleaving-string)|25%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|98|[Validate Binary Search Tree](./Algorithms/98.validate-binary-search-tree)|24%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|99|[Recover Binary Search Tree](./Algorithms/99.recover-binary-search-tree)|31%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|100|[Same Tree](./Algorithms/100.same-tree)|47%|Easy||
|101|[Symmetric Tree](./Algorithms/101.symmetric-tree)|40%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|102|[Binary Tree Level Order Traversal](./Algorithms/102.binary-tree-level-order-traversal)|42%|Medium||
|103|[Binary Tree Zigzag Level Order Traversal](./Algorithms/103.binary-tree-zigzag-level-order-traversal)|36%|Medium||
|104|[Maximum Depth of Binary Tree](./Algorithms/104.maximum-depth-of-binary-tree)|54%|Easy||
|105|[Construct Binary Tree from Preorder and Inorder Traversal](./Algorithms/105.construct-binary-tree-from-preorder-and-inorder-traversal)|34%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|106|[Construct Binary Tree from Inorder and Postorder Traversal](./Algorithms/106.construct-binary-tree-from-inorder-and-postorder-traversal)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|107|[Binary Tree Level Order Traversal II](./Algorithms/107.binary-tree-level-order-traversal-ii)|42%|Easy||
|108|[Convert Sorted Array to Binary Search Tree](./Algorithms/108.convert-sorted-array-to-binary-search-tree)|44%|Easy||
|109|[Convert Sorted List to Binary Search Tree](./Algorithms/109.convert-sorted-list-to-binary-search-tree)|35%|Medium||
|110|[Balanced Binary Tree](./Algorithms/110.balanced-binary-tree)|38%|Easy||
|111|[Minimum Depth of Binary Tree](./Algorithms/111.minimum-depth-of-binary-tree)|33%|Easy||
|112|[Path Sum](./Algorithms/112.path-sum)|34%|Easy||
|113|[Path Sum II](./Algorithms/113.path-sum-ii)|35%|Medium||
|114|[Flatten Binary Tree to Linked List](./Algorithms/114.flatten-binary-tree-to-linked-list)|36%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|115|[Distinct Subsequences](./Algorithms/115.distinct-subsequences)|32%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|118|[Pascal's Triangle](./Algorithms/118.pascals-triangle)|40%|Easy||
|119|[Pascal's Triangle II](./Algorithms/119.pascals-triangle-ii)|38%|Easy||
|120|[Triangle](./Algorithms/120.triangle)|34%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|121|[Best Time to Buy and Sell Stock](./Algorithms/121.best-time-to-buy-and-sell-stock)|42%|Easy||
|122|[Best Time to Buy and Sell Stock II](./Algorithms/122.best-time-to-buy-and-sell-stock-ii)|47%|Easy||
|123|[Best Time to Buy and Sell Stock III](./Algorithms/123.best-time-to-buy-and-sell-stock-iii)|30%|Hard||
|124|[Binary Tree Maximum Path Sum](./Algorithms/124.binary-tree-maximum-path-sum)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|125|[Valid Palindrome](./Algorithms/125.Valid-Palindrome)|27.0%|Easy||
|126|[Word Ladder II](./Algorithms/126.word-ladder-ii)|14%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|127|[Word Ladder](./Algorithms/127.word-ladder)|20%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|128|[Longest Consecutive Sequence](./Algorithms/128.longest-consecutive-sequence)|38%|Hard||
|129|[Sum Root to Leaf Numbers](./Algorithms/129.sum-root-to-leaf-numbers)|37%|Medium||
|130|[Surrounded Regions](./Algorithms/130.surrounded-regions)|19%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|131|[Palindrome Partitioning](./Algorithms/131.palindrome-partitioning)|35%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|132|[Palindrome Partitioning II](./Algorithms/132.palindrome-partitioning-ii)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|134|[Gas Station](./Algorithms/134.gas-station)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|135|[Candy](./Algorithms/135.candy)|25%|Hard||
|136|[Single Number](./Algorithms/136.single-number)|55%|Easy||
|137|[Single Number II](./Algorithms/137.single-number-ii)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|139|[Word Break](./Algorithms/139.word-break)|31%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|140|[Word Break II](./Algorithms/140.word-break-ii)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|141|[Linked List Cycle](./Algorithms/140.Linked-List-Cycle)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|142|[Linked List Cycle II](./Algorithms/140.Linked-List-Cycle-II)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|143|[Reorder List](./Algorithms/143.reorder-list)|26%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|144|[Binary Tree Preorder Traversal](./Algorithms/144.binary-tree-preorder-traversal)|46%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|145|[Binary Tree Postorder Traversal](./Algorithms/145.binary-tree-postorder-traversal)|42%|Hard||
|146|[LRU Cache](./Algorithms/146.lru-cache)|19%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|147|[Insertion Sort List](./Algorithms/147.insertion-sort-list)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|148|[Sort List](./Algorithms/148.sort-list)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|149|[Max Points on a Line](./Algorithms/149.max-points-on-a-line)|15%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|150|[Evaluate Reverse Polish Notation](./Algorithms/150.evaluate-reverse-polish-notation)|28%|Medium||
|152|[Maximum Product Subarray](./Algorithms/152.maximum-product-subarray)|26%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|153|[Find Minimum in Rotated Sorted Array](./Algorithms/153.find-minimum-in-rotated-sorted-array)|40%|Medium||
|154|[Find Minimum in Rotated Sorted Array II](./Algorithms/154.find-minimum-in-rotated-sorted-array-ii)|37%|Hard||
|155|[Min Stack](./Algorithms/155.min-stack)|31%|Easy||
|162|[Find Peak Element](./Algorithms/162.find-peak-element)|38%|Medium||
|164|[Maximum Gap](./Algorithms/164.maximum-gap)|30%|Hard||
|165|[Compare Version Numbers](./Algorithms/165.compare-version-numbers)|20%|Medium||
|166|[Fraction to Recurring Decimal](./Algorithms/166.fraction-to-recurring-decimal)|18%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|167|[Two Sum II - Input array is sorted](./Algorithms/167.two-sum-ii-input-array-is-sorted)|47%|Easy||
|168|[Excel Sheet Column Title](./Algorithms/168.excel-sheet-column-title)|27%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|169|[Majority Element](./Algorithms/169.majority-element)|48%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|171|[Excel Sheet Column Number](./Algorithms/171.excel-sheet-column-number)|48%|Easy||
|172|[Factorial Trailing Zeroes](./Algorithms/172.factorial-trailing-zeroes)|36%|Easy||
|174|[Dungeon Game](./Algorithms/174.dungeon-game)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|179|[Largest Number](./Algorithms/179.largest-number)|23%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|187|[Repeated DNA Sequences](./Algorithms/187.repeated-dna-sequences)|33%|Medium||
|188|[Best Time to Buy and Sell Stock IV](./Algorithms/188.best-time-to-buy-and-sell-stock-iv)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|189|[Rotate Array](./Algorithms/189.rotate-array)|25%|Easy||
|198|[House Robber](./Algorithms/198.house-robber)|39%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|199|[Binary Tree Right Side View](./Algorithms/199.binary-tree-right-side-view)|42%|Medium||
|200|[Number of Islands](./Algorithms/200.number-of-islands)|36%|Medium||
|201|[Bitwise AND of Numbers Range](./Algorithms/201.bitwise-and-of-numbers-range)|34%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|202|[Happy Number](./Algorithms/202.happy-number)|41%|Easy||
|203|[Remove Linked List Elements](./Algorithms/203.remove-linked-list-elements)|33%|Easy||
|204|[Count Primes](./Algorithms/204.count-primes)|26%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|205|[Isomorphic Strings](./Algorithms/205.isomorphic-strings)|34%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|206|[Reverse Linked List](./Algorithms/206.Reverse-Linked-List)|46%|Easy||
|207|[Course Schedule](./Algorithms/207.course-schedule)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|208|[Implement Trie (Prefix Tree)](./Algorithms/208.implement-trie-prefix-tree)|30%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|209|[Minimum Size Subarray Sum](./Algorithms/209.minimum-size-subarray-sum)|32%|Medium||
|210|[Course Schedule II](./Algorithms/210.course-schedule-ii)|30%|Medium||
|211|[Add and Search Word - Data structure design](./Algorithms/211.add-and-search-word-data-structure-design)|25%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|212|[Word Search II](./Algorithms/212.word-search-ii)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|213|[House Robber II](./Algorithms/213.house-robber-ii)|34%|Medium||
|214|[Shortest Palindrome](./Algorithms/214.shortest-palindrome)|25%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|215|[Kth Largest Element in an Array](./Algorithms/215.kth-largest-element-in-an-array)|40%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|216|[Combination Sum III](./Algorithms/216.combination-sum-iii)|47%|Medium||
|217|[Contains Duplicate](./Algorithms/217.contains-duplicate)|47%|Easy||
|218|[The Skyline Problem](./Algorithms/218.the-skyline-problem)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|219|[Contains Duplicate II](./Algorithms/219.contains-duplicate-ii)|32%|Easy||
|220|[Contains Duplicate III](./Algorithms/220.contains-duplicate-iii)|18%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|221|[Maximal Square](./Algorithms/221.maximal-square)|30%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|223|[Rectangle Area](./Algorithms/223.rectangle-area)|33%|Medium||
|224|[Basic Calculator](./Algorithms/224.basic-calculator)|28%|Hard||
|225|[Implement Stack using Queues](./Algorithms/225.implement-stack-using-queues)|34%|Easy||
|226|[Invert Binary Tree](./Algorithms/226.invert-binary-tree)|53%|Easy||
|227|[Basic Calculator II](./Algorithms/227.basic-calculator-ii)|30%|Medium||
|228|[Summary Ranges](./Algorithms/228.summary-ranges)|32%|Medium||
|229|[Majority Element II](./Algorithms/229.majority-element-ii)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|230|[Kth Smallest Element in a BST](./Algorithms/230.kth-smallest-element-in-a-bst)|45%|Medium||
|231|[Power of Two](./Algorithms/231.power-of-two)|40%|Easy||
|232|[Implement Queue using Stacks](./Algorithms/232.implement-queue-using-stacks)|38%|Easy||
|233|[Number of Digit One](./Algorithms/233.number-of-digit-one)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|234|[Palindrome Linked List](./Algorithms/234.palindrome-linked-list)|33%|Easy||
|238|[Product of Array Except Self](./Algorithms/238.product-of-array-except-self)|50%|Medium||
|239|[Sliding Window Maximum](./Algorithms/239.sliding-window-maximum)|34%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|240|[Search a 2D Matrix II](./Algorithms/240.search-a-2d-matrix-ii)|39%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|241|[Different Ways to Add Parentheses](./Algorithms/241.different-ways-to-add-parentheses)|46%|Medium||
|242|[Valid Anagram](./Algorithms/242.valid-anagram)|47%|Easy||
|257|[Binary Tree Paths](./Algorithms/257.binary-tree-paths)|41%|Easy||
|258|[Add Digits](./Algorithms/258.add-digits)|51%|Easy||
|260|[Single Number III](./Algorithms/260.single-number-iii)|53%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|263|[Ugly Number](./Algorithms/263.ugly-number)|39%|Easy||
|264|[Ugly Number II](./Algorithms/264.ugly-number-ii)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|268|[Missing Number](./Algorithms/268.missing-number)|45%|Easy||
|273|[Integer to English Words](./Algorithms/273.integer-to-english-words)|22%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|274|[H-Index](./Algorithms/274.h-index)|33%|Medium||
|275|[H-Index II](./Algorithms/275.h-index-ii)|34%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|279|[Perfect Squares](./Algorithms/279.perfect-squares)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|282|[Expression Add Operators](./Algorithms/282.expression-add-operators)|30%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|283|[Move Zeroes](./Algorithms/283.move-zeroes)|51%|Easy||
|287|[Find the Duplicate Number](./Algorithms/287.find-the-duplicate-number)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|289|[Game of Life](./Algorithms/289.game-of-life)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|290|[Word Pattern](./Algorithms/290.word-pattern)|33%|Easy||
|292|[Nim Game](./Algorithms/292.nim-game)|55%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|295|[Find Median from Data Stream](./Algorithms/295.find-median-from-data-stream)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|299|[Bulls and Cows](./Algorithms/299.bulls-and-cows)|35%|Medium||
|300|[Longest Increasing Subsequence](./Algorithms/300.longest-increasing-subsequence)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|301|[Remove Invalid Parentheses](./Algorithms/301.remove-invalid-parentheses)|35%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|303|[Range Sum Query - Immutable](./Algorithms/303.range-sum-query-immutable)|32%|Easy||
|304|[Range Sum Query 2D - Immutable](./Algorithms/304.range-sum-query-2d-immutable)|27%|Medium||
|306|[Additive Number](./Algorithms/306.additive-number)|27%|Medium||
|307|[Range Sum Query - Mutable](./Algorithms/307.range-sum-query-mutable)|22%|Medium||
|309|[Best Time to Buy and Sell Stock with Cooldown](./Algorithms/309.best-time-to-buy-and-sell-stock-with-cooldown)|41%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|310|[Minimum Height Trees](./Algorithms/310.minimum-height-trees)|28%|Medium||
|312|[Burst Balloons](./Algorithms/312.burst-balloons)|43%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|313|[Super Ugly Number](./Algorithms/313.super-ugly-number)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|315|[Count of Smaller Numbers After Self](./Algorithms/315.count-of-smaller-numbers-after-self)|34%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|316|[Remove Duplicate Letters](./Algorithms/316.remove-duplicate-letters)|30%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|318|[Maximum Product of Word Lengths](./Algorithms/318.maximum-product-of-word-lengths)|45%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|319|[Bulb Switcher](./Algorithms/319.bulb-switcher)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|321|[Create Maximum Number](./Algorithms/321.create-maximum-number)|24%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|322|[Coin Change](./Algorithms/322.coin-change)|26%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|324|[Wiggle Sort II](./Algorithms/324.wiggle-sort-ii)|26%|Medium||
|326|[Power of Three](./Algorithms/326.power-of-three)|40%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|327|[Count of Range Sum](./Algorithms/327.count-of-range-sum)|30%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|328|[Odd Even Linked List](./Algorithms/328.odd-even-linked-list)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|329|[Longest Increasing Path in a Matrix](./Algorithms/329.longest-increasing-path-in-a-matrix)|37%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|330|[Patching Array](./Algorithms/330.patching-array)|32%|Hard||
|331|[Verify Preorder Serialization of a Binary Tree](./Algorithms/331.verify-preorder-serialization-of-a-binary-tree)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|332|[Reconstruct Itinerary](./Algorithms/332.reconstruct-itinerary)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|334|[Increasing Triplet Subsequence](./Algorithms/334.increasing-triplet-subsequence)|39%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|335|[Self Crossing](./Algorithms/335.self-crossing)|26%|Hard||
|336|[Palindrome Pairs](./Algorithms/336.palindrome-pairs)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|337|[House Robber III](./Algorithms/337.house-robber-iii)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|338|[Counting Bits](./Algorithms/338.counting-bits)|62%|Medium||
|342|[Power of Four](./Algorithms/342.power-of-four)|39%|Easy||
|343|[Integer Break](./Algorithms/343.integer-break)|46%|Medium||
|344|[Reverse String](./Algorithms/344.reverse-string)|60%|Easy||
|345|[Reverse Vowels of a String](./Algorithms/345.reverse-vowels-of-a-string)|39%|Easy||
|347|[Top K Frequent Elements](./Algorithms/347.top-k-frequent-elements)|49%|Medium||
|349|[Intersection of Two Arrays](./Algorithms/349.intersection-of-two-arrays)|48%|Easy||
|350|[Intersection of Two Arrays II](./Algorithms/350.intersection-of-two-arrays-ii)|44%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|352|[Data Stream as Disjoint Intervals](./Algorithms/352.data-stream-as-disjoint-intervals)|41%|Hard||
|354|[Russian Doll Envelopes](./Algorithms/354.russian-doll-envelopes)|32%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|355|[Design Twitter](./Algorithms/355.design-twitter)|25%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|357|[Count Numbers with Unique Digits](./Algorithms/357.count-numbers-with-unique-digits)|46%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|363|[Max Sum of Rectangle No Larger Than K](./Algorithms/363.max-sum-of-rectangle-no-larger-than-k)|33%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|365|[Water and Jug Problem](./Algorithms/365.water-and-jug-problem)|28%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|367|[Valid Perfect Square](./Algorithms/367.valid-perfect-square)|38%|Easy||
|368|[Largest Divisible Subset](./Algorithms/368.largest-divisible-subset)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|371|[Sum of Two Integers](./Algorithms/371.sum-of-two-integers)|50%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|372|[Super Pow](./Algorithms/372.super-pow)|34%|Medium||
|373|[Find K Pairs with Smallest Sums](./Algorithms/373.find-k-pairs-with-smallest-sums)|31%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|375|[Guess Number Higher or Lower II](./Algorithms/375.guess-number-higher-or-lower-ii)|36%|Medium||
|376|[Wiggle Subsequence](./Algorithms/376.wiggle-subsequence)|36%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|377|[Combination Sum IV](./Algorithms/377.combination-sum-iv)|42%|Medium||
|378|[Kth Smallest Element in a Sorted Matrix](./Algorithms/378.kth-smallest-element-in-a-sorted-matrix)|45%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|380|[Insert Delete GetRandom O(1)](./Algorithms/380.insert-delete-getrandom-o1)|39%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|381|[Insert Delete GetRandom O(1) - Duplicates allowed](./Algorithms/381.insert-delete-getrandom-o1-duplicates-allowed)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|382|[Linked List Random Node](./Algorithms/382.linked-list-random-node)|47%|Medium||
|383|[Ransom Note](./Algorithms/383.ransom-note)|47%|Easy||
|384|[Shuffle an Array](./Algorithms/384.shuffle-an-array)|47%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|385|[Mini Parser](./Algorithms/385.mini-parser)|30%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|387|[First Unique Character in a String](./Algorithms/387.first-unique-character-in-a-string)|47%|Easy||
|388|[Longest Absolute File Path](./Algorithms/388.longest-absolute-file-path)|37%|Medium||
|389|[Find the Difference](./Algorithms/389.find-the-difference)|51%|Easy||
|390|[Elimination Game](./Algorithms/390.elimination-game)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|391|[Perfect Rectangle](./Algorithms/391.perfect-rectangle)|27%|Hard||
|392|[Is Subsequence](./Algorithms/392.is-subsequence)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|393|[UTF-8 Validation](./Algorithms/393.utf-8-validation)|34%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|394|[Decode String](./Algorithms/394.decode-string)|42%|Medium||
|395|[Longest Substring with At Least K Repeating Characters](./Algorithms/395.longest-substring-with-at-least-k-repeating-characters)|35%|Medium||
|396|[Rotate Function](./Algorithms/396.rotate-function)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|397|[Integer Replacement](./Algorithms/397.integer-replacement)|30%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|398|[Random Pick Index](./Algorithms/398.random-pick-index)|44%|Medium||
|399|[Evaluate Division](./Algorithms/399.evaluate-division)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|400|[Nth Digit](./Algorithms/400.nth-digit)|30%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|401|[Binary Watch](./Algorithms/401.binary-watch)|44%|Easy||
|402|[Remove K Digits](./Algorithms/402.remove-k-digits)|25%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|403|[Frog Jump](./Algorithms/403.frog-jump)|32%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|404|[Sum of Left Leaves](./Algorithms/404.sum-of-left-leaves)|47%|Easy||
|405|[Convert a Number to Hexadecimal](./Algorithms/405.convert-a-number-to-hexadecimal)|41%|Easy||
|406|[Queue Reconstruction by Height](./Algorithms/406.queue-reconstruction-by-height)|56%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|407|[Trapping Rain Water II](./Algorithms/407.trapping-rain-water-ii)|37%|Hard||
|409|[Longest Palindrome](./Algorithms/409.longest-palindrome)|45%|Easy||
|410|[Split Array Largest Sum](./Algorithms/410.split-array-largest-sum)|39%|Hard||
|412|[Fizz Buzz](./Algorithms/412.Fizz-Buzz)|58%|Easy||
|413|[Arithmetic Slices](./Algorithms/413.arithmetic-slices)|54%|Medium||
|414|[Third Maximum Number](./Algorithms/414.third-maximum-number)|28%|Easy||
|415|[Add Strings](./Algorithms/415.add-strings)|41%|Easy||
|416|[Partition Equal Subset Sum](./Algorithms/416.partition-equal-subset-sum)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|417|[Pacific Atlantic Water Flow](./Algorithms/417.pacific-atlantic-water-flow)|34%|Medium||
|419|[Battleships in a Board](./Algorithms/419.battleships-in-a-board)|63%|Medium||
|420|[Strong Password Checker](./Algorithms/420.strong-password-checker)|19%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|421|[Maximum XOR of Two Numbers in an Array](./Algorithms/421.maximum-xor-of-two-numbers-in-an-array)|47%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|423|[Reconstruct Original Digits from English](./Algorithms/423.reconstruct-original-digits-from-english)|44%|Medium||
|424|[Longest Repeating Character Replacement](./Algorithms/424.longest-repeating-character-replacement)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|432|[All O`one Data Structure](./Algorithms/432.all-oone-data-structure)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|434|[Number of Segments in a String](./Algorithms/434.number-of-segments-in-a-string)|36%|Easy||
|435|[Non-overlapping Intervals](./Algorithms/435.non-overlapping-intervals)|41%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|436|[Find Right Interval](./Algorithms/436.find-right-interval)|41%|Medium||
|437|[Path Sum III](./Algorithms/437.path-sum-iii)|40%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|438|[Find All Anagrams in a String](./Algorithms/438.find-all-anagrams-in-a-string)|33%|Easy||
|440|[K-th Smallest in Lexicographical Order](./Algorithms/440.k-th-smallest-in-lexicographical-order)|25%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|441|[Arranging Coins](./Algorithms/441.arranging-coins)|36%|Easy||
|442|[Find All Duplicates in an Array](./Algorithms/442.find-all-duplicates-in-an-array)|56%|Medium||
|443|[String Compression](./Algorithms/443.string-compression)|35%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|445|[Add Two Numbers II](./Algorithms/445.add-two-numbers-ii)|46%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|446|[Arithmetic Slices II - Subsequence](./Algorithms/446.arithmetic-slices-ii-subsequence)|28%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|447|[Number of Boomerangs](./Algorithms/447.number-of-boomerangs)|46%|Easy||
|448|[Find All Numbers Disappeared in an Array](./Algorithms/448.find-all-numbers-disappeared-in-an-array)|51%|Easy||
|450|[Delete Node in a BST](./Algorithms/450.delete-node-in-a-bst)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|451|[Sort Characters By Frequency](./Algorithms/451.sort-characters-by-frequency)|51%|Medium||
|452|[Minimum Number of Arrows to Burst Balloons](./Algorithms/452.minimum-number-of-arrows-to-burst-balloons)|44%|Medium||
|453|[Minimum Moves to Equal Array Elements](./Algorithms/453.minimum-moves-to-equal-array-elements)|48%|Easy||
|454|[4Sum II](./Algorithms/454.4sum-ii)|47%|Medium||
|455|[Assign Cookies](./Algorithms/455.assign-cookies)|47%|Easy||
|456|[132 Pattern](./Algorithms/456.132-pattern)|27%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|459|[Repeated Substring Pattern](./Algorithms/459.repeated-substring-pattern)|38%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|460|[LFU Cache](./Algorithms/460.lfu-cache)|25%|Hard||
|461|[Hamming Distance](./Algorithms/461.hamming-distance)|69%|Easy||
|462|[Minimum Moves to Equal Array Elements II](./Algorithms/462.minimum-moves-to-equal-array-elements-ii)|51%|Medium||
|463|[Island Perimeter](./Algorithms/463.island-perimeter)|57%|Easy||
|464|[Can I Win](./Algorithms/464.can-i-win)|25%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|466|[Count The Repetitions](./Algorithms/466.count-the-repetitions)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|467|[Unique Substrings in Wraparound String](./Algorithms/467.unique-substrings-in-wraparound-string)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|468|[Validate IP Address](./Algorithms/468.validate-ip-address)|20%|Medium||
|472|[Concatenated Words](./Algorithms/472.concatenated-words)|30%|Hard||
|473|[Matchsticks to Square](./Algorithms/473.matchsticks-to-square)|35%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|474|[Ones and Zeroes](./Algorithms/474.ones-and-zeroes)|38%|Medium||
|475|[Heaters](./Algorithms/475.heaters)|29%|Easy||
|476|[Number Complement](./Algorithms/476.number-complement)|61%|Easy||
|477|[Total Hamming Distance](./Algorithms/477.total-hamming-distance)|47%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|479|[Largest Palindrome Product](./Algorithms/479.largest-palindrome-product)|25%|Easy||
|480|[Sliding Window Median](./Algorithms/480.sliding-window-median)|30%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|481|[Magical String](./Algorithms/481.magical-string)|45%|Medium||
|482|[License Key Formatting](./Algorithms/482.license-key-formatting)|39%|Easy||
|483|[Smallest Good Base](./Algorithms/483.smallest-good-base)|33%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|485|[Max Consecutive Ones](./Algorithms/485.max-consecutive-ones)|53%|Easy||
|486|[Predict the Winner](./Algorithms/486.predict-the-winner)|45%|Medium||
|488|[Zuma Game](./Algorithms/488.zuma-game)|37%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|491|[Increasing Subsequences](./Algorithms/491.increasing-subsequences)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|492|[Construct the Rectangle](./Algorithms/492.construct-the-rectangle)|48%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|493|[Reverse Pairs](./Algorithms/493.reverse-pairs)|20%|Hard||
|494|[Target Sum](./Algorithms/494.target-sum)|43%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|495|[Teemo Attacking](./Algorithms/495.teemo-attacking)|51%|Medium||
|496|[Next Greater Element I](./Algorithms/496.next-greater-element-i)|56%|Easy||
|498|[Diagonal Traverse](./Algorithms/498.diagonal-traverse)|46%|Medium||
|500|[Keyboard Row](./Algorithms/500.keyboard-row)|59%|Easy||
|501|[Find Mode in Binary Search Tree](./Algorithms/501.find-mode-in-binary-search-tree)|37%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|502|[IPO](./Algorithms/502.ipo)|36%|Hard||
|503|[Next Greater Element II](./Algorithms/503.next-greater-element-ii)|48%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|504|[Base 7](./Algorithms/504.base-7)|43%|Easy||
|506|[Relative Ranks](./Algorithms/506.relative-ranks)|46%|Easy||
|507|[Perfect Number](./Algorithms/507.perfect-number)|32%|Easy||
|508|[Most Frequent Subtree Sum](./Algorithms/508.most-frequent-subtree-sum)|52%|Medium||
|513|[Find Bottom Left Tree Value](./Algorithms/513.find-bottom-left-tree-value)|56%|Medium||
|514|[Freedom Trail](./Algorithms/514.freedom-trail)|39%|Hard||
|515|[Find Largest Value in Each Tree Row](./Algorithms/515.find-largest-value-in-each-tree-row)|55%|Medium||
|516|[Longest Palindromic Subsequence](./Algorithms/516.longest-palindromic-subsequence)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|517|[Super Washing Machines](./Algorithms/517.super-washing-machines)|36%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|520|[Detect Capital](./Algorithms/520.detect-capital)|51%|Easy||
|521|[Longest Uncommon Subsequence I](./Algorithms/521.longest-uncommon-subsequence-i)|55%|Easy||
|522|[Longest Uncommon Subsequence II](./Algorithms/522.longest-uncommon-subsequence-ii)|31%|Medium||
|523|[Continuous Subarray Sum](./Algorithms/523.continuous-subarray-sum)|23%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|524|[Longest Word in Dictionary through Deleting](./Algorithms/524.longest-word-in-dictionary-through-deleting)|43%|Medium||
|525|[Contiguous Array](./Algorithms/525.contiguous-array)|41%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|526|[Beautiful Arrangement](./Algorithms/526.beautiful-arrangement)|53%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|529|[Minesweeper](./Algorithms/529.minesweeper)|49%|Medium||
|530|[Minimum Absolute Difference in BST](./Algorithms/530.minimum-absolute-difference-in-bst)|47%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|532|[K-diff Pairs in an Array](./Algorithms/532.k-diff-pairs-in-an-array)|28%|Easy||
|537|[Complex Number Multiplication](./Algorithms/537.complex-number-multiplication)|63%|Medium||
|538|[Convert BST to Greater Tree](./Algorithms/538.convert-bst-to-greater-tree)|48%|Easy||
|539|[Minimum Time Difference](./Algorithms/539.minimum-time-difference)|46%|Medium||
|540|[Single Element in a Sorted Array](./Algorithms/540.single-element-in-a-sorted-array)|55%|Medium||
|541|[Reverse String II](./Algorithms/541.reverse-string-ii)|43%|Easy||
|542|[01 Matrix](./Algorithms/542.01-matrix)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|543|[Diameter of Binary Tree](./Algorithms/543.diameter-of-binary-tree)|44%|Easy||
|546|[Remove Boxes](./Algorithms/546.remove-boxes)|35%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|547|[Friend Circles](./Algorithms/547.friend-circles)|49%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|551|[Student Attendance Record I](./Algorithms/551.student-attendance-record-i)|44%|Easy||
|552|[Student Attendance Record II](./Algorithms/552.student-attendance-record-ii)|31%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|553|[Optimal Division](./Algorithms/553.optimal-division)|55%|Medium||
|554|[Brick Wall](./Algorithms/554.brick-wall)|46%|Medium||
|556|[Next Greater Element III](./Algorithms/556.next-greater-element-iii)|28%|Medium||
|557|[Reverse Words in a String III](./Algorithms/557.reverse-words-in-a-string-iii)|60%|Easy||
|560|[Subarray Sum Equals K](./Algorithms/560.subarray-sum-equals-k)|39%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|561|[Array Partition I](./Algorithms/561.array-partition-i)|66%|Easy||
|563|[Binary Tree Tilt](./Algorithms/563.binary-tree-tilt)|47%|Easy||
|564|[Find the Closest Palindrome](./Algorithms/564.find-the-closest-palindrome)|17%|Hard||
|565|[Array Nesting](./Algorithms/565.array-nesting)|49%|Medium||
|566|[Reshape the Matrix](./Algorithms/566.reshape-the-matrix)|57%|Easy||
|567|[Permutation in String](./Algorithms/567.permutation-in-string)|36%|Medium||
|572|[Subtree of Another Tree](./Algorithms/572.subtree-of-another-tree)|40%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|575|[Distribute Candies](./Algorithms/575.distribute-candies)|57%|Easy||
|576|[Out of Boundary Paths](./Algorithms/576.out-of-boundary-paths)|30%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|581|[Shortest Unsorted Continuous Subarray](./Algorithms/581.shortest-unsorted-continuous-subarray)|29%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|583|[Delete Operation for Two Strings](./Algorithms/583.delete-operation-for-two-strings)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|587|[Erect the Fence](./Algorithms/587.erect-the-fence)|33%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|591|[Tag Validator](./Algorithms/591.tag-validator)|31%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|592|[Fraction Addition and Subtraction](./Algorithms/592.fraction-addition-and-subtraction)|46%|Medium||
|593|[Valid Square](./Algorithms/593.valid-square)|39%|Medium||
|594|[Longest Harmonious Subsequence](./Algorithms/594.longest-harmonious-subsequence)|40%|Easy||
|598|[Range Addition II](./Algorithms/598.range-addition-ii)|48%|Easy||
|599|[Minimum Index Sum of Two Lists](./Algorithms/599.minimum-index-sum-of-two-lists)|46%|Easy||
|600|[Non-negative Integers without Consecutive Ones](./Algorithms/600.non-negative-integers-without-consecutive-ones)|31%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|605|[Can Place Flowers](./Algorithms/605.can-place-flowers)|30%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|606|[Construct String from Binary Tree](./Algorithms/606.construct-string-from-binary-tree)|49%|Easy||
|609|[Find Duplicate File in System](./Algorithms/609.find-duplicate-file-in-system)|52%|Medium||
|611|[Valid Triangle Number](./Algorithms/611.valid-triangle-number)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|617|[Merge Two Binary Trees](./Algorithms/617.merge-two-binary-trees)|67%|Easy||
|621|[Task Scheduler](./Algorithms/621.task-scheduler)|42%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|623|[Add One Row to Tree](./Algorithms/623.add-one-row-to-tree)|46%|Medium||
|628|[Maximum Product of Three Numbers](./Algorithms/628.maximum-product-of-three-numbers)|44%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|629|[K Inverse Pairs Array](./Algorithms/629.k-inverse-pairs-array)|27%|Hard||
|630|[Course Schedule III](./Algorithms/630.course-schedule-iii)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|632|[Smallest Range](./Algorithms/632.smallest-range)|41%|Hard||
|633|[Sum of Square Numbers](./Algorithms/633.sum-of-square-numbers)|32%|Easy||
|636|[Exclusive Time of Functions](./Algorithms/636.exclusive-time-of-functions)|44%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|637|[Average of Levels in Binary Tree](./Algorithms/637.average-of-levels-in-binary-tree)|55%|Easy||
|638|[Shopping Offers](./Algorithms/638.shopping-offers)|45%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|639|[Decode Ways II](./Algorithms/639.decode-ways-ii)|24%|Hard||
|640|[Solve the Equation](./Algorithms/640.solve-the-equation)|38%|Medium||
|643|[Maximum Average Subarray I](./Algorithms/643.maximum-average-subarray-i)|37%|Easy||
|645|[Set Mismatch](./Algorithms/645.set-mismatch)|39%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|646|[Maximum Length of Pair Chain](./Algorithms/646.maximum-length-of-pair-chain)|47%|Medium||
|647|[Palindromic Substrings](./Algorithms/647.palindromic-substrings)|54%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|648|[Replace Words](./Algorithms/648.replace-words)|47%|Medium||
|649|[Dota2 Senate](./Algorithms/649.dota2-senate)|36%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|650|[2 Keys Keyboard](./Algorithms/650.2-keys-keyboard)|44%|Medium||
|652|[Find Duplicate Subtrees](./Algorithms/652.find-duplicate-subtrees)|36%|Medium||
|653|[Two Sum IV - Input is a BST](./Algorithms/653.two-sum-iv-input-is-a-bst)|49%|Easy||
|654|[Maximum Binary Tree](./Algorithms/654.maximum-binary-tree)|69%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|655|[Print Binary Tree](./Algorithms/655.print-binary-tree)|49%|Medium||
|657|[Judge Route Circle](./Algorithms/657.judge-route-circle)|68%|Easy||
|658|[Find K Closest Elements](./Algorithms/658.find-k-closest-elements)|34%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|659|[Split Array into Consecutive Subsequences](./Algorithms/659.split-array-into-consecutive-subsequences)|36%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|661|[Image Smoother](./Algorithms/661.image-smoother)|46%|Easy||
|662|[Maximum Width of Binary Tree](./Algorithms/662.maximum-width-of-binary-tree)|37%|Medium||
|664|[Strange Printer](./Algorithms/664.strange-printer)|34%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|665|[Non-decreasing Array](./Algorithms/665.non-decreasing-array)|20%|Easy||
|667|[Beautiful Arrangement II](./Algorithms/667.beautiful-arrangement-ii)|51%|Medium||
|668|[Kth Smallest Number in Multiplication Table](./Algorithms/668.kth-smallest-number-in-multiplication-table)|40%|Hard||
|669|[Trim a Binary Search Tree](./Algorithms/669.trim-a-binary-search-tree)|58%|Easy||
|670|[Maximum Swap](./Algorithms/670.maximum-swap)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|671|[Second Minimum Node In a Binary Tree](./Algorithms/671.second-minimum-node-in-a-binary-tree)|41%|Easy||
|672|[Bulb Switcher II](./Algorithms/672.bulb-switcher-ii)|49%|Medium||
|673|[Number of Longest Increasing Subsequence](./Algorithms/673.number-of-longest-increasing-subsequence)|31%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|674|[Longest Continuous Increasing Subsequence](./Algorithms/674.longest-continuous-increasing-subsequence)|42%|Easy||
|675|[Cut Off Trees for Golf Event](./Algorithms/675.cut-off-trees-for-golf-event)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|676|[Implement Magic Dictionary](./Algorithms/676.implement-magic-dictionary)|49%|Medium||
|677|[Map Sum Pairs](./Algorithms/677.map-sum-pairs)|51%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|678|[Valid Parenthesis String](./Algorithms/678.valid-parenthesis-string)|29%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|679|[24 Game](./Algorithms/679.24-game)|38%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|680|[Valid Palindrome II](./Algorithms/680.valid-palindrome-ii)|32%|Easy||
|682|[Baseball Game](./Algorithms/682.baseball-game)|58%|Easy||
|684|[Redundant Connection](./Algorithms/684.redundant-connection)|43%|Medium||
|685|[Redundant Connection II](./Algorithms/685.redundant-connection-ii)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|686|[Repeated String Match](./Algorithms/686.repeated-string-match)|32%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|687|[Longest Univalue Path](./Algorithms/687.longest-univalue-path)|32%|Easy||
|688|[Knight Probability in Chessboard](./Algorithms/688.knight-probability-in-chessboard)|39%|Medium||
|689|[Maximum Sum of 3 Non-Overlapping Subarrays](./Algorithms/689.maximum-sum-of-3-non-overlapping-subarrays)|41%|Hard||
|691|[Stickers to Spell Word](./Algorithms/691.stickers-to-spell-word)|34%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|692|[Top K Frequent Words](./Algorithms/692.top-k-frequent-words)|41%|Medium||
|693|[Binary Number with Alternating Bits](./Algorithms/693.binary-number-with-alternating-bits)|55%|Easy||
|695|[Max Area of Island](./Algorithms/695.max-area-of-island)|51%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|696|[Count Binary Substrings](./Algorithms/696.count-binary-substrings)|50%|Easy||
|697|[Degree of an Array](./Algorithms/697.degree-of-an-array)|46%|Easy||
|698|[Partition to K Equal Sum Subsets](./Algorithms/698.partition-to-k-equal-sum-subsets)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|699|[Falling Squares](./Algorithms/699.falling-squares)|37%|Hard||
|712|[Minimum ASCII Delete Sum for Two Strings](./Algorithms/712.minimum-ascii-delete-sum-for-two-strings)|51%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|713|[Subarray Product Less Than K](./Algorithms/713.subarray-product-less-than-k)|33%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|714|[Best Time to Buy and Sell Stock with Transaction Fee](./Algorithms/714.best-time-to-buy-and-sell-stock-with-transaction-fee)|45%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|715|[Range Module](./Algorithms/715.range-module)|31%|Hard||
|717|[1-bit and 2-bit Characters](./Algorithms/717.1-bit-and-2-bit-characters)|49%|Easy||
|718|[Maximum Length of Repeated Subarray](./Algorithms/718.maximum-length-of-repeated-subarray)|41%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|719|[Find K-th Smallest Pair Distance](./Algorithms/719.find-k-th-smallest-pair-distance)|27%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|720|[Longest Word in Dictionary](./Algorithms/720.longest-word-in-dictionary)|41%|Easy||
|721|[Accounts Merge](./Algorithms/721.accounts-merge)|32%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|722|[Remove Comments](./Algorithms/722.remove-comments)|27%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|724|[Find Pivot Index](./Algorithms/724.find-pivot-index)|39%|Easy||
|725|[Split Linked List in Parts](./Algorithms/725.split-linked-list-in-parts)|47%|Medium||
|726|[Number of Atoms](./Algorithms/726.number-of-atoms)|43%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|728|[Self Dividing Numbers](./Algorithms/728.self-dividing-numbers)|66%|Easy||
|729|[My Calendar I](./Algorithms/729.my-calendar-i)|42%|Medium||
|730|[Count Different Palindromic Subsequences](./Algorithms/730.count-different-palindromic-subsequences)|34%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|731|[My Calendar II](./Algorithms/731.my-calendar-ii)|37%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|732|[My Calendar III](./Algorithms/732.my-calendar-iii)|50%|Hard||
|733|[Flood Fill](./Algorithms/733.flood-fill)|47%|Easy||
|735|[Asteroid Collision](./Algorithms/735.asteroid-collision)|37%|Medium||
|736|[Parse Lisp Expression](./Algorithms/736.parse-lisp-expression)|42%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|738|[Monotone Increasing Digits](./Algorithms/738.monotone-increasing-digits)|40%|Medium||
|739|[Daily Temperatures](./Algorithms/739.daily-temperatures)|52%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|740|[Delete and Earn](./Algorithms/740.delete-and-earn)|43%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|741|[Cherry Pickup](./Algorithms/741.cherry-pickup)|23%|Hard||
|743|[Network Delay Time](./Algorithms/743.network-delay-time)|34%|Medium||
|744|[Find Smallest Letter Greater Than Target](./Algorithms/744.find-smallest-letter-greater-than-target)|44%|Easy||
|745|[Prefix and Suffix Search](./Algorithms/745.prefix-and-suffix-search)|25%|Hard||
|746|[Min Cost Climbing Stairs](./Algorithms/746.min-cost-climbing-stairs)|43%|Easy||
|747|[Largest Number At Least Twice of Others](./Algorithms/747.largest-number-at-least-twice-of-others)|41%|Easy||
|748|[Shortest Completing Word](./Algorithms/748.shortest-completing-word)|51%|Medium||
|749|[Contain Virus](./Algorithms/749.contain-virus)|39%|Hard||
|752|[Open the Lock](./Algorithms/752.open-the-lock)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|753|[Cracking the Safe](./Algorithms/753.cracking-the-safe)|39%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|754|[Reach a Number](./Algorithms/754.reach-a-number)|26%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|756|[Pyramid Transition Matrix](./Algorithms/756.pyramid-transition-matrix)|45%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|757|[Set Intersection Size At Least Two](./Algorithms/757.set-intersection-size-at-least-two)|34%|Hard||
|761|[Special Binary String](./Algorithms/761.special-binary-string)|41%|Hard||
|762|[Prime Number of Set Bits in Binary Representation](./Algorithms/762.prime-number-of-set-bits-in-binary-representation)|55%|Easy||
|763|[Partition Labels](./Algorithms/763.partition-labels)|64%|Medium||
|764|[Largest Plus Sign](./Algorithms/764.largest-plus-sign)|37%|Medium||
|765|[Couples Holding Hands](./Algorithms/765.couples-holding-hands)|48%|Hard||
|766|[Toeplitz Matrix](./Algorithms/766.toeplitz-matrix)|57%|Easy||
|767|[Reorganize String](./Algorithms/767.reorganize-string)|35%|Medium||
|768|[Max Chunks To Make Sorted II](./Algorithms/768.max-chunks-to-make-sorted-ii)|42%|Hard||
|769|[Max Chunks To Make Sorted](./Algorithms/769.max-chunks-to-make-sorted)|47%|Medium||
|770|[Basic Calculator IV](./Algorithms/770.basic-calculator-iv)|43%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|771|[Jewels and Stones](./Algorithms/771.jewels-and-stones)|82%|Easy||
|773|[Sliding Puzzle](./Algorithms/773.sliding-puzzle)|47%|Hard||
|775|[Global and Local Inversions](./Algorithms/775.global-and-local-inversions)|31%|Medium||
|777|[Swap Adjacent in LR String](./Algorithms/777.swap-adjacent-in-lr-string)|27%|Medium||
|778|[Swim in Rising Water](./Algorithms/778.swim-in-rising-water)|44%|Hard||
|779|[K-th Symbol in Grammar](./Algorithms/779.k-th-symbol-in-grammar)|35%|Medium||
|780|[Reaching Points](./Algorithms/780.reaching-points)|21%|Hard||
|781|[Rabbits in Forest](./Algorithms/781.rabbits-in-forest)|49%|Medium||
|782|[Transform to Chessboard](./Algorithms/782.transform-to-chessboard)|35%|Hard||
|783|[Minimum Distance Between BST Nodes](./Algorithms/783.minimum-distance-between-bst-nodes)|47%|Easy||
|784|[Letter Case Permutation](./Algorithms/784.letter-case-permutation)|52%|Easy|[❤](https://leetcode.com/list/oussv5j)|
|785|[Is Graph Bipartite?](./Algorithms/785.is-graph-bipartite)|38%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|786|[K-th Smallest Prime Fraction](./Algorithms/786.k-th-smallest-prime-fraction)|29%|Hard|[❤](https://leetcode.com/list/oussv5j)|
|787|[Cheapest Flights Within K Stops](./Algorithms/787.cheapest-flights-within-k-stops)|30%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|788|[Rotated Digits](./Algorithms/788.rotated-digits)|50%|Easy||
|789|[Escape The Ghosts](./Algorithms/789.escape-the-ghosts)|47%|Medium||
|790|[Domino and Tromino Tiling](./Algorithms/790.domino-and-tromino-tiling)|32%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|791|[Custom Sort String](./Algorithms/791.custom-sort-string)|60%|Medium||
|792|[Number of Matching Subsequences](./Algorithms/792.number-of-matching-subsequences)|35%|Medium|[❤](https://leetcode.com/list/oussv5j)|
|793| * Preimage Size of Factorial Zeroes Function|45%|Hard||
|794| * Valid Tic-Tac-Toe State|27%|Medium||
|795| * Number of Subarrays with Bounded Maximum|40%|Medium||
|796| * Rotate String|54%|Easy||
|797| * All Paths From Source to Target|69%|Medium||
|798| * Smallest Rotation with Highest Score|31%|Hard||
|799| * Champagne Tower|28%|Medium||
|801| * Minimum Swaps To Make Sequences Increasing|23%|Medium||
|802| * Find Eventual Safe States|35%|Medium||
|803| * Bricks Falling When Hit|20%|Hard||
|804| * Unique Morse Code Words|76%|Easy||
|805| * Split Array With Same Average|19%|Hard||
|806| * Number of Lines To Write String|64%|Easy||
|807| * Max Increase to Keep City Skyline|82%|Medium||
|808| * Soup Servings|30%|Medium||
|809| * Expressive Words|34%|Medium||
|810| * Chalkboard XOR Game|35%|Hard||
|811| * Subdomain Visit Count|65%|Easy||
|812| * Largest Triangle Area :new: |51%|Easy||
|813| * Largest Sum of Averages :new: |37%|Medium||
|814| * Binary Tree Pruning :new: |73%|Medium||
|815| * Bus Routes :new: |30%|Hard||


------------------------------------------------------------------

以下免费的算法题，暂时不能使用 Go 解答

- [116.Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)
- [117.Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/)
- [133.Clone Graph](https://leetcode.com/problems/clone-graph/)
- [138.Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)
- [141.Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)
- [142.Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)
- [151.Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)
- [160.Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)
- [173.Binary Search Tree Iterator](https://leetcode.com/problems/binary-search-tree-iterator/)
- [190.Reverse Bits](https://leetcode.com/problems/reverse-bits/)
- [191.Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)
- [222.Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/)
- [235.Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
- [236.Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
- [237.Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/)
- [278.First Bad Version](https://leetcode.com/problems/first-bad-version/)
- [284.Peeking Iterator](https://leetcode.com/problems/peeking-iterator/)
- [297.Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)
- [341.Flatten Nested List Iterator](https://leetcode.com/problems/flatten-nested-list-iterator/)
- [374.Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/)
- [386.Lexicographical Numbers](https://leetcode.com/problems/lexicographical-numbers/)
- [449.Serialize and Deserialize BST](https://leetcode.com/problems/serialize-and-deserialize-bst/)
- [535.Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/)
- [690.Employee Importance](https://leetcode.com/problems/employee-importance/)

------------------------------------------------------------------

## 二.分类

## Array
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/)| [Swift](./Array/MaxConsecutiveOnes.swift)| Easy| O(n)| O(1)|
[Heaters](https://leetcode.com/problems/heaters/)| [Swift](./Array/Heaters.swift)| Easy| O(nlogn)| O(1)|
[Number of Boomerangs](https://leetcode.com/problems/number-of-boomerangs/)| [Swift](./Array/NumberBoomerangs.swift)| Easy| O(n ^ 2)| O(n)|
[Island Perimeter](https://leetcode.com/problems/island-perimeter/)| [Swift](./Array/IslandPerimeter.swift)| Easy| O(nm)| O(1)|
[Majority Element](https://leetcode.com/problems/majority-element/)| [Swift](./Array/MajorityElement.swift)| Easy| O(n)| O(1)|
[Majority Element II](https://leetcode.com/problems/majority-element-ii/)| [Swift](./Array/MajorityElementII.swift)| Medium| O(n)| O(1)|
[Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)| [Swift](./Array/IntersectionTwoArrays.swift)| Easy| O(n)| O(n)|
[Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/)| [Swift](./Array/IntersectionTwoArraysII.swift)| Easy| O(n)| O(n)|
[Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)| [Swift](./Array/ContainsDuplicate.swift)| Easy| O(n)| O(n)|
[Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/)| [Swift](./Array/ContainsDuplicateII.swift)| Easy| O(n)| O(n)|
[Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)| [Swift](./Array/RemoveDuplicatesFromSortedArray.swift)| Easy| O(n)| O(1)|
[Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/)| [Swift](./Array/RemoveDuplicatesFromSortedArrayII.swift)| Medium| O(n)| O(1)|
[Move Zeroes](https://leetcode.com/problems/move-zeroes/)| [Swift](./Array/MoveZeroes.swift)| Easy| O(n)| O(1)|
[Remove Element](https://leetcode.com/problems/remove-element/)| [Swift](./Array/RemoveElement.swift)| Easy| O(n)| O(1)|
[Two Sum](https://leetcode.com/problems/two-sum/)| [Swift](./Array/TwoSum.swift)| Easy| O(n)| O(n)|
[3Sum](https://leetcode.com/problems/3sum/)| [Swift](./Array/ThreeSum.swift)| Medium| O(n^2)| O(nC3)|
[3Sum Closest](https://leetcode.com/problems/3sum-closest/)| [Swift](./Array/ThreeSum.swift)| Medium| O(n^2)| O(nC3)|
[4Sum](https://leetcode.com/problems/4sum/)| [Swift](./Array/FourSum.swift)| Medium| O(n^3)| O(nC4)|
[Summary Ranges](https://leetcode.com/problems/summary-ranges/)| [Swift](./Array/SummaryRanges.swift)| Medium| O(n)| O(n)|
[Shortest Word Distance](https://leetcode.com/problems/shortest-word-distance/)| [Swift](./Array/ShortestWordDistance.swift)| Easy| O(n)| O(1)|
[Shortest Word Distance III](https://leetcode.com/problems/shortest-word-distance-iii/)| [Swift](./Array/ShortestWordDistanceIII.swift)| Medium| O(n)| O(1)|
[Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)| [Swift](./Array/MinimumSizeSubarraySum.swift)| Medium| O(n)| O(1)|
[Maximum Size Subarray Sum Equals k](https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/)| [Swift](./Array/MaximumSizeSubarraySumEqualsK.swift)| Medium| O(n)| O(n)|
[Smallest Range](https://leetcode.com/problems/smallest-range/)| [Swift](./Array/SmallestRange.swift)| Hard | O(nm)| O(nm)|
[Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)| [Swift](./Array/ProductExceptSelf.swift)| Medium| O(n)| O(n)|
[Rotate Array](https://leetcode.com/problems/rotate-array/)| [Swift](./Array/RotateArray.swift)| Easy| O(n)| O(1)|
[Rotate Image](https://leetcode.com/problems/rotate-image/)| [Swift](./Array/RotateImage.swift)| Medium| O(n^2)| O(1)|
[Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)| [Swift](./Array/SpiralMatrix.swift)| Medium| O(n^2)| O(1)|
[Spiral Matrix II](https://leetcode.com/problems/spiral-matrix/)| [Swift](./Array/SpiralMatrixII.swift)| Medium| O(n^2)| O(1)|
[Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)| [Swift](./Array/ValidSudoku.swift)| Easy| O(n^2)| O(n)|
[Set Matrix Zero](https://leetcode.com/problems/set-matrix-zeroes/)| [Swift](./Array/SetMatrixZero.swift)| Medium| O(n^2)| O(1)|
[Next Permutation](https://leetcode.com/problems/next-permutation/)| [Swift](./Array/NextPermutation.swift)| Medium| O(n)| O(1)|
[Gas Station](https://leetcode.com/problems/gas-station/)| [Swift](./Array/GasStation.swift)| Medium| O(n)| O(1)|
[Game of Life](https://leetcode.com/problems/game-of-life/)| [Swift](./Array/GameLife.swift)| Medium| O(n)| O(1)|
[Task Scheduler](https://leetcode.com/problems/task-scheduler/)| [Swift](./Array/TaskScheduler.swift)| Medium| O(nlogn)| O(n)|
[Sliding Window Maximum ](https://leetcode.com/problems/sliding-window-maximum/)| [Swift](./Array/SlidingWindowMaximum.swift)| Hard| O(n)| O(n)|
[Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)| [Swift](./Array/LongestConsecutiveSequence.swift)| Hard| O(n)| O(n)|


## String
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Fizz Buzz](https://leetcode.com/problems/fizz-buzz/)| [Swift](./String/FizzBuzz.swift)| Easy| O(n)| O(1)|
[First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/)| [Swift](./String/FirstUniqueCharacterInString.swift)| Easy| O(n)| O(1)|
[Keyboard Row](https://leetcode.com/problems/keyboard-row/)| [Swift](./String/KeyboardRow.swift)| Easy| O(nm)| O(n)|
[Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)| [Swift](./String/ValidPalindrome.swift)| Easy| O(n)| O(n)|
[Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/)| [Swift](./String/ValidPalindromeII.swift)| Easy| O(n)| O(n)|
[Detect Capital](https://leetcode.com/problems/detect-capital/)| [Swift](./String/DetectCapital.swift)| Easy| O(n)| O(1)|
[Count and Say](https://leetcode.com/problems/count-and-say/)| [Swift](./String/CountAndSay.swift)| Easy| O(n^2)| O(n)|
[Flip Game](https://leetcode.com/problems/flip-game/)| [Swift](./String/FlipGame.swift)| Easy| O(n)| O(n)|
[Implement strStr()](https://leetcode.com/problems/implement-strstr/)| [Swift](./String/StrStr.swift)| Easy| O(nm)| O(n)|
[Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/)| [Swift](./String/IsomorphicStrings.swift)| Easy| O(n)| O(n)|
[Reverse String](https://leetcode.com/problems/reverse-string/)| [Swift](./String/ReverseString.swift)| Easy| O(n)| O(n)|
[Reverse String II](https://leetcode.com/problems/reverse-string-ii/)| [Swift](./String/ReverseStringII.swift)| Easy| O(n)| O(n)|
[Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/)| [Swift](./String/ReverseVowelsOfAString.swift)| Easy| O(n)| O(n)|
[Length of Last Word](https://leetcode.com/problems/length-of-last-word/)| [Swift](./String/AddStrings.swift)| Easy| O(n)| O(n)|
[Add Strings](https://leetcode.com/problems/add-strings/)| [Swift](./String/LengthLastWord.swift)| Easy| O(n)| O(1)|
[Multiply Strings](https://leetcode.com/problems/multiply-strings/)| [Swift](./String/MultiplyStrings.swift)| Medium| O(n)| O(1)|
[Palindrome Permutation](https://leetcode.com/problems/palindrome-permutation/)| [Swift](./String/PalindromePermutation.swift)| Easy| O(n)| O(n)|
[Valid Anagram](https://leetcode.com/problems/valid-anagram/)| [Swift](./String/ValidAnagram.swift)| Easy| O(nlogn)| O(1)|
[Ransom Note](https://leetcode.com/problems/ransom-note/)| [Swift](./String/RansomNote.swift)| Easy| O(n)| O(n)|
[Group Anagrams](https://leetcode.com/problems/anagrams/)| [Swift](./String/GroupAnagrams.swift)| Medium| O(nmlogm + nlogn)| O(n)
[Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)| [Swift](./String/LongestCommonPrefix.swift)| Easy| O(nm)| O(m)|
[Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)| [Swift](./String/LongestSubstringWithoutRepeatingCharacters.swift)| Medium| O(n)| O(n)|
[One Edit Distance](https://leetcode.com/problems/one-edit-distance/)| [Swift](./String/OneEditDistance.swift)| Medium| O(n)| O(n)|
[Word Pattern](https://leetcode.com/problems/word-pattern/)| [Swift](./String/WordPattern.swift)| Easy| O(n)| O(n)|
[Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)| [Swift](./Array/MinimumWindowSubstring.swift)| Hard| O(n^2)| O(n)|
[Text Justification](https://leetcode.com/problems/text-justification/)| [Swift](./String/TextJustification.swift)| Hard| O(n)| O(n)|

## Linked List
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Swift](./LinkedList/ReverseLinkedList.swift)| Easy| O(n)| O(1)|
[Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)| [Swift](./LinkedList/PalindromeLinkedList.swift)| Easy| O(n)| O(1)|
[Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)| [Swift](./LinkedList/SwapNodesInPairs.swift)| Easy| O(n)| O(1)|
[Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)| [Swift](./LinkedList/RemoveLinkedListElements.swift)| Easy| O(n)| O(1)|
[Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)| [Swift](./LinkedList/RemoveDuplicatesFromSortedList.swift)| Easy| O(n)| O(1)|
[Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/)| [Swift](./LinkedList/RemoveDuplicatesFromSortedListII.swift)| Medium| O(n)| O(1)|
[Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)| [Swift](./LinkedList/RemoveNthFromEnd.swift)| Easy| O(n)| O(1)|
[Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/)| [Swift](./LinkedList/OddEvenLinkedList.swift)| Medium| O(n)| O(1)|
[Rotate List](https://leetcode.com/problems/rotate-list/)| [Swift](./LinkedList/RotateList.swift)| Medium| O(n)| O(1)|
[Reorder List](https://leetcode.com/problems/reorder-list/)| [Swift](./LinkedList/ReorderList.swift)| Medium| O(n)| O(1)|
[Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)| [Swift](./LinkedList/MergeTwoSortedLists.swift)| Easy| O(n)| O(1)|
[Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)| [Swift](./LinkedList/MergeKSortedLists.swift)| Hard| O(mlogn)| O(1)|
[Partition List](https://leetcode.com/problems/partition-list/)| [Swift](./LinkedList/PartitionList.swift)| Medium| O(n)| O(1)|

## Stack
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)| [Swift](./Stack/ValidParentheses.swift)| Easy| O(n)| O(n)|
[Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)| [Swift](./Stack/LongestValidParentheses.swift)| Hard| O(n)| O(n)|
[Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)| [Swift](./Stack/EvaluateReversePolishNotation.swift)| Medium| O(n)| O(n)|
[Simplify Path](https://leetcode.com/problems/simplify-path/)| [Swift](./Stack/SimplifyPath.swift)| Medium| O(n)| O(n)|
[Remove K Digits](https://leetcode.com/problems/remove-k-digits/)| [Swift](./Stack/RemoveKDigits.swift)| Medium| O(n)| O(n)|
[Ternary Expression Parser](https://leetcode.com/problems/ternary-expression-parser/)| [Swift](./Stack/TernaryExpressionParser.swift)| Medium| O(n)| O(n)|
[Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)| [Swift](./Stack/PreorderTraversal.swift)| Medium| O(n)| O(n)|
[Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)| [Swift](./Stack/InorderTraversal.swift)| Medium| O(n)| O(n)|
[Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)| [Swift](./Stack/PostorderTraversal.swift)| Hard| O(n)| O(n)|


## Tree
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Same Tree](https://oj.leetcode.com/problems/same-tree/)| [Swift](./Tree/SameTree.swift)| Easy| O(n)| O(n)|
[Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)| [Swift](./Tree/SymmetricTree.swift)| Easy| O(n)| O(n)|
[Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)| [Swift](./Tree/InvertBinaryTree)| Easy| O(n)| O(n)|
[Binary Tree Upside Down](https://leetcode.com/problems/binary-tree-upside-down/)| [Swift](./Tree/BinaryTreeUpsideDown)| Medium| O(n)| O(1)|
[Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)| [Swift](./Tree/MinimumDepthOfBinaryTree.swift)| Easy| O(n)| O(1)|
[Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)| [Swift](./Tree/MaximumDepthOfBinaryTree.swift)| Easy| O(n)| O(1)|
[Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)| [Swift](./Tree/DiameterBinaryTree.swift)| Easy| O(n)| O(1)|
[Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)| [Swift](./Tree/BalancedBinaryTree.swift)| Easy|  O(n)| O(n)|
[Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/)| [Swift](./Tree/SumLeftLeaves.swift)| Easy|  O(n)| O(1)|
[Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)| [Swift](./Tree/FlattenBinaryTreeLinkedList.swift)| Medium| O(n)| O(1)|
[Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)| [Swift](./Tree/ValidateBinarySearchTree.swift)| Medium| O(n)| O(n)|
[Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)| [Swift](./Tree/BinaryTreeLevelOrderTraversal.swift)| Easy| O(n)| O(n)|
[Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)| [Swift](./Tree/BinaryTreeLevelOrderTraversalII.swift)| Easy| O(n)| O(n)|
[Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)| [Swift](./Tree/BinaryTreeZigzagLevelOrderTraversal.swift)| Medium| O(n)| O(n)|
[Binary Tree Vertical Order Traversal](https://leetcode.com/problems/binary-tree-vertical-order-traversal/)| [Swift](./Tree/BinaryTreeVerticalOrderTraversal.swift)| Medium| O(n)| O(n)|
[Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)| [Swift](./Tree/BinaryTreeRightSideView.swift)| Medium| O(n)| O(n)|
[Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)| [Swift](./Tree/ConstructBinaryTreePreorderInorder.swift)| Medium| O(n)| O(n)|
[Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)| [Swift](./Tree/ConstructBinaryTreeInorderPostorder.swift)| Medium| O(n)| O(n)|
[Path Sum](https://leetcode.com/problems/path-sum/)| [Swift](./Tree/PathSum.swift)| Easy| O(n)| O(n)|
[Path Sum II](https://leetcode.com/problems/path-sum-ii/)| [Swift](./Tree/PathSumII.swift)| Medium| O(n)| O(n)|
[Path Sum III](https://leetcode.com/problems/path-sum-iiI/)| [Swift](./Tree/PathSumIII.swift)| Easy| O(n^2)| O(1)|
[Bnary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)| [Swift](./Tree/BnaryTreePaths.swift)| Easy| O(n)| O(n)|
[Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/)| [Swift](./Tree/UniqueBinarySearchTrees.swift)| Medium| O(n^2)| O(n)|
[Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree/)| [Swift](./Tree/RecoverBinarySearchTree.swift)| Hard| O(n)| O(1)|
[Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/description/) | [Swift](./Tree/MergeTwoBinaryTrees.swift) | Easy | O(n) | O(n) |

## Dynamic programming
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Nested List Weight Sum](https://leetcode.com/problems/nested-list-weight-sum/)| [Swift](./DP/NestedListWeightSum.swift)| Easy| O(n)| O(1)|
[Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)| [Swift](./DP/ClimbingStairs.swift)| Easy| O(n)| O(1)|
[Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)| [Swift](./DP/MinCostClimbingStairs.swift)| Easy| O(n)| O(n)|
[Unique Paths](https://leetcode.com/problems/unique-paths/)| [Swift](./DP/UniquePaths.swift)| Medium| O(mn)| O(mn)|
[Unique Paths II](https://leetcode.com/problems/unique-paths-ii/)| [Swift](./DP/UniquePathsII.swift)| Medium| O(mn)| O(mn)|
[Decode Ways](https://leetcode.com/problems/decode-ways/)| [Swift](./DP/DecodeWays.swift) | O(n)|O(n)|
[Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/)| [Swift](./DP/MinimumPathSum.swift)| Medium| O(mn)| O(mn)|
[Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)| [Swift](./DP/GenerateParentheses.swift)| Medium| O(2^n)| O(n)|
[Different Ways to Add Parentheses](https://leetcode.com/problems/different-ways-to-add-parentheses/)| [Swift](./DP/DifferentWaysAddParentheses.swift)| Medium| O(n^n)| O(n)|
[Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)| [Swift](./DP/BestTimeBuySellStock.swift)| Easy| O(n)| O(1)|
[Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)| [Swift](./DP/BestTimeBuySellStockII.swift)| Medium| O(n)| O(1)|
[Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)| [Swift](./DP/BestTimeBuySellStockIII.swift)| Hard| O(n)| O(n)|
[Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)| [Swift](./DP/BestTimeBuySellStockIV.swift)| Hard| O(n^2)| O(n)|
[Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)| [Swift](./DP/BestTimeBuySellStockCooldown.swift)| Medium| O(n^2)| O(n)|
[Coin Change](https://leetcode.com/problems/coin-change/)| [Swift](./DP/CoinChange.swift)| Medium| O(n^2)| O(n)|
[Coin Change II](https://leetcode.com/problems/coin-change-ii/)| [Swift](./DP/CoinChangeII.swift)| Medium| O(n^2)| O(n)|
[Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)| [Swift](./DP/LongestIncreasingSubsequence.swift)| Medium| O(n^2)| O(n)|
[Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)| [Swift](./DP/LongestPalindromicSubstring.swift)| Medium| O(n^2)| O(n^2)|
[Perfect Squares](https://leetcode.com/problems/perfect-squares/)| [Swift](./DP/PerfectSquares.swift)| Medium| O(n^2)| O(n)|
[House Robber](https://leetcode.com/problems/house-robber/)| [Swift](./DP/HouseRobber.swift)| Easy| O(n)| O(1)|
[House Robber II](https://leetcode.com/problems/house-robber-ii/)| [Swift](./DP/HouseRobberII.swift)| Medium| O(n)| O(1)|
[Paint Fence](https://leetcode.com/problems/paint-fence/)| [Swift](./DP/PaintFence.swift)| Easy| O(n)| O(n)|
[Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)| [Swift](./DP/MaximumSubarray.swift)| Medium| O(n)| O(1)|
[Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)| [Swift](./DP/MaximumProductSubarray.swift)| Medium| O(n)| O(1)|
[Maximal Square](https://leetcode.com/problems/maximal-square/)| [Swift](./DP/MaximalSquare.swift)| Medium| O(mn)| O(mn)|
[Edit Distance](https://leetcode.com/problems/edit-distance/)| [Swift](./DP/EditDistance.swift)| Hard| O(mn)| O(mn)|
[Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/)| [Swift](./DP/CombinationSumIV.swift)| Medium| O(2^n)| O(n)|
[Triangle](https://leetcode.com/problems/triangle/)| [Swift](./DP/Triangle.swift)| Medium| O(2^n - 1)| O(m)|
[Guess Number Higher or Lower II](https://leetcode.com/problems/guess-number-higher-or-lower-ii/)| [Swift](./DP/GuessNumberHigherOrLowerII.swift)| Medium| O(nlogn)| O(n^2)|
[Burst Ballons](https://leetcode.com/problems/burst-balloons/)| [Swift](./DP/BurstBalloons.swift)| Hard| O(n^3)| O(n)|
[Frog Jump](https://leetcode.com/problems/frog-jump/)| [Swift](./DP/FrogJump.swift)| Hard| O(n^2)| O(n)|

## Depth-first search
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Permutations](https://leetcode.com/problems/permutations/)| [Swift](./DFS/Permutations.swift)| Medium| O(n!)| O(n)|
[Permutations II](https://leetcode.com/problems/permutations-ii/)| [Swift](./DFS/PermutationsII.swift)| Medium| O(n!)| O(n)|
[Subsets](https://leetcode.com/problems/subsets/)| [Swift](./DFS/Subsets.swift)| Medium| O(n!)| O(n)|
[Subsets II](https://leetcode.com/problems/subsets-ii/)| [Swift](./DFS/SubsetsII.swift)| Medium| O(n!)| O(n)|
[Combinations](https://leetcode.com/problems/combinations/)| [Swift](./DFS/Combinations.swift)| Medium| O(n!)| O(n)|
[Combination Sum](https://leetcode.com/problems/combination-sum/)| [Swift](./DFS/CombinationSum.swift)| Medium| O(n^n)| O(2^n - 1)|
[Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)| [Swift](./DFS/CombinationSumII.swift)| Medium| O(n!)| O(2^n - 2)|
[Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)| [Swift](./DFS/CombinationSumIII.swift)| Medium| O(n!)| O(nCk)|
[Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| [Swift](./DFS/LetterCombinationsPhoneNumber.swift)| Medium| O(mn)| O(n)|
[Factor Combinations](https://leetcode.com/problems/factor-combinations/)| [Swift](./DFS/FactorCombinations.swift)| Medium| O(n^n))| O(2^n - 1)|
[Generalized Abbreviation](https://leetcode.com/problems/generalized-abbreviation/)| [Swift](./DFS/GeneralizedAbbreviation.swift)| Medium| O(n!)| O(2^n)|
[Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)| [Swift](./DFS/PalindromePartitioning.swift)| Medium| O(n!)| O(n)|
[Number of Islands](https://leetcode.com/problems/number-of-islands/)| [Swift](./DFS/NumberofIslands.swift)| Medium| O((mn)^2)| O(1)|
[Walls and Gates](https://leetcode.com/problems/walls-and-gates/)| [Swift](./DFS/WallsGates.swift)| Medium| O(n!)| O(2^n)|
[Word Search](https://leetcode.com/problems/word-search/)| [Swift](./DFS/WordSearch.swift)| Medium| O((n^2)!)| O(n^2)|
[Word Search II](https://leetcode.com/problems/word-search-ii/)| [Swift](./DFS/WordSearchII.swift)| Hard| O(((mn)^2))| O(n^2)|
[Add and Search Word - Data structure design](https://leetcode.com/problems/add-and-search-word-data-structure-design/)| [Swift](./DFS/WordDictionary.swift)| Medium| O(n)| O(n)|
[N-Queens](https://leetcode.com/problems/n-queens/)| [Swift](./DFS/NQueens.swift)| Hard| O((n^4))| O(n^2)|
[N-Queens II](https://leetcode.com/problems/n-queens-ii/)| [Swift](./DFS/NQueensII.swift)| Hard| O((n^3))| O(n)|
[Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)| [Swift](./DFS/SudokuSolver.swift)| Hard| O(n^4)| O(1)|
[Remove Invalid Parentheses](https://leetcode.com/problems/remove-invalid-parentheses/)| [Swift](./DFS/RemoveInvalidParentheses.swift)| Hard| O(n!)| O(n)|
[Expression Add Operators](https://leetcode.com/problems/expression-add-operators/)| [Swift](./DFS/ExpressionAddOperators.swift)| Hard| O(n!)| O(n)|

## Math
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Add Binary](https://leetcode.com/problems/add-binary/)| [Swift](./Math/AddBinary.swift)| Easy| O(n)| O(n)|
[Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)| [Swift](./Math/AddTwoNumbers.swift)| Medium| O(n)| O(1)|
[Add Digits](https://leetcode.com/problems/add-digits/)| [Swift](./Math/AddDigits.swift)| Easy| O(1)| O(1)|
[Plus One](https://leetcode.com/problems/plus-one/)| [Swift](./Math/PlusOne.swift)| Easy| O(n)| O(1)|
[Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)| [Swift](./Math/DivideTwoIntegers.swift)| Medium| O(logn)| O(1)|
[Number Complement](https://leetcode.com/problems/number-complement/)| [Swift](./Math/NumberComplement.swift)| Easy| O(n)| O(1)|
[Hamming Distance](https://leetcode.com/problems/hamming-distance/)| [Swift](./Math/HammingDistance.swift)| Easy| O(n)| O(1)|
[Integer Break](https://leetcode.com/problems/integer-break/)| [Swift](./Math/IntegerBreak.swift)| Medium| O(logn)| O(1)|
[Happy Number](https://leetcode.com/problems/happy-number/)| [Swift](./Math/HappyNumber.swift)| Easy| O(n)| O(n)|
[Single Number](https://leetcode.com/problems/single-number/)| [Swift](./Math/SingleNumber.swift)| Medium| O(n)| O(1)|
[Ugly Number](https://leetcode.com/problems/ugly-number/)| [Swift](./Math/UglyNumber.swift)| Easy| O(logn)| O(1)|
[Ugly Number II](https://leetcode.com/problems/ugly-number-ii/)| [Swift](./Math/UglyNumberII.swift)| Medium| O(n)| O(n)|
[Super Ugly Number](https://leetcode.com/problems/super-ugly-number/)| [Swift](./Math/SuperUglyNumber.swift)| Medium| O(n^2)| O(n)|
[Count Primes](https://leetcode.com/problems/count-primes/)| [Swift](./Math/CountPrimes.swift)| Easy| O(n)| O(n)|
[String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)| [Swift](./Math/Atoi.swift)| Easy| O(n)| O(1)|
[Pow(x, n)](https://leetcode.com/problems/isomorphic-strings/)| [Swift](./Math/Pow.swift)| Medium| O(logn)| O(1)|
[Power of Two](https://leetcode.com/problems/power-of-two/)| [Swift](./Math/PowerTwo.swift)| Easy| O(1)| O(1)|
[Power of Three](https://leetcode.com/problems/power-of-three/)| [Swift](./Math/PowerThree.swift)| Easy| O(1)| O(1)|
[Super Power](https://leetcode.com/problems/super-pow/)| [Swift](./Math/SuperPow.swift)| Medium| O(n)| O(1)|
[Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)| [Swift](./Math/SumTwoIntegers.swift)| Easy| O(n)| O(1)|
[Reverse Integer](https://leetcode.com/problems/reverse-integer/)| [Swift](./Math/ReverseInteger.swift)| Easy| O(n)| O(1)|
[Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/)| [Swift](./Math/ExcelSheetColumnNumber.swift)| Easy| O(n)| O(1)|
[Integer to Roman](https://leetcode.com/problems/integer-to-roman/)| [Swift](./Math/IntegerToRoman.swift)| Medium| O(n)| O(1)|
[Roman to Integer](https://leetcode.com/problems/roman-to-integer/)| [Swift](./Math/RomanToInteger.swift)| Easy| O(n)| O(n)|
[Integer to English Words](https://leetcode.com/problems/integer-to-english-words/)| [Swift](./Math/IntegerEnglishWords.swift)| Hard| O(n)| O(1)|
[Rectangle Area](https://leetcode.com/problems/rectangle-area/)| [Swift](./Math/RectangleArea.swift)| Easy| O(1)| O(1)|
[Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)| [Swift](./Math/TrappingRainWater.swift)| Hard| O(n)| O(n)|
[Container With Most Water](https://leetcode.com/problems/container-with-most-water/)| [Swift](./Math/ContainerMostWater.swift)| Medium| O(n)| O(1)|
[Counting Bits](https://leetcode.com/problems/counting-bits/)| [Swift](./Math/CountingBits.swift)| Medium| O(n)| O(n)|
[K-th Smallest in Lexicographical Order](https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/)| [Swift](./Math/KthSmallestLexicographicalOrder.swift)| Hard| O(n)| O(1)|

## Search
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Closest Binary Search Tree Value](https://leetcode.com/problems/closest-binary-search-tree-value/)| [Swift](./Search/ClosestBinarySearchTreeValue.swift)| Easy| O(logn)| O(1)|
[Closest Binary Search Tree Value II](https://leetcode.com/problems/closest-binary-search-tree-value-ii/)| [Swift](./Search/ClosestBinarySearchTreeValueII.swift)| Hard| O(n)| O(n)|
[Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)| [Swift](./Search/SearchInRotatedSortedArray.swift)| Hard| O(logn)| O(1)|
[Search in Rotated Sorted Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)| [Swift](./Search/SearchInRotatedSortedArrayII.swift)| Medium| O(logn)| O(1)|
[Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)| [Swift](./Search/FindMinimumRotatedSortedArray.swift)| Medium| O(logn)| O(1)|
[Find Minimum in Rotated Sorted Array II](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)| [Swift](./Search/FindMinimumRotatedSortedArrayII.swift)| Hard| O(logn)| O(1)|
[Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)| [Swift](./Search/Search2DMatrix.swift)| Medium| O(log(m + n))| O(1)|
[Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/)| [Swift](./Search/Search2DMatrixII.swift)| Medium| O(m + n)| O(1)|
[Search for a Range](https://leetcode.com/problems/search-for-a-range/)| [Swift](./Search/SearchForARange.swift)| Medium| O(logn)| O(1)|
[Search Insert Position](https://leetcode.com/problems/search-insert-position/)| [Swift](./Search/SearchForARange.swift)| Medium| O(logn)| O(1)|
[Find Peak Element](https://leetcode.com/problems/find-peak-element/)| [Swift](./Search/FindPeakElement.swift)| Medium| O(logn)| O(1)|
[Sqrt(x)](https://leetcode.com/problems/sqrtx/)| [Swift](./Search/Sqrtx.swift)| Medium| O(logn)| O(1)|
[Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)| [Swift](./Search/MedianTwoSortedArrays.swift)| Hard| O(log(m + n))| O(1)|

## Sort
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)| [Swift](./Sort/MergeSortedArray.swift)| Easy| O(n)| O(1)|
[Sort Colors](https://leetcode.com/problems/sort-colors/)| [Swift](./Sort/SortColors.swift)| Medium| O(n)| O(1)|
[Wiggle Sort](https://leetcode.com/problems/wiggle-sort/)| [Swift](./Sort/WiggleSort.swift)| Medium| O(n)| O(1)|
[Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii/)| [Swift](./Sort/WiggleSortII.swift)| Medium| O(nlogn)| O(n)|
[Sort Transformed Array](https://leetcode.com/problems/sort-transformed-array/)| [Swift](./Sort/SortTransformedArray.swift)| Medium| O(n)| O(1)|
[Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)| [Swift](./Array/TopKFrequentElements.swift)| Medium| O(nlogn)| O(n)|
[Meeting Rooms](https://leetcode.com/problems/meeting-rooms/)| [Swift](./Sort/MeetingRooms.swift)| Easy| O(nlogn)| O(1)|
[Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/)| [Swift](./Sort/MeetingRoomsII.swift)| Medium| O(nlogn)| O(n)|
[Merge Intervals](https://leetcode.com/problems/merge-intervals/)| [Swift](./Sort/MergeIntervals.swift)| Hard| O(nlogn)| O(n)|
[Alien Dictionary](https://leetcode.com/problems/alien-dictionary/)| [Swift](./Sort/AlienDictionary.swift)| Hard| O(nm)| O(nm)|
[Array Partition I](https://leetcode.com/problems/array-partition-i/description/)| [Swift](./Sort/ArrayPartitionI.swift)|Easy| O(nlogn)| O(n)|

## Union Find
| Title | Solution | Difficulty | Time | Space |
| ----- | -------- | ---------- | ---- | ----- |
[Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/)| [Swift](./UnionFind/NumberConnectedComponentsUndirectedGraph.swift)| Medium| O(nlogn)| O(n)|
[Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/)| [Swift](./UnionFind/GraphValidTree.swift)| Medium| O(nlogn)| O(n)|

## Google
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Plus One](https://leetcode.com/problems/plus-one/)| [Swift](./Math/PlusOne.swift)| Easy| ★★★★★★|
[Number of Islands](https://leetcode.com/problems/number-of-islands/)| [Swift](./DFS/NumberofIslands.swift)| Medium| ★★★★|
[Summary Ranges](https://leetcode.com/problems/summary-ranges/)| [Swift](./Array/SummaryRanges.swift)| Medium| ★★★★|
[Perfect Squares](https://leetcode.com/problems/perfect-squares/)| [Swift](./DP/PerfectSquares.swift)| Medium| ★★★★|
[Merge Intervals](https://leetcode.com/problems/merge-intervals/)| [Swift](./Sort/MergeIntervals.swift)| Hard| ★★★|
[Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)| [Swift](./Stack/ValidParentheses.swift)| Easy| ★★★|
[Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)| [Swift](./Math/TrappingRainWater.swift)| Hard| ★★|
[Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)| [Swift](./LinkedList/MergeKSortedLists.swift)| Hard| ★★|
[Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)| [Swift](./Array/LongestConsecutiveSequence.swift)| Hard| ★★|
[Find Peak Element](https://leetcode.com/problems/find-peak-element/)| [Swift](./Search/FindPeakElement.swift)| Medium| ★★|
[Power of Two](https://leetcode.com/problems/power-of-two/)| [Swift](./Math/PowerTwo.swift)| Easy| ★★|
[Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)| [Swift](./Array/SpiralMatrix.swift)| Medium| ★★|
[Sliding Window Maximum ](https://leetcode.com/problems/sliding-window-maximum/)| [Swift](./Array/SlidingWindowMaximum.swift)| Hard| ★★|
[Pow(x, n)](https://leetcode.com/problems/isomorphic-strings/)| [Swift](./Math/Pow.swift)| Medium| ★★|
[Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| [Swift](./DFS/LetterCombinationsPhoneNumber.swift)| Medium| ★★|
[Heaters](https://leetcode.com/problems/heaters/)| [Swift](./Array/Heaters.swift)| Easy| ★|

## Facebook
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[3Sum](https://leetcode.com/problems/3sum/)| [Swift](./Array/ThreeSum.swift)| Medium| ★★★★★★|
[Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)| [Swift](./String/ValidPalindrome.swift)| Easy| ★★★★★★|
[Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/)| [Swift](./String/ValidPalindromeII.swift)| Easy| ★★★★★★|
[Move Zeroes](https://leetcode.com/problems/move-zeroes/)| [Swift](./Array/MoveZeroes.swift)| Easy| ★★★★★★|
[Remove Invalid Parentheses](https://leetcode.com/problems/remove-invalid-parentheses/)| [Swift](./DFS/RemoveInvalidParentheses.swift)| Hard| ★★★★★★|
[Add Binary](https://leetcode.com/problems/add-binary/)| [Swift](./Math/AddBinary.swift)| Easy| ★★★★★|
[Two Sum](https://leetcode.com/problems/two-sum/)| [Swift](./Array/TwoSum.swift)| Easy| ★★★★★|
[Bnary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)| [Swift](./Tree/BnaryTreePaths.swift)| Easy| ★★★★|
[Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| [Swift](./DFS/LetterCombinationsPhoneNumber.swift)| Medium| ★★★★|
[Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)| [Swift](./LinkedList/MergeKSortedLists.swift)| Hard| ★★★★|
[Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Swift](./LinkedList/ReverseLinkedList.swift)| Easy| ★★★|
[Merge Intervals](https://leetcode.com/problems/merge-intervals/)| [Swift](./Sort/MergeIntervals.swift)| Hard| ★★★|
[Number of Islands](https://leetcode.com/problems/number-of-islands/)| [Swift](./DFS/NumberofIslands.swift)| Medium| ★★★|
[Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Swift](./LinkedList/ReverseLinkedList.swift)| Easy| ★★★|
[Expression Add Operators](https://leetcode.com/problems/expression-add-operators/)| [Swift](./DFS/ExpressionAddOperators.swift)| Hard| ★★★|
[Subsets](https://leetcode.com/problems/subsets/)| [Swift](./DFS/Subsets.swift)| Medium| ★★★|
[Sort Colors](https://leetcode.com/problems/sort-colors/)| [Swift](./Sort/SortColors.swift)| Medium| ★★|

## Snapchat
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Game of Life](https://leetcode.com/problems/game-of-life/)	|	[Swift](./Array/GameLife.swift)| Medium| ★★★★★★|
[Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/)| [Swift](./Sort/MeetingRoomsII.swift)| Medium| ★★★★★★|
[Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)| [Swift](./Array/ValidSudoku.swift)| Easy| ★★★★★|
[Binary Tree Vertical Order Traversal](https://leetcode.com/problems/binary-tree-vertical-order-traversal/)| [Swift](./Tree/BinaryTreeVerticalOrderTraversal.swift)| Medium| ★★★★|
[Alien Dictionary](https://leetcode.com/problems/alien-dictionary/)| [Swift](./Sort/AlienDictionary.swift)| Hard| ★★★★|
[One Edit Distance](https://leetcode.com/problems/one-edit-distance/)| [Swift](./String/OneEditDistance.swift)| Medium| ★★★|
[Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)| [Swift](./Math/SudokuSolver.swift)| Hard| ★★★|
[Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Swift](./LinkedList/ReverseLinkedList.swift)| Easy| ★★|
[Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/)| [Swift](./Tree/UniqueBinarySearchTrees.swift)| Medium| ★★|
[Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)| [Swift](./Array/MinimumWindowSubstring.swift)| Hard| ★★|
[Remove K Digits](https://leetcode.com/problems/remove-k-digits/)| [Swift](./Stack/RemoveKDigits.swift)| Medium| ★|
[Ternary Expression Parser](https://leetcode.com/problems/ternary-expression-parser/)| [Swift](./Stack/TernaryExpressionParser.swift)| Medium| ★|

## Uber
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)| [Swift](./Array/ValidSudoku.swift)| Easy| ★★★★|
[Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)| [Swift](./Array/SpiralMatrix.swift)| Medium| ★★★★|
[Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| [Swift](./DFS/LetterCombinationsPhoneNumber.swift)| Medium| ★★★★|
[Group Anagrams](https://leetcode.com/problems/anagrams/)| [Swift](./String/GroupAnagrams.swift)| Medium| ★★★★|
[Word Pattern](https://leetcode.com/problems/word-pattern/)| [Swift](./String/WordPattern.swift)| Easy| ★★★|
[Roman to Integer](https://leetcode.com/problems/roman-to-integer/)| [Swift](./Math/RomanToInteger.swift)| Easy| ★★★|
[Combination Sum](https://leetcode.com/problems/combination-sum/)| [Swift](./DFS/CombinationSum.swift)| Medium| ★★|

## Airbnb
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Two Sum](https://leetcode.com/problems/two-sum/)| [Swift](./Array/TwoSum.swift)| Easy| ★★★★★|
[Text Justification](https://leetcode.com/problems/text-justification/)| [Swift](./String/TextJustification.swift)| Hard| ★★★★|
[House Robber](https://leetcode.com/problems/house-robber/)| [Swift](./DP/HouseRobber.swift)| Easy| ★★|
[Single Number](https://leetcode.com/problems/single-number/)| [Swift](./Math/SingleNumber.swift)| Medium| ★★|
[Word Search II](https://leetcode.com/problems/word-search-ii/)| [Swift](./DFS/WordSearchII.swift)| Hard| ★★|
[Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)| [Swift](./Math/AddTwoNumbers.swift)| Medium| ★★|

## LinkedIn
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)| [Swift](./DP/MaximumSubarray.swift)| Medium| ★★★★★★|
[Pow(x, n)](https://leetcode.com/problems/isomorphic-strings/)| [Swift](./Math/Pow.swift)| Medium| ★★★★★★|
[Merge Intervals](https://leetcode.com/problems/merge-intervals/)| [Swift](./Sort/MergeIntervals.swift)| Hard| ★★★★★★|
[Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/)| [Swift](./String/IsomorphicStrings.swift)| Easy| ★★★★★★|
[Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)| [Swift](./Search/SearchInRotatedSortedArray.swift)| Hard| ★★★★★|
[Search for a Range](https://leetcode.com/problems/search-for-a-range/)| [Swift](./Search/SearchForARange.swift)| Medium| ★★★★★|
[Two Sum](https://leetcode.com/problems/two-sum/)| [Swift](./Array/TwoSum.swift)| Easy| ★★★★|
[Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)| [Swift](./Tree/BinaryTreeLevelOrderTraversal.swift)| Easy| ★★★★|
[Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)| [Swift](./Stack/EvaluateReversePolishNotation.swift)| Medium| ★★★|
[Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)| [Swift](./DP/MaximumProductSubarray.swift)| Medium| ★★★|
[Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)| [Swift](./Array/ProductExceptSelf.swift)| Medium| ★★★|
[Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)| [Swift](./Tree/SymmetricTree.swift)| Easy| ★★|

## Amazon
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Two Sum](https://leetcode.com/problems/two-sum/)| [Swift](./Array/TwoSum.swift)| Easy| ★★★★★★|
[Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)| [Swift](./DP/MinCostClimbingStairs.swift)| Easy| ★★★★|
[Number of Islands](https://leetcode.com/problems/number-of-islands/)| [Swift](./DFS/NumberofIslands.swift)| Medium| ★★|
[Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)| [Swift](./Math/AddTwoNumbers.swift)| Medium| ★★|
[Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Swift](./LinkedList/ReverseLinkedList.swift)| Easy| ★★|
[Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)| [Swift](./Stack/ValidParentheses.swift)| Easy| ★★|
[Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)| [Swift](./DP/LongestPalindromicSubstring.swift)| Medium| ★★|
[Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)| [Swift](./Math/TrappingRainWater.swift)| Hard| ★★|
[Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)| [Swift](./String/LongestSubstringWithoutRepeatingCharacters.swift)| Medium| ★★|
[Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| [Swift](./DFS/LetterCombinationsPhoneNumber.swift)| Medium| ★★|
[Valid Anagram](https://leetcode.com/problems/valid-anagram/)| [Swift](./String/ValidAnagram.swift)| Easy| ★★|
[Rotate Image](https://leetcode.com/problems/rotate-image/)| [Swift](./Array/RotateImage.swift)| Medium| ★★|
[Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)| [Swift](./DP/BestTimeBuySellStock.swift)| Easy| ★★|
[3Sum](https://leetcode.com/problems/3sum/)| [Swift](./Array/ThreeSum.swift)| Medium| ★★|
[Sliding Window Maximum ](https://leetcode.com/problems/sliding-window-maximum/)| [Swift](./Array/SlidingWindowMaximum.swift)| Hard| ★★|

## Microsoft
| Title | Solution | Difficulty | Frequency |
| ----- | -------- | ---------- | --------- |
[Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Swift](./LinkedList/ReverseLinkedList.swift)| Easy| ★★★★★★|
[Two Sum](https://leetcode.com/problems/two-sum/)| [Swift](./Array/TwoSum.swift)| Easy| ★★★★★|
[String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)| [Swift](./Math/Atoi.swift)| Easy| ★★★★|
[Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)| [Swift](./Math/AddTwoNumbers.swift)| Medium| ★★★★|
[Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/)| [Swift](./Math/ExcelSheetColumnNumber.swift)| Easy| ★★★★|
[Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)| [Swift](./Tree/ValidateBinarySearchTree.swift)| Medium| ★★★|
[Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)| [Swift](./LinkedList/MergeTwoSortedLists.swift)| Easy| ★★★|
# Numbers Sprouted on The Binary Tree

In our garden there is a beautiful binary tree. In spring, many numbers germinate. Because we are very greedy, we always want to get the biggest one at the beginning.

Let's say we own the following binary tree:

```
              271
               |
       -----------------
       |               |
      422             688
       |               |
   ---------       ---------
   |       |       |       |
  389     872     754     147
   |       |       |       |
 -----   -----   -----   -----
 |   |   |   |   |   |   |   |
231 564 112 211 145 275 124 333
```

Here, the biggest number is 872.

To process the binary tree, it will be structured as an array: the root node is placed at index 0, the root’s left child at index 1, its right child at index 2, and so on, progressing from left to right along each level of the tree. In our example, the binary tree is represented by the array:

```
271 422 688 389 872 754 147 231 564 112 211 145 275 124 333
```

The root starts at depth 1 and a level starts at index 1. So here, our biggest number is located at depth 3 and is the 2nd of this level. Finally, the flag is the multiplication of these 3 values:

```
3 * 2 * 872 = 5232
```

Find the flag, given the tree from the input file.

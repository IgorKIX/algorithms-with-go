# Quick find / quick union / weighted quick unionf + path compression

## Problem to solve:
We have a group of consecutive numbers. Their number is known (the first number is 0 and we know the value of the last number). We want to combine these numbers into sets and then be able to answer the question of whether the two given numbers are linked together.

Create an algorithm to union numbers together into groups by `union(int, int)` method, and than be able to say if the given numbers are connected `find(int, int)`.

ex:
`union(1, 2)`, `union(2, 3)`, `union(4, 5)` will give us `{1, 2, 3} {4, 5}`
`find(1, 3)` will return `true`, but `find(1, 5)` will return `false`

## Ways of solving

| Algorithm | Worst case time |
| --- | --- |
| quick-find | $MN$ |
| quick-union | $MN$ |
| weighted q-u | $N + M\log{N}$ |
| q-u + path compression | $N + M\log{N}$ |
| weighted q-u + path compression | $N + M\log*{N}$ |

M - union/find operations on a set of N objects

### Quick find
Cause we know the lenght of numbers, so we can create an array. Array indexes represent a number, and the number under each index represents the "id" of the group to which that number belongs. 

To connect two numbers, we assign an "id" of the same group to both of numbers (We put the same value under both indexes). The same thing happens with other indexes that have kept the same id as those that were under given in union method.

To find out if two given numbers are connected, we check what content they contain. When the values are the same, it means they are in the same group .

### Quick union

We operate on arrays again. This time we will create a kind of tree.

Connecting numbers together is creating a relationship where one number becomes the root of the other. We are talking about trees because when we look at the root values of next roots, we will get to the number whose root points to that number.

The search for main root method is used to check if both numbers are in the same group. So we find the main root for both of the given numbers (that is, the number whose root value points to itself). When this number is identical, it means that both checked numbers are connected with each other.

ex. for better understanding :wink:
```
id   0 1 2 3 4 5 6 7
[id] 0 2 0 0 6 5 7 7

    0         7    5
   / \       /
  2   3     6
 /         /
1         4
```
### Weighted quick union + path compression

The idea is similar to the previous one, but this time we also keeps the track of how high (how many valuse) got this tree. Thanks to this information we are able to connect only smaller trees to the main root of the bigger one, during the `union` operation. This easy trick help us to avoid the problem of the to high trees.

#### path compression
With a help of the small line of code We are able to "teach" the branch what is the main root of it. During the search of the main root we will write to all numbers of the branch what was the main root. Thanks to this, after couple of searches we will be able to create almost flat trees, what will speed up a lot, the search for main root.

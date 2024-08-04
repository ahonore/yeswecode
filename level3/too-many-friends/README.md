# Too Many Friends

You travel a lot so you meet so many people. You know that you're going to meet a lot more, but your address book is already full. Suddenly, you think about writing an algorithm that can optimize entries, to save memory and increase the speed of search in your address book.

Instead of writing full first and last names line by line in your address book, you've decided to use some sort of indexing tree. To be faster to find someone, you also put letters in lowercase and remove diacritics of characters (accents, cedillas...). For example, instead of:

```
François Plo
Franck Dia
Ivo Geißler
```

You will write nodes (a node contains one character) in a tree as follow:

```
f ─ r ─ a ─ n ─ c ┬ o ─ i ─ s ─   ─ p ─ l ─ o
                  └─ k ─   ─ d ─ i ─ a
i ─ v ─ o ─   ─ g ─ e ─ i ─ ß ─ l ─ e ─ r
```

Here, you needed 28 nodes to write 3 entries. Given the first and last names of your friends in the input file, return the number of nodes used to write your whole address book.

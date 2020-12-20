# GO WaterSortPuzzleSolver

Solves WaterSortPuzzle games, you can find the game here:
[Android](https://play.google.com/store/apps/details?id=com.gma.water.sort.puzzle),
[IOS](https://apps.apple.com/us/app/water-sort-puzzle/id1514542157).

Download the game, it is fun if you like that sort of games. 

Got stuck in level 105, so I am implementing a solver. 

Solution is a trivial exhaustive search, but some details were considerated:
- relatively small data structure (56 bytes)
- no play-out will be tried twice, to avoid potential infinite loops and prune search space
- ...
- magic?

```shell script
$ make run
```

![](lvl105.jpg)



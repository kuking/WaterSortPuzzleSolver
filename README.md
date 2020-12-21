# GO Water Sort Puzzle Solver

Solves Water Sort Puzzle games, you can find the game here:
[Android](https://play.google.com/store/apps/details?id=com.gma.water.sort.puzzle),
[IOS](https://apps.apple.com/us/app/water-sort-puzzle/id1514542157).

Download the game, it is fun if you like that sort of games. 

Got stuck in level 105, so I implemented a solver. 

Solution is a trivial exhaustive search, details:
- relatively small data structure (56 bytes, L1 cache friendly)
- no play-out will be tried twice, to avoid potential infinite loops and prune search space
- avoids moving solved vials into empty vials (innocuous move)
- reasonably fast for v1, about 3µs per move (~350k moves/sec single-core), it will find instantly one solution; 
  it can take a couple of minutes to find the shortest solution for some levels.

Comments:
- It can search the shortest solution, but it is usually not worth the extra time, it takes longer and the greedy
  algorithm usually finds a similar solution with maybe one extra move.
  
  
The game:
  
![](lvl105.jpg)


```shell script
$ make 
go clean -testcache -cache
rm -f WSPZ
go test
PASS
ok  	github.com/kukino/WaterSortPuzzleSolver	0.262s
go build -o WSPZ
./WSPZ
Solving Level 105
Solution took: 27.920402ms, exploring 7730 moves, or 3.611µs/move
Move  1:  2 -> 13
Move  2:  3 -> 14
Move  3:  3 -> 13
Move  4:  4 -> 14
Move  5:  2 ->  4
Move  6:  9 ->  2
Move  7: 12 ->  2
Move  8: 11 ->  9
Move  9:  6 -> 11
Move 10:  5 ->  6
Move 11:  5 ->  3
Move 12: 11 ->  5
Move 13:  8 -> 11
Move 14:  8 -> 13
Move 15:  4 ->  8
Move 16:  4 ->  3
Move 17:  4 -> 14
Move 18:  3 ->  4
Move 19:  7 ->  3
Move 20: 10 ->  3
Move 21:  7 -> 13
Move 22:  7 ->  4
Move 23:  7 -> 11
Move 24:  5 ->  7
Move 25:  1 ->  5
Move 26: 10 ->  7
Move 27:  9 -> 10
Move 28:  9 -> 12
Move 29:  5 ->  9
Move 30:  1 ->  5
Move 31: 12 ->  1
Move 32:  6 ->  5
Move 33:  6 ->  3
Move 34:  5 ->  6
Move 35:  2 ->  5
Move 36: 10 ->  2
Move 37:  5 -> 10
Move 38:  8 ->  5
Move 39: 12 ->  5
Move 40: 11 ->  8
Move 41:  9 -> 11
Move 42: 12 -> 14
```



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
- reasonably fast about 1.35µs per move simulation (~720k moves/sec single-core), it will find instantly one solution; 
  it can take a couple of minutes to find the shortest solution for some levels.

Comments:
- It can search the shortest solution, but it is usually not worth the extra time, the greedy algorithm usually finds
  a similar solution with maybe one extra move.
- At least Go 1.14 is required.
 
The game:
  
![](lvl105.jpg)


```shell script
$ lscpu | grep "Model name"
Model name:          AMD Ryzen 7 3800X 8-Core Processor
$ make 
go clean -testcache -cache
rm -f WSPZ
go test
PASS
ok  	github.com/kukino/WaterSortPuzzleSolver	0.002s
go build -o WSPZ
./WSPZ
Solving Level 105, find shortest: true
mega moves: 1m, t: 1s, p: 1.629µs/m, d: 46, ts: {s: 14 2222 7777 8888 0444 9999 5555 6666 bbbb cccc aaaa 0000 0004 1111 3333}
mega moves: 2m, t: 3s, p: 1.552µs/m, d: 44, ts: {s: 14 00cc 7777 6666 3333 aaaa 8888 9999 00cc bbbb 5555 2222 0004 1111 0444}
mega moves: 3m, t: 4s, p: 1.499µs/m, d: 45, ts: {s: 14 cccc 0bbb 6666 4444 005b aaaa 3333 0555 8888 9999 0000 7777 1111 2222}
[...]
mega moves: 54m, t: 1m18s, p: 1.456µs/m, d: 48, ts: {s: 14 cccc 1111 0002 6666 bbbb 8888 aaaa 0000 0222 9999 7777 4444 5555 3333}
mega moves: 55m, t: 1m20s, p: 1.464µs/m, d: 47, ts: {s: 14 cccc 3333 0002 0444 7777 6666 5555 aaaa 1111 9999 bbbb 2224 0000 8888}
Solution took: 1m16.217403532s, exploring 55370061 moves, 1.376µs/move, 719091 mps (moves-per-second)
Move  1:  5 -> 13
Move  2:  6 -> 14
Move  3:  6 -> 13
Move  4:  7 ->  6
Move  5: 10 ->  6
Move  6: 10 -> 14
Move  7: 11 -> 10
Move  8: 11 -> 14
Move  9:  8 -> 11
Move 10:  7 ->  8
Move 11:  7 ->  5
Move 12:  7 -> 11
Move 13:  2 ->  7
Move 14:  8 ->  7
Move 15:  2 ->  8
Move 16:  9 ->  2
Move 17: 12 ->  2
Move 18:  9 -> 10
Move 19:  9 -> 12
Move 20:  1 ->  9
Move 21:  1 -> 13
Move 22: 12 ->  1
Move 23:  8 -> 12
Move 24: 11 ->  8
Move 25:  9 -> 11
Move 26: 12 ->  9
Move 27:  3 -> 12
Move 28:  3 ->  7
Move 29:  4 -> 12
Move 30:  4 ->  9
Move 31:  3 ->  4
Move 32:  6 ->  3
Move 33:  6 -> 13
Move 34:  4 ->  6
Move 35: 12 ->  4
Move 36:  2 -> 12
Move 37: 10 ->  2
Move 38:  5 ->  6
Move 39:  5 -> 14
Move 40: 10 -> 12
Move 41:  5 -> 11
```

Changelog:
- 2020/12/21: code cleanup, performance improved X2.9 (~4us to ~1.37us)

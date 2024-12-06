import Tasks

extension Day6 {
    var example1: String {
        return """
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
"""
    }
    var data: String {
        return """
.................#.#............................................#................#........................................#.......
.....#.............................................................#.......#...#........................#...........#.......##....
..........................................................#..#........#..........#...............#.#..................#...........
..#...#..........#....#.....#......................................#......#..#............................#..#....................
.........................#....#.............#....................................#............................................#...
.....................#.......#.............#...#..........#.................#.......#...............#............#......#.........
........#.........#.....................................#............#.....................##...#..#..........#...................
....#...................#.....##......#............#...#........................#...................................#.............
...............................................................................#.......#.....#...#................#.#......##.....
..........................................#............................#.#.................................#.........#............
........#...............#...............#...................##......#..................#...............#..........................
.........................#....................................#..........#.........................................##.............
.......#.................#..................................#......#....................#...............#..........#......#.......
.............#.#...#.............#.............................................#...............................##.....#...#.......
.....................................................#............................................#.........................#.....
.#...#..........................................................##.................................#............#...............#.
...........#..#........#.....#.............#.......#............................................#..............#.................#
......#.#.##...............#.......#......#......................................................................#...............#
...#........#.#.....................##...#...........................#..............................#...................#.........
..............#..#....................#...................#.................##..............#.#........................#..........
...........................................................#.........#.......................#......................#.........#...
................................................#............................................#....................................
..#.....................................................#..............#...#.........#.....#........................#.............
.#..............................#......#..............#........................................................#..................
#......#.....#..#..#........#....................#...........##..............................#..#.....#.....#.....................
.................#.....#.....................#.........#..............................#..............##......................#....
...........................................................#...##..............#..........................#....#..................
........#..#............................#.#........................#.....#................#...#...................................
..........................#........#.......................................................................##...................#.
............#....#.............................................#.......#....#.....#........................................#......
..........#............#....#.......................#.#............#..............................................................
......#...#......................................................#..#.............................................................
.#...#.....#.....#......................................................................................#..........#.....#........
...#..#.............................................................................#.............................................
............#...................#...........................................#..............#.....#.....................#.##.......
.................................................#................................................................................
.........................#.......................................#..................#..............................#..............
................................................................................................................................#.
...#...........................................#........#.................................................#......#................
..#...........................................#.......#.....#.................................#....#.............#...............#
........#..............#...................................................................................#......................
...........#......#............#............#..................................#.........#..............#.........#..........#....
.........................#.........#.....#.................................#..................#...............##..#...............
.....#........#.........................................................#....................................#......#....#........
...#..#.......#......#.#..................#..........................^#...........................................................
....................#.#..........................................................................#...#............................
...........................#..........................................................#...........................................
................#..............................................................................#..................................
...............................#............................................................#...................................#.
..............#..........................................#..............#..................#..........#...........................
................#...#..........................#.##..........#.......#..............................#........#..............#.....
...................#.............#....................#...........................................................................
....#....................................#....#................................................#..#...........#...................
....#.#..............#..........#...................#...................................#....................#........#......#....
...........................#...........#..#......#.......................................#.......................#......#.........
.....................................................................................................#.................#..........
...................................................#.............#..........#..............................#......................
................................#...........##....................................................................#....#..........
...........#.......#.......................................................................#............#.....................#...
.#........#.......#........................................#.#...............................#...#...........................#....
..#...............#................#....#............#................................................................#...........
........#..........#.........#....................................................................................#...............
............................................................................#.............................#...............#......#
......................#................................#.............#...#............#...........................................
................................#.............#.................................................................#..............#..
.........#.#..#...............................#....................................................#.#.......#..#........#........
.............................#.......#.....#......................................................#..#............................
..................#.............................................................................#..........#......................
.....................................#....##.................#......................................#....#........................
#.......................................................................................................................#.........
......#....#..#...........................#..#.....#........................................#..................................#..
.......#.....#...............#....#......#....#..........#.....................#....#..........#.....#..#......................##.
..................#..........#.....#........................................................#............##...................#...
........................#....#............#..................................#....#..#.......#...................................#
.....#............................#...............................#..#...................#.....#..#........#....#.#...............
......#...#..#......#..........................#........#....................#.......#...........#..........................#..#..
.....#.............................................................................................#......#......................#
...........................................#......................................................................................
.....#...........................#.#.#............................#..................#............................#.......#.......
........#.......................................................#...................#.......#.....................................
...............#..................................................#.................................................#.............
..##............................................................#..............................................................#..
.......#..................................#......................#.................................#...................#..........
......#.........#..............#....##........#....#................................#.....................................#......#
................#.....................................................................................#.............#.............
#..#..#......#.#.....#..#.................................................................................................#..#....
.....................#.....#................................................#.......#........#.................#........#.#...#...
..................................................................#....#..#...................#................................#..
........................#...........#.......................................................#..#..................................
.#.......#...................................................#..............#.......#...#.........................................
.......##...............#.............................#...........................#...............................................
....#............#....#...#........................#...........................#.#.....#..........................................
...........................#............................#..#..................#...........................#...#...................
...........#.#.............#..........#...........................................................................................
.............#..........#...............#...##.........................#..............#...#.............#..#......................
..................#.....#............#..#.........#...............................................................................
...........................................................#.....#..................#.......................#..#..................
....................................#....#.......##............#.............................#...............#...........#........
......#....................#..................#..........................................#...............................#.#......
......................#...........................................#.....#.................................................#....#..
.....................................#...............#....................................#.....#.............#......##...........
......................#...............................................................##.#..........#....#..................#.....
...........#.............#......................#....................#..................................#......#...........##.....
.....#.#...#......................................................................................................................
.........#.....................................#................#.............#....#......................#..........#............
......##.#.................#...............#..##.......................#...............#.......................................#..
....................#.........#.......#....#..................................#..........................#..............#.........
......................#....#..#..#................................#.......#.......................................................
.........................................##............................##..#..................#..................#...........#....
...................#..................................#................................................................#..........
.#.............#....................................................................................#.............................
......................................#.....#.............#..#............................................#...#............#......
.....................................#..................................#...#.......................##...................#........
...........................................................#...................................................#...........#......
.............................#.......#.......................................#.........................#.....................#...#
.....................#............................................................................................................
#....#.......................................#.........#........#...................................#.....#.........#........#....
.......................................................#......#................................##.#........#.........#..........#.
.........#........................##.....#.............................................................#.................#..#.....
....##....................................................................................................#...#...#...............
.......#.....#.......#.............#........#...............##.#......#...............#.................#..#.........#...#....#...
....................#.................................#........................................................................#..
#...............#....................................#.........#...............................#..#..........##...#...#......#....
............#........................#.....#...#.........#.........#..#......................................#....................
..........#.......#.#.....................#...........#...................#.##....................#..#...#........#..............#
.......................................#.......#........#.....#...........................................#.......#......#........
...............................#.........#..............#......................#..........................................#.#.....
............#........#.............................................#.....................................................##......#
......#.#............................................................................#......#.....#......#................#.......
.........#.............#................#.................#..###..........#.........................#.#....#......................
"""
    }
}

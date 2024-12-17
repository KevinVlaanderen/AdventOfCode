import Foundation
import Tasks
import Framework

// sourcery: cases
struct Cases {
    struct Day1 {
        struct Task1 {
            // sourcery: expected = 11
            let example1 = Tasks.Day1.task(loadData(file: Data.Input.Day1.example1), param: .task1)
            // sourcery: expected = 1151792
            let data = Tasks.Day1.task(loadData(file: Data.Input.Day1.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 31
            let example1 = Tasks.Day1.task(loadData(file: Data.Input.Day1.example1), param: .task2)
            // sourcery: expected = 21790168
            let data = Tasks.Day1.task(loadData(file: Data.Input.Day1.data), param: .task2)
        }
    }
    
    struct Day2 {
        struct Task1 {
            // sourcery: expected = 2
            let example1 = Tasks.Day2.task(loadData(file: Data.Input.Day2.example1), param: .task1)
            // sourcery: expected = 442
            let data = Tasks.Day2.task(loadData(file: Data.Input.Day2.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 4
            let example1 = Tasks.Day2.task(loadData(file: Data.Input.Day2.example1), param: .task2)
            // sourcery: expected = 493
            let data = Tasks.Day2.task(loadData(file: Data.Input.Day2.data), param: .task2)
        }
    }
    
    struct Day3 {
        struct Task1 {
            // sourcery: expected = 161
            let example1 = Tasks.Day3.task(loadData(file: Data.Input.Day3.example1), param: .task1)
            // sourcery: expected = 180233229
            let data = Tasks.Day3.task(loadData(file: Data.Input.Day3.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 48
            let example1 = Tasks.Day3.task(loadData(file: Data.Input.Day3.example2), param: .task2)
            // sourcery: expected = 95411583
            let data = Tasks.Day3.task(loadData(file: Data.Input.Day3.data), param: .task2)
        }
    }
    
    struct Day4 {
        struct Task1 {
            // sourcery: expected = 18
            let example1 = Tasks.Day4.task(loadData(file: Data.Input.Day4.example1), param: (.task1, "XMAS"))
            // sourcery: expected = 2573
            let data = Tasks.Day4.task(loadData(file: Data.Input.Day4.data), param: (.task1, "XMAS"))
        }
        
        struct Task2 {
            // sourcery: expected = 9
            let example1 = Tasks.Day4.task(loadData(file: Data.Input.Day4.example1), param: (.task2, "MAS"))
            // sourcery: expected = 1850
            let data = Tasks.Day4.task(loadData(file: Data.Input.Day4.data), param: (.task2, "MAS"))
        }
    }
    
    struct Day5 {
        struct Task1 {
            // sourcery: expected = 143
            let example1 = Tasks.Day5.task(loadData(file: Data.Input.Day5.example1), param: .task1)
            // sourcery: expected = 4959
            let data = Tasks.Day5.task(loadData(file: Data.Input.Day5.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 123
            let example1 = Tasks.Day5.task(loadData(file: Data.Input.Day5.example1), param: .task2)
            // sourcery: expected = 4655
            let data = Tasks.Day5.task(loadData(file: Data.Input.Day5.data), param: .task2)
        }
    }
    
    struct Day6 {
        struct Task1 {
            // sourcery: expected = 41
            let example1 = Tasks.Day6.task(loadData(file: Data.Input.Day6.example1), param: .task1)
            // sourcery: expected = 4973
            let data = Tasks.Day6.task(loadData(file: Data.Input.Day6.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 6
            let example1 = Tasks.Day6.task(loadData(file: Data.Input.Day6.example1), param: .task2)
            // sourcery: expected = 1482
            let data = Tasks.Day6.task(loadData(file: Data.Input.Day6.data), param: .task2)
        }
    }
    
    struct Day7 {
        struct Task1 {
            // sourcery: expected = 3749
            let example1 = Tasks.Day7.task(loadData(file: Data.Input.Day7.example1), param: [.add, .multiply])
            // sourcery: expected = 5702958180383
            let data = Tasks.Day7.task(loadData(file: Data.Input.Day7.data), param: [.add, .multiply])
        }
        
        struct Task2 {
            // sourcery: expected = 11387
            let example1 = Tasks.Day7.task(loadData(file: Data.Input.Day7.example1), param: Tasks.Day7.Operators.allCases)
            // sourcery: expected = 92612386119138
            let data = Tasks.Day7.task(loadData(file: Data.Input.Day7.data), param: Tasks.Day7.Operators.allCases)
        }
    }
    
    struct Day8 {
        struct Task1 {
            // sourcery: expected = 14
            let example1 = Tasks.Day8.task(loadData(file: Data.Input.Day8.example1), param: .task1)
            // sourcery: expected = 320
            let data = Tasks.Day8.task(loadData(file: Data.Input.Day8.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 34
            let example1 = Tasks.Day8.task(loadData(file: Data.Input.Day8.example1), param: .task2)
            // sourcery: expected = 1157
            let data = Tasks.Day8.task(loadData(file: Data.Input.Day8.data), param: .task2)
        }
    }
    
    struct Day9 {
        struct Task1 {
            // sourcery: expected = 1928
            let example1 = Tasks.Day9.task(loadData(file: Data.Input.Day9.example1), param: .task1)
            // sourcery: expected = 6356833654075
            let data = Tasks.Day9.task(loadData(file: Data.Input.Day9.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 2858
            let example1 = Tasks.Day9.task(loadData(file: Data.Input.Day9.example1), param: .task2)
            // sourcery: expected = 6389911791746
            let data = Tasks.Day9.task(loadData(file: Data.Input.Day9.data), param: .task2)
        }
    }
    
    struct Day10 {
        struct Task1 {
            // sourcery: expected = 1
            let example1 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example1), param: .task1)
            // sourcery: expected = 2
            let example2 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example2), param: .task1)
            // sourcery: expected = 4
            let example3 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example3), param: .task1)
            // sourcery: expected = 3
            let example4 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example4), param: .task1)
            // sourcery: expected = 36
            let example5 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example5), param: .task1)
            // sourcery: expected = 816
            let data = Tasks.Day10.task(loadData(file: Data.Input.Day10.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 3
            let example6 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example6), param: .task2)
            // sourcery: expected = 13
            let example7 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example7), param: .task2)
            // sourcery: expected = 227
            let example8 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example8), param: .task2)
            // sourcery: expected = 81
            let example5 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example5), param: .task2)
            // sourcery: expected = 1960
            let data = Tasks.Day10.task(loadData(file: Data.Input.Day10.data), param: .task2)
        }
    }
    
    struct Day11 {
        struct Task1 {
            // sourcery: expected = 55312
            let example1 = Tasks.Day11.task(loadData(file: Data.Input.Day11.example1), param: 25)
            // sourcery: expected = 213625
            let data = Tasks.Day11.task(loadData(file: Data.Input.Day11.data), param: 25)
        }
        
        struct Task2 {
            // sourcery: expected = 252442982856820
            let data = Tasks.Day11.task(loadData(file: Data.Input.Day11.data), param: 75)
        }
    }
    
    struct Day12 {
        struct Task1 {
            // sourcery: expected = 140
            let example1 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example1), param: .task1)
            // sourcery: expected = 772
            let example2 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example2), param: .task1)
            // sourcery: expected = 1930
            let example3 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example3), param: .task1)
            // sourcery: expected = 1457298
            let data = Tasks.Day12.task(loadData(file: Data.Input.Day12.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 80
            let example1 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example1), param: .task2)
            // sourcery: expected = 436
            let example2 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example2), param: .task2)
            // sourcery: expected = 236
            let example4 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example4), param: .task2)
            // sourcery: expected = 368
            let example5 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example5), param: .task2)
            // sourcery: expected = 921636
            let data = Tasks.Day12.task(loadData(file: Data.Input.Day12.data), param: .task2)
        }
    }
    
    struct Day13 {
        struct Task1 {
            // sourcery: expected = 480
            let example1 = Tasks.Day13.task(loadData(file: Data.Input.Day13.example1), param: (0, 100))
            // sourcery: expected = 35729
            let data = Tasks.Day13.task(loadData(file: Data.Input.Day13.data), param: (0, 100))
        }
        
        struct Task2 {
            // sourcery: expected = 875318608908
            let example1 = Tasks.Day13.task(loadData(file: Data.Input.Day13.example1), param: (10_000_000_000_000, -1))
            // sourcery: expected = 88584689879723
            let data = Tasks.Day13.task(loadData(file: Data.Input.Day13.data), param: (10_000_000_000_000, -1))
        }
    }
    
    struct Day14 {
        struct Task1 {
            // sourcery: expected = 12
            let example1 = Tasks.Day14.task(loadData(file: Data.Input.Day14.example1), param: (.task1, false))
            // sourcery: expected = 208437768
            let data = Tasks.Day14.task(loadData(file: Data.Input.Day14.data), param: (.task1, false))
        }
        
        struct Task2 {
            // sourcery: expected = 7492
            let data = Tasks.Day14.task(loadData(file: Data.Input.Day14.data), param: (.task2, false))
        }
    }
    
    struct Day15 {
        struct Task1 {
            // sourcery: expected = 10092
            let example1 = Tasks.Day15.task(loadData(file: Data.Input.Day15.example1), param: 1)
            // sourcery: expected = 2028
            let example2 = Tasks.Day15.task(loadData(file: Data.Input.Day15.example2), param: 1)
            // sourcery: expected = 1412971
            let data = Tasks.Day15.task(loadData(file: Data.Input.Day15.data), param: 1)
        }
        
        struct Task2 {
            // sourcery: expected = 9021
            let example1 = Tasks.Day15.task(loadData(file: Data.Input.Day15.example1), param: 2)
            // sourcery: expected = 1429299
            let data = Tasks.Day15.task(loadData(file: Data.Input.Day15.data), param: 2)
        }
    }
    
    struct Day16 {
        struct Task1 {
            // sourcery: expected = 7036
            let example1 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example1), param: .task1)
            // sourcery: expected = 11048
            let example2 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example2), param: .task1)
            // sourcery: expected = 143580
            let data = Tasks.Day16.task(loadData(file: Data.Input.Day16.data), param: .task1)
        }
        
        struct Task2 {
            // sourcery: expected = 45
            let example1 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example1), param: .task2)
            // sourcery: expected = 64
            let example2 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example2), param: .task2)
            // sourcery: expected = 645
            let data = Tasks.Day16.task(loadData(file: Data.Input.Day16.data), param: .task2)
        }
    }
}

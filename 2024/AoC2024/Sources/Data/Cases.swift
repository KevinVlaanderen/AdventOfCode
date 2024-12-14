import Foundation
import Tasks
import Framework

// sourcery: cases
struct Cases: Sendable {
    struct Day1 {
        let day = Tasks.Day1()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day1.P = ()
            
            // sourcery: expected = 11
            let example1 = loadData(file: Data.Input.Day1.example1)
            // sourcery: expected = 1151792
            let data = loadData(file: Data.Input.Day1.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day1.P = ()
            
            // sourcery: expected = 31
            let example1 = loadData(file: Data.Input.Day1.example1)
            // sourcery: expected = 21790168
            let data = loadData(file: Data.Input.Day1.data)
        }
    }
    
    struct Day2 {
        let day = Tasks.Day2()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day2.P = ()
            
            // sourcery: expected = 2
            let example1 = loadData(file: Data.Input.Day2.example1)
            // sourcery: expected = 442
            let data = loadData(file: Data.Input.Day2.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day2.P = ()
            
            // sourcery: expected = 4
            let example1 = loadData(file: Data.Input.Day2.example1)
            // sourcery: expected = 493
            let data = loadData(file: Data.Input.Day2.data)
        }
    }
    
    struct Day3 {
        let day = Tasks.Day3()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day3.P = ()
            
            // sourcery: expected = 161
            let example1 = loadData(file: Data.Input.Day3.example1)
            // sourcery: expected = 180233229
            let data = loadData(file: Data.Input.Day3.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day3.P = ()
            
            // sourcery: expected = 48
            let example1 = loadData(file: Data.Input.Day3.example2)
            // sourcery: expected = 95411583
            let data = loadData(file: Data.Input.Day3.data)
        }
    }
    
    struct Day4 {
        let day = Tasks.Day4()
        
        struct Task1 {
            let task = Task.task1
            let param = "XMAS"
            
            // sourcery: expected = 18
            let example1 = loadData(file: Data.Input.Day4.example1)
            // sourcery: expected = 2573
            let data = loadData(file: Data.Input.Day4.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param = "MAS"
            
            // sourcery: expected = 9
            let example1 = loadData(file: Data.Input.Day4.example1)
            // sourcery: expected = 1850
            let data = loadData(file: Data.Input.Day4.data)
        }
    }
    
    struct Day5 {
        let day = Tasks.Day5()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day5.P = ()
            
            // sourcery: expected = 143
            let example1 = loadData(file: Data.Input.Day5.example1)
            // sourcery: expected = 4959
            let data = loadData(file: Data.Input.Day5.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day5.P = ()
            
            // sourcery: expected = 123
            let example1 = loadData(file: Data.Input.Day5.example1)
            // sourcery: expected = 4655
            let data = loadData(file: Data.Input.Day5.data)
        }
    }
    
    struct Day6 {
        let day = Tasks.Day6()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day6.P = ()
            
            // sourcery: expected = 41
            let example1 = loadData(file: Data.Input.Day6.example1)
            // sourcery: expected = 4973
            let data = loadData(file: Data.Input.Day6.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day6.P = ()
            
            // sourcery: expected = 6
            let example1 = loadData(file: Data.Input.Day6.example1)
            // sourcery: expected = 1482
            let data = loadData(file: Data.Input.Day6.data)
        }
    }
    
    struct Day7 {
        let day = Tasks.Day7()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day7.P = [.add, .multiply]
            
            // sourcery: expected = 3749
            let example1 = loadData(file: Data.Input.Day7.example1)
            // sourcery: expected = 5702958180383
            let data = loadData(file: Data.Input.Day7.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param = Tasks.Day7.Operators.allCases
            
            // sourcery: expected = 11387
            let example1 = loadData(file: Data.Input.Day7.example1)
            // sourcery: expected = 92612386119138
            let data = loadData(file: Data.Input.Day7.data)
        }
    }
    
    struct Day8 {
        let day = Tasks.Day8()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day8.P = ()
            
            // sourcery: expected = 14
            let example1 = loadData(file: Data.Input.Day8.example1)
            // sourcery: expected = 320
            let data = loadData(file: Data.Input.Day8.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day8.P = ()
            
            // sourcery: expected = 34
            let example1 = loadData(file: Data.Input.Day8.example1)
            // sourcery: expected = 1157
            let data = loadData(file: Data.Input.Day8.data)
        }
    }
    
    struct Day9 {
        let day = Tasks.Day9()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day9.P = ()
            
            // sourcery: expected = 1928
            let example1 = loadData(file: Data.Input.Day9.example1)
            // sourcery: expected = 6356833654075
            let data = loadData(file: Data.Input.Day9.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day9.P = ()
            
            // sourcery: expected = 2858
            let example1 = loadData(file: Data.Input.Day9.example1)
            // sourcery: expected = 6389911791746
            let data = loadData(file: Data.Input.Day9.data)
        }
    }
    
    struct Day10 {
        let day = Tasks.Day10()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day10.P = ()
            
            // sourcery: expected = 1
            let example1 = loadData(file: Data.Input.Day10.example1)
            // sourcery: expected = 2
            let example2 = loadData(file: Data.Input.Day10.example2)
            // sourcery: expected = 4
            let example3 = loadData(file: Data.Input.Day10.example3)
            // sourcery: expected = 3
            let example4 = loadData(file: Data.Input.Day10.example4)
            // sourcery: expected = 36
            let example5 = loadData(file: Data.Input.Day10.example5)
            // sourcery: expected = 816
            let data = loadData(file: Data.Input.Day10.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day10.P = ()
            
            // sourcery: expected = 3
            let example6 = loadData(file: Data.Input.Day10.example6)
            // sourcery: expected = 13
            let example7 = loadData(file: Data.Input.Day10.example7)
            // sourcery: expected = 227
            let example8 = loadData(file: Data.Input.Day10.example8)
            // sourcery: expected = 81
            let example5 = loadData(file: Data.Input.Day10.example5)
            // sourcery: expected = 1960
            let data = loadData(file: Data.Input.Day10.data)
        }
    }
    
    struct Day11 {
        let day = Tasks.Day11()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day11.P = 25
            
            // sourcery: expected = 55312
            let example1 = loadData(file: Data.Input.Day11.example1)
            // sourcery: expected = 213625
            let data = loadData(file: Data.Input.Day11.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day11.P = 75
            
            // sourcery: expected = 252442982856820
            let data = loadData(file: Data.Input.Day11.data)
        }
    }
    
    struct Day12 {
        let day = Tasks.Day12()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day12.P = ()
            
            // sourcery: expected = 140
            let example1 = loadData(file: Data.Input.Day12.example1)
            // sourcery: expected = 772
            let example2 = loadData(file: Data.Input.Day12.example2)
            // sourcery: expected = 1930
            let example3 = loadData(file: Data.Input.Day12.example3)
            // sourcery: expected = 1457298
            let data = loadData(file: Data.Input.Day12.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day12.P = ()
            
            // sourcery: expected = 80
            let example1 = loadData(file: Data.Input.Day12.example1)
            // sourcery: expected = 436
            let example2 = loadData(file: Data.Input.Day12.example2)
            // sourcery: expected = 236
            let example4 = loadData(file: Data.Input.Day12.example4)
            // sourcery: expected = 368
            let example5 = loadData(file: Data.Input.Day12.example5)
            // sourcery: expected = 921636
            let data = loadData(file: Data.Input.Day12.data)
        }
    }
    
    struct Day13 {
        let day = Tasks.Day13()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day13.P = (0, 100)
            
            // sourcery: expected = 480
            let example1 = loadData(file: Data.Input.Day13.example1)
            // sourcery: expected = 35729
            let data = loadData(file: Data.Input.Day13.data)
        }
        
        struct Task2 {
            let task = Task.task2
            let param: Tasks.Day13.P = (10_000_000_000_000, -1)
            
            // sourcery: expected = 875318608908
            let example1 = loadData(file: Data.Input.Day13.example1)
            // sourcery: expected = 88584689879723
            let data = loadData(file: Data.Input.Day13.data)
        }
    }
    
    struct Day14 {
        let day = Tasks.Day14()
        
        struct Task1 {
            let task = Task.task1
            let param: Tasks.Day14.P = 100
            
            // sourcery: expected = 12
            let example1 = loadData(file: Data.Input.Day14.example1)
            // sourcery: expected = 208437768
            let data = loadData(file: Data.Input.Day14.data)
        }
    }
}

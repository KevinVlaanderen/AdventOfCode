import Foundation
import Tasks
import Framework

// sourcery: cases
struct Cases {
    struct Day1 {
        struct Task1 {
            // sourcery: expected = 11
            let example1 = Tasks.Day1.task(loadData(file: Data.Input.Day1.example1), task: .task1, param: ())
            // sourcery: expected = 1151792
            let data = Tasks.Day1.task(loadData(file: Data.Input.Day1.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 31
            let example1 = Tasks.Day1.task(loadData(file: Data.Input.Day1.example1), task: .task2, param: ())
            // sourcery: expected = 21790168
            let data = Tasks.Day1.task(loadData(file: Data.Input.Day1.data), task: .task2, param: ())
        }
    }
    
    struct Day2 {
        struct Task1 {
            // sourcery: expected = 2
            let example1 = Tasks.Day2.task(loadData(file: Data.Input.Day2.example1), task: .task1, param: ())
            // sourcery: expected = 442
            let data = Tasks.Day2.task(loadData(file: Data.Input.Day2.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 4
            let example1 = Tasks.Day2.task(loadData(file: Data.Input.Day2.example1), task: .task2, param: ())
            // sourcery: expected = 493
            let data = Tasks.Day2.task(loadData(file: Data.Input.Day2.data), task: .task2, param: ())
        }
    }
    
    struct Day3 {
        struct Task1 {
            // sourcery: expected = 161
            let example1 = Tasks.Day3.task(loadData(file: Data.Input.Day3.example1), task: .task1, param: [Tasks.Day3.InstructionParam.mulInst])
            // sourcery: expected = 180233229
            let data = Tasks.Day3.task(loadData(file: Data.Input.Day3.data), task: .task1, param: [Tasks.Day3.InstructionParam.mulInst])
        }
        
        struct Task2 {
            // sourcery: expected = 48
            let example1 = Tasks.Day3.task(loadData(file: Data.Input.Day3.example2), task: .task2, param: Tasks.Day3.InstructionParam.allCases)
            // sourcery: expected = 95411583
            let data = Tasks.Day3.task(loadData(file: Data.Input.Day3.data), task: .task2, param: Tasks.Day3.InstructionParam.allCases)
        }
    }
    
    struct Day4 {
        struct Task1 {
            // sourcery: expected = 18
            let example1 = Tasks.Day4.task(loadData(file: Data.Input.Day4.example1), task: .task1, param: "XMAS")
            // sourcery: expected = 2573
            let data = Tasks.Day4.task(loadData(file: Data.Input.Day4.data), task: .task1, param: "XMAS")
        }
        
        struct Task2 {
            // sourcery: expected = 9
            let example1 = Tasks.Day4.task(loadData(file: Data.Input.Day4.example1), task: .task2, param: "MAS")
            // sourcery: expected = 1850
            let data = Tasks.Day4.task(loadData(file: Data.Input.Day4.data), task: .task2, param: "MAS")
        }
    }
    
    struct Day5 {
        struct Task1 {
            // sourcery: expected = 143
            let example1 = Tasks.Day5.task(loadData(file: Data.Input.Day5.example1), task: .task1, param: ())
            // sourcery: expected = 4959
            let data = Tasks.Day5.task(loadData(file: Data.Input.Day5.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 123
            let example1 = Tasks.Day5.task(loadData(file: Data.Input.Day5.example1), task: .task2, param: ())
            // sourcery: expected = 4655
            let data = Tasks.Day5.task(loadData(file: Data.Input.Day5.data), task: .task2, param: ())
        }
    }
    
    struct Day6 {
        struct Task1 {
            // sourcery: expected = 41
            let example1 = Tasks.Day6.task(loadData(file: Data.Input.Day6.example1), task: .task1, param: ())
            // sourcery: expected = 4973
            let data = Tasks.Day6.task(loadData(file: Data.Input.Day6.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 6
            let example1 = Tasks.Day6.task(loadData(file: Data.Input.Day6.example1), task: .task2, param: ())
            // sourcery: expected = 1482
            let data = Tasks.Day6.task(loadData(file: Data.Input.Day6.data), task: .task2, param: ())
        }
    }
    
    struct Day7 {
        struct Task1 {
            // sourcery: expected = 3749
            let example1 = Tasks.Day7.task(loadData(file: Data.Input.Day7.example1), task: .task1, param: [.add, .multiply])
            // sourcery: expected = 5702958180383
            let data = Tasks.Day7.task(loadData(file: Data.Input.Day7.data), task: .task1, param: [.add, .multiply])
        }
        
        struct Task2 {
            // sourcery: expected = 11387
            let example1 = Tasks.Day7.task(loadData(file: Data.Input.Day7.example1), task: .task2, param: Tasks.Day7.OperatorParam.allCases)
            // sourcery: expected = 92612386119138
            let data = Tasks.Day7.task(loadData(file: Data.Input.Day7.data), task: .task2, param: Tasks.Day7.OperatorParam.allCases)
        }
    }
    
    struct Day8 {
        struct Task1 {
            // sourcery: expected = 14
            let example1 = Tasks.Day8.task(loadData(file: Data.Input.Day8.example1), task: .task1, param: ())
            // sourcery: expected = 320
            let data = Tasks.Day8.task(loadData(file: Data.Input.Day8.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 34
            let example1 = Tasks.Day8.task(loadData(file: Data.Input.Day8.example1), task: .task2, param: ())
            // sourcery: expected = 1157
            let data = Tasks.Day8.task(loadData(file: Data.Input.Day8.data), task: .task2, param: ())
        }
    }
    
    struct Day9 {
        struct Task1 {
            // sourcery: expected = 1928
            let example1 = Tasks.Day9.task(loadData(file: Data.Input.Day9.example1), task: .task1, param: ())
            // sourcery: expected = 6356833654075
            let data = Tasks.Day9.task(loadData(file: Data.Input.Day9.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 2858
            let example1 = Tasks.Day9.task(loadData(file: Data.Input.Day9.example1), task: .task2, param: ())
            // sourcery: expected = 6389911791746
            let data = Tasks.Day9.task(loadData(file: Data.Input.Day9.data), task: .task2, param: ())
        }
    }
    
    struct Day10 {
        struct Task1 {
            // sourcery: expected = 1
            let example1 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example1), task: .task1, param: ())
            // sourcery: expected = 2
            let example2 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example2), task: .task1, param: ())
            // sourcery: expected = 4
            let example3 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example3), task: .task1, param: ())
            // sourcery: expected = 3
            let example4 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example4), task: .task1, param: ())
            // sourcery: expected = 36
            let example5 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example5), task: .task1, param: ())
            // sourcery: expected = 816
            let data = Tasks.Day10.task(loadData(file: Data.Input.Day10.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 3
            let example6 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example6), task: .task2, param: ())
            // sourcery: expected = 13
            let example7 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example7), task: .task2, param: ())
            // sourcery: expected = 227
            let example8 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example8), task: .task2, param: ())
            // sourcery: expected = 81
            let example5 = Tasks.Day10.task(loadData(file: Data.Input.Day10.example5), task: .task2, param: ())
            // sourcery: expected = 1960
            let data = Tasks.Day10.task(loadData(file: Data.Input.Day10.data), task: .task2, param: ())
        }
    }
    
    struct Day11 {
        struct Task1 {
            // sourcery: expected = 55312
            let example1 = Tasks.Day11.task(loadData(file: Data.Input.Day11.example1), task: .task1, param: 25)
            // sourcery: expected = 213625
            let data = Tasks.Day11.task(loadData(file: Data.Input.Day11.data), task: .task1, param: 25)
        }
        
        struct Task2 {
            // sourcery: expected = 252442982856820
            let data = Tasks.Day11.task(loadData(file: Data.Input.Day11.data), task: .task2, param: 75)
        }
    }
    
    struct Day12 {
        struct Task1 {
            // sourcery: expected = 140
            let example1 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example1), task: .task1, param: ())
            // sourcery: expected = 772
            let example2 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example2), task: .task1, param: ())
            // sourcery: expected = 1930
            let example3 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example3), task: .task1, param: ())
            // sourcery: expected = 1457298
            let data = Tasks.Day12.task(loadData(file: Data.Input.Day12.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 80
            let example1 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example1), task: .task2, param: ())
            // sourcery: expected = 436
            let example2 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example2), task: .task2, param: ())
            // sourcery: expected = 236
            let example4 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example4), task: .task2, param: ())
            // sourcery: expected = 368
            let example5 = Tasks.Day12.task(loadData(file: Data.Input.Day12.example5), task: .task2, param: ())
            // sourcery: expected = 921636
            let data = Tasks.Day12.task(loadData(file: Data.Input.Day12.data), task: .task2, param: ())
        }
    }
    
    struct Day13 {
        struct Task1 {
            // sourcery: expected = 480
            let example1 = Tasks.Day13.task(loadData(file: Data.Input.Day13.example1), task: .task1, param: (0, 100))
            // sourcery: expected = 35729
            let data = Tasks.Day13.task(loadData(file: Data.Input.Day13.data), task: .task1, param: (0, 100))
        }
        
        struct Task2 {
            // sourcery: expected = 875318608908
            let example1 = Tasks.Day13.task(loadData(file: Data.Input.Day13.example1), task: .task2, param: (10_000_000_000_000, -1))
            // sourcery: expected = 88584689879723
            let data = Tasks.Day13.task(loadData(file: Data.Input.Day13.data), task: .task2, param: (10_000_000_000_000, -1))
        }
    }
    
    struct Day14 {
        struct Task1 {
            // sourcery: expected = 12
            let example1 = Tasks.Day14.task(loadData(file: Data.Input.Day14.example1), task: .task1, param: false)
            // sourcery: expected = 208437768
            let data = Tasks.Day14.task(loadData(file: Data.Input.Day14.data), task: .task1, param: false)
        }
        
        struct Task2 {
            // sourcery: expected = 7492
            let data = Tasks.Day14.task(loadData(file: Data.Input.Day14.data), task: .task2, param: false)
        }
    }
    
    struct Day15 {
        struct Task1 {
            // sourcery: expected = 10092
            let example1 = Tasks.Day15.task(loadData(file: Data.Input.Day15.example1), task: .task1, param: 1)
            // sourcery: expected = 2028
            let example2 = Tasks.Day15.task(loadData(file: Data.Input.Day15.example2), task: .task1, param: 1)
            // sourcery: expected = 1412971
            let data = Tasks.Day15.task(loadData(file: Data.Input.Day15.data), task: .task1, param: 1)
        }
        
        struct Task2 {
            // sourcery: expected = 9021
            let example1 = Tasks.Day15.task(loadData(file: Data.Input.Day15.example1), task: .task2, param: 2)
            // sourcery: expected = 1429299
            let data = Tasks.Day15.task(loadData(file: Data.Input.Day15.data), task: .task2, param: 2)
        }
    }
    
    struct Day16 {
        struct Task1 {
            // sourcery: expected = 7036
            let example1 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example1), task: .task1, param: ())
            // sourcery: expected = 11048
            let example2 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example2), task: .task1, param: ())
            // sourcery: expected = 143580
            let data = Tasks.Day16.task(loadData(file: Data.Input.Day16.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 45
            let example1 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example1), task: .task2, param: ())
            // sourcery: expected = 64
            let example2 = Tasks.Day16.task(loadData(file: Data.Input.Day16.example2), task: .task2, param: ())
            // sourcery: expected = 645
            let data = Tasks.Day16.task(loadData(file: Data.Input.Day16.data), task: .task2, param: ())
        }
    }
    
    struct Day17 {
        struct Task1 {
            // sourcery: expected = "4,6,3,5,6,3,5,2,1,0"
            let example1 = Tasks.Day17.task(loadData(file: Data.Input.Day17.example1), task: .task1, param: ())
            // sourcery: expected = "7,3,1,3,6,3,6,0,2"
            let data = Tasks.Day17.task(loadData(file: Data.Input.Day17.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = "117440"
            let example2 = Tasks.Day17.task(loadData(file: Data.Input.Day17.example2), task: .task2, param: ())
            // sourcery: expected = "105843716614554"
            let data = Tasks.Day17.task(loadData(file: Data.Input.Day17.data), task: .task2, param: ())
        }
    }
    
    struct Day18 {
        struct Task1 {
            // sourcery: expected = "22"
            let example1 = Tasks.Day18.task(loadData(file: Data.Input.Day18.example1), task: .task1, param: (7, 12))
            // sourcery: expected = "270"
            let data = Tasks.Day18.task(loadData(file: Data.Input.Day18.data), task: .task1, param: (71, 1024))
        }
        
        struct Task2 {
            // sourcery: expected = "6,1"
            let example1 = Tasks.Day18.task(loadData(file: Data.Input.Day18.example1), task: .task2, param: (7, 12))
            // sourcery: expected = "51,40"
            let data = Tasks.Day18.task(loadData(file: Data.Input.Day18.data), task: .task2, param: (71, 1024))
        }
    }
    
    struct Day19 {
        struct Task1 {
            // sourcery: expected = 6
            let example1 = Tasks.Day19.task(loadData(file: Data.Input.Day19.example1), task: .task1, param: ())
            // sourcery: expected = 220
            let data = Tasks.Day19.task(loadData(file: Data.Input.Day19.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 16
            let example1 = Tasks.Day19.task(loadData(file: Data.Input.Day19.example1), task: .task2, param: ())
            // sourcery: expected = 565600047715343
            let data = Tasks.Day19.task(loadData(file: Data.Input.Day19.data), task: .task2, param: ())
        }
    }
    
    struct Day20 {
        struct Task1 {
            // sourcery: expected = 5
            let example1 = Tasks.Day20.task(loadData(file: Data.Input.Day20.example1), task: .task1, param: (2, 20))
            // sourcery: expected = 1429
            let data = Tasks.Day20.task(loadData(file: Data.Input.Day20.data), task: .task1, param: (2, 100))
        }
        
        struct Task2 {
            // sourcery: expected = 41
            let example1 = Tasks.Day20.task(loadData(file: Data.Input.Day20.example1), task: .task2, param: (20, 70))
            // sourcery: expected = 988931
            let data = Tasks.Day20.task(loadData(file: Data.Input.Day20.data), task: .task2, param: (20, 100))
        }
    }
    
    struct Day21 {
        struct Task1 {
            // sourcery: expected = 126384
            let example1 = Tasks.Day21.task(loadData(file: Data.Input.Day21.example1), task: .task1, param: 2)
            // sourcery: expected = 134120
            let data = Tasks.Day21.task(loadData(file: Data.Input.Day21.data), task: .task1, param: 2)
        }
        
        struct Task2 {
            // sourcery: expected = 167389793580400
            let data = Tasks.Day21.task(loadData(file: Data.Input.Day21.data), task: .task2, param: 25)
        }
    }
    
    struct Day22 {
        struct Task1 {
            // sourcery: expected = 37327623
            let example1 = Tasks.Day22.task(loadData(file: Data.Input.Day22.example1), task: .task1, param: ())
            // sourcery: expected = 15006633487
            let data = Tasks.Day22.task(loadData(file: Data.Input.Day22.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = 23
            let example2 = Tasks.Day22.task(loadData(file: Data.Input.Day22.example2), task: .task2, param: ())
            // sourcery: expected = 1710
            let data = Tasks.Day22.task(loadData(file: Data.Input.Day22.data), task: .task2, param: ())
        }
    }
    
    struct Day23 {
        struct Task1 {
            // sourcery: expected = "7"
            let example1 = Tasks.Day23.task(loadData(file: Data.Input.Day23.example1), task: .task1, param: ())
            // sourcery: expected = "1352"
            let data = Tasks.Day23.task(loadData(file: Data.Input.Day23.data), task: .task1, param: ())
        }
        
        struct Task2 {
            // sourcery: expected = "co,de,ka,ta"
            let example1 = Tasks.Day23.task(loadData(file: Data.Input.Day23.example1), task: .task2, param: ())
            // sourcery: expected = "dm,do,fr,gf,gh,gy,iq,jb,kt,on,rg,xf,ze"
            let data = Tasks.Day23.task(loadData(file: Data.Input.Day23.data), task: .task2, param: ())
        }
    }
    
    struct Day24 {
        struct Task1 {
            // sourcery: expected = "4"
            let example1 = Tasks.Day24.task(loadData(file: Data.Input.Day24.example1), task: .task1, param: true)
            // sourcery: expected = "2024"
            let example2 = Tasks.Day24.task(loadData(file: Data.Input.Day24.example2), task: .task1, param: true)
            // sourcery: expected = "60714423975686"
            let data = Tasks.Day24.task(loadData(file: Data.Input.Day24.data), task: .task1, param: false)
        }
        
        struct Task2 {
            // sourcery: expected = "z00,z01,z02,z05"
            let example3 = Tasks.Day24.task(loadData(file: Data.Input.Day24.example3), task: .task2, param: true)
            // sourcery: expected = "cgh,frt,pmd,sps,tst,z05,z11,z23"
            let data = Tasks.Day24.task(loadData(file: Data.Input.Day24.data), task: .task2, param: false)
        }
    }
    
    struct Day25 {
        struct Task1 {
            // sourcery: expected = 3
            let example1 = Tasks.Day25.task(loadData(file: Data.Input.Day25.example1), task: .task1, param: ())
            // sourcery: expected = 3269
            let data = Tasks.Day25.task(loadData(file: Data.Input.Day25.data), task: .task1, param: ())
        }
    }
}

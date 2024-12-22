import Foundation
internal import Algorithms
import Framework

public final class Day4: Day<(task: Task, word: String), Int> {
    public override func perform(data: String, param: P) throws -> R {
        let grid = parse(data)

        return switch param.task {
        case .task1: try task1(grid: grid, word: param.word)
        case .task2: try task2(grid: grid, word: param.word)
        }
    }
    
    private func parse(_ data: String) -> some Grid<Character> {
        ArrayGrid(data.lines.map(\.characters))
    }
    
    private func task1(grid: any Grid<Character>, word: String) throws -> Int {
        guard let firstCharacter = word.first else {
            throw AoCError.taskError("word is empty")
        }
        
        var count = 0
        
        for item in grid {
            if item.value != firstCharacter {
                continue
            }
            
            for heading in Heading.allCases {
                if wordFound(grid: grid, word: word, position: item.position, heading: heading) {
                    count += 1
                }
            }
        }

        return count
    }
    
    private func task2(grid: any Grid<Character>, word: String) throws -> Int {
        if word.count.isEven {
            throw AoCError.taskError("word should have an odd length")
        }
        guard let firstCharacter = word.first else {
            throw AoCError.taskError("word is empty")
        }
        
        var crossed: [Point: Int] = [:]
        
        for item in grid {
            if item.value != firstCharacter {
                continue
            }
            
            for heading in [Heading.northEast, Heading.southEast, Heading.southWest, Heading.northWest] {
                if wordFound(grid: grid, word: word, position: item.position, heading: heading) {
                    let center = item.position.neighbour(heading: heading, distance: word.count / 2)
                    let current = crossed[center] ?? 0
                    crossed[center] = current+1
                }
            }
        }

        return crossed.count(where: { $0.value >= 2 })
    }
    
    private func wordFound(grid: any Grid<Character>, word: String, position: Point, heading: Heading) -> Bool {
        guard let currentCharacter = word.first else {
            return true
        }
        guard let gridCharacter = grid[position] else {
            return false
        }
        
        if currentCharacter != gridCharacter {
            return false
        }
        
        let newWord = String(word.dropFirst())
        let newPosition = position.neighbour(heading: heading)
        
        return wordFound(grid: grid, word: newWord, position: newPosition, heading: heading)
    }
}

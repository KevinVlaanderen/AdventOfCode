import Foundation
import Framework
internal import Algorithms

public struct Day9: Day {
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        var disk = parse(data)

        switch task {
        case .task1:
            task1(&disk)
        case .task2:
            task2(&disk)
        }
        
        return disk.checksum()
    }
    
    private func task1(_ disk: inout Disk) {
        var endIndex = disk.endIndex-1
        
        for startIndex in 0...endIndex {
            if case .file(_) = disk[startIndex] {
                continue
            }
            
            while case .emptySpace = disk[endIndex] {
                endIndex -= 1
            }
            
            if endIndex <= startIndex {
                break
            }
            
            disk[startIndex] = disk[endIndex]
            disk[endIndex] = .emptySpace
        }
    }
    
    private func task2(_ disk: inout Disk) {
        var currentContent: Content = .emptySpace
        var currentLength = 0
        
        let originalDisk = disk
        
        for endIndex in (0..<originalDisk.endIndex).reversed() {
            let content = originalDisk[endIndex]
            
            if currentContent == content {
                currentLength += 1
                continue
            }
            
            if currentContent == .emptySpace {
                currentContent = content
                currentLength = 1
                continue
            }
            
            guard case let .file(id) = currentContent else {
                fatalError()
            }
            
            var startIndex = 0
            var emptySpaceLength = 0
            while startIndex <= endIndex {
                if case .file(_) = disk[startIndex] {
                    emptySpaceLength = 0
                    startIndex += 1
                    continue
                }
                
                emptySpaceLength += 1
                if emptySpaceLength == currentLength {
                    for i in startIndex-currentLength+1...startIndex {
                        disk[i] = .file(id)
                    }
                    for i in endIndex+1...endIndex+currentLength {
                        disk[i] = .emptySpace
                    }
                    break
                } else {
                    startIndex += 1
                }
            }
            
            currentContent = content
            currentLength = 1
        }
    }

    private func parse(_ data: String) -> Disk {
        data.enumerated().flatMap { offset, character in
            let id = offset/2
            let size = character.wholeNumberValue!
            let content: Content = offset%2 == 0 ? .file(id) : .emptySpace
            return Disk(repeating: content, count: size)
        }
    }
    
    typealias Disk = [Content]
    
    enum Content: Equatable {
        case file(Int), emptySpace
    }
}

extension Day9.Disk {
    func checksum() -> Int {
        self.enumerated().reduce(0) { result, content in
            if case let .file(id) = content.element {
                return result + content.offset*id
            }
            return result
        }
    }
}

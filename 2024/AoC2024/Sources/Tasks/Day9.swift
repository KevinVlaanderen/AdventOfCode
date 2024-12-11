import Foundation
import Framework
internal import Algorithms

public struct Day9: Day {
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        var disk = parse(data)

        switch task {
        case .task1:
            return checksum(files: task1(&disk))
        case .task2:
            return checksum(files: task2(&disk))
        }
    }
    
    private func task1(_ disk: inout Disk) -> [Content] {
        var files: [Content] = []
        
        var emptySpaceIndex = 0
        var fileIndex = disk.files.endIndex - 1

        var emptySpace: Content? = disk.emptySpaces[emptySpaceIndex]
        var file: Content? = disk.files[fileIndex]
        
        let emptySpaceCount = disk.emptySpaces.count
        
        func increaseEmptySpace() {
            emptySpaceIndex += 1
            emptySpace = emptySpaceIndex < emptySpaceCount ? disk.emptySpaces[emptySpaceIndex] : nil
            file = fileIndex >= 0 ? disk.files[fileIndex] : nil
        }
        
        func increaseFile() {
            fileIndex -= 1
            file = fileIndex >= 0 ? disk.files[fileIndex] : nil
            emptySpace = emptySpaceIndex < emptySpaceCount ? disk.emptySpaces[emptySpaceIndex] : nil
        }
        
        while var localEmptySpace = emptySpace, let file = file {
            if localEmptySpace.position >= file.position {
                break
            }
            
            if localEmptySpace.size > file.size {
                files.append(Content(position: localEmptySpace.position, size: file.size, type: file.type))
                localEmptySpace.size -= file.size
                localEmptySpace.position += file.size
                disk.files[fileIndex].size = 0

                emptySpace = localEmptySpace
                disk.emptySpaces[emptySpaceIndex] = localEmptySpace
                
                increaseFile()
            } else if localEmptySpace.size < file.size {
                files.append(Content(position: localEmptySpace.position, size: localEmptySpace.size, type: file.type))
                disk.files[fileIndex].size -= localEmptySpace.size
                localEmptySpace.size = 0
                
                emptySpace = localEmptySpace
                disk.emptySpaces[emptySpaceIndex] = localEmptySpace
                
                increaseEmptySpace()
            } else {
                files.append(Content(position: localEmptySpace.position, size: localEmptySpace.size, type: file.type))
                localEmptySpace.size = 0
                disk.files[fileIndex].size = 0
                
                emptySpace = localEmptySpace
                disk.emptySpaces[emptySpaceIndex] = localEmptySpace
                
                increaseEmptySpace()
                increaseFile()
            }
        }

        return files + disk.files.filter({ $0.type != .emptySpace && $0.size > 0 })
    }
    
    private func task2(_ disk: inout Disk) -> [Content] {
        var files: [Content] = []
        
        var emptySpaceIndex = 0
        var fileIndex = disk.files.endIndex - 1
        
        var emptySpace: Content? = disk.emptySpaces[emptySpaceIndex]
        var file: Content? = disk.files[fileIndex]

        let emptySpaceCount = disk.emptySpaces.count
        
        func increaseEmptySpace() {
            emptySpaceIndex += 1
            if emptySpaceIndex >= emptySpaceCount {
                increaseFile()
            } else {
                emptySpace = disk.emptySpaces[emptySpaceIndex]
            }
        }
        
        func increaseFile() {
            fileIndex -= 1
            if fileIndex >= 0 {
                emptySpaceIndex = 0
                emptySpace = disk.emptySpaces[emptySpaceIndex]
                file = disk.files[fileIndex]
            } else {
                file = nil
                emptySpace = nil
            }
        }
        
        while var localEmptySpace = emptySpace, let file = file {
            if localEmptySpace.size == 0 {
                increaseEmptySpace()
                continue
            }
            
            if localEmptySpace.position >= file.position {
                increaseFile()
                continue
            }
            
            if localEmptySpace.size > file.size {
                files.append(Content(position: localEmptySpace.position, size: file.size, type: file.type))
                localEmptySpace.size -= file.size
                localEmptySpace.position += file.size
                disk.files[fileIndex].size = 0
                
                emptySpace = localEmptySpace
                disk.emptySpaces[emptySpaceIndex] = localEmptySpace
                
                increaseFile()
            } else if localEmptySpace.size < file.size {
                increaseEmptySpace()
            } else {
                files.append(Content(position: localEmptySpace.position, size: file.size, type: file.type))
                disk.emptySpaces[emptySpaceIndex].size = 0
                disk.files[fileIndex].size = 0
                
                increaseFile()
            }
        }

        return files + disk.files.filter({ $0.type != .emptySpace && $0.size > 0 })
    }

    private func parse(_ data: String) -> Disk {
        var position = 0
        return data.enumerated().reduce(into: Disk()) { disk, entry in
            let size = entry.element.wholeNumberValue!
            let id = entry.offset/2
            let content = Content(position: position, size: size, type: entry.offset%2 == 0 ? .file(id) : .emptySpace)
            
            switch content.type {
            case .file(_):
                disk.files.append(content)
            case .emptySpace:
                disk.emptySpaces.append(content)
            }
            
            position += size
        }
    }
    
    private func checksum(files: [Content]) -> Int {
        files.reduce(0) { result, content in
            if case let .file(id) = content.type {
                return result + (content.position..<content.position+content.size).reduce(0) { $0 + $1*id }
            }
            return result
        }
    }
    
    private struct Disk {
        var files: [Content] = []
        var emptySpaces: [Content] = []
    }
    
    private struct Content: Equatable {
        var position: Position
        var size: Size
        let type: ContentType
    }
    
    private enum ContentType: Equatable {
        case file(ID), emptySpace
    }
    
    private typealias ID = Int
    private typealias Size = Int
    private typealias Position = Int
}

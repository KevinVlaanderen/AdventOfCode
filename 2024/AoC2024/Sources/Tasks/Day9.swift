import Foundation
import Framework
internal import Algorithms

public struct Day9: Day {
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        var disk = parse(data)

        switch task {
        case .task1:
            return task1(&disk)
        case .task2:
            return task2(&disk)
        }
    }
    
    private func task1(_ disk: inout Disk) -> Int {
        var inserted = 0
        var lowestEmptySpaceIndex = 0
        
        fileLoop: for (fileIndex, file) in disk.enumerated().filter({ $0.element.isFile }).reversed() {
            let adjustedFileIndex = fileIndex + inserted
            var remaining = file.size
            
            for contentIndex in lowestEmptySpaceIndex..<disk.endIndex {
                let content = disk[contentIndex]
                
                if case .file(_) =  content.type {
                    continue
                }
                
                lowestEmptySpaceIndex = contentIndex
                
                if contentIndex >= adjustedFileIndex {
                    disk[adjustedFileIndex].size = remaining
                    break fileLoop
                }
                if content.size > remaining {
                    disk.insert(file, at: contentIndex)
                    inserted += 1
                    disk[contentIndex].size = remaining
                    disk[contentIndex+1].size -= remaining
                    disk.remove(at: adjustedFileIndex+1)
                    continue fileLoop
                } else if content.size < remaining {
                    disk[contentIndex] = file
                    disk[contentIndex].size = content.size
                    remaining -= content.size
                } else {
                    disk[contentIndex] = file
                    disk[contentIndex].size = remaining
                    disk.remove(at: adjustedFileIndex)
                    continue fileLoop
                }
            }
        }
        
        return disk.checksum()
    }
    
    private func task2(_ disk: inout Disk) -> Int {
        var insertedAt: [Int] = []
        fileLoop: for (fileIndex, file) in disk.enumerated().filter({ $0.element.isFile }).reversed() {
            for contentIndex in 0..<disk.endIndex {
                let content = disk[contentIndex]
                
                if case .file(_) = content.type {
                    continue
                }
                
                let adjustedFileIndex = fileIndex + insertedAt.count { $0 < contentIndex}
                
                if contentIndex >= adjustedFileIndex {
                    continue fileLoop
                }
                if content.size > file.size {
                    disk.insert(file, at: contentIndex)
                    insertedAt.append(contentIndex)
                    disk[contentIndex+1].size -= file.size
                    disk[adjustedFileIndex+1] = Content(type: .emptySpace, size: file.size)
                    continue fileLoop
                } else if content.size == file.size {
                    disk[contentIndex] = file
                    disk[adjustedFileIndex] = Content(type: .emptySpace, size: file.size)
                    continue fileLoop
                }
            }
        }

        return disk.checksum()
    }

    private func parse(_ data: String) -> Disk {
        data.enumerated().map { offset, character in
            let id = offset/2
            let size = character.wholeNumberValue!
            let contentType: ContentType = offset%2 == 0 ? .file(id) : .emptySpace
            return Content(type: contentType, size: size)
        }
    }
    
    typealias ID = Int
    typealias Disk = [Content]
    
    struct Content {
        let type: ContentType
        var size: Int
    }
    
    
    
    enum ContentType: Equatable {
        case file(Int), emptySpace
    }
}

extension Day9.Disk {
    func checksum() -> Int {
        var index: Int = 0
        return self.reduce(0) { result, content in
            var value = 0
            if case let .file(id) = content.type {
                value = (index..<index+content.size).reduce(0) { $0 + $1*id }
            }
            index += content.size
            return result + value
        }
    }
}

extension Day9.Content {
  var isFile: Bool {
      switch self.type {
    case .file(_): return true
    default: return false
    }
  }
}

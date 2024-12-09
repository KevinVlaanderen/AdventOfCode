import Foundation
import Framework
internal import Algorithms

public struct Day9: Day {
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let diskMap = parse(data)

        return diskMap.enumerated().reduce(0) { $0 + $1.offset*$1.element }
    }

    private func parse(_ data: String) -> DiskMap {
        data.enumerated().reduce(into: DiskMap()) {
            let size = $1.element.wholeNumberValue!
            if $1.offset%2 == 0 {
                let id = $1.offset/2
                $0.files.append(Content(type: .file(id), size: size))
            } else {
                $0.emptySpaces.append(Content(type: .emptySpace, size: size))
            }
        }
    }
    
    typealias ID = Int
    
    struct DiskMap: Sequence {
        var files: [Content] = []
        var emptySpaces: [Content] = []

        func makeIterator() -> BlocksIterator {
            return BlocksIterator(diskMap: self)
        }
        
        struct BlocksIterator: IteratorProtocol {
            let diskMap: DiskMap
            
            var contentIndex = 0
            var currentContentIndex = 0
            
            var reversedFileIndex = 0
            var reversedCurrentFileIndex = 0
            
            init(diskMap: DiskMap) {
                self.diskMap = diskMap
            }
            
            mutating func next() -> ID? {
                while contentIndex < diskMap.files.count + diskMap.emptySpaces.count {
                    if contentIndex % 2 == 0 {
                        let fileIndex = contentIndex/2
                        if currentContentIndex < diskMap.files[fileIndex].size {
                            guard case let .file(id) = diskMap.files[fileIndex].type else {
                                fatalError("unexpected type")
                            }
                            if fileIndex == diskMap.files.count - reversedFileIndex - 1 && diskMap.files[fileIndex].size - currentContentIndex - reversedCurrentFileIndex <= 0  {
                                return nil
                            }
                            currentContentIndex += 1
                            return id
                        } else {
                            contentIndex += 1
                            currentContentIndex = 0
                        }
                    } else {
                        if currentContentIndex < diskMap.emptySpaces[contentIndex/2].size {
                            currentContentIndex += 1
                            
                            while reversedFileIndex < diskMap.files.count {
                                let fileIndex = diskMap.files.count - reversedFileIndex - 1
                                if reversedCurrentFileIndex < diskMap.files[fileIndex].size {
                                    guard case let .file(id) = diskMap.files[fileIndex].type else {
                                        fatalError("unexpected type")
                                    }
                                    if id == currentContentIndex/2 {
                                        return nil
                                    }
                                    reversedCurrentFileIndex += 1
                                    return id
                                } else {
                                    reversedFileIndex += 1
                                    reversedCurrentFileIndex = 0
                                }
                            }
                        } else {
                            contentIndex += 1
                            currentContentIndex = 0
                        }
                    }
                }
                
                return nil
            }
        }
    }
    
    struct Content {
        let type: ContentType
        let size: Int
    }
    
    enum ContentType: Equatable {
        case file(Int), emptySpace
    }
}

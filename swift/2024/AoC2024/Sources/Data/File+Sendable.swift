import Framework

extension File: @unchecked Sendable {}

public func loadData(file: File) -> DataDescriptor {
    DataDescriptor(url: file.path, load: {
        try String(contentsOfFile: file.path).trimmingCharacters(in: .newlines)
    })
}

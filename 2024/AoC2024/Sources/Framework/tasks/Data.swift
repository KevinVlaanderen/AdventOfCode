public struct DataDescriptor: Sendable {
    public init(url: String, load: @escaping DataLoader) {
        self.url = url
        self.load = load
    }
    
    public let url: String
    public let load: DataLoader
}

public typealias DataLoader = @Sendable () throws -> String

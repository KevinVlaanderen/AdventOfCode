public func toInt(_ value: any StringProtocol) throws -> Int {
    guard let result = Int(value) else {
        throw AoCError.parseError("could not convert string \(value) to Int")
    }
    return result
}

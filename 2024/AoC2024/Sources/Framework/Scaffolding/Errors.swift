public enum AoCError: Error {
    case notImplemented
    case dataNotFound
    case invalidState(String)
    case invalidTask(String)
    case parseError(String)
}

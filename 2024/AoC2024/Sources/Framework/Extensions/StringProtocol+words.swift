import Foundation

public extension StringProtocol {
    @inlinable
    var words: [Self.SubSequence] {
        self.split(whereSeparator: \.isWhitespace)
    }
}

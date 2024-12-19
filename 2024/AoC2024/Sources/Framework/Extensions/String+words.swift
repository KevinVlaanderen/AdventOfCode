import Foundation

public extension StringProtocol {
    var words: [Self.SubSequence] {
        self.split(whereSeparator: \.isWhitespace)
    }
}

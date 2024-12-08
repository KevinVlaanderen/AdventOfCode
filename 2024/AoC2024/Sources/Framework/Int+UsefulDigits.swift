public extension Int {
    var usefulDigits: Int {
        var num = self
        var count = 0
        while num != 0 {
            let digit = abs(num % 10)
            if digit != 0 && self % digit == 0 {
                count += 1
            }
            num = num / 10
        }
        return count
    }
}

public extension Int {
    var usefulDigits: Int {
        var num = self
        var count = 0
        if (num == 0){
              return 1
           }
        while num > 0 {
            num = num / 10
            count += 1
        }
        return count
    }
}

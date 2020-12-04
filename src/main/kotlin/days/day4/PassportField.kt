package days.day4

enum class PassportField(val acronym: String, val required: Boolean = true) {
    BIRTH_YEAR("byr"),
    ISSUE_YEAR("iyr"),
    EXPIRATION_YEAR("eyr"),
    HEIGHT("hgt"),
    HAIR_COLOR("hcl"),
    EYE_COLOR("ecl"),
    PASSPORT_ID("pid"),
    COUNTRY_ID("cid", false)
}
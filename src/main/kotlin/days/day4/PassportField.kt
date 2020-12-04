package days.day4

import shared.IValidate

val yearPattern = """\d{4}""".toRegex()
val heightPattern = """(\d+)(cm|in)""".toRegex()
val hairColorPattern = """#[0-9a-f]{6}""".toRegex()
val eyeColorPattern = """amb|blu|brn|gry|grn|hzl|oth""".toRegex()
val passportIdPattern = """\d{9}""".toRegex()

enum class PassportField(val acronym: String, val required: Boolean = true) : IValidate {
    BIRTH_YEAR("byr") {
        override fun validate(value: String): Boolean = yearPattern.matches(value) && value.toInt() in 1920..2002
    },
    ISSUE_YEAR("iyr") {
        override fun validate(value: String): Boolean = yearPattern.matches(value) && value.toInt() in 2010..2020
    },
    EXPIRATION_YEAR("eyr") {
        override fun validate(value: String): Boolean = yearPattern.matches(value) && value.toInt() in 2020..2030
    },
    HEIGHT("hgt") {
        override fun validate(value: String): Boolean {
            val matchResult = heightPattern.find(value) ?: return false

            val (number, unit) = matchResult.destructured

            return when (unit) {
                "cm" -> number.toInt() in 150..193
                "in" -> number.toInt() in 59..76
                else -> false
            }
        }
    },
    HAIR_COLOR("hcl") {
        override fun validate(value: String): Boolean = hairColorPattern.matches(value)
    },
    EYE_COLOR("ecl") {
        override fun validate(value: String): Boolean = eyeColorPattern.matches(value)
    },
    PASSPORT_ID("pid") {
        override fun validate(value: String): Boolean = passportIdPattern.matches(value)
    },
    COUNTRY_ID("cid", false) {
        override fun validate(value: String): Boolean = true
    };

    companion object {
        fun getPassportFieldByAcronym(acronym: String) = values().find { it.acronym == acronym }!!
    }
}
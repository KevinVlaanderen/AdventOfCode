import days.day4.PassportField
import org.assertj.core.api.Assertions
import org.junit.jupiter.api.DynamicTest
import org.junit.jupiter.api.TestFactory

class Day4 {
    @TestFactory
    fun birthYear() = listOf(
        Pair(1920, true),
        Pair(2002, true),
        Pair(1986, true),
        Pair(1919, false),
        Pair(2003, false))
        .map { (year, expected) ->
            DynamicTest.dynamicTest("The year $year is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.BIRTH_YEAR.validate(year.toString())
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }

    @TestFactory
    fun issueYear() = listOf(
        Pair(2010, true),
        Pair(2020, true),
        Pair(2015, true),
        Pair(2009, false),
        Pair(2021, false))
        .map { (year, expected) ->
            DynamicTest.dynamicTest("The year $year is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.ISSUE_YEAR.validate(year.toString())
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }

    @TestFactory
    fun expirationYear() = listOf(
        Pair(2020, true),
        Pair(2030, true),
        Pair(2025, true),
        Pair(2019, false),
        Pair(2031, false))
        .map { (year, expected) ->
            DynamicTest.dynamicTest("The year $year is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.EXPIRATION_YEAR.validate(year.toString())
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }

    @TestFactory
    fun height() = listOf(
        Pair("150cm", true),
        Pair("193cm", true),
        Pair("186cm", true),
        Pair("149cm", false),
        Pair("194cm", false),
        Pair("59in", true),
        Pair("76in", true),
        Pair("70in", true),
        Pair("58in", false),
        Pair("77in", false),
        Pair("160x", false),
        Pair("170", false),
        Pair("x", false))
        .map { (height, expected) ->
            DynamicTest.dynamicTest("A height of $height is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.HEIGHT.validate(height)
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }

    @TestFactory
    fun hairColor() = listOf(
        Pair("#000000", true),
        Pair("#ababab", true),
        Pair("000000", false),
        Pair("#00000", false),
        Pair("000000#", false),
        Pair("$000000", false))
        .map { (color, expected) ->
            DynamicTest.dynamicTest("The color $color is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.HAIR_COLOR.validate(color)
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }

    @TestFactory
    fun eyeColor() = listOf(
        Pair("amb", true),
        Pair("blu", true),
        Pair("brn", true),
        Pair("gry", true),
        Pair("grn", true),
        Pair("hzl", true),
        Pair("oth", true),
        Pair("aaa", false),
        Pair("", false))
        .map { (color, expected) ->
            DynamicTest.dynamicTest("The color $color is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.EYE_COLOR.validate(color)
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }

    @TestFactory
    fun passportId() = listOf(
        Pair("123456789", true),
        Pair("12345678", false),
        Pair("1234567890", false),
        Pair("12345678a", false))
        .map { (id, expected) ->
            DynamicTest.dynamicTest("The id $id is ${if(expected) "valid" else "not valid"}") {
                val result = PassportField.PASSPORT_ID.validate(id)
                Assertions.assertThat(result).isEqualTo(expected)
            }
        }
}
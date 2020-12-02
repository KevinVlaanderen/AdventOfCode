plugins {
    java
    kotlin("jvm") version "1.4.10"
}

group = "com.kevinvlaanderen"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    testImplementation("junit", "junit", "4.12")
    testImplementation("org.assertj", "assertj-core", "3.12.2")
    testImplementation( "org.junit.jupiter", "junit-jupiter-api", "5.4.2")
    testRuntimeOnly("org.junit.jupiter","junit-jupiter-engine", "5.4.2")
}

tasks {
    test {
        useJUnitPlatform()
    }
}
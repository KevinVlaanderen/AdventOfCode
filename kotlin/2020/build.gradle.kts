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

    testImplementation("org.assertj", "assertj-core", "3.18.1")
    testImplementation("org.junit.jupiter", "junit-jupiter-api", "5.7.0")

    testRuntimeOnly("org.junit.jupiter", "junit-jupiter-engine", "5.7.0")
}

tasks.withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = freeCompilerArgs + "-Xallow-result-return-type"
        jvmTarget = "14"
    }
}

tasks {
    test {
        useJUnitPlatform()
    }
}
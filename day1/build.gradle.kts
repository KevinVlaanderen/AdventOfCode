plugins {
    application
    kotlin("jvm")
}

group = "com.kevinvlaanderen"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    testImplementation("junit", "junit", "4.12")

    implementation(project(":shared"))
}

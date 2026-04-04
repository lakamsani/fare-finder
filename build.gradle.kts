plugins {
    java
    application
}

application {
    mainClass.set("com.fareFinder.Main")
}

java {
    sourceCompatibility = JavaVersion.VERSION_17
    targetCompatibility = JavaVersion.VERSION_17
}

dependencies {
    implementation(files("/opt/gradle-8.13/lib/gson-2.10.jar"))
    testImplementation(files(
        "/opt/gradle-8.13/lib/junit-4.13.2.jar",
        "/opt/gradle-8.13/lib/hamcrest-core-1.3.jar"
    ))
}

tasks.test {
    useJUnit()
}

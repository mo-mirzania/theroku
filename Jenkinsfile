pipeline {
    agent any
    
    stages {

        stage("Build") {
            steps {
                sh 'go build main.go'
            }
        }
        stage("Deploy") {
            steps {
                sh 'script.sh'
            }
        }
    }
}
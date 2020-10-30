pipeline {
    agent any
    
    stages {

        stage("Build") {
            steps {
                sh 'go build main.go'
                sh 'mv main /usr/local/bin/main'
            }
        }
        stage("Deploy") {
            steps {
                sh 'sh script.sh'
            }
        }
    }
}
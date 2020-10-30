pipeline {
    agent any
    
    stages {
        stage("Deploy") {
            steps {
                sh 'nohup go run main.go&'
            }
        }
    }
}
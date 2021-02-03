pipeline {
    agent any
    tools {
        go 'Go1.15.7'
    }
    parameters {
        string(name: 'RECIPIENTS', defaultValue: 'venkata.koripalli@gmail.com', description: 'Email for the build result')
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Pre Test') {
              when {
                anyOf { branch 'master'; branch 'staging' }  
            }
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            when {
                anyOf { branch 'master'; branch 'staging' }  
            }
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    // echo 'Running vetting'
                    // sh 'go vet .'
                    // echo 'Running linting'
                    // sh 'golint .'
                    echo 'Running test'
                    sh 'cd tests && go test -json > json-report.txt'
                    sh '~/go/bin/gotest2allure -f json-report.txt'
                }
            }
        }

        stage('Deploy') {
            when {
                branch 'master'
            }
            steps {
                echo 'Deploy the app'
            }
        }
    }
    post {
        always {
            script {
                allure([
                        includeProperties: false,
                        jdk: '',
                        properties: [],
                        reportBuildPolicy: 'ALWAYS',
                        results: [[path: 'tests/allure-results']]
                ])
            }
        }
    } 
}
pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.20.1'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        // stage("unit-test") {
        //     steps {
        //         echo 'UNIT TEST EXECUTION STARTED'
        //         sh 'make unit-tests'
        //     }
        // }
        // stage("functional-test") {
        //     steps {
        //         echo 'FUNCTIONAL TEST EXECUTION STARTED'
        //         sh 'make functional-tests'
        //     }
        // }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                echo $PAY_WEB_CC_REGISTRY_TCP_HOST
                sh 'go version'
                sh 'go get ./...'
                sh 'docker build . -t reinhardjs/go-epp-rest'
            }
        }
        stage('deliver') {
            agent any
            steps {
                withCredentials([usernamePassword(credentialsId: 'reinhardjs-dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                sh 'docker push reinhardjs/go-epp-rest'
                }
            }
        }
    }
}

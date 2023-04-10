pipeline {
    
    agent any
    
    tools {
        go 'go1.20.1'
    }
    
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        DEPLOYMENT_NAME = "go-epp-rest"
        REGISTRY_HOST = "sccs.webnic.tl:5000"
        DOCKER_LABEL = "reinhardjs/go-epp-rest"
        DOCKER_TAG = "latest"
    }

    stages {
        
        stage("unit-test") {
            steps {
                script {
                    try {
                        echo 'UNIT TEST EXECUTION STARTED'
                        sh 'make unit-tests'
                    } catch (err){
                        echo err.toString()
                    }
                }
            }
        }
        
        stage("functional-test") {
            steps {
                script {
                    try {
                        echo 'FUNCTIONAL TEST EXECUTION STARTED'
                        sh 'make functional-tests'
                    } catch (err){
                        echo err.toString()
                    }
                }
            }
        }
        
        stage("git-clone-secrets") {
            steps {   
                git credentialsId: 'gitlab-cred', url: 'http://gitlab/merekmu/k8s-secrets', branch: 'main'   
            }
        }
        
        stage('deploy-kubernetes-secret') {
            steps {
                // Install kubectl
                sh 'curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl'
                sh 'chmod +x ./kubectl'
                
                withKubeConfig([credentialsId: 'k8s-cred', serverUrl: 'https://10.10.3.206:6443']) {
                    sh './kubectl config view'
                    
                    script {
                        try {
                            def result = sh(script: './kubectl delete secret ${DEPLOYMENT_NAME}-secret', returnStdout: true)
                        } catch(err){
                            echo err.toString()
                        }
                        
                        try {
                            def result = sh(script: './kubectl create secret generic ${DEPLOYMENT_NAME}-secret --from-env-file=./${DEPLOYMENT_NAME}/.env', returnStdout: true)
                        } catch(err){
                            echo err.toString()
                        }
                    }
                }
            }
        }

        stage("git-clone-project") {
            steps { 
                git credentialsId: 'gitlab-cred', url: 'http://gitlab/merekmu/go-epp-rest', branch: 'feat/nojira-cicd'   
            }
        }
        
        stage('cleanup-old-image') {
            steps {
                script {
                    def image = "$REGISTRY_HOST/${DOCKER_LABEL}:${DOCKER_TAG}"
                    def imageUsed  = sh(returnStdout: true, script: "docker ps -a --filter ancestor=${image} --format {{.ID}}").trim()
                    if (imageUsed ) {
                        echo "Cannot remove image ${image} as it is being used by container \n${imageUsed}"
                    } else {
                        try {
                            sh "docker rmi ${image}"
                        } catch (Exception ex) {
                            // block of code to handle exception
                            echo "Failed to delete image: ${ex.getMessage()}"
                        }
                    }
                }
            }
        }
        
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                sh 'docker build -t $REGISTRY_HOST/$DOCKER_LABEL:$DOCKER_TAG .'
            }
        }
        
        stage('deliver') {
            agent any
            steps {
                withCredentials([usernamePassword(credentialsId: 'reinhardjs-dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                    sh "docker login $REGISTRY_HOST -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                    sh 'docker push $REGISTRY_HOST/$DOCKER_LABEL:$DOCKER_TAG'
                }
            }
        }
        
        stage('deploy-to-kubernetes') {
            steps {
                // Downloading kubectl
                sh 'curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl'
                sh 'chmod +x ./kubectl'
                
                withKubeConfig([credentialsId: 'k8s-cred', serverUrl: 'https://10.10.3.206:6443']) {
                    sh './kubectl config view'
                    
                    script {
                        def result = sh(script: './kubectl apply -f Deployment.yaml', returnStdout: true)
                        
                        // inspecting the error message output, nb: this is not a good practice
                        if (result.contains("unchanged")){
                            // Rolling update using `kubectl rollout restart`
                            echo "Deployment unchanged, need to restart deployment"
                            sh './kubectl rollout restart deployment $DEPLOYMENT_NAME'
                        }
                        
                        // exposing as service
                        try {
                            sh './kubectl apply -f Service.yaml'
                        } catch (err){}
                    }
                }
            }
        }
    }
}

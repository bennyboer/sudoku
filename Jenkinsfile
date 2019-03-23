pipeline {
    agent none
    stages {
        stage('Test') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }
            steps {
                sh 'go test ./... -coverprofile cover.out -v'
                sh 'go test ./... -bench=.'

                // Check that code coverage was > 90 %
                sh '''
                    LAST_LINE=$(go tool cover -func cover.out | tail -1);
                    REVERSED=$(echo $LAST_LINE | rev);
                    LAST_PART=$(echo $REVERSED | cut -d ' ' -f 1);
                    COVERAGE=$(echo $LAST_PART | rev);
                    echo $COVERAGE;
                '''
                // sh "go tool cover -func cover.out | tail -1 | rev | cut -d ' ' -f 1 | rev | cat"
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }   
            steps {
                sh 'golangci-lint run --enable-all --disable goimports'
            }
        }
        stage('Build Docker Image') {
            agent {
                label 'master'
            }
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME}"
            }
        }
    }
    post {
        changed {
            script {
                if (currentBuild.currentResult == 'FAILURE') { // Other values: SUCCESS, UNSTABLE
                    // Send an email only if the build status has changed from green/unstable to red
                    emailext subject: '$DEFAULT_SUBJECT',
                        body: '$DEFAULT_CONTENT',
                        recipientProviders: [
                            [$class: 'DevelopersRecipientProvider']
                        ], 
                        replyTo: '$DEFAULT_REPLYTO'
                }
            }
        }
    }
}

def pipeline() {
    pipeline {
        agent {
            docker {
                image 'foo/bar'
            }
        }
        
        stages {
            stage('Build') {
                steps {
                    // Perform the build steps here
                    sh 'mvn clean install'
                }
            }
    
            stage('Test') {
                steps {
                    // Run the tests here
                    sh 'mvn test'
                }
            }
    
            stage('Deploy') {
                steps {
                    // Deploy the application here
                    sh 'mvn deploy'
                }
            }
        }
    
        post {
            success {
                // Actions to perform when the pipeline succeeds
                echo 'Pipeline succeeded!'
            }
    
            failure {
                // Actions to perform when the pipeline fails
                echo 'Pipeline failed!'
            }
    
            always {
                // Actions to perform always, regardless of the pipeline result
                echo 'Pipeline finished!'
            }
        }
    }
}

return this

node {
    try{
        ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
            withEnv(["GOPATH=."]) {
                env.PATH="."
                
                stage('Checkout'){
                    echo 'Checking out SCM'
                    checkout scm
                }
                
                stage('Pre Test'){
                    echo 'Pulling Dependencies'
                    sh 'go version'
                    sh 'go get -u github.com/golang/dep/cmd/dep'
                    sh 'go get -u github.com/golang/lint/golint'
                    sh 'go get github.com/tebeka/go2xunit'
                    sh 'go get -v'
                    //or -update
                    sh 'dep ensure' 
                }
        
                stage('Test'){
                    
                    //List all our project files with 'go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org'
                    //Push our project files relative to ./src
                    sh 'go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org > projectPaths'
                    
                    //Print them with 'awk '$0="./src/"$0' projectPaths' in order to get full relative path to $GOPATH
                    def paths = sh returnStdout: true, script: """awk '\$0="./"\$0' projectPaths"""
                    
                    echo 'Vetting'

                    sh """go tool vet ${paths}"""

                    echo 'Linting'
                    sh """golint ${paths}"""
                    
                    echo 'Testing'
                    sh """go test -race -cover ${paths}"""
                }
            
                stage('Build'){
                    echo 'Building Executable'
                
                    //Produced binary is $GOPATH/src/cmd/project/project
                    sh """go build -ldflags '-s'"""
                }
                
                stage('BitBucket Publish'){
                
                    //Find out commit hash
                    sh 'git rev-parse HEAD > commit'
                    def commit = readFile('commit').trim()
                
                    //Find out current branch
                    sh 'git name-rev --name-only HEAD > GIT_BRANCH'
                    def branch = readFile('GIT_BRANCH').trim()
                    
                    //strip off repo-name/origin/ (optional)
                    branch = branch.substring(branch.lastIndexOf('/') + 1)
                
                    def archive = "./project-${branch}-${commit}.tar.gz"

                    echo "Building Archive ${archive}"
                    
                    sh """tar -cvzf ${archive} ."""

                    echo "Uploading ${archive} to BitBucket Downloads"
                    withCredentials([string(credentialsId: 'trunglv', variable: 'KEY')]) { 
                        sh """curl -s -u 'user:${KEY}' -X POST 'downloads-page-url' --form files=@'${archive}' --fail"""
                    }
                }
            }
        }
    }catch (e) {
        currentBuild.result = "FAILED"
    } finally {
        def bs = currentBuild.result ?: 'SUCCESSFUL'
        if(bs == 'SUCCESSFUL'){
            echo ${bs}
        }
    }
}
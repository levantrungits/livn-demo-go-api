// 1. Jenkinsfile (Declarative Pipeline) - pipeline
// 2. Jenkinsfile (Scripted Pipeline) - node
pipeline {
  // Execute this Pipeline or any of its stages, on any available agent.
  agent any
  stages {
    stage('Checkout') {
      steps {
        echo 'Checking out SCM'
        checkout scm
      }
    }
    stage('Pre Test') {  // Defines the "Build" stage.
        steps {
          // 	Perform some steps related to the "Build" stage.
          sh 'go version'
          sh 'go get -v'
        }
    }
    stage('Test') { // Defines the "Test" stage.
        steps {
          // Perform some steps related to the "Test" stage.
          sh 'go list ./... | grep -v /vendor/ | grep -v github.com | grep -v golang.org > projectPaths'
          echo 'Vetting'
          sh """go tool vet ."""
          echo 'Linting'
          sh """golint ."""
          echo 'Testing'
          sh """go test -race -cover ."""
        }
    }
    stage('Build'){
      steps {
        // Perform some steps related to the "Test" stage.
        // Produced binary
        sh """go build -ldflags '-s'"""
      }
    }
    stage('BitBucket Publish (Docker Image)') { // Defines the "BitBucket Publish Docker Image".
      steps {
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
        withCredentials([string(credentialsId: 'docker-hub', variable: 'KEY')]) { 
          sh """curl -s -u 'user:${KEY}' -X POST 'Downloads Page URL' --form files=@'${archive}' --fail"""
        }
      }
    }

  }
}
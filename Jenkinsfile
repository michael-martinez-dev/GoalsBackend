pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      parallel {
        stage('Checkout Code') {
          steps {
            git(url: 'https://github.com/MixedMachine/GoalsBackend', branch: 'prod')
            sh 'cd api && go mod tidy'
            sh 'cd ..'
            sh 'cd recommender && go mod tidy'
          }
        }

        stage('Log') {
          steps {
            sh 'ls -la'
            sh 'go version'
            sh 'docker version'
          }
        }
        stage('Env vars set-up') {
          steps {
            sh 'echo "LOG_LEVEL=$GB_LOG_LEVEL" >> .env'
            sh 'echo "API_HOST=$GB_API_HOST" >> .env'
            sh 'echo "API_PORT=$GB_API_PORT" >> .env'
            sh 'echo "POSTGRES_USER=$GB_POSTGRES_USER" >> .env'
            sh 'echo "POSTGRES_PASS=$GB_POSTGRES_PASS" >> .env'
            sh 'echo "POSTGRES_HOST=$GB_POSTGRES_HOST" >> .env'
            sh 'echo "POSTGRES_PORT=$GB_POSTGRES_PORT" >> .env'
            sh 'echo "POSTGRES_DB=$GB_POSTGRES_DB" >> .env'
            sh 'echo "NATS_URL=$GB_NATS_URL" >> .env'
            sh 'echo "NATS_HOST=$GB_NATS_HOST" >> .env'
            sh 'echo "NATS_PORT=$GB_NATS_PORT" >> .env'
            sh 'echo "NATS_GOALS_REC_TOPIC=$GB_NATS_GOALS_REC_TOPIC" >> .env'
            sh 'echo "JWT_SECRET_KEY=$SAB_JWT_SECRET_KEY" >> .env'
            sh 'echo "GPT3_TOKEN=$GB_GPT3_TOKEN" >> .env'
            sh 'echo "GPT3_MODEL=$GB_GPT3_MODEL" >> .env'
            sh 'echo "" >> .env'
          }
        }
      }
    }

    stage('Unit tests') {
      steps {
        echo 'Running Unit tests...'
        sh 'cd api && go test ./tests/unit/... && cd ..'
        sh 'cd recommender && go test ./tests/unit/... && cd ..'
      }
    }

    stage('Build images') {
      parallel {
        stage('Build images') {
          steps {
            echo 'Building docker images & pushing them to repo...'
            sh 'make image'
          }
        }

        stage('Build resources') {
          steps {
            echo 'Building Databases & Storage resources...'
            sh 'make db'
          }
        }

        stage('Log into Docker') {
          steps {
            sh 'docker login -u $DOCKER_HUB_USER -p $DOCKER_HUB_PW'
          }
        }

      }
    }

    stage('Run service') {
      steps {
        echo 'Running service with docker to run functional testing...'
      }
    }

    stage('Functional tests') {
      steps {
        echo 'Running functional tests with postman...'
      }
    }

    stage('Docker Hub push') {
      steps {
        echo 'Pushing to Dockerhub...'
        sh 'make image-push'
      }
    }

    stage('Prod env set-up') {
      steps {
        echo 'Setting up production environment...'
        sh 'ssh ${PROD_USER}@${PROD_IP} "cd code/GoalsApi && git pull --rebase --autostash && make prod"'
      }
    }

  }
  post {
      always {
          sh 'make clean'
      }
      success {
          echo 'The Pipeline was successful! 🎉'
      }
      failure {
          echo'The Pipeline failed 😔'
      }
  }
}
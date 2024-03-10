pipeline {
    agent any

    environment {
        // 设置服务名称
        SERVER_NAME = "payment"
        // 配置名称
        CONFIG_NAME = "${SERVER_NAME}"
        // 服务类型，默认http服务
        SERVER_TYPE = "http"
        // 环境，默认dev
        ENV = "dev"
        // 镜像名称
        IMAGE_NAME = "${env.HARBOR_HOST}/${env.SERVER_NAME}/${env.JOB_BASE_NAME}"
        // 副本数
        REPLICAS = 1
    }

    stages {

        stage('pull') {
            steps {
                echo '代码拉取完成'

                script {
                    // 判断环境
                    if (env.JOB_BASE_NAME =~ /test/) {
                        ENV = "test"
                    }

                    if (env.JOB_BASE_NAME =~ /master/) {
                        ENV = "master"
                        REPLICAS = 2

                        if (env.JOB_BASE_NAME =~ /job/) {
                            REPLICAS = 1
                        }
                    }

                    // 如果发布job，则变更服务类型为job
                    if (env.JOB_BASE_NAME =~ /job/) {
                        SERVER_TYPE = "job"
                        SERVER_NAME = env.JOB_BASE_NAME
                    }

                    // 镜像名称
                    IMAGE_NAME = "${env.HARBOR_HOST}/${env.SERVER_NAME}/${env.ENV}:v${env.BUILD_NUMBER}"
                }
            }
        }

        stage('build-image') {
            steps {
                sh "docker build -t $IMAGE_NAME -f ./deploy/http/Dockerfile ."
            }
        }

        stage('push-image') {
            steps {
                sh "docker push $IMAGE_NAME"
            }
        }


        stage('deploy-k8s'){
            steps {
                sh "sed -i 's|IMAGE_ADDRESS|$IMAGE_NAME|g' ./deploy/$SERVER_TYPE/deployment.yaml"
                sh "sed -i 's|SERVER_NAME|$SERVER_NAME|g' ./deploy/$SERVER_TYPE/deployment.yaml"
                sh "sed -i 's|ENV|$ENV|g' ./deploy/$SERVER_TYPE/deployment.yaml"
                sh "sed -i 's|REPLICAS|$REPLICAS|g' ./deploy/$SERVER_TYPE/deployment.yaml"
                sh "sed -i 's|CONFIG_NAME|$CONFIG_NAME|g' ./deploy/$SERVER_TYPE/deployment.yaml"

                script{
                    withKubeConfig(caCertificate: '', clusterName: '', contextName: '', credentialsId: 'k8s-config', namespace: 'nameSpaceofchildProject3', restrictKubeConfigAccess: false, serverUrl: "${env.K8S_HOST}"){
                        sh "kubectl apply -f ./deploy/$SERVER_TYPE/deployment.yaml"
                    }
                }
            }
        }
    }

}
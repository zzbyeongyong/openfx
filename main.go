package main

import (
	"log"

	"github.com/keti-openfx/openfx-gateway/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	//kubernetes가 pod에 제공하는 config 객체를 반환
	//kubernetes에서 실행중인 Pod 내부에서 실행하지 않으면, 에러 발생.
	conf, err := rest.InClusterConfig()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("kubernetest host: %s\n", conf.Host)

	//주어진 config을 위한 clienteset 반환
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		log.Panic(err)
	}

	//FxGateway config 생성
	c := config.NewFxGatewayConfig()

	//FxGateway 생성
	s := NewFxGateway(c, clientset)
	log.Printf("[fxgateway] service start\n")

	//FxGateway 실행
	s.Start()
}

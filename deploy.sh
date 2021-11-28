cat /etc/os-release | grep 20.04
if [ $? -eq 0 ]; then
    echo "###############  it's ubuntu 20.04 machine ###############"
        echo "###############  git clone: ###############"
        git clone https://github.com/EldanHamdani/go-web.git
        echo "############### starting kubernetes eployment  ###############"
        helm install ./go-web --generate-name
        sleep 20
        echo "###############  kubectl get deployment ###############"
        kubectl get deployment
        echo "############### kubectl get services  ###############"
        kubectl get services
        echo "###############  you can access to go website using the following url: ###############"
        minikube service --url web-service
        echo "############### create systemctl service: ###############"
        FILE="/lib/systemd/system/go-web.service"
        echo "[Unit]" >$FILE
    echo "Description=go-web" >>$FILE
        echo "[Service]" >>$FILE
        echo "Type=simple" >>$FILE
        echo "Restart=always" >>$FILE
        echo "RestartSec=5s" >>$FILE
        echo "ExecStart=/home/ubuntu/go-web" >>$FILE
        echo "[Install]" >>$FILE
        echo "WantedBy=multi-user.target " >>$FILE

        echo "############### service go-web start  ###############"
        service go-web start
        echo "############### service go-web enable  ###############"
        service go-web enable
        echo "############### service go-web status  ###############"
        service go-web status

else
    echo "it's not ubuntu machine - can't start the deployment"
fi

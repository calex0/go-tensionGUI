En este fichero se guardan los paths:

~/.bashrc

En Terminal hacer:

export PATH=$PATH:/usr/local/go/bin

export GOPATH=$HOME/go

export PATH=$PATH:$GOPATH/bin

------

Abrimos 2 Terminal; cada uno en src y bin. Es lo más cómodo.

Ejecución del programa sin compilar:

alex@X230:~/go/src/hipoteca$ go run hipoteca.go

Crea binario:
crea un compilado en /go/bin/ 

alex@X230:~/go/src/hipoteca$ go install hipoteca

Ejecuta binario desde carpeta /go/bin/

alex@X230:~/go/bin$ ./hipoteca
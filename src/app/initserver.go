package app

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	logserver = "./log/server/logserver.log"
	logerror = "./log/error/logerror.log"
)


func InitServer(){
	fmt.Print("Validando conexão dos Logs...\n")
	//time.Sleep(8000*time.Millisecond)
	errServer := validLog(logserver)
	if errServer != nil{
		gravaLog([]byte(errServer.Error()))
	}

	errSys := validLog(logerror)
	if errSys != nil{
		gravaLog([]byte(errSys.Error()))
	}

	fmt.Println("Testando Conexão com o Banco de dados...\n")
	//time.Sleep(8000*time.Millisecond)
	fmt.Println("Disponibilizando servidor...\n")
	//time.Sleep(8000*time.Millisecond)
	fmt.Println("Servidor completo...")



}

func validLog(log string) error{
	if _,err := os.Stat(log);
	os.IsNotExist(err){
		fmt.Println("File not Found")
		return err
	}else {
		fmt.Println("Logs Online")
		return err
	}
}

func gravaLog(log []byte) string{
	err := ioutil.WriteFile(logserver, log, 0644)
	if err != nil{
		fmt.Println("Erro na gravação do arquivo",err)
		return "erro"
	}else{
		return "OK"
	}

}
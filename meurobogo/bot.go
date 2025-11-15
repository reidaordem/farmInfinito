package main
import (
	"fmt"
	"time"
	"github.com/go-vgo/robotgo"
)


func farmVoidbagsInfinito() {

	
 estadoAtual:= EstadoLobby
  
for{
	
	robotgo.ActiveName("Roblox")
	
	err:=capturarTela("screenshot.png")
	
	if err!= nil {
		fmt.Println("erro ao capturar a tela",err)
		return
	}

		switch estadoAtual{
		case EstadoLobby:
			moveEClica(closeDailyRewards.X,closeDailyRewards.Y)
			//sempre que estiver no lobby eu quero que ele abra o inventario para ver se tem voidbag
		  abrirInventario()
		  if detectarVoidbag(){
			processarIventario()

		  } 
		  if !detectarVoidbag(){
			estadoAtual = EstadoEntrarJogo
		  }
		case EstadoEntrarJogo:
			entrarInfinito()
			time.Sleep(60*time.Second)
            if carregandoModoInfinito() || detectarAbilityManager() || detectarMapInfo() || detectarUnitManager(){
				estadoAtual = EstadoJogando
			}

		case EstadoJogando:
			if detectarAbilityManager() || detectarMapInfo() || detectarUnitManager(){
				robotgo.KeyDown("w")
				time.Sleep(2200*time.Millisecond)
				robotgo.KeyUp("w")
				ajustarCamera()
				jogandoFarmando()
				
				if detectarLobby() || detectarInventario(){
					
					estadoAtual = EstadoLobby
					time.Sleep(60*time.Second)
				}
			}
		}
		
	

		time.Sleep(400*time.Millisecond)
	}


}

func farmExpedicoes(){

	for {

	robotgo.ActiveName("Roblox")
	
	hyperbolicTraining()
	time.Sleep(1*time.Second)
	arenaTowerTraining()
	time.Sleep(5*time.Second)
	  	}
}
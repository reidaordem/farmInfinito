package main
import (
	"fmt"
	"strconv"
	"time"
	
	"os/exec"
	"github.com/go-vgo/robotgo"
)

func ajustarCamera() {
    // 1. Arrastar a câmera com AutoHotkey
    ahkPath := "C:\\Users\\Usuario\\AppData\\Local\\Programs\\AutoHotkey\\v2\\AutoHotkey64.exe"
    scriptPath := "C:\\meus_scripts\\arrastar_camera.ahk"
	robotgo.Move(960,250)
    cmd := exec.Command(ahkPath, scriptPath)
   err:= cmd.Run()
	if err!= nil {
		println("Erro ao executar o script AHK:", err.Error())
	}

    time.Sleep(300 * time.Millisecond)

   robotgo.KeyDown("o")
   time.Sleep(3 * time.Second)
   robotgo.KeyUp("o")

}

func processarIventario()(bool){
time.Sleep(500*time.Millisecond)
	
	if detectarVoidbag(){
		
		moveEClica(posVoidbag.X,posVoidbag.Y)
		time.Sleep(500*time.Millisecond)
		
		moveEClica(btnUseVoidbag.X,btnUseVoidbag.Y)
		time.Sleep(500*time.Millisecond)
		//confirma para abrir as bolsas
		moveEClica(qtdVoidbag.X,qtdVoidbag.Y)
		time.Sleep(500*time.Millisecond)

		moveEClica(openVoidbag.X,openVoidbag.Y)
		time.Sleep(500*time.Millisecond)
		fmt.Println("usou as voidbags ")
		moveEClica(960,240)
		moveEClica(960,230)
		//retorna a função pois pode ter mais bolsas para abrir ainda
		return processarIventario()
	}
	moveEClica(closeInventory.X,closeInventory.Y)

	return false
}


func moveEClica(x,y int){
	ahkPath := "C:\\Users\\Usuario\\AppData\\Local\\Programs\\AutoHotkey\\v2\\AutoHotkey64.exe"
	scriptPath := "C:\\meus_scripts\\mover_clicar.ahk"

	cmd := exec.Command(ahkPath,scriptPath,strconv.Itoa(x),strconv.Itoa(y))

	err:= cmd.Run()
	if err!= nil {
		println("Erro ao executar o script AHK:", err.Error())
	}
	
}

func clicaApenas(x,y int){
	ahkPath := "C:\\Users\\Usuario\\AppData\\Local\\Programs\\AutoHotkey\\v2\\AutoHotkey64.exe"
	scriptPath := "C:\\meus_scripts\\clicar.ahk"

	cmd := exec.Command(ahkPath,scriptPath,strconv.Itoa(x),strconv.Itoa(y))

	err:= cmd.Run()
	if err!= nil {
		println("Erro ao executar o script AHK:", err.Error())
	}
}
 
// vulgo clicarCordenadas
func colocarPersonagem(coords [][2]int,slot string, delay time.Duration) {
	for _, c := range coords {
		robotgo.KeyPress(slot)
		time.Sleep(delay)
		clicaApenas(c[0], c[1]) // usa X e Y da lista
		time.Sleep(delay)      // espera entre cliques
	}
}


func abrirInventario() bool{
	if detectarInventario(){
		moveEClica(btnInventario.X,btnInventario.Y)
		
		if inventarioAberto(){

			fmt.Println("abriu inventario")
			moveEClica(barPesquisa.X,barPesquisa.Y)
			time.Sleep(100*time.Millisecond)
			clicaApenas(barPesquisa.X,barPesquisa.Y)
			time.Sleep(100*time.Millisecond)
			clicaApenas(barPesquisa.X,barPesquisa.Y)

			robotgo.TypeStrDelay("Void Bag",100)
			fmt.Println("escrevendo void bag")
			robotgo.KeyPress("enter")


        
			fmt.Println("se apareceu é pq escreveu ")
			return true
		 
		}
		
		if !inventarioAberto(){
			clicaApenas(960,450)
			abrirInventario()
		}
	}
	return false
}

func entrarInfinito(){
	if detectarWarp(){
		fmt.Println("detectou botão")
		moveEClica(btnWarp.X,btnWarp.Y)
		time.Sleep(100*time.Millisecond)
		fmt.Println("clicou botão")

		if detectarTpinfinito(){
		moveEClica(tpInfinitoMode.X,tpInfinitoMode.Y)

		time.Sleep(2*time.Second)
		fmt.Println("deu teleporte")

		moveEClica(closeWarp.X,closeWarp.Y)
		time.Sleep(500*time.Millisecond)
		logDebug("fechou o tp")

		robotgo.KeyDown("a")
		fmt.Println("andou pra esquerda")
		time.Sleep(800*time.Millisecond)
		robotgo.KeyUp("a")
		logDebug("parou de andar")

		robotgo.KeyDown("e")
		fmt.Println("entrando na sala")
		robotgo.KeyUp("e")

		if detectarInfinitoSala(){
			fmt.Println("entrou na sala")
			
			moveEClica(startInfinitoMode.X,startInfinitoMode.Y)
			//cliques de precaução
			for range 5{clicaApenas(startInfinitoMode.X,startInfinitoMode.Y)}

				if detectarInfinitoSala2(){
			 		moveEClica(startInfinitoMode2.X,startInfinitoMode2.Y)
					for range 5{clicaApenas(startInfinitoMode2.X,startInfinitoMode2.Y)}
			 	}

			}
		}
		if !detectarTpinfinito(){
			clicaApenas(960,450)
			entrarInfinito()
		}
	
	}
	
}

// 1.akainu | 2.Farm(tio do ramen) | 3.Sukuna | 4.Mihawk |5.Stark |6.Frieren
func jogandoFarmando(){
	
	
	logDebug("iniciando partida")
	for{
		robotgo.Move(960,250)
	if detectarTelaDerrota() || detectarReturnLobby() || detectarReplay() {
      moveEClica(960,480)
	  moveEClica(960,490)
	  moveEClica(960,500)

		if qtdpartidas >=5 {
			for range 12 {
			detectarReturnLobby()
			clicaApenas(btnReturnLobby.X,btnReturnLobby.Y)
			time.Sleep(500*time.Millisecond)
				}
		  qtdpartidas = 0
		  break
		  
			  
		}

		qtdpartidas ++
		fmt.Println("partida n° ",qtdpartidas)
		for range 12{
			detectarReplay()
			clicaApenas(btnReplay.X,btnReplay.Y)
			time.Sleep(500*time.Millisecond)
			
		  
		}
			
			
		
	}
	moveEClica(btnSpeedUp.X,btnSpeedUp.Y)
	colocarPersonagem(coordsFarmUnit,"2",300*time.Millisecond)
	

	moveEClica(btnStartGame.X,btnStartGame.Y)
	

	colocarPersonagem(coordsGround,"5",300*time.Millisecond)
	

	colocarPersonagem(coordsGround,"1",300*time.Millisecond)
	

	colocarPersonagem(coordsAir,"6",300*time.Millisecond)
	colocarPersonagem(coordsAir,"3",300*time.Millisecond)
	colocarPersonagem(coordsAir,"4",300*time.Millisecond)
	

	time.Sleep(700*time.Millisecond)
	robotgo.KeyPress("t")
	time.Sleep(700*time.Millisecond)
	robotgo.KeyPress("k")
	
	time.Sleep(700*time.Millisecond)
	robotgo.KeyPress("t")
	time.Sleep(700*time.Millisecond)

	
	}
	
}


func dialogoExpedition(){
	
	robotgo.KeyPress("e")
	time.Sleep(700*time.Millisecond)
	clicaApenas(960,540)
	
	moveEClica(btnExpedition.X,btnExpedition.Y)
		

}

func hyperbolicTraining()bool{
	dialogoExpedition()
	time.Sleep(150*time.Millisecond)
	moveEClica(btn1stExpedition.X,btn1stExpedition.Y)
	logDebug("clicou na expedicao 1")
	
		logDebug("detectou com sucesso primeira missao")
		moveEClica(btnClaimandConfExpedition.X,btnClaimandConfExpedition.Y)
		logDebug("primeiro botao confirm")
		if detectarPersoExpedition() || detectarPersoExpedition2(){
		comecarExpedition()
		return true
		}
		for range 8{clicaApenas(btnCloseMission.X,btnCloseMission.Y)
		time.Sleep(100*time.Millisecond)}
		moveEClica(btnClaimandConfExpedition.X+200,btnClaimandConfExpedition.Y)//fechando a janela
		return false
}

func arenaTowerTraining()bool{
	dialogoExpedition()
	moveEClica(btn2stExpedition.X,btn2stExpedition.Y)
	logDebug("clicou na expedicao 2")
		logDebug("detectou com sucesso segunda missao")
		moveEClica(btnClaimandConfExpedition.X,btnClaimandConfExpedition.Y)
		logDebug("segundo botao confirm")
		if detectarPersoExpedition() || detectarPersoExpedition2(){
		comecarExpedition()
		return true
		}
		for range 8{clicaApenas(btnCloseMission.X,btnCloseMission.Y)
		time.Sleep(100*time.Millisecond)}
		moveEClica(btnClaimandConfExpedition.X+200,btnClaimandConfExpedition.Y)
		return false
	
}

func comecarExpedition() bool{
		
			logDebug("detectou os personagens")
			moveEClica(posPersoExpedition.X,posPersoExpedition.Y)
			time.Sleep(50*time.Millisecond)
			moveEClica(posPersoExpedition.X+110,posPersoExpedition.Y+115)
			time.Sleep(50*time.Millisecond)
			moveEClica(posPersoExpedition.X+220,posPersoExpedition.Y+230)
			time.Sleep(50*time.Millisecond)
			moveEClica(posPersoExpedition.X+330,posPersoExpedition.Y)
			time.Sleep(50*time.Millisecond)
			moveEClica(posPersoExpedition.X,posPersoExpedition.Y+115)
			time.Sleep(50*time.Millisecond)
			moveEClica(posPersoExpedition.X+110,posPersoExpedition.Y+230)
			time.Sleep(50*time.Millisecond)
			moveEClica(btnConfirmPerso.X,btnConfirmPerso.Y)
			time.Sleep(1000*time.Millisecond)
			moveEClica(btnStartExpedition.X,btnStartExpedition.Y)
			time.Sleep(1000*time.Millisecond)

			if detectarMissionHyperbolic() || detectarMissionArenaTower(){
			moveEClica(realstartexpedition.X,realstartexpedition.Y)
			logDebug("começou a expedição")
			return true
			
		}
		logDebug("ja tem uma expedição em andamento ou a detecção falhou")
		return false
}
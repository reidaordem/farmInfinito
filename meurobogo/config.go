package main
import "image"


type Estado string

const (
	EstadoLobby      Estado = "lobby" // vou abrir o inventario pesquisar voidbag e usar todas, dps vou dar tp pro modo de jogo
	EstadoEntrarJogo Estado = "entrarJogo"
	EstadoJogando    Estado = "jogando"
	EstadoErro       Estado = "erro"
	
)

var(
		qtdpartidas = 0
 btnInventario = image.Point{X: 63, Y: 470}
 btnWarp = image.Point{X: 63, Y: 570}
 barPesquisa = image.Point{X: 1035, Y: 320}
 posVoidbag = image.Point{X: 713, Y: 415}
 btnUseVoidbag = image.Point{X: 1342, Y: 782}
 qtdVoidbag = image.Point{X: 1125, Y: 577}
 closeInventory = image.Point{X: 1137, Y: 240}
 tpInfinitoMode = image.Point{X: 762, Y: 645}
 startInfinitoMode = image.Point{X: 1212, Y: 831}
 startInfinitoMode2 = image.Point{X: 1500, Y: 715}
 openVoidbag = image.Point{X: 1120, Y: 660}
 closeWarp = image.Point{X: 1300, Y: 297}
 closeDailyRewards = image.Point{X: 1325, Y: 245}
 btnReplay = image.Point{X: 1057, Y: 757}
 btnReturnLobby = image.Point{X: 1332, Y: 760}
 btnStartGame = image.Point{X: 946, Y: 190}
 btnSpeedUp = image.Point{X: 1793, Y: 658}
 btnExpedition = image.Point{X:726 , Y:800}
 btn1stExpedition = image.Point{X:480, Y:337}
 btn2stExpedition = image.Point{X:480, Y:473}
 btn3stExpedition = image.Point{X:480, Y:602}
 btnClaimandConfExpedition = image.Point{X:1210, Y:831}
 posPersoExpedition = image.Point{X:640, Y:345} // distancia X= 110 ,distancia Y=115
 btnConfirmPerso = image.Point{X:1197, Y:715}
 btnStartExpedition = image.Point{X:845, Y:647}
 realstartexpedition = image.Point{X:1500, Y:745}
 btnCloseMission = image.Point{X:950, Y:757}
 modoDebug = true
coordsAir = [][2]int{
	{919,332},{948,332},
	{981,332},{941,371},
	{981,376},{976,402},
	{958,399},{984,402},
	{1001,427},{931,255},
	{921,279},{1001,250},
	{1001,270},{1001,290},
	{1001,310},{1001,330},

}
coordsGround = [][2]int{
	{771,374},{751,374},
	{721,374},{776,429},
	{756,429},{726,429},
	{756,402},{726,402},
	{706,402},
}

coordsFarmUnit = [][2]int{
	{1494,185},{1500,250},
	{1500,315},{1500,380},
	{1500,445},{1500,510},
	{1127,774},{1500,250},
	{207,820},
}
)
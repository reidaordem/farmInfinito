package main
import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
)

func esperarAte(screenPath, targetPath string, limiar float32, timeout time.Duration) bool {
    inicio := time.Now()
    for time.Since(inicio) < timeout {
        capturarTela(screenPath)
        if encontrado, _, _ := detectarImagemApenas(screenPath, targetPath, limiar); encontrado {
            return true
        }
        time.Sleep(500 * time.Millisecond)
    }
    return false
}

func capturarTela(nomeArquivo string) error{
	bounds := screenshot.GetDisplayBounds(0)
	imagem, erro := screenshot.CaptureRect(bounds)
	if erro != nil{
		fmt.Println("Erro ao capturar a tela:")
		return erro
	}
	arquivo, erro := os.Create(nomeArquivo)
	if erro != nil{
		return erro
	}
	defer arquivo.Close()

	return png.Encode(arquivo,imagem)
}

func detectarImagemEClicar(screenPath,targetPath string, limiar float32)(bool,image.Point,error){
	tela:= gocv.IMRead(screenPath,gocv.IMReadColor)
	defer tela.Close()

	alvo:= gocv.IMRead(targetPath,gocv.IMReadColor)
	defer alvo.Close()

	if tela.Empty() || alvo.Empty() {
		return false, image.Point{}, fmt.Errorf("erro ao carregar imagem")
	}

	resultado:= gocv.NewMat()
	defer resultado.Close()

	gocv.MatchTemplate(tela,alvo,&resultado,gocv.TmCcoeffNormed,gocv.NewMat())

	_ ,maxCorrelacao, _ ,posicaoAlvo := gocv.MinMaxLoc(resultado)

	if maxCorrelacao >= limiar{
		
		x := posicaoAlvo.X + alvo.Cols()/2
		y := posicaoAlvo.Y + alvo.Rows()/2
		//robotgo.MoveSmooth(x, y)
		//robotgo.Click("left", false)
		moveEClica(x,y)

		return true, posicaoAlvo, nil
	}
	return false, image.Point{}, nil
}


func detectarImagemApenas(screenPath, targetPath string, limiar float32) (bool, image.Point, error) {
	tela := gocv.IMRead(screenPath, gocv.IMReadColor)
	defer tela.Close()

	alvo := gocv.IMRead(targetPath, gocv.IMReadColor)
	defer alvo.Close()

	if tela.Empty() || alvo.Empty() {
		return false, image.Point{}, fmt.Errorf("erro ao carregar imagens")
	}

	resultado := gocv.NewMat()
	defer resultado.Close()

	gocv.MatchTemplate(tela, alvo, &resultado, gocv.TmCcoeffNormed, gocv.NewMat())

	_, maxVal, _, maxLoc := gocv.MinMaxLoc(resultado)

	if maxVal >= limiar {
		// Retorna o centro da imagem-alvo
		x := maxLoc.X + alvo.Cols()/2
		y := maxLoc.Y + alvo.Rows()/2
		return true, image.Point{X: x, Y: y}, nil
	}

	return false, image.Point{}, nil
}

func tentarDetectarImagem(screen, target string, limiar float32, tentativas int) bool {
    for i := 0; i < tentativas; i++ {
        capturarTela(screen)
        if encontrado, _, _ := detectarImagemApenas(screen, target, limiar); encontrado {
            return true
        }
        time.Sleep(500 * time.Millisecond)
    }
    return false
}

func detectarInventario() bool{
		return esperarAte("screenshot.png","imagens\\items.png",0.85,7*time.Second)
}

func detectarUnitManager() bool{
		return esperarAte("screenshot.png","imagens\\unitmanager.png",0.85,7*time.Second)
}

func detectarAbilityManager() bool{
		return esperarAte("screenshot.png","imagens\\abilityManager.png",0.85,7*time.Second)
}

func detectarMapInfo() bool{
		return esperarAte("screenshot.png","imagens\\mapinfo.png",0.85,7*time.Second)
}

func detectarVoidbag() bool{
	return esperarAte("screenshot.png","imagens\\VoidBag.png",0.75,10*time.Second)
}

func detectarWarp() bool{
		return esperarAte("screenshot.png","imagens\\warp.png",0.85,7*time.Second)
}

func detectarTpinfinito() bool{
		return esperarAte("screenshot.png","imagens\\infinite.png",0.85,7*time.Second)
}

func detectarInfinitoSala() bool{
		return esperarAte("screenshot.png","imagens\\entrarinfinito.png",0.85,7*time.Second)
}

func detectarInfinitoSala2() bool{
	return esperarAte("screenshot.png","imagens\\entrarinfinito2.png",0.85,7*time.Second)
}

func detectarLobby() bool{
	return esperarAte("screenshot.png","imagens\\telaCarregamentoLobby.png",0.85,7*time.Second)
}

func carregandoModoInfinito() bool{
	return esperarAte("screenshot.png","imagens\\FoodIsland.png",0.65,7*time.Second)
}

func inventarioAberto() bool{
	return esperarAte("screenshot.png","imagens\\abriuInventario.png",0.65,7*time.Second)
}

func detectarReplay() bool{
	return esperarAte("screenshot.png","imagens\\replay.png",0.85,7*time.Second)
}

func detectarReturnLobby() bool{
	return esperarAte("screenshot.png","imagens\\returnLobby.png",0.85,7*time.Second)
}

func detectarTelaDerrota() bool{
	return esperarAte("screenshot.png","imagens\\telaDerrota.png",0.7,7*time.Second)
}

func detectarPersoExpedition() bool{
	return esperarAte("screenshot.png","imagens\\SelecPersoExpedition.png",0.6,3*time.Second)
}

func detectarPersoExpedition2() bool{
	return esperarAte("screenshot.png","imagens\\SelecPersoExpedition2.png",0.6,3*time.Second)
}

func detectarMissionArenaTower() bool{
	return esperarAte("screenshot.png","imagens\\StartMissionArenaTower.png",0.85,4*time.Second)
}

func detectarMissionHyperbolic() bool{
	return esperarAte("screenshot.png","imagens\\StartMissionHyperbolic.png",0.85,4*time.Second)
}
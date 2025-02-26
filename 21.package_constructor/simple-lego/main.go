package main

import (
	"fmt"
	"sync"
	"time"
)

// это наш интерфейс для всех этапов сборки Lego, он нужен для полиморфизма
type legoBuilder interface {
	legoBuild() time.Duration // время на сборку
	legoUpdate()              // обновление состояния
}

// Структура для хранения состояния сборки
type LegoGo struct {
	legoBuilder legoBuilder
	detail      string
}

func (lg *LegoGo) legoBuild() {
	/*
		Метод legoBuild в legoGo делегирует вызовы методов legoBuild и legoUpdate объекту,
		который реализует интерфейс legoBuilder.
		Это позволяет динамически менять логику сборки в зависимости от текущего этапа.
	*/
	lg.legoBuilder.legoBuild()
	lg.legoBuilder.legoUpdate()
}

// структура для детали ================================================================================================
type Detail struct {
	id int
}

func (d *Detail) legoBuild() time.Duration {
	time.Sleep(100 * time.Millisecond) // имитация времени сборки детали

	return 100 * time.Millisecond
}

func (d *Detail) legoUpdate() {
	fmt.Printf("Деталь %d выбрана\n", d.id)
}

// структура для брика (1 брик = 3 детали) =============================================================================

type Brick struct {
	details []*Detail
}

func (br *Brick) legoBuild() time.Duration {
	var totalTime time.Duration

	if len(br.details) == 0 {
		fmt.Println("Ошибка: Brick без деталей!")
		return 0
	}

	for _, detail := range br.details {
		totalTime += detail.legoBuild()
	}

	time.Sleep(200 * time.Millisecond) // Имитация времени сборки брика
	return totalTime + 200*time.Millisecond
}

func (br *Brick) legoUpdate() {
	fmt.Println("Брик собран")
}

// структура для блока (1 блок = 3 брика) ==============================================================================
type Block struct {
	bricks []*Brick
}

func (bl *Block) legoBuild() time.Duration {
	var totalTime time.Duration

	for _, brick := range bl.bricks {
		totalTime += brick.legoBuild()
	}

	time.Sleep(300 * time.Millisecond) // имитация времени сборки блока
	return totalTime + 300*time.Millisecond
}

func (bl *Block) legoUpdate() {
	fmt.Println("Блок собран")
}

// структура для конструкции (1 конструкция = 3 блока) =================================================================
type Construction struct {
	blocks []*Block
}

func (c *Construction) legoBuild() time.Duration {
	var totalTime time.Duration

	for _, block := range c.blocks {
		totalTime += block.legoBuild()
	}

	time.Sleep(500 * time.Millisecond) // имитация времени сборки конструкции
	return totalTime + 500*time.Millisecond
}

func (c *Construction) legoUpdate() {
	fmt.Println("Конструкция собрана")
}

// структура для модели (1 модель = 3 конструкции) =====================================================================
type FullModel struct {
	constructions []*Construction
}

func (fm *FullModel) legoBuild() time.Duration {
	var totalTime time.Duration

	for _, construction := range fm.constructions {
		totalTime += construction.legoBuild()
	}

	time.Sleep(1 * time.Second) // имитация времени сборки модели
	return totalTime + 1*time.Second
}

func (fm *FullModel) legoUpdate() {
	fmt.Println("Модель полностью собрана!")
}

// =====================================================================================================================
// Функция для создания деталей
func createDetails(count int) []*Detail {
	details := make([]*Detail, count)

	for i := 0; i < count; i++ {
		details[i] = &Detail{id: i + 1}
	}

	return details
}

// Функция для создания бриков из деталей
func createBricks(details []*Detail) []*Brick {
	bricks := make([]*Brick, 0)

	for i := 0; i < len(details); i += 3 {
		end := i + 3
		if end > len(details) {
			end = len(details) // если не хватает деталей, создаём неполный Brick
		}
		bricks = append(bricks, &Brick{details: details[i:end]})
	}

	return bricks
}

// Функция для создания блоков из бриков
func createBlocks(bricks []*Brick) []*Block {
	blocks := make([]*Block, 0)

	for i := 0; i < len(bricks); i += 3 {
		end := i + 3
		if end > len(bricks) {
			end = len(bricks)
		}
		blocks = append(blocks, &Block{bricks: bricks[i:end]})
	}

	return blocks
}

// Функция для создания конструкций из блоков
func createConstructions(blocks []*Block) []*Construction {
	constructions := make([]*Construction, 0)

	for i := 0; i < len(blocks); i += 2 {
		end := i + 2
		if end <= len(blocks) {
			end = len(blocks)
		}
		constructions = append(constructions, &Construction{blocks: blocks[i : i+2]})
	}

	return constructions
}

// Функция для создания полной модели
func createFullModel(constructions []*Construction) *FullModel {
	return &FullModel{constructions: constructions}
}

// CONSTRUCTOR AHEAD ===================================================================================================

// start of the constructor, но это тип, а не структура!!!
type legoOption func(*LegoGo)

func withBuild(builder legoBuilder) legoOption {
	return func(lg *LegoGo) {
		lg.legoBuilder = builder
	}
}

func withUpdate(update bool) legoOption {
	return func(lg *LegoGo) {
		if update {
			fmt.Println("Lego update")
		}
	}
}

// HERE IS OUR CONSTRUCTOR NewLego
func NewLego(detail string, options ...legoOption) *LegoGo {

	legoConstructor := &LegoGo{
		detail: detail,
	}

	for _, option := range options {
		option(legoConstructor)
	}

	return legoConstructor
}

// =====================================================================================================================

func main() {
	startTime := time.Now()

	// Создаем детали, брики, блоки, конструкции и модель
	details := createDetails(50) // 50 деталей
	bricks := createBricks(details)
	blocks := createBlocks(bricks)
	constructions := createConstructions(blocks)
	fullModel := createFullModel(constructions)

	// Конструктор с функциональными опциями
	lg := NewLego(
		"legoDetail",
		withBuild(fullModel),
		withUpdate(true),
	)

	// WaitGroup для синхронизации работы горутин
	var wg sync.WaitGroup

	// Каналы для передачи данных между горутинами
	detailChan := make(chan *Detail, len(details))
	brickChan := make(chan *Brick, len(bricks))
	blockChan := make(chan *Block, len(blocks))
	constructionChan := make(chan *Construction, len(constructions))
	completeModelChan := make(chan *FullModel, 1)

	// Горутина для сборки деталей
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, detail := range details {
			detail.legoBuild()
			detail.legoUpdate()
			detailChan <- detail
		}
		close(detailChan)
	}()

	// Горутина для сборки бриков
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			detail1, ok1 := <-detailChan
			detail2, ok2 := <-detailChan
			detail3, ok3 := <-detailChan

			if !ok1 || !ok2 || !ok3 {
				break // Если канал закрыт или данных недостаточно, завершаем работу
			}

			brick := &Brick{details: []*Detail{detail1, detail2, detail3}}
			brick.legoBuild()
			brick.legoUpdate()
			brickChan <- brick
		}
		close(brickChan)
	}()

	// Горутина для сборки блоков
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			brick1, ok1 := <-brickChan
			brick2, ok2 := <-brickChan
			brick3, ok3 := <-brickChan

			if !ok1 || !ok2 || !ok3 {
				break // Если канал закрыт или данных недостаточно, завершаем работу
			}

			block := &Block{bricks: []*Brick{brick1, brick2, brick3}}
			block.legoBuild()
			block.legoUpdate()
			blockChan <- block
		}
		close(blockChan)
	}()

	// Горутина для сборки конструкций
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			block1, ok1 := <-blockChan
			block2, ok2 := <-blockChan

			if !ok1 || !ok2 {
				break // Если канал закрыт или данных недостаточно, завершаем работу
			}

			construction := &Construction{blocks: []*Block{block1, block2}}
			construction.legoBuild()
			construction.legoUpdate()
			constructionChan <- construction
		}
		close(constructionChan)
	}()

	// Горутина для сборки полной модели
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			construction1, ok1 := <-constructionChan
			construction2, ok2 := <-constructionChan

			if !ok1 || !ok2 {
				break // Если канал закрыт или данных недостаточно, завершаем работу
			}

			completeModelChan <- &FullModel{constructions: []*Construction{construction1, construction2}}
		}
		close(completeModelChan)
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Собираем финальную модель
	finalModel := <-completeModelChan
	finalModel.legoBuild()
	finalModel.legoUpdate()

	fmt.Println("Hello Lego User!", lg)

	// Выводим общее время сборки
	fmt.Printf("Общее время сборки: %v\n", time.Since(startTime))
}

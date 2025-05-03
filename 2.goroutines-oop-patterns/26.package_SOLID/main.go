package main

import (
	"fmt"
)

// S: Single Responsibility Principle (Принцип единственной ответственности)
// Каждый тип отвечает за одну задачу.

// Brick представляет собой Lego-кирпич.
type Brick struct {
	Color string
}

// NewBrick создает новый Lego-кирпич.
func NewBrick(color string) *Brick {
	return &Brick{Color: color}
}

// O: Open/Closed Principle (Принцип открытости/закрытости)
// Мы можем добавлять новые типы кирпичей, не изменяя существующий код.

// SpecialBrick представляет собой специальный Lego-кирпич.
type SpecialBrick struct {
	Brick
	Effect string
}

// NewSpecialBrick создает новый специальный Lego-кирпич.
func NewSpecialBrick(color, effect string) *SpecialBrick {
	return &SpecialBrick{
		Brick:  Brick{Color: color},
		Effect: effect,
	}
}

// L: Liskov Substitution Principle (Принцип подстановки Барбары Лисков)
// SpecialBrick может использоваться везде, где используется Brick.

// I: Interface Segregation Principle (Принцип разделения интерфейса)
// Создаем небольшие интерфейсы для разных задач.

// Placer определяет поведение для размещения кирпича.
type Placer interface {
	Place() string
}

// Place реализует интерфейс Placer для Brick.
func (b *Brick) Place() string {
	return fmt.Sprintf("Placed a %s brick", b.Color)
}

// Place реализует интерфейс Placer для SpecialBrick.
func (sb *SpecialBrick) Place() string {
	return fmt.Sprintf("Placed a %s brick with %s effect", sb.Color, sb.Effect)
}

// D: Dependency Inversion Principle (Принцип инверсии зависимостей)
// Модули верхнего уровня зависят от абстракций, а не от конкретных реализаций.

// Builder собирает Lego-модель.
type Builder struct {
	Placer Placer
}

// NewBuilder создает новый Builder.
func NewBuilder(placer Placer) *Builder {
	return &Builder{Placer: placer}
}

// BuildModel строит модель, используя Placer.
func (b *Builder) BuildModel() string {
	return b.Placer.Place()
}

func main() {
	// Создаем обычный кирпич
	redBrick := NewBrick("red")
	builder := NewBuilder(redBrick)
	fmt.Println(builder.BuildModel()) // Placed a red brick

	// Создаем специальный кирпич
	glowingBrick := NewSpecialBrick("blue", "glowing")
	builder = NewBuilder(glowingBrick)
	fmt.Println(builder.BuildModel()) // Placed a blue brick with glowing effect
}

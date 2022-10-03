package main

import (
	"fmt"
	"patterns/builder"
	chain_of_resp "patterns/chain-of-resp"
	"patterns/command"
	"patterns/facade"
	"patterns/factory"
	"patterns/state"
	"patterns/strategy"
	"patterns/visitor"
)

func main() {
	fmt.Println("============== [Паттерны проектирования] ==============")
	fmt.Println("\n==================== [Порождающие] ====================")
	fmt.Println("================== 1. [Строитель] =====================")
	builder.Builder()
	fmt.Println("================== 2. [Фабрика] =====================")
	factory.Factory()

	fmt.Println("\n==================== [Структурные] ====================")
	fmt.Println("===================== 1. [Фасад] ======================")
	facade.Facade()
	fmt.Println("\n==================== [Поведенческие] ====================")
	fmt.Println("==================== 1. [Комманда] ======================")
	command.Command()
	fmt.Println("================ 2. [Цепочка вызовов] ===================")
	chain_of_resp.ChainOfResp()
	fmt.Println("=================== 3. [Посетитель] =====================")
	visitor.Visitor()
	fmt.Println("==================== 4. [Состояние] =====================")
	state.State()
	fmt.Println("==================== 5. [Стратегия] =====================")
	strategy.Strategy()
}

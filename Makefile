all: todo

todo: contracts/Todo.sol
	@mkdir -p gen
	@solc --overwrite --bin --abi contracts/Todo.sol -o build
	@abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go

clean:
	@rm -r gen
	@rm -r build
package gojawalker

import (
	"fmt"
	"testing"

	"github.com/dop251/goja"
)

func TestWalk(t *testing.T) {
	test_data := []string{
		`function add(a, b) {
        if (a > 0) {
            return a + b;
        } else {
            return b - a;
        }
    }
    const result = add(5, 3);
    const array = [1, 2, 3];
    const obj = { name: 'John', age: 30 };
    for (let i = 0; i < array.length; i++) {
        console.log(array[i]);
    }
    while (result > 0) {
        result--;
    }
    if (obj.age > 25) {
        console.log('Adult');
    } else {
        console.log('Young');
    }
    const arrowFunc = (x) => x * 2;
    class Person {
        constructor(name) {
            this.name = name;
        }
        sayHello() {
            console.log('Hello, ' + this.name);
        }
    }
    const person = new Person('Alice');
    person.sayHello();
    const asyncFunc = async () => {
        await new Promise((resolve) => setTimeout(resolve, 1000));
        console.log('Async function completed');
    };
    asyncFunc();`,
		`// 函数和条件语句
    function max(a, b) {
        return (a > b)? a : b;
    }
    let value = max(10, 20);

    // 数组和对象
    const arr = [1, 2, 3, 4];
    const obj = { name: 'Jane', age: 28 };

    // 循环
    for (let i in arr) {
        console.log(arr[i]);
    }
    for (let item of arr) {
        console.log(item);
    }

    // 箭头函数
    const square = x => x * x;
    console.log(square(5));

    // 类和继承
    class Animal {
        constructor(name) {
            this.name = name;
        }
        speak() {
            console.log(this.name +'makes a noise.');
        }
    }
    class Dog extends Animal {
        speak() {
            super.speak();
            console.log(this.name +'barks.');
        }
    }
    let dog = new Dog('Rex');
    dog.speak();

    // 模板字符串
    let greeting = "Hello, ${dog.name}";
    console.log(greeting);

    // 解构赋值
    const [first,, third] = arr;
    console.log(first, third);

    // 异常处理
    try {
        throw new Error('Something went wrong');
    } catch (e) {
        console.error(e);
    }

    // 异步函数
    async function fetchData() {
        const response = await fetch('https://jsonplaceholder.typicode.com/posts/1');
        const data = await response.json();
        console.log(data);
    }
    fetchData();`,
	}
	ast_tree, err := goja.Parse("test.js", test_data[1])
	if err != nil {
		fmt.Println(ast_tree, err)
		t.Fail()
	}

	visitor := ExampleVisitor{}
	Walk(visitor, ast_tree)
}

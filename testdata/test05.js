const arrowFunc = (x) => x * 2;
asyncFunc();

const asyncFunc2 = async () => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    console.log('Async function completed');
};
asyncFunc2(); `,`// 函数和条件语句
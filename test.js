function test1(){
    console.log(1)
    for (let i = 0; i < 1000000000; i++) {
        i+1*10/2-1
    }
    console.log(2)
}

function test2(){
    console.log(3)
    for (let i = 0; i < 10000; i++) {
        i+1*10/2-1
    }
    console.log(4)
}

setTimeout(() => {
    test1()
}, 1);

setTimeout(() => {
    test2()
}, 1);
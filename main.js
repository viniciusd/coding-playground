(async (global) => {
    function wasmFactory(id) {
        let tempContext = {
            Module: {
                print: function() {
                    let text = Array.prototype.slice.call(arguments).join(' ');
                    console.log(id, text);
                    document.getElementById("result").innerHTML += "<br>\n" + text;
                },
                printErr: console.error
            }
        };
        return loadWasm.call(tempContext);
    }

    global.wasmFactory = wasmFactory;

    let wasmWrapper = await (await fetch('hello.js')).text();
    let loadWasm = new Function(`
        var Module = this.Module;
        let _resolve;
        let promise = new Promise((resolve, reject) => _resolve = resolve);
        Module.postRun = (Module.postRun || []).concat([
            _resolve.bind(promise, Module)
        ]); 
        ${wasmWrapper};
        return promise; 
    `);

})(window);

var hello1, hello2;

function click1() {
    if (typeof hello1 === 'undefined') {
        wasmFactory(1).then(m => hello1 = m);
    } else {
        hello1._hello()
    }
}

function click2() {
    if (typeof hello2 === 'undefined') {
        wasmFactory(1).then(m => hello2 = m);
    } else {
        hello2._hello()
    }
}

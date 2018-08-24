(async (window) => { 
    var wasmWrapper = await (await fetch('hello.js')).text();

    let tempContext = Object.assign(new (function x() {})(), {
        Module: {
            print: (function() {
                return function(text) {
                    if (arguments.length > 1) text = Array.prototype.slice.call(arguments).join(' ');
                    console.log(text);
                    document.getElementById("result").innerHTML += "<br>\n" + text;
                };
            })(),
            printErr: console.error
        } 
    });

    (function(str) {
      return eval('var Module = this.Module;'+str);
    }.call(tempContext, wasmWrapper));

    window.hello = tempContext.Module
})(window)


function click1() {
    hello._hello()
}

function click2() {
    hello._hello()
}

async function moduleFactory() { 

    var wasmWrapper = await (await fetch('hello.js')).text();

    let tempContext = Object.assign(new (function _thisContext() {})(), {
        Module: {
            print: function() {
                let text = Array.prototype.slice.call(arguments).join(' ');
                console.log(text);
                document.getElementById("result").innerHTML += "<br>\n" + text;
            },
            printErr: console.error
        } 
    });

    (function(str) {
      return eval('var Module = this.Module;'+str);
    }.call(tempContext, wasmWrapper));

    return tempContext.Module;
}

var hello1, hello2;
moduleFactory().then((m) => hello1 = m);
moduleFactory().then((m) => hello2 = m);

function click1() {
    hello1._hello()
}

function click2() {
    hello2._hello()
}

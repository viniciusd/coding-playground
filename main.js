async function moduleFactory(id) {

    var wasmWrapper = await (await fetch('hello.js')).text();

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

    (function(str) {
      return eval('var Module = this.Module;'+str);
    }.call(tempContext, wasmWrapper));

    return tempContext.Module;
}

var hello1, hello2;
moduleFactory(1).then((m) => hello1 = m);
moduleFactory(2).then((m) => hello2 = m);

function click1() {
    hello1._hello()
}

function click2() {
    hello2._hello()
}

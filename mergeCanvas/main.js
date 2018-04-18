function onLoad() {
    var img = new Image();
    img.onload = start;
    img.src = "js.png";

    function start(){

        var mainCanvas = document.getElementById('theCanvas'),
            ctx = mainCanvas.getContext('2d'),
            canvas1 = document.createElement('canvas'),
            ctx1 = canvas1.getContext('2d'),
            canvas2 = document.createElement('canvas'),
            ctx2 = canvas2.getContext('2d');

        canvas1.width = mainCanvas.width/2;
        canvas1.height = mainCanvas.height;
        canvas2.width = mainCanvas.width/2;
        canvas2.height = mainCanvas.height;

        ctx1.drawImage(img, 0, 0, canvas1.width, canvas1.height);
        ctx2.drawImage(img, 0, 0, canvas2.width, canvas2.height);

        ctx.drawImage(canvas1, 0, 0, canvas1.width, canvas1.height);
        ctx.drawImage(canvas2, canvas1.width, 0,
                      canvas2.width, canvas2.height);
    }
}

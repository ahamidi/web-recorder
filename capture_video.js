var system = require('system');
var args = system.args;

var page = require('webpage').create(); 
page.clipRect = { top: 0, left: 0, width: 1600, height: 900};
page.viewportSize = { width: 1600, height: 900};
page.zoomFactor = 0.75;

var fps = 10;

var frames = args[2]*fps;

if (args.length === 1) {
  console.log('Not enough parameters');
} else {
  page.open(args[1], function () {
    page.scrollPosition = {
      top: 0,
      left: 0
    };

    // Play
    page.evaluate(function(){
      return document.getElementsByClassName('play-btn')[0].click();
    })
    delay = 1000/fps;
    setInterval(function() {
      page.render('/dev/stdout', { format: "png", quality: "50" });
      if (frames === 0) {
        phantom.exit();
      }
      frames--
    }, delay);

  });
}



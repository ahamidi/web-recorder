var system = require('system');
var args = system.args;

var page = require('webpage').create();
page.viewportSize = { width: 1024, height: 800};

if (args.length === 1) {
  console.log('Try to pass some arguments when invoking this script!');
} else {
  //page.onResourceRequested = function (request) {
    //console.log('Request ' + JSON.stringify(request, undefined, 4));
  //}
  page.onError = function (msg, trace) {
    console.log(msg);
    trace.forEach(function(item) {
      console.log('   ', item.file, ':', items.line);
    })
  }
  page.open(args[1], function(status) {
    if (status !== 'success') {
      console.log('Unable to load the address!');
      phantom.exit(1);
    } else {
      //console.log('Content: ' + page.content);
      page.render('screenshot.png');
      phantom.exit();
    }
  });
}



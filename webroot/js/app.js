/**
 * @author denghb
 * @since 2021-09-11
 */
var el = function(id) {
  return document.getElementById(id);
}

el("last-time").onchange = function() {
  if (this.value === '') {
    el("day-range").style.display = "inline";
  } else {
    el("day-range").style.display = "none";
  }
}

var now = new Date();
var month = now.getMonth() + 1;
if (month < 10) {
    month = "0" + month;
}
el("day").value = now.getFullYear() + "-" + month + "-" + now.getDate();
el("time-start").value = "00:00";
el("time-end").value = "23:59";


var getHeatmap = function() {
  var heatmap = window.heatmap_k104;
  if (!heatmap) {
    heatmap = h337.create({
      radius: 60,
      container: el("k104")
    });
    window.heatmap_k104 = heatmap;
  }
  return heatmap;
}

var httpGet = function(url, callback) {
  // request server
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      var res = JSON.parse(this.responseText);
      callback(res);
    }
  };
  xhr.open("GET", url, true);
  xhr.send();
}
var formatNumber = function (num){
  return (num || 0).toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,');
}
var doDraw = function() {
  // request server
  var url = "/km?lastTime=" + el('last-time').value;
  url += "&day=" + el('day').value;
  url += "&timeStart=" + el('time-start').value;
  url += "&timeEnd=" + el('time-end').value;

  httpGet(url, function(res) {
    // console.log(res);
    el('total-keyboard').innerText = formatNumber(res.keyboard.total);
    el('total-mouse').innerText = formatNumber(res.mouse.total);

    var stat = res.keyboard.stat;
    var data = [];
    var max = 0, max2 = 0;
    for (var key in stat) {
      if (key.length === 1) {
        key = key.toUpperCase();
      }
      var point = KEYBOARD.k104[key];
      if (undefined === point) {
        console.log('"%o" undefined', key);
        continue;
      }
      point.value = stat[key];
      if (max < point.value) {
        max2 = max;
        max = point.value;
      }
      data.push(point);
    }

    //  > 20%
    if (max2 > 0 && max > max2 * 1.2) {
      max = max2 * 1.2;
    }
    var heatmap = getHeatmap();
    heatmap.setData({
      max: max,
      data: data
    });
  });
}
// doDraw()
setInterval("doDraw()", 200);

// show all
var showAll = function() {
  var heatmap = getHeatmap();
  var data = [];
  for (var key in arr) {
    var point = KEYBOARD.k104[key];
    point.value = 5;
    data.push(point);
  }

  heatmap.setData({
    max: 5,
    data: data
  });
}
// showAll();

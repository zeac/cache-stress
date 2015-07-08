$(window).load(function() {
  makeRequest = function (seed, size, i, count) {
    if (i < count) {
      var url = "cache/3600/" + size + "/" + seed + i;
      $.get (url, function (data) {
        console.log("Received " + url);
        makeRequest(seed, size, i + 1, count);
      });
    }
  }

  var form = $("#form");
  var seedElement = form.find("input[name='seed']");
  var sizeElement = form.find("input[name='size']");
  var countElement = form.find("input[name='count']");

  updateSubmitValue = function () {
    var count = countElement.val();
    var size = sizeElement.val();
    var seed = seedElement.val();

    $("#submit").val("Request " + count + " files with " + size + "Kb size and seed = " + seed);
  }

  updateSubmitValue();

  form.find("input").change(updateSubmitValue);
  form.submit(function(ev) {
    ev.preventDefault();

    var count = countElement.val();
    var size = sizeElement.val();
    var seed = seedElement.val();
    makeRequest(seed, size * 1024, 0, count);
  });
});

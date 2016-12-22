$(document).ready(function() {
  makeSomething()
})

var data = "this is a test".split("")

function makeSomething() {
  var vis = d3.select("#vis")

  var poly = d3.select(".hex")
  poly.attr("fill", "#666")

  var polyName = poly.property("nodeName");

  vis.append(polyName)
  .attr("fill", "#00ff00")
  .attr("x", 100)
}

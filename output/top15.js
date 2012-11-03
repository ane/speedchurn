var width = 500;
var height = 500;
var radius =175;
var color = d3.scale.category20b()

var arc = d3.svg.arc()
    .outerRadius(radius)

var pie = d3.layout.pie()
    .sort(null)
    .value(function(d) { return d.stats.lines });

var svg = d3.select("#top15").append("svg")
    .attr("width", width)
    .attr("height", height)
    .append("g")
    .attr("transform", "translate(" + width / 2 + "," + height / 2 + ")");

d3.json("data/" + chan + "_top15.json", function(data) {

  var g = svg.selectAll(".arc")
      .data(pie(data))
    .enter().append("g")
      .attr("class", "arc");

  g.append("path")
      .attr("d", arc)
      .style("fill", function(d,i) { return color(i); });

  g.append("text")
      .attr("transform", function(d) { 
        //we have to make sure to set these before calling arc.centroid
        d.outerRadius = radius + 75; // Set Outer Coordinate
        d.innerRadius = radius + 75; // Set Inner Coordinate
        return "translate(" + arc.centroid(d) + ")";
	  })
      .attr("dy", ".35em")
      .style("text-anchor", "middle")
      .style("fill", "#330570")
      .style("font", "bold 13px Arial")
      .text(function(d, i) { return (i+1) + ". " + d.data.nick });
});


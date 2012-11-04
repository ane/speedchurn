google.load("visualization", "1", {packages:["corechart"]});
google.setOnLoadCallback(drawChart);

function drawChart() {
	var data = new google.visualization.DataTable();

	var jsonData = $.ajax({ 
		url: "data/" + chan + "_top15.json",
		dataType:"json",
		async: false
	}).responseText;

	var options = {
		title: 'User Statistics',
		width: 550,
		chartArea: {'width': '90%', 'height': '80%'},
		height: 400,
	};

	data.addColumn("string", "Nick")
	data.addColumn("number", "Lines")

	var blob = eval(jsonData);
	blob.forEach(function(e, idx, arr) {
		data.addRow([e.nick, e.stats.lines]);
	})

	var chart = new google.visualization.PieChart(document.getElementById('top15'));
	chart.draw(data, options);
}


google.setOnLoadCallback(drawActivity);

function drawActivity() {
	var data = new google.visualization.DataTable();

	var jsonData = $.ajax({ 
		url: "data/" + chan + "_daily_activity.json",
		dataType:"json",
		async: false
	}).responseText;

	var options = {
		title: 'Daily activity',
		width: 940,
		chartArea: {'width': '80%', 'height': '80%'},
		height: 400,
	};

	data.addColumn("number", "Day")
	data.addColumn("number", "Lines")

	var blob = eval(jsonData);
	blob.forEach(function(e, idx, arr) {
		data.addRow([e.order, e.lines]);
	})

	var chart = new google.visualization.AreaChart(document.getElementById('daily'));
	chart.draw(data, options);
}


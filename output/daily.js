google.setOnLoadCallback(drawActivity);

function drawActivity() {
	var data = new google.visualization.DataTable();

	var jsonData = $.ajax({ 
		url: "data/" + chan + "_daily_activity.json",
		dataType:"json",
		async: false
	}).responseText;

	var options = {
		title: 'Daily activity (total number of lines per day)',
		width: '100%',
		chartArea: {'width': '85%', 'height': '80%'},
		height: 400,
	};

	data.addColumn("date", "Date")
	data.addColumn("number", "Lines")

	var blob = eval(jsonData);
	blob.forEach(function(e, idx, arr) {
		data.addRow([new Date(e.date), e.lines]);
	})

	var chart = new google.visualization.AreaChart(document.getElementById('daily'));
	chart.draw(data, options);
}


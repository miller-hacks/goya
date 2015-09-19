$(function() {
	$('#detectButton').on('click', function() {
		var imageUrl = $('#imageUrlInput').val();
		console.log(imageUrl);
		if (imageUrl.length) {
			var data = JSON.stringify({url: imageUrl});

			$.ajax({
				type: 'POST',
				url: '/',
				contentType: "application/json; charset=utf-8",
				data: data,
				success: function(data) {
					console.log(data);
				}
			})
		}
	});
});
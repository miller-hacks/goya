$(function() {
	$('#detectButton').on('click', function() {
		var imageUrl = $('#imageUrlInput').val();
		console.log(imageUrl);
		if (imageUrl.length) {
			$.ajax('/', {imageUrl: imageUrl}, function(data) {
				console.log(data);
			})
		}
	});
});
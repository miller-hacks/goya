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
					(function draw(imageUrl, faces) {
						var ctx = document.getElementById('canvas').getContext('2d');
						var img = new Image();

						img.onload = function() {
							console.log('adada');
							console.log(imageUrl);
							for (var idx=0; idx<=faces.len; idx++) {
								console.log(idx);
								var obj = faces[idx];
								ctx.drawImage(img,0,0);
								ctx.beginPath();
								ctx.moveTo(obj.x, obj.y);
								
								ctx.lineTo(obj.x + obj.height, obj.y);
								ctx.lineTo(obj.x, obj.y + obj.width);

								ctx.stroke();
							}
						}
						img.src = imageUrl;							
					})(imageUrl, data.faces)
				}
			})
		}
	});
});
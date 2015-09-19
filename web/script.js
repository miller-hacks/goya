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
					$('#image').attr('src', imageUrl);

					(function draw(imageUrl, faces) {
						var canvas = document.getElementById('canvas');
						var ctx = canvas.getContext('2d');
						var img = new Image();
						
						img.onload = function() {
							canvas.width = 1200;
							canvas.height = 1200;

							ctx.drawImage(img, 0, 0);
							ctx.strokeStyle = 'white';
							ctx.lineWidth=2.5;

							for (var idx=0; idx<faces.length; idx++) {							
								var obj = faces[idx];
								ctx.beginPath();
								ctx.moveTo(obj.x, obj.y);
								ctx.lineTo(obj.x + obj.width, obj.y);
								ctx.lineTo(obj.x + obj.width, obj.y + obj.height);
								ctx.lineTo(obj.x, obj.y + obj.height);
								ctx.lineTo(obj.x, obj.y);

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
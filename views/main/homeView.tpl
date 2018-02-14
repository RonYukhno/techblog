{{ template "templates/basic-layout.tpl" .}}

{{ define "content" }}
<script>
	function add_item() {
		var content = document.getElementById('content');
		var div = document.createElement('div');
		div.className = "wrapper";
		div.innerHTML = `<a href="#">
						<div class="item_img">
							<img src="http://via.placeholder.com/350x200" alt="">
							<div class="img-cover"></div>
							<div class="category">Новости</div>
							<h2 class="item_title">Заголовок</h2>
						</div>
						<div class="item_wrap">
							<p class="item_description">Lorem ipsum dolor sit amet consectetur adipisicing elit. Placeat harum iure commodi animi inventore. Officiis explicabo rerum possimus deserunt repudiandae, iure laudantium distinctio Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni obcaecati consequatur temporibus explicabo laboriosam, labore totam, earum, cum neque, consequuntur odit sed magnam illo in quos deserunt eaque dolor. Suscipit?</p>
							<p class="item_date">12.02.2018 <span id="time">13:21</span></p>
							<p class="item_author">Автор</p>
						</div>
						</a>`;
		content.appendChild(div);
	}
	for (i = 0; i < 6; i++) {
		add_item();
	}
</script>
{{ end }}

$(function () {

	$('.aside h4').click(function () {

		//		$(this).toggleClass('active');


		$(this).siblings('ul').slideToggle();
	})

	$('.deleteConfirm').click(function () {
		return confirm("确认删除" + this.title + "吗？")
	})

	$('.changeStatusButton').click(function () {
		var id = $(this).attr("data-id")
		var table_name = $(this).attr("data-table_name")
		var status_value = $(this).attr("data-status_value")
		var el = $(this)
		$.get("changeStatus", { id: id, table_name: table_name, status_value: status_value }, function (response) {
			if (response.success) {
				el.attr("data-status_value", Math.abs(el.attr("data-status_value") - 1))
				if (el.attr("src").indexOf("yes") != -1) {
					el.attr("src", "/static/images/no.gif")
				} else {
					el.attr("src", "/static/images/yes.gif")
				}
			} else {
				console.log(response)
			}
		})
	})

})

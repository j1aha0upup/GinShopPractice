(function ($) {

    var app = {
        init: function () {

            this.initSwiper();

            this.initNavSlide();
            this.initVersion();
            this.initColor();
        },
        initSwiper: function () {
            new Swiper('.swiper-container', {
                loop: true,
                navigation: {
                    nextEl: '.swiper-button-next',
                    prevEl: '.swiper-button-prev'
                },
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true
                }

            });
        },
        initNavSlide: function () {
            $("#nav_list>li").hover(function () {

                $(this).find('.children-list').show();
            }, function () {
                $(this).find('.children-list').hide();
            })

        },
        initVersion: function () {
            $('.detail_info .detail_info_item:first').addClass('active');
            $('.detail_list li:first').addClass('active');
            $('.detail_list li').click(function () {
                var index = $(this).index();
                $(this).addClass('active').siblings().removeClass('active');
                $('.detail_info .detail_info_item').removeClass('active').eq(index).addClass('active');

            })
        },
        initColor: function () {
            var that = this;
            $('#color_list .banben').first().addClass('active');
            $('#color_list .banben').click(function () {
                $(this).addClass('active').siblings().removeClass('active');
                var goodsId = $(this).attr("goodsId");
                var colorId = $(this).attr("colorId");
                $('#GoodsColor').val($(this).find(".yanse").text());
                $.get("/product/getImageList",
                    { "goodsId": goodsId, "colorId": colorId },
                    function (response) {
                        if (response["status"] == "success") {
                            var newFocus = "";
                            for (var i = 0; i < response["imageList"].length; i++) {
                                newFocus += `<div class="swiper-slide"><img src="` + response["imageList"][i]["ImageUrl"] + `" style="width: 600px;height: 600px;"></div>`
                            }
                            $("#item-focus").html(newFocus);
                            that.initSwiper();
                        }
                    }
                )
            })
        },
    }
    $(function () {
        app.init();
    })
})($)

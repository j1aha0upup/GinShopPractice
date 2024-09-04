var app = {
    init: function () {
        this.getCaptcha();
        this.captchaClick();
    },
    getCaptcha: function () {
        $.get("get_captcha?t=" + Math.random()).then(function (result) {
            $("#captchaId").attr("value", result.captchaId);
            $("#captchaImg").attr("src", result.captchaB64s);
        })
    },
    captchaClick: function () {
        that = this;
        $("#captchaImg").click(function () {
            that.getCaptcha();
        });
    }

}
$(function () {
    app.init()
})
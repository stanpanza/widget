!function(i){"function"==typeof define&&define.amd?define(["jquery"],i):i("object"==typeof exports?require("jquery"):jQuery)}(function(i){"use strict";function t(e,n){this.$element=i(e),this.options=i.extend({},t.DEFAULTS,i.isPlainObject(n)&&n),this.init()}var e="qor.widget",n="enable."+e,o="disable."+e,r="click."+e,d=".qor-widget-button, .qor-slideout__close",a="";return t.prototype={constructor:t,init:function(){this.bind(),this.initStatus()},bind:function(){this.$element.on(r,i.proxy(this.click,this))},initStatus:function(){i("body").append('<iframe id="qor-widget-iframe" src="'+a+'"></iframe>'),i("body").append('<iframe id="qor-widget-inline-iframe" src="http://localhost:7000/admin/widgets/BannerEditor"></iframe>'),i("#qor-widget-inline-iframe").load(function(){var t=i("#qor-widget-inline-iframe").contents().find(".qor-form-container"),e=t.outerHeight();i("#qor-widget-inline-iframe").height(e),i("#qor-widget-inline-iframe").contents().find("header").remove(),t.css("margin",0),i("#qor-widget-inline-iframe").css({border:"1px solid #eee"}),i("#qor-widget-inline-iframe").contents().find("body").css({overflow:"hidden"})})},click:function(t){var e=i(t.target);if(t.stopPropagation(),e.is(d))if(e.data("is-inline-edit")){var n=e.parents(".qor-widget");n.find("*").hide(),e.parents(".qor-widget").append(i("#qor-widget-inline-iframe")),i("#qor-widget-inline-iframe").width("100%")}else i("#qor-widget-iframe").contents().find(".js-widget-edit-link").attr("data-url",e.data("url")),i("#qor-widget-iframe").addClass("show"),i("body").addClass("open-widget-editor")}},t.plugin=function(n){return this.each(function(){var o,r=i(this),d=r.data(e);if(!d){if(/destroy/.test(n))return;r.data(e,d=new t(this,n))}"string"==typeof n&&i.isFunction(o=d[n])&&o.apply(d)})},t.isScrollToBottom=function(i){return i.clientHeight+i.scrollTop===i.scrollHeight},i(function(){i("body").attr("data-toggle","qor.widgets"),i(".qor-widget").each(function(t,e){var n=i(e).find("*").eq(0);a=i(e).data("widget-inline-edit-url"),n.css("position","relative").addClass("qor-widget").attr("data-url",i(e).data("url")).unwrap(),n.append('<div class="qor-widget-embed-wrapper"><button data-is-inline-edit="'+i(e).data("is-inline-edit")+'" data-url="'+i(e).data("url")+'" class="qor-widget-button">Edit</button></div>')}),window.closeWidgetEditBox=function(){i("#qor-widget-iframe").removeClass("show"),i("#qor-widget-iframe")[0].contentWindow.location.reload(),i("body").removeClass("open-widget-editor")};var e='[data-toggle="qor.widgets"]';i(document).on(o,function(n){t.plugin.call(i(e,n.target),"destroy")}).on(n,function(n){t.plugin.call(i(e,n.target))}).triggerHandler(n)}),t});
//# sourceMappingURL=widget_inline_edit.js.map

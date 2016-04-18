var incr = require('./incr');

$(function () {
    (function () {
        var $button = $('#incr-button');
        var $field = $('#incr-field');
        var count = parseInt($field.data('default'));
        var render = function () {
            $field.text(count);
        };
        $button.click(function () {
            count = incr(count);
            render();
        });
        render();
    })();
});

var incr = require('./incr');

$(function () {
    (function () {
        var count = 0;
        var $button = $('#incr-button');
        var $field = $('#incr-field');
        $button.click(function () {
            count = incr(count);
            $field.text(count);
        });
    })();
});

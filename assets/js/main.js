var incr = require('./incr');

$(function () {
    (function () {
        var count = 0;
        var $button = $('#incr-button');
        var $field = $('#incr-field');
        $button.click(function () {
                count = incr(parseInt($field.text()));
            $field.text(count);
        });
    })();
});

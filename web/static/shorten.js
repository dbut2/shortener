function shorten() {

    $.post('api/shorten', {"url":$('#url').val()}, function(data) {
        console.log(data);
        $('#url').val(window.location.href + "ort/" + JSON.parse(data)['code']).select();
    });
    return false
}
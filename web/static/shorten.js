function shorten() {

    $.post('api/shorten', {"url":$('#url').val(),"code":$('#code').val()}, function(data) {
        $('#link').val(window.location.href + "ort/" + $('#code').val());
    });

    return false;
}
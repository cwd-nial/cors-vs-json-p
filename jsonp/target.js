function callback(json) {
    console.log(json);
}

$.ajax({
    url: "http://localhost:8091/info",
    crossDomain: true,
    dataType: 'jsonp',
    type: 'GET',
    cache: false,
    jsonpCallback: "callback",
    error: function (xhr, textStatus, errorThrown) {
        console.log("nay!")
        console.log(xhr)
        console.log(textStatus)
        console.log(errorThrown)
    }
});

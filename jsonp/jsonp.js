function aliasGatewayCallback(json) {
    console.log(json);
    window.location = "/info";
}

$.post({
    url: "http://[::1]:8091/info",
    crossDomain: true,
    dataType: 'jsonp',
    cache: false,
    jsonpCallback: "aliasGatewayCallback",
    withCredentials: true,
    xhrFields: {
        'withCredentials': true // tell the client to send the cookies if any for the requested domain
    },
    error: function (xhr, textStatus, errorThrown) {
        console.log("nay!")
        console.log(xhr)
        console.log(textStatus)
        console.log(errorThrown)
    }
});

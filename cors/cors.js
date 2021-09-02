$.post({
    url: "http://[::1]:8091/cors-target",
    crossDomain: true,
    dataType: 'json',
    cache: false,
    withCredentials: true,
    xhrFields: {
        'withCredentials': true // tell the client to send the cookies if any for the requested domain
    },
    success: function (data) {
        console.log("yay!")
        console.log(data)
    },
    error: function (xhr, textStatus, errorThrown) {
        console.log("nay!")
        console.log(xhr)
        console.log(textStatus)
        console.log(errorThrown)
    }
});

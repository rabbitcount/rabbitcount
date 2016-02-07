function createUser() {
    //alert("rabbit: " + $("#taste").serialize());

    //$.post("/api/user/create", $("#taste").serialize());

    var formData = new FormData($("#taste")[0])
    //$.post("/api/user/create", formData);
    $.ajax({
        url: '/api/user/create',
        type: 'POST',
        data: formData,
        async: false,
        cache: false,
        contentType: false,
        processData: false,
        success: function (returndata) {
            alert(returndata);
        },
        error: function (returndata) {
            alert(returndata);
        }
    });
}